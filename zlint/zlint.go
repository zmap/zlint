/* zlint.go
 * Used to check parsed info from certificate for compliance
 */

package zlint

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/lints"
	"time"
)

type PrettyOutput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Providence  string `json:"providence"`
}

//Pretty Print lint outputs
func PrettyPrintZLint() {
	for _, l := range lints.Lints {
		p := PrettyOutput{}
		p.Name = l.Name
		p.Description = l.Description
		p.Providence = l.Providence

		buffer := new(bytes.Buffer)
		enc := json.NewEncoder(buffer)
		enc.SetEscapeHTML(false)
		enc.Encode(p)
		fmt.Print(buffer.String())
	}
}

//Calls all other checks on parsed certs producing version
func ZLintResultTestHandler(cert *x509.Certificate) *lints.ZLintResult {
	if cert == nil {
		return nil
	}
	//run all tests
	ZLintResult := lints.ZLintResult{}
	ZLintReport := lints.LintReport{}
	InitLintReport(&ZLintReport)
	ZLintResult.ZLint = &ZLintReport
	ZLintResult.Execute(cert)
	ZLintResult.ZLintVersion = lints.ZLintVersion
	ZLintResult.Timestamp = time.Now().Unix()
	return &ZLintResult
}

func InitLintReport(report *lints.LintReport) {
	report.EBasicConstraintsNotCritical = lints.ResultStruct{Result: lints.NA}
	report.EIanBareWildcard = lints.ResultStruct{Result: lints.NA}
	report.EIanWildcardNotFirst = lints.ResultStruct{Result: lints.NA}
	report.ESanBareWildcard = lints.ResultStruct{Result: lints.NA}
	report.ESanWildcardNotFirst = lints.ResultStruct{Result: lints.NA}
	report.ECaCountryNameInvalid = lints.ResultStruct{Result: lints.NA}
	report.ECaCountryNameMissing = lints.ResultStruct{Result: lints.NA}
	report.ECaCrlSignNotSet = lints.ResultStruct{Result: lints.NA}
	report.NCaDigitalSignatureNotSet = lints.ResultStruct{Result: lints.NA}
	report.ECaKeyCertSignNotSet = lints.ResultStruct{Result: lints.NA}
	report.ECaKeyUsageMissing = lints.ResultStruct{Result: lints.NA}
	report.ECaKeyUsageNotCritical = lints.ResultStruct{Result: lints.NA}
	report.ECaOrganizationNameMissing = lints.ResultStruct{Result: lints.NA}
	report.ECaSubjectFieldEmpty = lints.ResultStruct{Result: lints.NA}
	report.ECertContainsUniqueIdentifier = lints.ResultStruct{Result: lints.NA}
	report.ECertExtensionsVersionNot_3 = lints.ResultStruct{Result: lints.NA}
	report.ECabDvConflictsWithLocality = lints.ResultStruct{Result: lints.NA}
	report.ECabDvConflictsWithOrg = lints.ResultStruct{Result: lints.NA}
	report.ECabDvConflictsWithPostal = lints.ResultStruct{Result: lints.NA}
	report.ECabDvConflictsWithProvince = lints.ResultStruct{Result: lints.NA}
	report.ECabDvConflictsWithStreet = lints.ResultStruct{Result: lints.NA}
	report.ECertPolicyIvRequiresCountry = lints.ResultStruct{Result: lints.NA}
	report.ECertPolicyIvRequiresProvinceOrLocality = lints.ResultStruct{Result: lints.NA}
	report.ECertPolicyOvRequiresCountry = lints.ResultStruct{Result: lints.NA}
	report.ECertPolicyOvRequiresProvinceOrLocality = lints.ResultStruct{Result: lints.NA}
	report.ECabOvRequiresOrg = lints.ResultStruct{Result: lints.NA}
	report.ECabIvRequiresPersonalName = lints.ResultStruct{Result: lints.NA}
	report.ECertUniqueIdentifierVersionNot_2Or_3 = lints.ResultStruct{Result: lints.NA}
	report.EDhParamsMissing = lints.ResultStruct{Result: lints.NA}
	report.EDistributionPointIncomplete = lints.ResultStruct{Result: lints.NA}
	report.WDistributionPointMissingLdapOrUri = lints.ResultStruct{Result: lints.NA}
	report.EDsaImproperModulusOrDivisorSize = lints.ResultStruct{Result: lints.NA}
	report.EDsaShorterThan_2048Bits = lints.ResultStruct{Result: lints.NA}
	report.EEcImproperCurves = lints.ResultStruct{Result: lints.NA}
	report.WEkuCriticalImproperly = lints.ResultStruct{Result: lints.NA}
	report.EEvBusinessCategoryMissing = lints.ResultStruct{Result: lints.NA}
	report.EEvCountryNameMissing = lints.ResultStruct{Result: lints.NA}
	report.EEvLocalityNameMissing = lints.ResultStruct{Result: lints.NA}
	report.EEvOrganizationNameMissing = lints.ResultStruct{Result: lints.NA}
	report.EEvSerialNumberMissing = lints.ResultStruct{Result: lints.NA}
	report.EEvValidTimeTooLong = lints.ResultStruct{Result: lints.NA}
	report.WExtAiaAccessLocationMissing = lints.ResultStruct{Result: lints.NA}
	report.EExtAiaMarkedCritical = lints.ResultStruct{Result: lints.NA}
	report.EExtAuthorityKeyIdentifierCritical = lints.ResultStruct{Result: lints.NA}
	report.EExtAuthorityKeyIdentifierMissing = lints.ResultStruct{Result: lints.NA}
	report.EExtAuthorityKeyIdentifierNoKeyIdentifier = lints.ResultStruct{Result: lints.NA}
	report.WExtCertPolicyContainsNoticeref = lints.ResultStruct{Result: lints.NA}
	report.EExtCertPolicyDisallowedAnyPolicyQualifier = lints.ResultStruct{Result: lints.NA}
	report.EExtCertPolicyDuplicate = lints.ResultStruct{Result: lints.NA}
	report.EExtCertPolicyExplicitTextIa5String = lints.ResultStruct{Result: lints.NA}
	report.WExtCertPolicyExplicitTextIncludesControl = lints.ResultStruct{Result: lints.NA}
	report.WExtCertPolicyExplicitTextNotNfc = lints.ResultStruct{Result: lints.NA}
	report.WExtCertPolicyExplicitTextNotUtf8 = lints.ResultStruct{Result: lints.NA}
	report.EExtCertPolicyExplicitTextTooLong = lints.ResultStruct{Result: lints.NA}
	report.WExtCrlDistributionMarkedCritical = lints.ResultStruct{Result: lints.NA}
	report.EExtDuplicateExtension = lints.ResultStruct{Result: lints.NA}
	report.EExtFreshestCrlMarkedCritical = lints.ResultStruct{Result: lints.NA}
	report.WExtIanCritical = lints.ResultStruct{Result: lints.NA}
	report.EExtIanDnsNotIa5String = lints.ResultStruct{Result: lints.NA}
	report.EExtIanEmptyName = lints.ResultStruct{Result: lints.NA}
	report.EExtIanNoEntries = lints.ResultStruct{Result: lints.NA}
	report.EExtIanRfc822FormatInvalid = lints.ResultStruct{Result: lints.NA}
	report.EExtIanSpaceDnsName = lints.ResultStruct{Result: lints.NA}
	report.EExtIanUriFormatInvalid = lints.ResultStruct{Result: lints.NA}
	report.EExtIanUriHostNotFqdnOrIp = lints.ResultStruct{Result: lints.NA}
	report.EExtIanUriNotIa5 = lints.ResultStruct{Result: lints.NA}
	report.EExtIanUriRelative = lints.ResultStruct{Result: lints.NA}
	report.EExtKeyUsageCertSignWithoutCa = lints.ResultStruct{Result: lints.NA}
	report.WExtKeyUsageNotCritical = lints.ResultStruct{Result: lints.NA}
	report.EExtKeyUsageWithoutBits = lints.ResultStruct{Result: lints.NA}
	report.EExtNameConstraintsNotCritical = lints.ResultStruct{Result: lints.NA}
	report.EExtNameConstraintsNotInCa = lints.ResultStruct{Result: lints.NA}
	report.EExtPolicyConstraintsEmpty = lints.ResultStruct{Result: lints.NA}
	report.EExtPolicyConstraintsNotCritical = lints.ResultStruct{Result: lints.NA}
	report.EExtPolicyMapAnyPolicy = lints.ResultStruct{Result: lints.NA}
	report.WExtPolicyMapNotCritical = lints.ResultStruct{Result: lints.NA}
	report.WExtPolicyMapNotInCertPolicy = lints.ResultStruct{Result: lints.NA}
	report.EExtSanContainsReservedIp = lints.ResultStruct{Result: lints.NA}
	report.WExtSanCriticalWithSubjectDn = lints.ResultStruct{Result: lints.NA}
	report.EExtSanDirectoryNamePresent = lints.ResultStruct{Result: lints.NA}
	report.EExtSanDnsNotIa5String = lints.ResultStruct{Result: lints.NA}
	report.EExtSanDnsnameNotFqdn = lints.ResultStruct{Result: lints.NA}
	report.EExtSanEdiPartyNamePresent = lints.ResultStruct{Result: lints.NA}
	report.EExtSanEmptyName = lints.ResultStruct{Result: lints.NA}
	report.EExtSanMissing = lints.ResultStruct{Result: lints.NA}
	report.EExtSanNoEntries = lints.ResultStruct{Result: lints.NA}
	report.EExtSanNotCriticalWithoutSubject = lints.ResultStruct{Result: lints.NA}
	report.EExtSanOtherNamePresent = lints.ResultStruct{Result: lints.NA}
	report.EExtSanRegisteredIdPresent = lints.ResultStruct{Result: lints.NA}
	report.EExtSanRfc822FormatInvalid = lints.ResultStruct{Result: lints.NA}
	report.EExtSanRfc822NamePresent = lints.ResultStruct{Result: lints.NA}
	report.EExtSanSpaceDnsName = lints.ResultStruct{Result: lints.NA}
	report.EExtSanUniformResourceIdentifierPresent = lints.ResultStruct{Result: lints.NA}
	report.EExtSanUriFormatInvalid = lints.ResultStruct{Result: lints.NA}
	report.EExtSanUriHostNotFqdnOrIp = lints.ResultStruct{Result: lints.NA}
	report.EExtSanUriNotIa5 = lints.ResultStruct{Result: lints.NA}
	report.EExtSanUriRelative = lints.ResultStruct{Result: lints.NA}
	report.EExtSubjectDirectoryAttrCritical = lints.ResultStruct{Result: lints.NA}
	report.EExtSubjectKeyIdentifierCritical = lints.ResultStruct{Result: lints.NA}
	report.EExtSubjectKeyIdentifierMissingCa = lints.ResultStruct{Result: lints.NA}
	report.WExtSubjectKeyIdentifierMissingSubCert = lints.ResultStruct{Result: lints.NA}
	report.EGeneralizedTimeDoesNotIncludeSeconds = lints.ResultStruct{Result: lints.NA}
	report.EGeneralizedTimeIncludesFractionSeconds = lints.ResultStruct{Result: lints.NA}
	report.EGeneralizedTimeNotInZulu = lints.ResultStruct{Result: lints.NA}
	report.WGtldUnderConsideration = lints.ResultStruct{Result: lints.NA}
	report.EIanDnsNameIncludesNullChar = lints.ResultStruct{Result: lints.NA}
	report.EIanDnsNameStartsWithPeriod = lints.ResultStruct{Result: lints.NA}
	report.WIanIanaPubSuffixEmpty = lints.ResultStruct{Result: lints.NA}
	report.EInhibitAnyPolicyNotCritical = lints.ResultStruct{Result: lints.NA}
	report.EInvalidCertificateVersion = lints.ResultStruct{Result: lints.NA}
	report.EIssuerFieldEmpty = lints.ResultStruct{Result: lints.NA}
	report.ENameConstraintEmpty = lints.ResultStruct{Result: lints.NA}
	report.ENameConstraintMaximumNotAbsent = lints.ResultStruct{Result: lints.NA}
	report.ENameConstraintMinimumNonZero = lints.ResultStruct{Result: lints.NA}
	report.WNameConstraintOnEdiPartyName = lints.ResultStruct{Result: lints.NA}
	report.WNameConstraintOnRegisteredId = lints.ResultStruct{Result: lints.NA}
	report.WNameConstraintOnX400 = lints.ResultStruct{Result: lints.NA}
	report.EOldRootCaRsaModLessThan_2048Bits = lints.ResultStruct{Result: lints.NA}
	report.EOldSubCaRsaModLessThan_1024Bits = lints.ResultStruct{Result: lints.NA}
	report.EOldSubCertRsaModLessThan_1024Bits = lints.ResultStruct{Result: lints.NA}
	report.EPathLenConstraintImproperlyIncluded = lints.ResultStruct{Result: lints.NA}
	report.EPathLenConstraintZeroOrLess = lints.ResultStruct{Result: lints.NA}
	report.EPublicKeyTypeNotAllowed = lints.ResultStruct{Result: lints.NA}
	report.WRootCaBasicConstraintsPathLenConstraintFieldPresent = lints.ResultStruct{Result: lints.NA}
	report.WRootCaContainsCertPolicy = lints.ResultStruct{Result: lints.NA}
	report.ERootCaExtendedKeyUsagePresent = lints.ResultStruct{Result: lints.NA}
	report.ERsaExpNegative = lints.ResultStruct{Result: lints.NA}
	report.WRsaModFactorsSmallerThan_752 = lints.ResultStruct{Result: lints.NA}
	report.ERsaModLessThan_2048Bits = lints.ResultStruct{Result: lints.NA}
	report.WRsaModNotOdd = lints.ResultStruct{Result: lints.NA}
	report.WRsaPublicExponentNotInRange = lints.ResultStruct{Result: lints.NA}
	report.ERsaPublicExponentNotOdd = lints.ResultStruct{Result: lints.NA}
	report.ERsaPublicExponentTooSmall = lints.ResultStruct{Result: lints.NA}
	report.ESanDnsNameIncludesNullChar = lints.ResultStruct{Result: lints.NA}
	report.ESanDnsNameStartsWithPeriod = lints.ResultStruct{Result: lints.NA}
	report.WSanIanaPubSuffixEmpty = lints.ResultStruct{Result: lints.NA}
	report.ESerialNumberLongerThan_20Octets = lints.ResultStruct{Result: lints.NA}
	report.ESerialNumberNotPositive = lints.ResultStruct{Result: lints.NA}
	report.WSubCaAiaDoesNotContainIssuingCaUrl = lints.ResultStruct{Result: lints.NA}
	report.ESubCaAiaDoesNotContainOcspUrl = lints.ResultStruct{Result: lints.NA}
	report.ESubCaAiaMissing = lints.ResultStruct{Result: lints.NA}
	report.WSubCaCertificatePoliciesMarkedCritical = lints.ResultStruct{Result: lints.NA}
	report.ESubCaCertificatePoliciesMissing = lints.ResultStruct{Result: lints.NA}
	report.ESubCaCrlDistributionPointsDoesNotContainUrl = lints.ResultStruct{Result: lints.NA}
	report.ESubCaCrlDistributionPointsMarkedCritical = lints.ResultStruct{Result: lints.NA}
	report.ESubCaCrlDistributionPointsMissing = lints.ResultStruct{Result: lints.NA}
	report.WSubCaEkuCritical = lints.ResultStruct{Result: lints.NA}
	report.WSubCaNameConstraintsNotCritical = lints.ResultStruct{Result: lints.NA}
	report.ESubCaNoDnsNameConstraints = lints.ResultStruct{Result: lints.NA}
	report.ESubCaNoIpNameConstraints = lints.ResultStruct{Result: lints.NA}
	report.ESubCertAiaDoesNotContainIssuingCaUrl = lints.ResultStruct{Result: lints.NA}
	report.ESubCertAiaDoesNotContainOcspUrl = lints.ResultStruct{Result: lints.NA}
	report.ESubCertAiaMissing = lints.ResultStruct{Result: lints.NA}
	report.ESubCertCertPolicyEmpty = lints.ResultStruct{Result: lints.NA}
	report.WSubCertCertificatePoliciesMarkedCritical = lints.ResultStruct{Result: lints.NA}
	report.ESubCertCrlDistributionPointsDoesNotContainUrl = lints.ResultStruct{Result: lints.NA}
	report.ESubCertCrlDistributionPointsMarkedCritical = lints.ResultStruct{Result: lints.NA}
	report.WSubCertEkuExtraValues = lints.ResultStruct{Result: lints.NA}
	report.ESubCertEkuMissing = lints.ResultStruct{Result: lints.NA}
	report.ESubCertEkuServerAuthClientAuthMissing = lints.ResultStruct{Result: lints.NA}
	report.ESubCertKeyUsageCertSignBitSet = lints.ResultStruct{Result: lints.NA}
	report.ESubCertOrSubCaUsingSha1 = lints.ResultStruct{Result: lints.NA}
	report.WSubCertSha1ExpirationTooLong = lints.ResultStruct{Result: lints.NA}
	report.ESubjectCommonNameDisallowed = lints.ResultStruct{Result: lints.NA}
	report.NSubjectCommonNameIncluded = lints.ResultStruct{Result: lints.NA}
	report.ESubjectCommonNameNotFromSan = lints.ResultStruct{Result: lints.NA}
	report.ESubjectContainsNoninformationalValue = lints.ResultStruct{Result: lints.NA}
	report.ESubjectContainsReservedIp = lints.ResultStruct{Result: lints.NA}
	report.ESubjectCountryNotIso = lints.ResultStruct{Result: lints.NA}
	report.ESubjectEmptyWithoutSan = lints.ResultStruct{Result: lints.NA}
	report.ESubjectInfoAccessMarkedCritical = lints.ResultStruct{Result: lints.NA}
	report.ESubjectLocalityWithoutOrg = lints.ResultStruct{Result: lints.NA}
	report.ESubjectNotDn = lints.ResultStruct{Result: lints.NA}
	report.ESubjectOrgWithoutCountry = lints.ResultStruct{Result: lints.NA}
	report.ESubjectOrgWithoutLocalityOrProvince = lints.ResultStruct{Result: lints.NA}
	report.ESubjectPostalWithoutOrg = lints.ResultStruct{Result: lints.NA}
	report.ESubjectProvinceWithoutOrg = lints.ResultStruct{Result: lints.NA}
	report.ESubjectStreetWithoutOrg = lints.ResultStruct{Result: lints.NA}
	report.EUtcTimeDoesNotIncludeSeconds = lints.ResultStruct{Result: lints.NA}
	report.EUtcTimeNotInZulu = lints.ResultStruct{Result: lints.NA}
	report.EValidityTimeNotPositive = lints.ResultStruct{Result: lints.NA}
	report.EWrongTimeFormatPre2050 = lints.ResultStruct{Result: lints.NA}
	report.ERsaNoPublicKey = lints.ResultStruct{Result: lints.NA}
	report.ESubCertCertificatePoliciesMissing = lints.ResultStruct{Result: lints.NA}
	report.ESubCertKeyUsageCrlSignBitSet = lints.ResultStruct{Result: lints.NA}
	report.ESubjectCommonNameMaxLength = lints.ResultStruct{Result: lints.NA}
	report.ESubjectLocalityNameMaxLength = lints.ResultStruct{Result: lints.NA}
	report.ESubjectOrganizationNameMaxLength = lints.ResultStruct{Result: lints.NA}
	report.ESubjectOrganizationalUnitNameMaxLength = lints.ResultStruct{Result: lints.NA}
	report.ESubjectStateNameMaxLength = lints.ResultStruct{Result: lints.NA}
}

//Calls all other checks on parsed certs
func ParsedTestHandler(cert *x509.Certificate) (map[string]string, error) {
	if cert == nil {
		return nil, errors.New("zlint: nil pointer passed in, no data returned")
	}
	//run all tests
	var out map[string]string = make(map[string]string)
	for _, l := range lints.Lints {
		result, err := l.ExecuteTest(cert)
		if err != nil {
			return out, err
		}
		out[l.Name] = lints.EnumToString(result.Result)
	}

	return out, nil
}

//expects Base64 encoded ASN.1 DER string, wrapper for ParsedTestHandler
func Lint64(certIn string) (map[string]string, error) {

	der, err := base64.StdEncoding.DecodeString(certIn) //decode string
	if err != nil {
		return nil, err //parsing error, no tests performed
	}

	cert, err := x509.ParseCertificate(der)
	if err != nil {
		return nil, err //parsing error, no tests performed
	}

	return ParsedTestHandler(cert) //return available reports & error from main testing function
}
