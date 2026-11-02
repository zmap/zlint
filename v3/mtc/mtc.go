package mtc

import (
	"bytes"
	"crypto/sha256"
	"errors"

	"github.com/zmap/zcrypto/encoding/asn1"
	"github.com/zmap/zcrypto/x509/pkix"
	"math/bits"

	"github.com/zmap/zcrypto/x509"
	"golang.org/x/crypto/cryptobyte"
	cbasn1 "golang.org/x/crypto/cryptobyte/asn1"
)

var ProofAlgorithmOID = asn1.ObjectIdentifier{1, 3, 6, 1, 5, 5, 7, 6, 67}
var CAExtensionOID = asn1.ObjectIdentifier{1, 3, 6, 1, 5, 5, 7, 1, 38}

var proofAlgorithmDER = []byte{0x30, 0x0a, 0x06, 0x08, 0x2b, 0x06, 0x01, 0x05, 0x05, 0x07, 0x06, 0x43}

type Proof struct {
	EntryExtensions []byte
	Start           uint64
	End             uint64
	InclusionProof  []byte
}

type Certificate struct {
	Proof          Proof
	Serial         uint64
	TBSCertificate []byte
}

func IsCACertificate(c *x509.Certificate) bool {
	return c.IsCA && extension(c, CAExtensionOID) != nil
}

func ValidateCA(c *x509.Certificate) error {
	extension := extension(c, CAExtensionOID)
	if extension == nil {
		return errors.New("MTC CA extension is absent")
	}
	if !extension.Critical {
		return errors.New("MTC CA extension is not critical")
	}
	if !c.BasicConstraintsValid || !c.IsCA {
		return errors.New("MTC CA basic constraints are invalid")
	}
	if c.KeyUsage&x509.KeyUsageCertSign == 0 {
		return errors.New("MTC CA key usage lacks keyCertSign")
	}
	var sequence cryptobyte.String
	contents := cryptobyte.String(extension.Value)
	if !contents.ReadASN1(&sequence, cbasn1.SEQUENCE) || !contents.Empty() {
		return errors.New("malformed MTC CA extension")
	}
	var algorithm cryptobyte.String
	var minSerial, maxSerial uint64
	if !sequence.ReadASN1Element(&algorithm, cbasn1.SEQUENCE) ||
		!sequence.ReadASN1Integer(&minSerial) ||
		!sequence.ReadASN1Integer(&maxSerial) || !sequence.Empty() {
		return errors.New("malformed MTC CA extension")
	}
	if len(algorithm) == 0 {
		return errors.New("MTC CA signature algorithm is absent")
	}
	if minSerial < 1<<48 || maxSerial < 1<<48 || minSerial > maxSerial {
		return errors.New("MTC CA serial range is invalid")
	}
	return nil
}

func extension(c *x509.Certificate, oid asn1.ObjectIdentifier) *pkix.Extension {
	for _, extension := range c.Extensions {
		if extension.Id.Equal(oid) {
			return &extension
		}
	}
	return nil
}

func IsProofCertificate(c *x509.Certificate) bool {
	input := cryptobyte.String(c.Raw)
	var cert, tbs, alg cryptobyte.String
	return input.ReadASN1(&cert, cbasn1.SEQUENCE) && input.Empty() &&
		cert.ReadASN1(&tbs, cbasn1.SEQUENCE) &&
		cert.ReadASN1Element(&alg, cbasn1.SEQUENCE) &&
		cert.SkipASN1(cbasn1.BIT_STRING) && cert.Empty() &&
		bytes.Equal(alg, proofAlgorithmDER)
}

func ParseCertificate(c *x509.Certificate) (*Certificate, error) {
	signature, ok := signatureAlgorithm(c.Raw)
	if !ok {
		return nil, errors.New("invalid MTC signature algorithm")
	}
	proof, err := parseProof(signature)
	if err != nil {
		return nil, err
	}
	serial, err := parseTBS(c.RawTBSCertificate)
	if err != nil {
		return nil, err
	}
	return &Certificate{Proof: *proof, Serial: serial, TBSCertificate: c.RawTBSCertificate}, nil
}

func ValidateSelfContained(c *x509.Certificate, parsed *Certificate) error {
	if parsed.Serial>>48 == 0 {
		return errors.New("serial number has log number zero")
	}
	index := parsed.Serial & (1<<48 - 1)
	if index < parsed.Proof.Start || index >= parsed.Proof.End {
		return errors.New("serial number index is outside the proof subtree")
	}
	entry, err := reconstructEntry(c.RawTBSCertificate, parsed.Proof.EntryExtensions)
	if err != nil {
		return err
	}
	leaf := sha256.Sum256(entry)
	_, err = EvaluateInclusionProof(index, parsed.Proof.Start, parsed.Proof.End, leaf, parsed.Proof.InclusionProof)
	return err
}

func signatureAlgorithm(raw []byte) ([]byte, bool) {
	input := cryptobyte.String(raw)
	var cert, tbs, alg cryptobyte.String
	var signature []byte
	if !input.ReadASN1(&cert, cbasn1.SEQUENCE) || !input.Empty() ||
		!cert.ReadASN1(&tbs, cbasn1.SEQUENCE) ||
		!cert.ReadASN1Element(&alg, cbasn1.SEQUENCE) ||
		!cert.ReadASN1BitStringAsBytes(&signature) || !cert.Empty() ||
		!bytes.Equal(alg, proofAlgorithmDER) {
		return nil, false
	}
	return signature, true
}

func parseProof(in []byte) (*Proof, error) {
	s := cryptobyte.String(in)
	var extensions, inclusionProof, signatures cryptobyte.String
	var p Proof
	if !s.ReadUint16LengthPrefixed(&extensions) || !s.ReadUint48(&p.Start) || !s.ReadUint48(&p.End) ||
		!s.ReadUint16LengthPrefixed(&inclusionProof) || !s.ReadUint24LengthPrefixed(&signatures) || !s.Empty() {
		return nil, errors.New("malformed MTC proof")
	}
	p.EntryExtensions = append([]byte(nil), extensions...)
	lastType := -1
	for !extensions.Empty() {
		var typ uint16
		var value cryptobyte.String
		if !extensions.ReadUint16(&typ) || !extensions.ReadUint16LengthPrefixed(&value) || int(typ) <= lastType {
			return nil, errors.New("malformed MTC entry extensions")
		}
		lastType = int(typ)
	}
	var previous []byte
	for !signatures.Empty() {
		var id, value cryptobyte.String
		if !signatures.ReadUint8LengthPrefixed(&id) || len(id) == 0 || !signatures.ReadUint16LengthPrefixed(&value) ||
			(previous != nil && compareIDs(previous, id) >= 0) {
			return nil, errors.New("malformed MTC cosignatures")
		}
		previous = id
	}
	if !validSubtree(p.Start, p.End) {
		return nil, errors.New("invalid MTC proof subtree")
	}
	p.InclusionProof = append([]byte(nil), inclusionProof...)
	return &p, nil
}

func parseTBS(raw []byte) (uint64, error) {
	s := cryptobyte.String(raw)
	var tbs cryptobyte.String
	if !s.ReadASN1(&tbs, cbasn1.SEQUENCE) || !s.Empty() || !tbs.SkipOptionalASN1(cbasn1.Tag(0).Constructed().ContextSpecific()) {
		return 0, errors.New("malformed TBSCertificate")
	}
	var serial uint64
	var algorithm cryptobyte.String
	if !tbs.ReadASN1Integer(&serial) || !tbs.ReadASN1Element(&algorithm, cbasn1.SEQUENCE) || !bytes.Equal(algorithm, proofAlgorithmDER) {
		return 0, errors.New("invalid MTC TBSCertificate")
	}
	return serial, nil
}

func reconstructEntry(rawTBS, extensions []byte) ([]byte, error) {
	s := cryptobyte.String(rawTBS)
	var tbs cryptobyte.String
	if !s.ReadASN1(&tbs, cbasn1.SEQUENCE) || !s.Empty() {
		return nil, errors.New("malformed TBSCertificate")
	}
	entry := make([]byte, 0, len(rawTBS)+len(extensions))
	entry = append(entry, byte(len(extensions)>>8), byte(len(extensions)))
	entry = append(entry, extensions...)
	entry = append(entry, 0, 1)
	if tbs.PeekASN1Tag(cbasn1.Tag(0).Constructed().ContextSpecific()) {
		var version cryptobyte.String
		if !tbs.ReadASN1Element(&version, cbasn1.Tag(0).Constructed().ContextSpecific()) {
			return nil, errors.New("malformed TBSCertificate version")
		}
		entry = append(entry, version...)
	}
	if !tbs.SkipASN1(cbasn1.INTEGER) || !tbs.SkipASN1(cbasn1.SEQUENCE) {
		return nil, errors.New("malformed TBSCertificate serial or signature")
	}
	for range 3 {
		var field cryptobyte.String
		if !tbs.ReadASN1Element(&field, cbasn1.SEQUENCE) {
			return nil, errors.New("malformed TBSCertificate name or validity")
		}
		entry = append(entry, field...)
	}
	var spki, algorithm cryptobyte.String
	if !tbs.ReadASN1(&spki, cbasn1.SEQUENCE) || !spki.ReadASN1Element(&algorithm, cbasn1.SEQUENCE) {
		return nil, errors.New("malformed TBSCertificate subject public key")
	}
	entry = append(entry, algorithm...)
	h := sha256.Sum256(rawSubjectPublicKeyInfo(rawTBS))
	entry = append(entry, 0x04, 32)
	entry = append(entry, h[:]...)
	entry = append(entry, tbs...)
	return entry, nil
}

func rawSubjectPublicKeyInfo(rawTBS []byte) []byte {
	s := cryptobyte.String(rawTBS)
	var tbs, field cryptobyte.String
	s.ReadASN1(&tbs, cbasn1.SEQUENCE)
	tbs.SkipOptionalASN1(cbasn1.Tag(0).Constructed().ContextSpecific())
	tbs.SkipASN1(cbasn1.INTEGER)
	tbs.SkipASN1(cbasn1.SEQUENCE)
	tbs.ReadASN1Element(&field, cbasn1.SEQUENCE)
	tbs.ReadASN1Element(&field, cbasn1.SEQUENCE)
	tbs.ReadASN1Element(&field, cbasn1.SEQUENCE)
	tbs.ReadASN1Element(&field, cbasn1.SEQUENCE)
	return field
}

func EvaluateInclusionProof(index, start, end uint64, leaf [sha256.Size]byte, proof []byte) ([sha256.Size]byte, error) {
	if !validSubtree(start, end) || index < start || index >= end {
		return [sha256.Size]byte{}, errors.New("invalid MTC inclusion proof coordinates")
	}
	fn, sn := index-start, end-start-1
	result := leaf
	for len(proof) > 0 {
		if len(proof) < sha256.Size || sn == 0 {
			return [sha256.Size]byte{}, errors.New("invalid MTC inclusion proof length")
		}
		var sibling [sha256.Size]byte
		copy(sibling[:], proof[:sha256.Size])
		proof = proof[sha256.Size:]
		if fn&1 == 1 || fn == sn {
			result = hashNode(sibling, result)
			for fn&1 == 0 {
				fn >>= 1
				sn >>= 1
			}
		} else {
			result = hashNode(result, sibling)
		}
		fn >>= 1
		sn >>= 1
	}
	if sn != 0 {
		return [sha256.Size]byte{}, errors.New("MTC inclusion proof is too short")
	}
	return result, nil
}

func hashNode(left, right [sha256.Size]byte) [sha256.Size]byte {
	h := sha256.New()
	h.Write([]byte{1})
	h.Write(left[:])
	h.Write(right[:])
	var result [sha256.Size]byte
	h.Sum(result[:0])
	return result
}

func validSubtree(start, end uint64) bool {
	if start > end || end >= 1<<48 {
		return false
	}
	if start == end {
		return true
	}
	ceil := uint64(1) << (64 - bits.LeadingZeros64(end-start-1))
	return start&(ceil-1) == 0
}

func compareIDs(a, b []byte) int {
	if len(a) != len(b) {
		return len(a) - len(b)
	}
	return bytes.Compare(a, b)
}
