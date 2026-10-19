package mtc

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	mtcparser "github.com/zmap/zlint/v3/mtc"
	"github.com/zmap/zlint/v3/util"
)

func init() {
	lint.RegisterProfile(lint.Profile{
		Name:        "mtc_plants_07",
		Description: "Merkle Tree Certificates draft-plants-07 structural checks",
		Citation:    "draft-ietf-plants-merkle-tree-certs-07",
		Source:      lint.MerkleTreeCertificates,
		LintNames: []string{
			"e_mtc_ca_structure_valid",
			"e_mtc_proof_structure_valid",
			"e_mtc_proof_self_contained_valid",
		},
	})
	lint.RegisterCertificateLint(&lint.CertificateLint{
		LintMetadata: lint.LintMetadata{
			Name:          "e_mtc_ca_structure_valid",
			Description:   "MTC CA certificates use the draft-plants-07 CA extension structure",
			Citation:      "draft-ietf-plants-merkle-tree-certs-07, Section 5.5",
			Source:        lint.MerkleTreeCertificates,
			EffectiveDate: util.ZeroDate,
		},
		Lint: NewCAStructure,
	})
	lint.RegisterCertificateLint(&lint.CertificateLint{
		LintMetadata: lint.LintMetadata{
			Name:          "e_mtc_proof_structure_valid",
			Description:   "MTC proof certificates use the draft-plants-07 structure",
			Citation:      "draft-ietf-plants-merkle-tree-certs-07, Sections 4 and 5",
			Source:        lint.MerkleTreeCertificates,
			EffectiveDate: util.ZeroDate,
		},
		Lint: NewProofStructure,
	})
	lint.RegisterCertificateLint(&lint.CertificateLint{
		LintMetadata: lint.LintMetadata{
			Name:          "e_mtc_proof_self_contained_valid",
			Description:   "MTC proof certificates contain a valid self-contained log entry and inclusion proof",
			Citation:      "draft-ietf-plants-merkle-tree-certs-07, Section 5",
			Source:        lint.MerkleTreeCertificates,
			EffectiveDate: util.ZeroDate,
		},
		Lint: NewProofSelfContained,
	})
}

type caStructure struct{}

type proofStructure struct{}

type proofSelfContained struct{}

func NewCAStructure() lint.LintInterface {
	return &caStructure{}
}

func NewProofStructure() lint.LintInterface {
	return &proofStructure{}
}

func NewProofSelfContained() lint.LintInterface {
	return &proofSelfContained{}
}

func (l *caStructure) CheckApplies(c *x509.Certificate) bool {
	return mtcparser.IsCACertificate(c)
}

func (l *proofStructure) CheckApplies(c *x509.Certificate) bool {
	return mtcparser.IsProofCertificate(c)
}

func (l *proofSelfContained) CheckApplies(c *x509.Certificate) bool {
	return mtcparser.IsProofCertificate(c)
}

func (l *caStructure) Execute(c *x509.Certificate) *lint.LintResult {
	if err := mtcparser.ValidateCA(c); err != nil {
		return &lint.LintResult{Status: lint.Error, Details: err.Error()}
	}
	return &lint.LintResult{Status: lint.Pass}
}

func (l *proofStructure) Execute(c *x509.Certificate) *lint.LintResult {
	if _, err := mtcparser.ParseCertificate(c); err != nil {
		return &lint.LintResult{Status: lint.Error, Details: err.Error()}
	}
	return &lint.LintResult{Status: lint.Pass}
}

func (l *proofSelfContained) Execute(c *x509.Certificate) *lint.LintResult {
	parsed, err := mtcparser.ParseCertificate(c)
	if err != nil {
		return &lint.LintResult{Status: lint.Error, Details: err.Error()}
	}
	if err := mtcparser.ValidateSelfContained(c, parsed); err != nil {
		return &lint.LintResult{Status: lint.Error, Details: err.Error()}
	}
	return &lint.LintResult{Status: lint.Pass}
}
