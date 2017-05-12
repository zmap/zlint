package lints

import (
	"github.com/zmap/zcrypto/x509"
	"time"
)

// global
var (
	Lints map[string]*Lint = make(map[string]*Lint)
)

type lintReportUpdater func(*LintReport, ResultStruct)

const ZLintVersion = 1

type ZLintResult struct {
	ZLintVersion    int64       `json:"version"`
	Timestamp       int64       `json:"timestamp"`
	ZLint           *LintReport `json:"lints"`
	NoticesPresent  bool        `json:"notices_present"`
	WarningsPresent bool        `json:"warnings_present"`
	ErrorsPresent   bool        `json:"errors_present"`
	FatalsPresent   bool        `json:"fatals_present"`
}

type LintReport struct {
	EBasicConstraintsNotCritical                         ResultStruct `json:"e_basic_constraints_not_critical,omitempty"`
	EIanBareWildcard                                     ResultStruct `json:"e_ian_bare_wildcard,omitempty"`
	EIanWildcardNotFirst                                 ResultStruct `json:"e_ian_wildcard_not_first,omitempty"`
	ESanBareWildcard                                     ResultStruct `json:"e_san_bare_wildcard,omitempty"`
	ESanWildcardNotFirst                                 ResultStruct `json:"e_san_wildcard_not_first,omitempty"`
	ECaCountryNameInvalid                                ResultStruct `json:"e_ca_country_name_invalid,omitempty"`
	ECaCountryNameMissing                                ResultStruct `json:"e_ca_country_name_missing,omitempty"`
	ECaCrlSignNotSet                                     ResultStruct `json:"e_ca_crl_sign_not_set,omitempty"`
	NCaDigitalSignatureNotSet                            ResultStruct `json:"n_ca_digital_signature_not_set,omitempty"`
	ECaKeyCertSignNotSet                                 ResultStruct `json:"e_ca_key_cert_sign_not_set,omitempty"`
	ECaKeyUsageMissing                                   ResultStruct `json:"e_ca_key_usage_missing,omitempty"`
	ECaKeyUsageNotCritical                               ResultStruct `json:"e_ca_key_usage_not_critical,omitempty"`
	ECaOrganizationNameMissing                           ResultStruct `json:"e_ca_organization_name_missing,omitempty"`
	ECaSubjectFieldEmpty                                 ResultStruct `json:"e_ca_subject_field_empty,omitempty"`
	ECertContainsUniqueIdentifier                        ResultStruct `json:"e_cert_contains_unique_identifier,omitempty"`
	ECertExtensionsVersionNot_3                          ResultStruct `json:"e_cert_extensions_version_not_3,omitempty"`
	ECabDvConflictsWithLocality                          ResultStruct `json:"e_cab_dv_conflicts_with_locality,omitempty"`
	ECabDvConflictsWithOrg                               ResultStruct `json:"e_cab_dv_conflicts_with_org,omitempty"`
	ECabDvConflictsWithPostal                            ResultStruct `json:"e_cab_dv_conflicts_with_postal,omitempty"`
	ECabDvConflictsWithProvince                          ResultStruct `json:"e_cab_dv_conflicts_with_province,omitempty"`
	ECabDvConflictsWithStreet                            ResultStruct `json:"e_cab_dv_conflicts_with_street,omitempty"`
	ECertPolicyIvRequiresCountry                         ResultStruct `json:"e_cert_policy_iv_requires_country,omitempty"`
	ECertPolicyIvRequiresProvinceOrLocality              ResultStruct `json:"e_cert_policy_iv_requires_province_or_locality,omitempty"`
	ECertPolicyOvRequiresCountry                         ResultStruct `json:"e_cert_policy_ov_requires_country,omitempty"`
	ECertPolicyOvRequiresProvinceOrLocality              ResultStruct `json:"e_cert_policy_ov_requires_province_or_locality,omitempty"`
	ECabOvRequiresOrg                                    ResultStruct `json:"e_cab_ov_requires_org,omitempty"`
	ECabIvRequiresPersonalName                           ResultStruct `json:"e_cab_iv_requires_personal_name,omitempty"`
	ECertUniqueIdentifierVersionNot_2Or_3                ResultStruct `json:"e_cert_unique_identifier_version_not_2_or_3,omitempty"`
	EDhParamsMissing                                     ResultStruct `json:"e_dh_params_missing,omitempty"`
	EDistributionPointIncomplete                         ResultStruct `json:"e_distribution_point_incomplete,omitempty"`
	WDistributionPointMissingLdapOrUri                   ResultStruct `json:"w_distribution_point_missing_ldap_or_uri,omitempty"`
	EDsaImproperModulusOrDivisorSize                     ResultStruct `json:"e_dsa_improper_modulus_or_divisor_size,omitempty"`
	EDsaShorterThan_2048Bits                             ResultStruct `json:"e_dsa_shorter_than_2048_bits,omitempty"`
	EEcImproperCurves                                    ResultStruct `json:"e_ec_improper_curves,omitempty"`
	WEkuCriticalImproperly                               ResultStruct `json:"w_eku_critical_improperly,omitempty"`
	EEvBusinessCategoryMissing                           ResultStruct `json:"e_ev_business_category_missing,omitempty"`
	EEvCountryNameMissing                                ResultStruct `json:"e_ev_country_name_missing,omitempty"`
	EEvLocalityNameMissing                               ResultStruct `json:"e_ev_locality_name_missing,omitempty"`
	EEvOrganizationNameMissing                           ResultStruct `json:"e_ev_organization_name_missing,omitempty"`
	EEvSerialNumberMissing                               ResultStruct `json:"e_ev_serial_number_missing,omitempty"`
	EEvValidTimeTooLong                                  ResultStruct `json:"e_ev_valid_time_too_long,omitempty"`
	WExtAiaAccessLocationMissing                         ResultStruct `json:"w_ext_aia_access_location_missing,omitempty"`
	EExtAiaMarkedCritical                                ResultStruct `json:"e_ext_aia_marked_critical,omitempty"`
	EExtAuthorityKeyIdentifierCritical                   ResultStruct `json:"e_ext_authority_key_identifier_critical,omitempty"`
	EExtAuthorityKeyIdentifierMissing                    ResultStruct `json:"e_ext_authority_key_identifier_missing,omitempty"`
	EExtAuthorityKeyIdentifierNoKeyIdentifier            ResultStruct `json:"e_ext_authority_key_identifier_no_key_identifier,omitempty"`
	WExtCertPolicyContainsNoticeref                      ResultStruct `json:"w_ext_cert_policy_contains_noticeref,omitempty"`
	EExtCertPolicyDisallowedAnyPolicyQualifier           ResultStruct `json:"e_ext_cert_policy_disallowed_any_policy_qualifier,omitempty"`
	EExtCertPolicyDuplicate                              ResultStruct `json:"e_ext_cert_policy_duplicate,omitempty"`
	EExtCertPolicyExplicitTextIa5String                  ResultStruct `json:"e_ext_cert_policy_explicit_text_ia5_string,omitempty"`
	WExtCertPolicyExplicitTextIncludesControl            ResultStruct `json:"w_ext_cert_policy_explicit_text_includes_control,omitempty"`
	WExtCertPolicyExplicitTextNotNfc                     ResultStruct `json:"w_ext_cert_policy_explicit_text_not_nfc,omitempty"`
	WExtCertPolicyExplicitTextNotUtf8                    ResultStruct `json:"w_ext_cert_policy_explicit_text_not_utf8,omitempty"`
	EExtCertPolicyExplicitTextTooLong                    ResultStruct `json:"e_ext_cert_policy_explicit_text_too_long,omitempty"`
	WExtCrlDistributionMarkedCritical                    ResultStruct `json:"w_ext_crl_distribution_marked_critical,omitempty"`
	EExtDuplicateExtension                               ResultStruct `json:"e_ext_duplicate_extension,omitempty"`
	EExtFreshestCrlMarkedCritical                        ResultStruct `json:"e_ext_freshest_crl_marked_critical,omitempty"`
	WExtIanCritical                                      ResultStruct `json:"w_ext_ian_critical,omitempty"`
	EExtIanDnsNotIa5String                               ResultStruct `json:"e_ext_ian_dns_not_ia5_string,omitempty"`
	EExtIanEmptyName                                     ResultStruct `json:"e_ext_ian_empty_name,omitempty"`
	EExtIanNoEntries                                     ResultStruct `json:"e_ext_ian_no_entries,omitempty"`
	EExtIanRfc822FormatInvalid                           ResultStruct `json:"e_ext_ian_rfc822_format_invalid,omitempty"`
	EExtIanSpaceDnsName                                  ResultStruct `json:"e_ext_ian_space_dns_name,omitempty"`
	EExtIanUriFormatInvalid                              ResultStruct `json:"e_ext_ian_uri_format_invalid,omitempty"`
	EExtIanUriHostNotFqdnOrIp                            ResultStruct `json:"e_ext_ian_uri_host_not_fqdn_or_ip,omitempty"`
	EExtIanUriNotIa5                                     ResultStruct `json:"e_ext_ian_uri_not_ia5,omitempty"`
	EExtIanUriRelative                                   ResultStruct `json:"e_ext_ian_uri_relative,omitempty"`
	EExtKeyUsageCertSignWithoutCa                        ResultStruct `json:"e_ext_key_usage_cert_sign_without_ca,omitempty"`
	WExtKeyUsageNotCritical                              ResultStruct `json:"w_ext_key_usage_not_critical,omitempty"`
	EExtKeyUsageWithoutBits                              ResultStruct `json:"e_ext_key_usage_without_bits,omitempty"`
	EExtNameConstraintsNotCritical                       ResultStruct `json:"e_ext_name_constraints_not_critical,omitempty"`
	EExtNameConstraintsNotInCa                           ResultStruct `json:"e_ext_name_constraints_not_in_ca,omitempty"`
	EExtPolicyConstraintsEmpty                           ResultStruct `json:"e_ext_policy_constraints_empty,omitempty"`
	EExtPolicyConstraintsNotCritical                     ResultStruct `json:"e_ext_policy_constraints_not_critical,omitempty"`
	EExtPolicyMapAnyPolicy                               ResultStruct `json:"e_ext_policy_map_any_policy,omitempty"`
	WExtPolicyMapNotCritical                             ResultStruct `json:"w_ext_policy_map_not_critical,omitempty"`
	WExtPolicyMapNotInCertPolicy                         ResultStruct `json:"w_ext_policy_map_not_in_cert_policy,omitempty"`
	EExtSanContainsReservedIp                            ResultStruct `json:"e_ext_san_contains_reserved_ip,omitempty"`
	WExtSanCriticalWithSubjectDn                         ResultStruct `json:"w_ext_san_critical_with_subject_dn,omitempty"`
	EExtSanDirectoryNamePresent                          ResultStruct `json:"e_ext_san_directory_name_present,omitempty"`
	EExtSanDnsNotIa5String                               ResultStruct `json:"e_ext_san_dns_not_ia5_string,omitempty"`
	EExtSanDnsSyntaxIncorrect                            ResultStruct `json:"e_ext_san_dns_syntax_incorrect,omitempty"`
	EExtSanDnsnameNotFqdn                                ResultStruct `json:"e_ext_san_dnsname_not_fqdn,omitempty"`
	EExtSanEdiPartyNamePresent                           ResultStruct `json:"e_ext_san_edi_party_name_present,omitempty"`
	EExtSanEmptyName                                     ResultStruct `json:"e_ext_san_empty_name,omitempty"`
	EExtSanMissing                                       ResultStruct `json:"e_ext_san_missing,omitempty"`
	EExtSanNoEntries                                     ResultStruct `json:"e_ext_san_no_entries,omitempty"`
	EExtSanNotCriticalWithoutSubject                     ResultStruct `json:"e_ext_san_not_critical_without_subject,omitempty"`
	EExtSanOtherNamePresent                              ResultStruct `json:"e_ext_san_other_name_present,omitempty"`
	EExtSanRegisteredIdPresent                           ResultStruct `json:"e_ext_san_registered_id_present,omitempty"`
	EExtSanRfc822FormatInvalid                           ResultStruct `json:"e_ext_san_rfc822_format_invalid,omitempty"`
	EExtSanRfc822NamePresent                             ResultStruct `json:"e_ext_san_rfc822_name_present,omitempty"`
	EExtSanSpaceDnsName                                  ResultStruct `json:"e_ext_san_space_dns_name,omitempty"`
	EExtSanUniformResourceIdentifierPresent              ResultStruct `json:"e_ext_san_uniform_resource_identifier_present,omitempty"`
	EExtSanUriFormatInvalid                              ResultStruct `json:"e_ext_san_uri_format_invalid,omitempty"`
	EExtSanUriHostNotFqdnOrIp                            ResultStruct `json:"e_ext_san_uri_host_not_fqdn_or_ip,omitempty"`
	EExtSanUriNotIa5                                     ResultStruct `json:"e_ext_san_uri_not_ia5,omitempty"`
	EExtSanUriRelative                                   ResultStruct `json:"e_ext_san_uri_relative,omitempty"`
	EExtSubjectDirectoryAttrCritical                     ResultStruct `json:"e_ext_subject_directory_attr_critical,omitempty"`
	EExtSubjectKeyIdentifierCritical                     ResultStruct `json:"e_ext_subject_key_identifier_critical,omitempty"`
	EExtSubjectKeyIdentifierMissingCa                    ResultStruct `json:"e_ext_subject_key_identifier_missing_ca,omitempty"`
	WExtSubjectKeyIdentifierMissingSubCert               ResultStruct `json:"w_ext_subject_key_identifier_missing_sub_cert,omitempty"`
	EGeneralizedTimeDoesNotIncludeSeconds                ResultStruct `json:"e_generalized_time_does_not_include_seconds,omitempty"`
	EGeneralizedTimeIncludesFractionSeconds              ResultStruct `json:"e_generalized_time_includes_fraction_seconds,omitempty"`
	EGeneralizedTimeNotInZulu                            ResultStruct `json:"e_generalized_time_not_in_zulu,omitempty"`
	WGtldUnderConsideration                              ResultStruct `json:"w_gtld_under_consideration,omitempty"`
	EIanDnsNameIncludesNullChar                          ResultStruct `json:"e_ian_dns_name_includes_null_char,omitempty"`
	EIanDnsNameStartsWithPeriod                          ResultStruct `json:"e_ian_dns_name_starts_with_period,omitempty"`
	WIanIanaPubSuffixEmpty                               ResultStruct `json:"w_ian_iana_pub_suffix_empty,omitempty"`
	EInhibitAnyPolicyNotCritical                         ResultStruct `json:"e_inhibit_any_policy_not_critical,omitempty"`
	EInvalidCertificateVersion                           ResultStruct `json:"e_invalid_certificate_version,omitempty"`
	EIssuerFieldEmpty                                    ResultStruct `json:"e_issuer_field_empty,omitempty"`
	ENameConstraintEmpty                                 ResultStruct `json:"e_name_constraint_empty,omitempty"`
	ENameConstraintMaximumNotAbsent                      ResultStruct `son:"e_name_constraint_maximum_not_absent,omitempty"`
	ENameConstraintMinimumNonZero                        ResultStruct `json:"e_name_constraint_minimum_non_zero,omitempty"`
	WNameConstraintOnEdiPartyName                        ResultStruct `json:"w_name_constraint_on_edi_party_name,omitempty"`
	WNameConstraintOnRegisteredId                        ResultStruct `json:"w_name_constraint_on_registered_id,omitempty"`
	WNameConstraintOnX400                                ResultStruct `json:"w_name_constraint_on_x400,omitempty"`
	EOldRootCaRsaModLessThan_2048Bits                    ResultStruct `json:"e_old_root_ca_rsa_mod_less_than_2048_bits,omitempty"`
	EOldSubCaRsaModLessThan_1024Bits                     ResultStruct `json:"e_old_sub_ca_rsa_mod_less_than_1024_bits,omitempty"`
	EOldSubCertRsaModLessThan_1024Bits                   ResultStruct `json:"e_old_sub_cert_rsa_mod_less_than_1024_bits,omitempty"`
	EPathLenConstraintImproperlyIncluded                 ResultStruct `json:"e_path_len_constraint_improperly_included,omitempty"`
	EPathLenConstraintZeroOrLess                         ResultStruct `json:"e_path_len_constraint_zero_or_less,omitempty"`
	EPublicKeyTypeNotAllowed                             ResultStruct `json:"e_public_key_type_not_allowed,omitempty"`
	WRootCaBasicConstraintsPathLenConstraintFieldPresent ResultStruct `json:"w_root_ca_basic_constraints_path_len_constraint_field_present,omitempty"`
	WRootCaContainsCertPolicy                            ResultStruct `json:"w_root_ca_contains_cert_policy,omitempty"`
	ERootCaExtendedKeyUsagePresent                       ResultStruct `json:"e_root_ca_extended_key_usage_present,omitempty"`
	ERsaExpNegative                                      ResultStruct `json:"e_rsa_exp_negative,omitempty"`
	WRsaModFactorsSmallerThan_752                        ResultStruct `json:"w_rsa_mod_factors_smaller_than_752,omitempty"`
	ERsaModLessThan_2048Bits                             ResultStruct `json:"e_rsa_mod_less_than_2048_bits,omitempty"`
	WRsaModNotOdd                                        ResultStruct `json:"w_rsa_mod_not_odd,omitempty"`
	WRsaPublicExponentNotInRange                         ResultStruct `json:"w_rsa_public_exponent_not_in_range,omitempty"`
	ERsaPublicExponentNotOdd                             ResultStruct `json:"e_rsa_public_exponent_not_odd,omitempty"`
	ERsaPublicExponentTooSmall                           ResultStruct `json:"e_rsa_public_exponent_too_small,omitempty"`
	ESanDnsNameIncludesNullChar                          ResultStruct `json:"e_san_dns_name_includes_null_char,omitempty"`
	ESanDnsNameStartsWithPeriod                          ResultStruct `json:"e_san_dns_name_starts_with_period,omitempty"`
	WSanIanaPubSuffixEmpty                               ResultStruct `json:"w_san_iana_pub_suffix_empty,omitempty"`
	ESerialNumberLongerThan_20Octets                     ResultStruct `json:"e_serial_number_longer_than_20_octets,omitempty"`
	ESerialNumberNotPositive                             ResultStruct `json:"e_serial_number_not_positive,omitempty"`
	WSubCaAiaDoesNotContainIssuingCaUrl                  ResultStruct `json:"w_sub_ca_aia_does_not_contain_issuing_ca_url,omitempty"`
	ESubCaAiaDoesNotContainOcspUrl                       ResultStruct `json:"e_sub_ca_aia_does_not_contain_ocsp_url,omitempty"`
	ESubCaAiaMissing                                     ResultStruct `json:"e_sub_ca_aia_missing,omitempty"`
	WSubCaCertificatePoliciesMarkedCritical              ResultStruct `json:"w_sub_ca_certificate_policies_marked_critical,omitempty"`
	ESubCaCertificatePoliciesMissing                     ResultStruct `json:"e_sub_ca_certificate_policies_missing,omitempty"`
	ESubCaCrlDistributionPointsDoesNotContainUrl         ResultStruct `json:"e_sub_ca_crl_distribution_points_does_not_contain_url,omitempty"`
	ESubCaCrlDistributionPointsMarkedCritical            ResultStruct `json:"e_sub_ca_crl_distribution_points_marked_critical,omitempty"`
	ESubCaCrlDistributionPointsMissing                   ResultStruct `json:"e_sub_ca_crl_distribution_points_missing,omitempty"`
	WSubCaEkuCritical                                    ResultStruct `json:"w_sub_ca_eku_critical,omitempty"`
	WSubCaNameConstraintsNotCritical                     ResultStruct `json:"w_sub_ca_name_constraints_not_critical,omitempty"`
	ESubCaNoDnsNameConstraints                           ResultStruct `json:"e_sub_ca_no_dns_name_constraints,omitempty"`
	ESubCaNoIpNameConstraints                            ResultStruct `json:"e_sub_ca_no_ip_name_constraints,omitempty"`
	ESubCertAiaDoesNotContainIssuingCaUrl                ResultStruct `json:"e_sub_cert_aia_does_not_contain_issuing_ca_url,omitempty"`
	ESubCertAiaDoesNotContainOcspUrl                     ResultStruct `json:"e_sub_cert_aia_does_not_contain_ocsp_url,omitempty"`
	ESubCertAiaMissing                                   ResultStruct `json:"e_sub_cert_aia_missing,omitempty"`
	ESubCertCertPolicyEmpty                              ResultStruct `json:"e_sub_cert_cert_policy_empty,omitempty"`
	WSubCertCertificatePoliciesMarkedCritical            ResultStruct `json:"w_sub_cert_certificate_policies_marked_critical,omitempty"`
	ESubCertCrlDistributionPointsDoesNotContainUrl       ResultStruct `json:"e_sub_cert_crl_distribution_points_does_not_contain_url,omitempty"`
	ESubCertCrlDistributionPointsMarkedCritical          ResultStruct `json:"e_sub_cert_crl_distribution_points_marked_critical,omitempty"`
	WSubCertEkuExtraValues                               ResultStruct `json:"w_sub_cert_eku_extra_values,omitempty"`
	ESubCertEkuMissing                                   ResultStruct `json:"e_sub_cert_eku_missing,omitempty"`
	ESubCertEkuServerAuthClientAuthMissing               ResultStruct `json:"e_sub_cert_eku_server_auth_client_auth_missing,omitempty"`
	ESubCertKeyUsageCertSignBitSet                       ResultStruct `json:"e_sub_cert_key_usage_cert_sign_bit_set,omitempty"`
	ESubCertOrSubCaUsingSha1                             ResultStruct `json:"e_sub_cert_or_sub_ca_using_sha1,omitempty"`
	WSubCertSha1ExpirationTooLong                        ResultStruct `json:"w_sub_cert_sha1_expiration_too_long,omitempty"`
	ESubjectCommonNameDisallowed                         ResultStruct `json:"e_subject_common_name_disallowed,omitempty"`
	NSubjectCommonNameIncluded                           ResultStruct `json:"n_subject_common_name_included,omitempty"`
	ESubjectCommonNameNotFromSan                         ResultStruct `json:"e_subject_common_name_not_from_san,omitempty"`
	ESubjectContainsNoninformationalValue                ResultStruct `json:"e_subject_contains_noninformational_value,omitempty"`
	ESubjectContainsReservedIp                           ResultStruct `json:"e_subject_contains_reserved_ip,omitempty"`
	ESubjectCountryNotIso                                ResultStruct `json:"e_subject_country_not_iso,omitempty"`
	ESubjectEmptyWithoutSan                              ResultStruct `json:"e_subject_empty_without_san,omitempty"`
	ESubjectInfoAccessMarkedCritical                     ResultStruct `json:"e_subject_info_access_marked_critical,omitempty"`
	ESubjectLocalityWithoutOrg                           ResultStruct `json:"e_subject_locality_without_org,omitempty"`
	ESubjectNotDn                                        ResultStruct `json:"e_subject_not_dn,omitempty"`
	ESubjectOrgWithoutCountry                            ResultStruct `json:"e_subject_org_without_country,omitempty"`
	ESubjectOrgWithoutLocalityOrProvince                 ResultStruct `json:"e_subject_org_without_locality_or_province,omitempty"`
	ESubjectPostalWithoutOrg                             ResultStruct `json:"e_subject_postal_without_org,omitempty"`
	ESubjectProvinceWithoutOrg                           ResultStruct `json:"e_subject_province_without_org,omitempty"`
	ESubjectStreetWithoutOrg                             ResultStruct `json:"e_subject_street_without_org,omitempty"`
	EUtcTimeDoesNotIncludeSeconds                        ResultStruct `json:"e_utc_time_does_not_include_seconds,omitempty"`
	EUtcTimeNotInZulu                                    ResultStruct `json:"e_utc_time_not_in_zulu,omitempty"`
	EValidityTimeNotPositive                             ResultStruct `json:"e_validity_time_not_positive,omitempty"`
	EWrongTimeFormatPre2050                              ResultStruct `json:"e_wrong_time_format_pre2050,omitempty"`
	ERsaNoPublicKey                                      ResultStruct `json:"e_rsa_no_public_key,omitempty"`
	ESubCertCertificatePoliciesMissing                   ResultStruct `json:"e_sub_cert_certificate_policies_missing,omitempty"`
	ESubCertKeyUsageCrlSignBitSet                        ResultStruct `json:"e_sub_cert_key_usage_crl_sign_bit_set,omitempty"`
}

func (result *ZLintResult) Execute(cert *x509.Certificate) error {
	for _, l := range Lints {
		res, err := l.ExecuteTest(cert)
		if err != nil {
			return err
		}
		l.updateReport(result.ZLint, res)
		result.updateErrorStatePresent(res)
	}
	return nil
}

func (zlintResult *ZLintResult) updateErrorStatePresent(result ResultStruct) {
	switch result.Result {
	case Notice:
		zlintResult.NoticesPresent = true
	case Warn:
		zlintResult.WarningsPresent = true
	case Error:
		zlintResult.ErrorsPresent = true
	case Fatal:
		zlintResult.FatalsPresent = true
	}
}

type LintTest interface {
	// runs once globally
	Initialize() error
	CheckApplies(c *x509.Certificate) bool
	RunTest(c *x509.Certificate) (ResultStruct, error)
}

type Lint struct {
	Name          string
	Description   string
	Providence    string
	EffectiveDate time.Time
	Test          LintTest
	updateReport  lintReportUpdater
}

func (l *Lint) ExecuteTest(cert *x509.Certificate) (ResultStruct, error) {
	if !l.Test.CheckApplies(cert) {
		return ResultStruct{Result: NA}, nil
	} else if !l.EffectiveDate.IsZero() && l.EffectiveDate.After(cert.NotBefore) {
		return ResultStruct{Result: NE}, nil
	}

	return l.Test.RunTest(cert)
}

func RegisterLint(l *Lint) {
	if Lints == nil {
		Lints = make(map[string]*Lint)
	}
	l.Test.Initialize()
	Lints[l.Name] = l
}
