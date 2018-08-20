/*
 * ZLint Copyright 2018 Regents of the University of Michigan
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

package zlint

import (
	"encoding/pem"
	"testing"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/lints"
)

var (
	globalLintResult       *ResultSet
	globalSingleLintResult *lints.LintResult
)

const bigCertificatePem = `-----BEGIN CERTIFICATE-----
MIILajCCClKgAwIBAgIMOp/m5bdkZ2+oPevRMA0GCSqGSIb3DQEBCwUAMGIxCzAJ
BgNVBAYTAkJFMRkwFwYDVQQKExBHbG9iYWxTaWduIG52LXNhMTgwNgYDVQQDEy9H
bG9iYWxTaWduIEV4dGVuZGVkIFZhbGlkYXRpb24gQ0EgLSBTSEEyNTYgLSBHMzAe
Fw0xNzA2MjIwNjU2MDNaFw0xOTA2MjMwNjU2MDNaMIH9MR0wGwYDVQQPDBRQcml2
YXRlIE9yZ2FuaXphdGlvbjEPMA0GA1UEBRMGNTc4NjExMRMwEQYLKwYBBAGCNzwC
AQMTAlVTMR4wHAYLKwYBBAGCNzwCAQITDU5ldyBIYW1wc2hpcmUxCzAJBgNVBAYT
AlVTMRYwFAYDVQQIEw1OZXcgSGFtcHNoaXJlMRMwEQYDVQQHEwpQb3J0c21vdXRo
MSAwHgYDVQQJExdUd28gSW50ZXJuYXRpb25hbCBEcml2ZTEdMBsGA1UEChMUR01P
IEdsb2JhbFNpZ24sIEluYy4xGzAZBgNVBAMTEnd3dy5nbG9iYWxzaWduLmNvbTCC
ASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKaVk8nelrMqQmTSBju68D8B
MO7GGHtuQU8bfvuGNTUe6HiAxHYRB+LfCVAoTMXRtKgiI2YnTQ7xedKCaGTo2ZLH
y58Ya4ASpFgGLS3sPLjIHCP68ck126efksscXl2vBVWGGS7a0oTGLaaonFkz4FFy
0SkSwCL9UPPKkpVoQQ48kOF+tKZx1RimoZbZC9BwXtZYjdIbL9EzineymyJGsMi4
5utV7zQfcbZj3V9j4TAcx6UwqdwlmF7FVQ3Q1YmFtOZy6/U44us/Oz4SJ2+FIWS3
fZ6oGXBh3qq3L4n7ixiNpuj+CZmAJP8VM7w1dSquJ9ndw6Lid0jKIpY6nlDfflkC
AwEAAaOCB4Iwggd+MA4GA1UdDwEB/wQEAwIFoDCBlgYIKwYBBQUHAQEEgYkwgYYw
RwYIKwYBBQUHMAKGO2h0dHA6Ly9zZWN1cmUuZ2xvYmFsc2lnbi5jb20vY2FjZXJ0
L2dzZXh0ZW5kdmFsc2hhMmczcjMuY3J0MDsGCCsGAQUFBzABhi9odHRwOi8vb2Nz
cDIuZ2xvYmFsc2lnbi5jb20vZ3NleHRlbmR2YWxzaGEyZzNyMzBVBgNVHSAETjBM
MEEGCSsGAQQBoDIBATA0MDIGCCsGAQUFBwIBFiZodHRwczovL3d3dy5nbG9iYWxz
aWduLmNvbS9yZXBvc2l0b3J5LzAHBgVngQwBATAJBgNVHRMEAjAAMEUGA1UdHwQ+
MDwwOqA4oDaGNGh0dHA6Ly9jcmwuZ2xvYmFsc2lnbi5jb20vZ3MvZ3NleHRlbmR2
YWxzaGEyZzNyMy5jcmwwggPRBgNVHREEggPIMIIDxIISd3d3Lmdsb2JhbHNpZ24u
Y29tghVzeXN0ZW0uZ2xvYmFsc2lnbi5jb22CF3N5c3RlbWV1Lmdsb2JhbHNpZ24u
Y29tghdzeXN0ZW11cy5nbG9iYWxzaWduLmNvbYISZ2NjLmdsb2JhbHNpZ24uY29t
ghpjdGwxLnN5c3RlbS5nbG9iYWxzaWduLmNvbYIaY3RsMi5zeXN0ZW0uZ2xvYmFs
c2lnbi5jb22CEmhjcy5nbG9iYWxzaWduLmNvbYIXY3RsMS5oY3MuZ2xvYmFsc2ln
bi5jb22CF2N0bDIuaGNzLmdsb2JhbHNpZ24uY29tghVjbGllbnQuZ2xvYmFsc2ln
bi5jb22CFmVwa2lwcm8uZ2xvYmFsc2lnbi5jb22CG2N0bDEuZXBraXByby5nbG9i
YWxzaWduLmNvbYIYb3BlcmF0aW9uLmdsb2JhbHNpZ24uY29tghVyZWdpc3QuZ2xv
YmFsc2lnbi5jb22CE3NlYWwuZ2xvYmFsc2lnbi5jb22CFHNzaWYxLmdsb2JhbHNp
Z24uY29tghZwcm9maWxlLmdsb2JhbHNpZ24uY29tgiByZmMzMTYxLXRpbWVzdGFt
cC5nbG9iYWxzaWduLmNvbYIfcmZjMzE2MXRpbWVzdGFtcC5nbG9iYWxzaWduLmNv
bYIiY2VydGlmaWVkLXRpbWVzdGFtcC5nbG9iYWxzaWduLmNvbYIRY24uZ2xvYmFs
c2lnbi5jb22CEWhrLmdsb2JhbHNpZ24uY29tghF0aC5nbG9iYWxzaWduLmNvbYIT
YXBhYy5nbG9iYWxzaWduLmNvbYISZWRpLmdsb2JhbHNpZ24uY29tghRvY25ncy5n
bG9iYWxzaWduLmNvbYIRZXYuZ2xvYmFsc2lnbi5jb22CEWpwLmdsb2JhbHNpZ24u
Y29tghVlLXNpZ24uZ2xvYmFsc2lnbi5jb22CF3NzbGNoZWNrLmdsb2JhbHNpZ24u
Y29tghZjc3JoZWxwLmdsb2JhbHNpZ24uY29tghZzdGF0aWMxLmdsb2JhbHNpZ24u
Y29tghZzdGF0aWMyLmdsb2JhbHNpZ24uY29tghNibG9nLmdsb2JhbHNpZ24uY29t
ghNpbmZvLmdsb2JhbHNpZ24uY29tghVzZWN1cmUuZ2xvYmFsc2lnbi5jb22CFmFy
Y2hpdmUuZ2xvYmFsc2lnbi5jb22CFXN0YXR1cy5nbG9iYWxzaWduLmNvbYIWc3Vw
cG9ydC5nbG9iYWxzaWduLmNvbYIOZ2xvYmFsc2lnbi5jb20wHQYDVR0lBBYwFAYI
KwYBBQUHAwEGCCsGAQUFBwMCMB0GA1UdDgQWBBRUTciSxFJzJeFvq8WcPxoBQUKf
GzAfBgNVHSMEGDAWgBTds+dtqC7oxU5uz3TmdTyUFc7oHTCCAfQGCisGAQQB1nkC
BAIEggHkBIIB4AHeAHUA3esdK3oNT6Ygi4GtgWhwfi6OnQHVXIiNPRHEzbbsvswA
AAFczpYhfgAABAMARjBEAiAhJrXOLs31S6LkFx6xPmf3F2wckkQZK4cCygJXvOJ8
QwIgapfp6Kal4+/un4yLjQJee1swP+LTYIhXK0vBHARXhfoAdgBWFAaaL9fC7NP1
4b1Esj7HRna5vJkRXMDvlJhV1onQ3QAAAVzOliGfAAAEAwBHMEUCIQDCI99WIuKT
+kVmLBvMlxQi9fHtjUJuKTmRUEic2YYtdAIgT81iWIFUFTDZzH365JnoUMgkoUm0
W9ORqqTKYgb3/iwAdgCkuQmQtBhYFIe7E6LMZ3AKPDWYBPkb37jjd80OyA3cEAAA
AVzOliRsAAAEAwBHMEUCIQDiruypdLDxo/3TisqFXxxXxDbwR8VSjrfmQJ1aqvy0
OwIgaeeWftYP2eNNnwEkgJEhfCfbZZxthhUJ/Xxtqx+WleEAdQDuS723dc5guuFC
aR+r4Z5mow9+X7By2IMAxHuJeqj9ywAAAVzOlid6AAAEAwBGMEQCIFTgSc6vU/n3
Xf29uuatcVDxaiy37JX6XubsnowOU8PrAiBiTgjJ6LelCJq7xCv02fYdoMNOQqFy
a/zh9QwsFs7mmzANBgkqhkiG9w0BAQsFAAOCAQEAliaxkGO3qX15z6WN1RkwwTnH
ngJ5nDTrMscQ3rMGnfEYFW9uudfUVRnNLS49IR/V01nVML5Ex+Bz8CENw6ms7pHa
eVcCW12pFbLQxLns+dhExFvZBfy2iewouKo8Q41tolmEv4A3ADNuv+3r1bYhTnzE
55e0GMvnRIz5zQ7JWBTuamWNFI4OccJVh7vt0dnrSgiXs+XJ89qmgDyc/DikdM4q
psw2SW2R/SwSnkgvaLM/o0tw77aapxlaAs29Y4SE/RvRR2CJ0V/gvq9GUorY4OF2
2HEky394KiGDZDYUUDArx2+9w+yPikV5llF7lm2o84kZifnBO6SE9+4zdBExzg==
-----END CERTIFICATE-----
`

func BenchmarkZlint(b *testing.B) {
	certDerBlock, _ := pem.Decode([]byte(bigCertificatePem))
	x509Cert, err := x509.ParseCertificate(certDerBlock.Bytes)
	if err != nil {
		b.Fatalf("Error parsing certificate: %s", err.Error())
	}

	b.ResetTimer()
	b.Run("All lints", func(b *testing.B) {
		var lintResult *ResultSet
		for i := 0; i < b.N; i++ {
			lintResult = LintCertificate(x509Cert)
		}

		globalLintResult = lintResult
	})

	b.Run("Fast lints", func(b *testing.B) {
		globalLintResult = &ResultSet{}
		globalLintResult.Results = make(map[string]*lints.LintResult, len(lints.Lints))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for key, value := range lints.Lints {
				switch key {
				case "w_dnsname_underscore_in_trd", "e_dnsname_underscore_in_sld", "e_dnsname_hyphen_in_sld",
					"w_dnsname_wildcard_left_of_public_suffix", "w_san_iana_pub_suffix_empty":
					continue
				}

				if !value.Lint.CheckApplies(x509Cert) {
					continue
				}
				globalLintResult.Results[key] = value.Lint.Execute(x509Cert)
			}
		}
	})

	b.Run("Fastest lints", func(b *testing.B) {
		globalLintResult = &ResultSet{}
		globalLintResult.Results = make(map[string]*lints.LintResult, len(lints.Lints))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for key, value := range lints.Lints {
				switch key {
				case "w_dnsname_underscore_in_trd", "e_dnsname_underscore_in_sld", "e_dnsname_hyphen_in_sld",
					"w_dnsname_wildcard_left_of_public_suffix", "w_san_iana_pub_suffix_empty",
					"w_rsa_mod_factors_smaller_than_752", "e_dnsname_bad_character_in_label", "w_subject_dn_leading_whitespace",
					"w_subject_dn_trailing_whitespace", "w_multiple_subject_rdn", "e_ext_san_dns_not_ia5_string",
					"e_ext_san_empty_name", "e_dnsname_not_valid_tld", "e_dnsname_contains_bare_iana_suffix",
					"e_dnsname_wildcard_only_in_left_label", "e_international_dns_name_not_nfc", "e_dnsname_left_label_wildcard_correct",
					"e_international_dns_name_not_unicode", "w_issuer_dn_trailing_whitespace", "w_issuer_dn_leading_whitespace",
					"w_multiple_issuer_rdn", "e_dnsname_empty_label", "e_dnsname_label_too_long", "e_distribution_point_incomplete",
					"e_wrong_time_format_pre2050", "e_utc_time_does_not_include_seconds", "e_sub_cert_not_is_ca", "w_rsa_mod_not_odd",
					"e_path_len_constraint_zero_or_less", "e_san_dns_name_includes_null_char":
					continue
				}

				if !value.Lint.CheckApplies(x509Cert) {
					continue
				}
				globalLintResult.Results[key] = value.Lint.Execute(x509Cert)
			}
		}
	})

	for key, value := range lints.Lints {
		b.Run(key, func(b *testing.B) {
			if !value.Lint.CheckApplies(x509Cert) {
				b.Skip("Check doesn't apply")
			}

			var result *lints.LintResult
			for i := 0; i < b.N; i++ {
				result = value.Lint.Execute(x509Cert)
			}

			globalSingleLintResult = result
		})
	}
}
