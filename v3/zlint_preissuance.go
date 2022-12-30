/*
 * ZLint Copyright 2022 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

// Used to check parsed info from certificate for compliance

package zlint

import (
	"crypto"
	"crypto/dsa" //lint:ignore SA1019 Support DSA even though it's deprecated upstream
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509/pkix"

	"encoding/asn1"
	"errors"
	"fmt"
	"math/big"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

var dummyRSA2048PrivateKey = &rsa.PrivateKey{
	PublicKey: rsa.PublicKey{
		N: fromHexString("BDDEF6A3335C16F06DDD9216E958DC6723754C29DDF44D65C49C0FCC06EBA93F65F93FE7D4CFBF23543C085C506E40C2C486CEE8848AF306AF66C55791182AA398C55CDB0F95EBCC1261B6B2BB0E427B434845817B5347476895A91ABFC257BB6829F0C3D08E93A3F19097C9262EFA1F18F0F6007B3ED25656BACDFDCCAB634B278AF337E5CEB486E17AD0FA7EFD84069BFF6CD7A9D492C470DDABD6FE5775764034519A503CB3453D850AD565D09299415D58A2AFA68768A9854909242E332731A387208AFEB6D5DDD6ADC30F18ED4B8B80AD34152E805B7741200138FFA817DB35C866FF66DC6F1D882CD1963612D76514C5A45D3B4AD9EFCBAA500AC412D1"),
		E: 65537,
	},
	D: fromHexString("082836FFAA0573DA3CF216E46364EB78531280D81C9E8BA07DCAD19CB86064C8153D1904E6C48047BB208CED339131D25C47C95F8A69CBFC29CA3E8F12458FDE314F11BBEE2688C491E9CBDD832878B935403AA89CC34C9A7E6A33FCFFC11D52E245C8FFC16B0DE007DDCEC05E62BB944389B2357DFC63FDD97BB98A1E5BD601981AF8A22F1E0F73B779C0FBFD499B559EAD72C6A0C224F213F316207AC37EEC0B6A49DF07223F5B03552EF2E8435DEAC71AB961FB3848C75E2C9A964AC559E9DDED934A5601BA95B9EA1C8752FE37AF41452B59C0C2F5FEB0A6FFB6E2AC37367D5560D95D577C27CF0877544BB768BCD9A5EE21F6F4099B97D4F9EC33D1CE35"),
	Primes: []*big.Int{
		fromHexString("C179B2FFB709C574C682522A54063B964D5729D206DAA1268E5756B20CE1D8FF5659623644D5732F7C310D1FEF0B924A813B39FE38E06A485B94CD0CD117D77507883193A7AEEF096BAD3AD72DE4DB2987B7B5DAE83B18FF387A74C9DBA7910CC81297646385B37F8D4927105F7A9C0D214BFE0A67740F3863E983F369F03625"),
		fromHexString("FB3B1150E79CEAA7676ADA8870D4FE61059B8A9142A7856CB90FF044846D11F99E0B6D096257994B325534E0AA4B33A3980CA75F750DD34A9D9053B747EE7E7FF7606BBAE9EA7836742FA30B460E7FA97698BBDF8845F19195F7E00A2E69ED0DE3A0D4CF4ADDDAE6B30EFED9173A0FA792FC2F62FAD68424F6BC2E02AC74BC3D"),
	},
}

var dummyDSA2048PrivateKey = &dsa.PrivateKey{
	PublicKey: dsa.PublicKey{
		Parameters: dsa.Parameters{
			P: fromHexString("B16B4090C8057D0BB285B8B641D931149A890530DA44E5374EBAB645EAD7B6F6B83D57D0B6DF27AE9C619E84B1F1444FED1DD30967A896CFC61D26756AB9E53FED5DC8C31A0494382626B3519BE7A55614143F655C4FE6C6EB3D96F11D8A26320A341A9BFA85AEE2B8B1CE709D55B9F74F1732135030156A0B0FF1AEB6412A592DA6AB1F000525A81C5E979E40532432932B593E8438212B9C04B6CF9D09BE45FD82E24EEF37A67318B7E59E8FE28A0F03F1ADA28BAA8689A85DE0A65215C7188CA655DD005814D2D6DE92D6FBF639A698C077E1954302F26D8D511A377EDC3B19353A85395D30D97830B126E9DDF0A0E85056EF6833C3CC7BA29B9DFFCFA595"),
			Q: fromHexString("A67D6F933F2CF69667C49617765672FCD0AAB1308D7D0A193D721409"),
			G: fromHexString("0094CA379347D760681A0A0F91B8AEA9FB197F354B8BF3CD301B2D1CD3B134C598160E6BFDD23A86F9F3074621184B55E5A60766A05081BEC4B4A5EF7CD98A88FF17F17699521ADDDDCED6580EF2E5BEE884FFA131B839C09E923E48342954794C66426D777F9792FBFC7EC7F3C1E9291C210B226D51BD25DF3C0DE00CE47A77DAC719F3850E01DCBC7EAE1682D15C9FD909208721C21353F987B37C152842E4875B8B5643639ABF128289538DA4630EA2B31D915A553BF04D12135419B79E146B7B48BFAA13B15704DE20951EFE44B1E5BB6226AF9A6064BE04B8187AEAFB721EB78F5FF3866DB6550311F5881AD704571AD3691E110E038A1B48DCE7B269EE5D"),
		},
		Y: fromHexString("62416A1C98C55E000AAEA4FCCF6232538894AF744935F74AABB50A24C6E167E27BF36F0F344AB7C87AD985059B7522059DC80BF35AD1FE26E5E3C584E11BBEF75B2E551774B956EB1BFB3F6750CDA859CAECC26C672AB729275810ED6D1BD31E53F3B33BE9DE23D6DBD9B64160D96B1DE4DD8A3F9266E9CB40A1C09E913AF9F335B42FA405B8CFD594E2E866BB3261D29E55D5EFC32C131183449E87CB72C925F38973805E8E421BCFE5105F267E025600633D1318CBCF9E40A0D2D14A62890E215EE510D25D8CB33AEADF02A9E7561BB222CCCF36A5E1E0596483A423D7E922E7A0AFF9108DE95EED99EDE0EB2E0089F5227098A035C7E1723AD9095271B082"),
	},
	X: fromHexString("1EF7CD6E6775B23D1E83DFBF2D907409D2669C66C1445E1DA527F591"),
}

var dummyP256PrivateKey = &ecdsa.PrivateKey{
	PublicKey: ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     fromHexString("929C15ED0F20840245F15EE932B39B99345A4FF084770E259E24BDE5C234D76C"),
		Y:     fromHexString("80977CF7298B3F1DA2847D29FCA2C8EC5583DA51440D4FCBC9F5CF27E899F26A"),
	},
	D: fromHexString("8BB629C966D4316B1933066F89C6AB4CBA5D225FDA812967CBA0709217FB93FE"),
}

var dummyP384PrivateKey = &ecdsa.PrivateKey{
	PublicKey: ecdsa.PublicKey{
		Curve: elliptic.P384(),
		X:     fromHexString("9BF32C9966B517C85420D872D1FCFE50ED4632D691C2DB52ECD2BAE8B4FC4696C60187401371D4ECB848C8DD7BFA8183"),
		Y:     fromHexString("981DAE8095278A2B1BC0B56F533A94BFDC2E9AE8EFCFB864E54F506D46EB0C49EBA08E20E5118633A7F261472DE41659"),
	},
	D: fromHexString("B3DB20536D5D6DB6E5931ED878EEFF45450842BFE30DCACBD34373896A92DD5A8CE73F53559AAC44A8C59B7A750369BE"),
}

var dummyP521PrivateKey = &ecdsa.PrivateKey{
	PublicKey: ecdsa.PublicKey{
		Curve: elliptic.P521(),
		X:     fromHexString("019006447EB477FE4D5CBD39325A08A300B929D1461281A590C606645079783285DB0CD25430C9B306BCA2C7AAB7E3B3FA0D27B9A81F2C305E680FCEDC7DC11FD0F5"),
		Y:     fromHexString("C6123D2E0D94AB3256B5B333165666D33F4EDAAC4A58A4259B1DED7FA9050D2F8E1A8684449E4B50D5816D0A69F3A8D6810662547B31FF50EB9E27B259C6152B70"),
	},
	D: fromHexString("83409F627B7B6D4039C564B14DE961F9B3498754E0C308C3DD49C60A03D6010FA0A89DB9FE35CBFF7AE22DAC6C479FC0E7C85A4B64126E4AA046367518115D61F8"),
}

var dummyEd25519PrivateKey = ed25519.PrivateKey{
	0x50, 0xB3, 0xE2, 0xF6, 0x72, 0xA5, 0x9C, 0x18, 0xAB, 0x1C, 0x20, 0x5C, 0x9D, 0x00, 0x5A, 0x0E,
	0x8F, 0x29, 0x45, 0xB5, 0xB7, 0x02, 0x14, 0xF6, 0xD3, 0xCD, 0x30, 0x96, 0x41, 0xB5, 0xB9, 0x69,
	0xDF, 0x7D, 0x26, 0xAB, 0x2A, 0xE1, 0x67, 0x82, 0x3D, 0x56, 0xE2, 0x98, 0xFA, 0x66, 0x6B, 0x10,
	0x73, 0x7E, 0xCE, 0xD4, 0xDD, 0x50, 0x06, 0xB8, 0xB1, 0xC3, 0x19, 0x08, 0xC2, 0xE8, 0x89, 0xD6,
}

func init() {
	dummyRSA2048PrivateKey.Precompute()
}

func fromHexString(base16 string) *big.Int {
	i, ok := new(big.Int).SetString(base16, 16)
	if !ok {
		panic("bad number: " + base16)
	}
	return i
}

type tbsCertificatePartial struct {
	Version            int `asn1:"optional,explicit,default:0,tag:0"`
	SerialNumber       *big.Int
	SignatureAlgorithm pkix.AlgorithmIdentifier
}

type signed struct {
	ToBeSigned         asn1.RawValue
	SignatureAlgorithm pkix.AlgorithmIdentifier
	SignatureValue     asn1.BitString
}

type pssParameters struct {
	Hash         pkix.AlgorithmIdentifier `asn1:"optional,explicit,tag:0"`
	MGF          pkix.AlgorithmIdentifier `asn1:"optional,explicit,tag:1"`
	SaltLength   int                      `asn1:"optional,explicit,tag:2,default:20"`
	TrailerField int                      `asn1:"optional,explicit,tag:3,default:1"`
}

type dsaSignature struct {
	R, S *big.Int
}

type SigAlgoFunction int

const (
	UnknownSigAlgoFunction SigAlgoFunction = iota
	rsa_SignPKCS1v15
	rsa_SignPSS
	dsa_Sign
	ecdsa_SignASN1_P256
	ecdsa_SignASN1_P384
	ecdsa_SignASN1_P521
	ed25519_Sign
)

var sigAlgoDetails = []struct {
	sigAlgoOID  asn1.ObjectIdentifier
	sigAlgoFunc SigAlgoFunction
	hash        crypto.Hash
}{
	{asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 4}, rsa_SignPKCS1v15, crypto.MD5},
	{asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 5}, rsa_SignPKCS1v15, crypto.SHA1},
	{asn1.ObjectIdentifier{1, 3, 14, 3, 2, 29}, rsa_SignPKCS1v15, crypto.SHA1},
	{asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 11}, rsa_SignPKCS1v15, crypto.SHA256},
	{asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 12}, rsa_SignPKCS1v15, crypto.SHA384},
	{asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 13}, rsa_SignPKCS1v15, crypto.SHA512},
	{asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 10}, rsa_SignPSS, crypto.SHA1},
	{asn1.ObjectIdentifier{1, 2, 840, 10040, 4, 3}, dsa_Sign, crypto.SHA1},
	{asn1.ObjectIdentifier{2, 16, 840, 1, 101, 3, 4, 3, 2}, dsa_Sign, crypto.SHA256},
	{asn1.ObjectIdentifier{1, 2, 840, 10045, 4, 1}, ecdsa_SignASN1_P256, crypto.SHA1},
	{asn1.ObjectIdentifier{1, 2, 840, 10045, 4, 3, 2}, ecdsa_SignASN1_P256, crypto.SHA256},
	{asn1.ObjectIdentifier{1, 2, 840, 10045, 4, 3, 3}, ecdsa_SignASN1_P384, crypto.SHA384},
	{asn1.ObjectIdentifier{1, 2, 840, 10045, 4, 3, 4}, ecdsa_SignASN1_P521, crypto.SHA512},
	{asn1.ObjectIdentifier{1, 3, 101, 112}, ed25519_Sign, crypto.Hash(0)},
}

func getSignatureAlgorithmComponentsFromAI(ai pkix.AlgorithmIdentifier) (SigAlgoFunction, crypto.Hash, int) {
	pssSaltLength := 0
	for _, details := range sigAlgoDetails {
		if ai.Algorithm.Equal(details.sigAlgoOID) {
			if details.sigAlgoFunc == rsa_SignPSS {
				pssSaltLength = 20
				var params pssParameters
				if _, err := asn1.Unmarshal(ai.Parameters.Bytes, &params); err == nil {
					if params.Hash.Algorithm.Equal(asn1.ObjectIdentifier{2, 16, 840, 1, 101, 3, 4, 2, 1}) {
						details.hash = crypto.SHA256
					} else if params.Hash.Algorithm.Equal(asn1.ObjectIdentifier{2, 16, 840, 1, 101, 3, 4, 2, 2}) {
						details.hash = crypto.SHA384
					} else if params.Hash.Algorithm.Equal(asn1.ObjectIdentifier{2, 16, 840, 1, 101, 3, 4, 2, 3}) {
						details.hash = crypto.SHA512
					}
					pssSaltLength = params.SaltLength
				}
			}

			return details.sigAlgoFunc, details.hash, pssSaltLength
		}
	}

	return UnknownSigAlgoFunction, crypto.Hash(0), pssSaltLength
}

// LintTBSCertificate runs all registered lints on rawTBSCertificate using default options,
// producing a ResultSet.
//
// Using LintTBSCertificate(rawTBSCertificate) is equivalent to calling LintCertificateEx(rawTBSCertificate, nil).
func LintTBSCertificate(rawTBSCertificate []byte) *ResultSet {
	// Run all lints from the global registry
	return LintTBSCertificateEx(rawTBSCertificate, nil)
}

// LintTBSCertificateEx runs lints from the provided registry on rawTBSCertificate producing
// a ResultSet. Providing an explicit registry allows the caller to filter the
// lints that will be run. (See lint.Registry.Filter())
//
// If registry is nil then the global registry of all lints is used and this
// function is equivalent to calling LintTBSCertificate(rawTBSCertificate).
func LintTBSCertificateEx(rawTBSCertificate []byte, registry lint.Registry) *ResultSet {
	if rawTBSCertificate == nil {
		return nil
	}

	var err error
	var tbs tbsCertificatePartial
	var signature []byte
	// Decode enough of the TBSCertificate to discover the signature algorithm
	if _, err = asn1.Unmarshal(rawTBSCertificate, &tbs); err == nil {
		var pubKeyAlgo SigAlgoFunction
		var hash crypto.Hash
		var pssSaltLength int
		pubKeyAlgo, hash, pssSaltLength = getSignatureAlgorithmComponentsFromAI(tbs.SignatureAlgorithm)
		// Calculate the required input for the signature function
		var hashed []byte
		switch hash {
		case crypto.MD5:
			temp := md5.Sum(rawTBSCertificate)
			hashed = temp[:]
		case crypto.SHA1:
			temp := sha1.Sum(rawTBSCertificate)
			hashed = temp[:]
		case crypto.SHA256:
			temp := sha256.Sum256(rawTBSCertificate)
			hashed = temp[:]
		case crypto.SHA384:
			temp := sha512.Sum384(rawTBSCertificate)
			hashed = temp[:]
		case crypto.SHA512:
			temp := sha512.Sum512(rawTBSCertificate)
			hashed = temp[:]
		}

		// Generate a dummy signature
		switch pubKeyAlgo {
		case UnknownSigAlgoFunction:
			err = errors.New("unknown signature algorithm function")
		case rsa_SignPKCS1v15:
			signature, err = rsa.SignPKCS1v15(rand.Reader, dummyRSA2048PrivateKey, hash, hashed)
		case rsa_SignPSS:
			opts := rsa.PSSOptions{SaltLength: pssSaltLength}
			signature, err = rsa.SignPSS(rand.Reader, dummyRSA2048PrivateKey, hash, hashed, &opts)
		case dsa_Sign:
			var dsaSig dsaSignature
			if dsaSig.R, dsaSig.S, err = dsa.Sign(rand.Reader, dummyDSA2048PrivateKey, hashed); err == nil {
				signature, err = asn1.Marshal(dsaSig)
			}
		case ecdsa_SignASN1_P256:
			signature, err = ecdsa.SignASN1(rand.Reader, dummyP256PrivateKey, hashed)
		case ecdsa_SignASN1_P384:
			signature, err = ecdsa.SignASN1(rand.Reader, dummyP384PrivateKey, hashed)
		case ecdsa_SignASN1_P521:
			signature, err = ecdsa.SignASN1(rand.Reader, dummyP521PrivateKey, hashed)
		case ed25519_Sign:
			signature = ed25519.Sign(dummyEd25519PrivateKey, rawTBSCertificate)
		}

		if err == nil {
			// Package the TBSCertificate in a dummy, yet syntactically valid, X.509 certificate that LintCertificateEx will be able to parse
			dummy := signed{
				ToBeSigned:         asn1.RawValue{FullBytes: rawTBSCertificate},
				SignatureAlgorithm: tbs.SignatureAlgorithm,
				SignatureValue:     asn1.BitString{Bytes: signature},
			}

			// DER encode the dummy certificate, then decode it again into an object that can be passed to LintCertificateEx
			var certDER []byte
			if certDER, err = asn1.Marshal(dummy); err == nil {
				var c *x509.Certificate
				if c, err = x509.ParseCertificate(certDER); err == nil {
					return LintCertificateEx(c, registry)
				}
			}
		}
	}

	// Return the TBSCertificate parsing error
	res := new(ResultSet)
	res.Results = make(map[string]*lint.LintResult, 1)
	res.Results["e_lint_tbs_certificate_ex"] = &lint.LintResult{Status: lint.Fatal, Details: fmt.Sprintf("%v", err)}
	return res
}
