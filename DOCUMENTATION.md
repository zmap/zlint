# DOCUMENTATION

The following is a list of lints run on each certificate at this time.

**Lint Name** | **Description** | **Providence**
------------------- | :-------------------------------------------------------------------: | ---:
basic_constraints_not_critical | Conforming CAs must mark Basic Constraints as critical when it is included in CA certs | RFC 5280: 4.2.1.9
br_san_bare_wildcard | Wildcard MUST be accompanied by other data to it's right (Only checks DNSName) | -
br_san_wildcard_not_first | Wildcard MUST be in the first label of FQDN, ie not: www.\*.com (Only checks DNSName) | -
ca_country_name_invalid  | Root & Subordinate CA certificates must have a two-letter country code that is in ISO 3166-1 | CAB: 7.1.2.1
ca_country_name_missing | Root & Subordinate CA certificates must have a countryName present in subject information | CAB: 7.1.2.1
ca_crl_sign_not_set | Root & Subordinate CA certificate keyUsage extension's crlSign bit must be set | CAB: 7.1.2.1
ca_dig_sign_not_set | Root & Subordinate CA Certificates that wish to use their private key for signing OCSP responses will not be able to with out digital signature set | CAB: 7.1.2.1, 7.1.6.2
ca_key_cert_sign_not_set | Root & Subordinate CA certificate keyUsage extension's keyCertSign bit must be set | CAB: 7.1.2.1
ca_key_usage_missing | Root & Subordinate CA certificate keyUsage extension must be present | CAB: 7.1.2.1, RFC 5280: 4.2.1.3
ca_key_usage_not_critical | Root & Subordinate CA certificate keyUsage extension must be marked as critical | CAB: 7.1.2.1
ca_organization_name_missing | Root & Subordinate CA certificates must have a organizationName present in subject information | CAB: 7.1.2.1
ca_subject_field_empty | CA Certificates subject field must not be empty and must have a non-empty distingushed name | RFC 5280: 4.1.2.6
cert_contains_unique_identifier | CAs must not generate certificates with unique identifiers. | RFC 5280: 4.1.2.8
cert_extensions_verson_not_3 | The extensions field must only appear in version 3 certificates. | RFC 5280: 4.1.2.9
cert_policy_conflicts_with_locality | If certificate policy 2.23.140.1.2.1 is included, locality name must not be included in subject. | CAB: 7.1.6.1
cert_policy_conflicts_with_org | If certificate policy 2.23.140.1.2.1 is included, organization name must not be included in subject. | CAB: 7.1.6.1
cert_policy_conflicts_with_postal | If certificate policy 2.23.140.1.2.1 is included, postal code must not be included in subject. | CAB: 7.1.6.1
cert_policy_conflicts_with_province | If certificate policy 2.23.140.1.2.1 is included, stateOrProvince name must not be included in subject. | CAB: 7.1.6.1
cert_policy_conflicts_with_street | If certificate policy 2.23.140.1.2.1 is included, street address must not be included in subject. | CAB: 7.1.6.1
cert_policy_iv_requires_country | If certificate policy 2.23.140.1.2.3 is included, country name must be included in subject. | CAB: 7.1.6.1
cert_policy_iv_requires_province_or_locality | If certificate policy 2.23.140.1.2.3 is included, locality or stateOrProvince must be included in subject. | CAB: 7.1.6.1
cert_policy_ov_requires_country | If certificate policy 2.23.140.1.2.2 is included, country name must be included in subject. | CAB: 7.1.6.1
cert_policy_ov_requires_province_or_locality | If certificate policy 2.23.140.1.2.2 is included, locality or stateOrProvince must be included in subject. | CAB: 7.1.6.1
cert_policy_requires_org | If certificate policy 2.23.140.1.2.2 is included, organization name must be included in subject. | CAB: 7.1.6.1
cert_policy_requires_personal_name | If certificate policy 2.23.140.1.2.3 is included, organization or givenName AND surname must be included in subject. | CAB: 7.1.6.1
cert_unique_identifier_version_not_2_or_3 | Unique identifiers must only appear if the version is 2 or 3. | RFC 5280: 4.1.2.8
dh_params_missing | DH keys must have parameters | -
distribution_point_incomplete | A DistributionPoint from the CRLDistributionPoints extension MUST NOT consist of only the reasons field; either distributionPoint or CRLIssuer must be present | RFC 5280: 4.2.1.13
distribution_point_missing_ldap_or_uri | When present in the CRLDistributionPoints extension, DistributionPointName SHOULD include at least one LDAP or HTTP URI | RFC 5280: 4.2.1.13
dns_name_empty | DNSNames MUST NOT be empty if included | -
dns_name_includes_null_char | DNSNames MUST NOT include a null character  | -
dns_name_starts_with_period | DNSName MUST NOT start with a period | -
dsa_improper_modulus_or_divisor_size | minimum DSA modulus and divisor size is either L= 2048, N= 224 or L= 2048, N= 256 | CAB: 6.1.5
dsa_shorter_than_2048_bits | DSA modulus size must be at least 2048 bits | CAB: 6.1.5
ec_improper_curves | only one of NIST P‐256, P‐384, or P‐521 can be used for all types of certificate | CAB: 6.1.5
eku_critical_improperly | Conforming CAs SHOULD NOT mark extended key usage extension as critical if the anyExtendedKeyUsage KeyPurposedID is present | RFC 5280: 4.2.1.12
ev_business_category_missing | EV certificates must include businessCategory in subject | -
ev_country_name_missing | EV certificates must include countryName in subject | -
ev_locality_name_missing | EV certificates must include localityName in subject | -
ev_organization_name_missing | EV certificates must include organizationName in subject | -
ev_serial_number_missing | EV certificates must include serialNumber in subject | -
ev_valid_time_too_long | EV certificates must be 27 months in validity or less | -
ext_aia_access_location_missing | When the id-ad-caIssuers accessMethod is used, at least one instance SHOULD specify an accessLocation that is an HTTP or LDAP URI | RFC 5280: 4.2.2.1
ext_aia_marked_critical | Conforming CAs must mark the Authority Information Access extension as non-critical | RFC 5280: 4.2.2.1
ext_authority_key_identifier_critical | The authority key identifier extension must be non-critical. | RFC 5280: 4.2.1.1
ext_authority_key_identifier_missing | CAs must support key identifiers and include them in all certs | RFC 5280: 4.2 & 4.2.1.1
ext_authority_key_identifier_no_key_identifier | CAs must include keyIdentifier field of aki in all non-self-issued certs | RFC 5280: 4.2.1.1
ext_cert_policy_contains_noticeref | should not use the noticeRef option | RFC 5280: 4.2.1.4
ext_cert_policy_disallowed_any_policy_qualifier | when qualifiers are used with the special policy anyPolicy, the must be limited to qualifiers identified in this section (4.2.1.4) | RFC 5280: 4.2.1.4
ext_cert_policy_duplicate | a cert policy OID must not appear more than once in the extension | RFC 5280: 4.2.1.4
ext_cert_policy_explicit_text_ia5_string | must not encode explicitTest as IA5String | RFC 6818: 3
ext_cert_policy_explicit_text_includes_control | should not include any control charaters (such as U+009F) | RFC 6818: 3
ext_cert_policy_explicit_text_not_nfc | should be normalized according to unicode normalization form C if utf8 or bmpstring | RFC 6818: 3
ext_cert_policy_explicit_text_not_utf8 | should use the utf8string encoding for explicitText | RFC 6818: 3
ext_cert_policy_explicit_text_too_long | explicit text has a maximum size of 200 characters | RFC 6818: 3
ext_crl_distribution_marked_critical | If included, the CRL Distribution Points extension SHOULD NOT be marked critical | RFC 5280: 4.2.1.13
ext_duplicate_extension | A cert must not contain duplicate extensions. | RFC 5280: 4.2
ext_freshest_crl_marked_critical | Conforming CAs must mark this extension as non-critical | RFC 5280: 4.2.1.15
ext_ian_critical | issuer alternate name should be marked as non-critical | RFC 5280: 4.2.1.7
ext_ian_dns_not_ia5_string | dNSNames are IA5 strings | RFC 5280: 4.2.1.7
ext_ian_empty_name | general name fields must not be empty in issuer alternate name | RFC 5280: 4.2.1.7
ext_ian_no_entries | if present, the issuer alternative name extension must contain at least one entry | RFC 5280: 4.2.1.7
ext_ian_rfc822_format_invalid | email must be stored as a mailbox@domain.tld rfc822 name without "<>()" and "." before the "@" | RFC 5280: 4.2.1.7
ext_ian_space_dns_name | the dNSName " " must not be used | RFC 5280: 4.2.1.7
ext_ian_uri_format_invalid | The name MUST include both a  scheme (e.g., "http" or "ftp") and a scheme-specific-part. | RFC 5280: 4.2.1.7
ext_ian_uri_host_not_fqdn_or_ip | URIs that  include an authority ([RFC3986], Section 3.2) MUST include a fully  qualified domain name or IP address as the host. | RFC 5280: 4.2.1.7
ext_ian_uri_not_ia5 | URIs must be encoded as IA5 string in uniformResourceIdentifier | RFC 5280: 4.2.1.7
ext_ian_uri_relative | the name must not be a relative uri | RFC 5280: 4.2.1.7
ext_key_usage_cert_sign_without_ca | if the keyCertSign bit is asserted, then the cA bit MUST also be asserted | RFC 5280: 4.2.1.3 & 4.2.1.9
ext_key_usage_not_critical | SHOULD mark this extension as critical | RFC 5280: 4.2.1.3
ext_key_usage_without_bits | when included, at least one bit must be set to 1 | RFC 5280: 4.2.1.3
ext_name_constraints_not_critical | If it is included, conforming CAs must mark the name constrains extension as critical. | RFC 5280: 4.2.1.10
ext_name_constraints_not_in_ca | The name constraints extension must only be used in CA certificates | RFC 5280: 4.2.1.10
ext_policy_constraints_empty | Conforming CAs MUST NOT issue certificates where policy constraints is an empty sequence. That is, either the inhibitPolicyMapping field or the requireExplicityPolicy field MUST be present | RFC 5280: 4.2.1.11
ext_policy_constraints_not_critical | Conforming CAs must mark the policy constraints extension as critical. | RFC 5280: 4.2.1.11
ext_policy_map_any_policy | policies must not be mapped to or from the anyPolicy value | RFC 5280: 4.2.1.5
ext_policy_map_not_critical | Policy mappings should be marked as critical | RFC 5280: 4.2.1.5
ext_policy_map_not_in_cert_policy | each issuerDomainPolicy named in policy mappings extension should also be asserted in a certificate policies extension | RFC 5280: 4.2.1.5
ext_san_contains_reserved_ip | Certs issued after 2012-07-01 and expireing after 2015-11-01 must not contain a reserved ip address in the SAN extension. | CAB: 7.1.4.2.1
ext_san_critical_with_subject_dn | if the subject contains a distinguished name SAN should be non-critical | RFC 5280: 4.2.1.6
ext_san_directory_name_present | The Subject Alternate Name extension must contain only dnsName and ipaddress name types. | CAB: 7.1.4.2.1
ext_san_dnsname_not_fqdn | SAN dnsnames must be a fully qualified domain name. | CAB: 7.1.4.2.1
ext_san_dns_not_ia5_string | dNSNames are IA5 strings | RFC 5280: 4.2.1.6
ext_san_dns_syntax_incorrect | dns name must be in preferred name syntax | RFC 5280: 4.2.1.6
ext_san_edi_party_name_present | The Subject Alternate Name extension must contain only dnsName and ipaddress name types. | CAB: 7.1.4.2.1
ext_san_empty_name | general name fields must not be empty in san | RFC 5280: 4.2.1.6
ext_san_missing | Certificates must contain the Subject Alternate Name extension. | CAB: 7.1.4.2.1
ext_san_no_entries | if present, the san extension must contain at least one entry | RFC 5280: 4.2.1.6
ext_san_not_critical_without_subject | if there is an empty subject feild, then the san extension must be critical | RFC 5280: 4.2.1.6 & 4.1.2.6
ext_san_other_name_present | The Subject Alternate Name extension must contain only dnsName and ipaddress name types. | CAB: 7.1.4.2.1
ext_san_registered_id_present | The Subject Alternate Name extension must contain only dnsName and ipaddress name types. | CAB: 7.1.4.2.1
ext_san_rfc822_format_invalid | email must be stored as a mailbox@domain.tld rfc822 name without "<>()" and "." before the "@" | RFC 5280: 4.2.1.6
ext_san_rfc822_name_present | The Subject Alternate Name extension must contain only dnsName and ipaddress name types. | CAB: 7.1.4.2.1
ext_san_space_dns_name | the dNSName " " must not be used | RFC 5280: 4.2.1.6
ext_san_uniform_resource_identifier_present | The Subject Alternate Name extension must contain only dnsName and ipaddress name types. | CAB: 7.1.4.2.1
ext_san_uri_format_invalid | The name MUST include both a  scheme (e.g., "http" or "ftp") and a scheme-specific-part. | RFC 5280: 4.2.1.6
ext_san_uri_host_not_fqdn_or_ip | URIs that include an authority ([RFC3986], Section 3.2) MUST include a fully qualified domain name or IP address as the host. | RFC 5280: 4.2.1.6
ext_san_uri_not_ia5 | URIs must be encoded as IA5 string in uniformResourceIdentifier | RFC 5280: 4.2.1.6
ext_san_uri_relative | the name must not be a relative uri | RFC 5280: 4.2.1.6
ext_subject_directory_attr_critical | Conforming CAs must mark the Subject Directory Attributes extension as not critical | RFC 5280: 4.2.1.8
ext_subject_key_identifier_critical | The subject key identifier extension must be non-critical. | RFC 5280: 4.2.1.2
ext_subject_key_identifier_missing_ca | CAs must include ski in all CA certificates. | RFC 5280: 4.2 & 4.2.1.2
ext_subject_key_identifier_missing_sub_cert | CAs should include ski in end entity certs | RFC 5280: 4.2 & 4.2.1.2
generalized_time_does_not_include_seconds | Generalized time values must include seconds | RFC 5280: 4.1.2.5.2
generalized_time_includes_fraction_seconds | Generalized time values must not include fraction seconds | RFC 5280: 4.1.2.5.2
generalized_time_not_in_zulu | Generalized time values must be expressed in Greenwich Mean Time (Zulu) | RFC 5280: 4.1.2.5.2
gtld_under_consideration | CAs SHOULD NOT issue Certificates containing a new gTLD under consideration by ICANN. | CAB: 4.2.2
iana_pub_suffix_empty | Domain SHOULD NOT have bare public suffix | -
inhibit_any_policy_not_critical | Conforming CAs must mark this extension as critical | RFC 5280: 4.2.1.14
invalid_certificate_version | Certificate must be version 3 | CAB: 7.1.1
issuer_field_empty | Certificates issuer field must not be empty and must have a non-empty distingushed name | RFC 5280: 4.1.2.4
name_constraint_empty | Conforming CAs must not issue certificates where name constraints is an empty sequence. That is, either the permittedSubtree or excludedSubtree fields must be present | RFC 5280: 4.2.1.10
name_constraint_maximum_not_absent | In the name constraints name forms the maximum is not used and therefore MUST be absent | RFC 5280: 4.2.1.10
name_constraint_minimum_non_zero | In the name constraints name forms the minimum is not used and therefore MUST be zero | RFC 5280: 4.2.1.10
name_constraint_on_edi_party_name | The name constraints extension SHOULD NOT impose constraints on the ediPartyName name form | RFC 5280: 4.2.1.10
name_constraint_on_registered_id | The name constraints extension SHOULD NOT impose constraints on the registeredID name form | RFC 5280: 4.2.1.10
name_constraint_on_x400 | The name constraints extension SHOULD NOT impose constraints on the x400Address name form | RFC 5280: 4.2.1.10
old_root_ca_rsa_mod_less_than_2048_bits | In a validity period beginning on or before 31 dec 2010, root CA certificates using RSA public key algorithm must have 2048 bits of modulus | CAB: 6.1.5
old_sub_ca_rsa_mod_less_than_1024_bits | In a validity period beginning on or before 31 dec 2010 and ending on or before 31 dec 2013, subordinate CA certificates using RSA public key algorithm must have 1024 bits of modulus | CAB: 6.1.5
old_sub_cert_rsa_mod_less_than_1024_bits | In a validity period ending on or before 31 dec 2013, subscriber certificates using RSA public key algorithm must have 1024 bits of modulus | CAB: 6.1.5
path_len_constraint_improperly_included | CAs MUST NOT include the pathLenConstraint field unless the CA boolean is asserted and the keyCertSign bit is set. | RFC 5280: 4.2.1.9
path_len_constraint_zero_or_less | Where it appears, the pathLenConstraint field must be greater than or equal to zero | RFC 5280: 4.2.1.9
public_key_type_not_allowed | Certificates must have RSA, DSA, or ECDSA public key type. | CAB: 6.1.5
root_ca_basic_constraints_path_len_constraint_field_present | Root CA certificate basicConstraint extension pathLenConstraint field should not be present | CAB: 7.1.2.1
root_ca_contains_cert_policy | Root CA certs SHOULD NOT contain the certificate policies extension. | CAB: 7.1.2.1, 7.1.6.2
root_ca_extended_key_usage_present | Root CA certificates must not have the extendedKeyUsage extension present | CAB: 7.1.2.1
rsa_exp_negative | “RSA public key exponent must be positive” (negative exponent found) | -
rsa_mod_factors_smaller_than_752 | The modulus of a RSA public key should have no factors smaller than 752 | CAB: 6.1.6
rsa_mod_less_than_2048_bits | in validity period beginning after 31 Dec 2010, all certificates using RSA public key algorithm must have 2048 bits of modulus | CAB: 6.1.5
rsa_mod_not_odd | The modulus of a RSA public key should be an odd number | CAB: 6.1.6
rsa_public_exponent_not_in_range | The RSA public exponent SHOULD be in the range between 2^16 + 1 and 2^256 - 1 | CAB: 6.1.6
rsa_public_exponent_not_odd | RSA public key has to be an odd number | CAB: 6.1.5
rsa_public_exponent_too_small | RSA public key has to be greater or equal to 3 | CAB: 6.1.5
serial_number_longer_than_20_octets | Certificates must not have a serial number longer than 20 octets | RFC 5280: 4.1.2.2
serial_number_not_positive | Certificates must have a positive serial number | RFC 5280: 4.1.2.2
sub_ca_aia_does_not_contain_issuing_ca_url | Subordinate CA certificates authorityInformationAccess extension should contain the HTTP URL of the Issuing CA’s certificate | CAB: 7.1.2.2
sub_ca_aia_does_not_contain_oscp_URL | Subordinate CA certificates authorityInformationAccess extension must contain the HTTP URL of the Issuing CA’s OCSP responder | CAB: 7.1.2.2
sub_ca_aia_missing | Subordinate CA certificates must have a authorityInformationAccess extension | CAB: 7.1.2.2
sub_ca_certificate_policies_marked_critical | Subordinate CA certificates certificatePolicies extension should not be marked as critical | CAB: 7.1.2.2
sub_ca_certificate_policies_missing | Subordinate CA certificates must have a certificatePolicies extension | CAB: 7.1.2.2
sub_ca_crl_distribution_points_does_not_contain_url | Subordinate CA certificates cRLDistributionPoints extension must contain the HTTP URL of the CA’s CRL service. | CAB: 7.1.2.2
sub_ca_crl_distribution_points_marked_critical | Subordinate CA certificates must not mark the cRLDistributionPoints extension as critical | CAB: 7.1.2.2
sub_ca_crl_distribution_points_missing | Subordinate CA certificates must have a cRLDistributionPoints extension | CAB: 7.1.2.2
sub_ca_eku_critical | Subordinate CA certificate extkeyUsage extension should be marked non-critical if present | CAB: 7.1.2.2
sub_ca_name_constraints_not_critical | Subordinate CA certificate nameConstraints extension should be marked critical if present | CAB: 7.1.2.2
sub_ca_no_dns_name_contstraints | Subordanate CA certs must include in the name contraints extension either premitted dns names or prohibit the empty DNS name. | CAB: 7.1.5
sub_ca_no_ip_name_contstraints | Subordanate CA certs must include in the name contraints extension either premitted ip ranges or prohibit all ip addresses. | CAB: 7.1.5
sub_cert_aia_does_not_contain_issuing_CA_url | Subscriber certificates authorityInformationAccess extension should contain the HTTP URL of the Issuing CA’s certificate | CAB: 7.1.2.3
sub_cert_aia_does_not_contain_ocsp_URL | Subscriber certificates authorityInformationAccess extension must contain the HTTP URL of the Issuing CA’s OCSP responder | CAB: 7.1.2.3
sub_cert_aia_missing | Subscriber certificates must have a authorityInformationAccess extension | CAB: 7.1.2.3
sub_cert_certificate_policies_marked_critical | Subscriber certificate certificatePolicies extension should not be marked critical | CAB: 7.1.2.3
sub_cert_certificate_policies_missing | Subscriber certificate must include the certificatePolicies extension | CAB: 7.1.2.3
sub_cert_cert_policy_empty | Subscriber certificates must contain at least one policy identifier that indicates adherance to CAB standards | CAB: 7.1.6.4
sub_cert_crl_distribution_points_does_not_contain_url | Subscriber certificate cRLDistributionPoints extension must contain the HTTP URL of the CA’s CRL service. | CAB: 7.1.2.3
sub_cert_crl_distribution_points_marked_critical | Subscriber certificate cRLDistributionPoints extension must not be marked critical if present | CAB: 7.1.2.3
sub_cert_eku_extra_values | Subscriber certificates should have only id-kp-serverAuth, id-kp-clientAuth, or id-kp-emailProtection in extKeyUsage. Anything should not be included. | CAB: 7.1.2.3
sub_cert_eku_missing | Subscriber certificates must have the extended key usage extension present | CAB: 7.1.2.3
sub_cert_eku_server_auth_client_auth_missing | Subscriber certificates must have have either id-kp-serverAuth or id-kp-clientAuth or both present in extKeyUsage | CAB: 7.1.2.3
sub_cert_key_usage_cert_sign_bit_set | Subscriber certificates keyUsage extension keyCertSign bit must not be set | CAB: 7.1.2.3
sub_cert_key_usage_crl_sign_bit_set | Subscriber certificates keyUsage extension cRLSign bit must not be set | CAB: 7.1.2.3
sub_cert_or_sub_ca_using_sha1 | Subscriber certificates and subordinate CA certificates must not use the SHA-1 hash algorithm on a certificate with a NotBefore date after 1 Jan 2016 | CAB: 7.1.3
sub_cert_sha1_expiration_too_long | Subscriber certificates using the SHA1 algorithm should not have an expiration date greater than 1 Jan 2017 | CAB: 7.1.3
subject_common_name_disallowed | If present Common name MUST contain a single IP address or Fully‐Qualified Domain Name that is one of the values contained in the Certificate’s subjectAltName extension | CAB: 7.1.4.2.2
subject_common_name_included | Use of the common name field is discouraged. | CAB: 7.1.4.2.2
subject_common_name_not_from_san | The common name field must include only names from the SAN extension. | CAB: 7.1.4.2.2
subject_contains_noninformational_value | Subject name fields must not contain ".","-"," " or any other indication that the field has been ommited. | CAB: 7.1.4.2.2
subject_contains_reserved_ip | Certs issued after 2012-07-01 and expireing after 2015-11-01 must not conatain a reserved ip address in the common name field. | CAB: 7.1.4.2.1
subject_country_not_iso | The country name field must contain the two-letter ISO code for the country or XX. | CAB: 7.1.4.2.2
subject_empty_without_san | CAs must include subject alternative name if the subject field is an empty sequence. | RFC 5280: 4.2 & 4.2.1.6
subject_info_access_marked_critical | Conforming CAs must mark the Subject Info Access extension as non-critical | RFC 5280: 4.2.2.2
subject_locality_without_org | The locality field must not be included without an organization name. | CAB: 7.1.4.2.2
subject_org_without_country | The organization name field must not be included without a country name. | CAB: 7.1.4.2.2
subject_org_without_locality_or_province | If organiation is included, either stateOrProvince or locality must be included. | CAB: 7.1.4.2.2 (d&e)
subject_postal_without_org | The postal code must not be included without an organization name. | CAB: 7.1.4.2.2
subject_province_without_org | The stateOrProvince name must not be included without an organization name. | CAB: 7.1.4.2.2
subject_street_without_org | The street address field must not be included without an organization name. | CAB: 7.1.4.2.2
UTC_time_does_not_include_seconds | UTCTime values must include seconds | RFC 5280: 4.1.2.5.1
UTC_time_not_in_zulu | UTCTime values must be expressed in Greenwich Mean Time (Zulu) | RFC 5280: 4.1.2.5.1
validity_time_not_positive | Certificates MUST have a positive time for which they are valid | -
wrong_time_format_pre2050 | Certificates with validity through the year 2049 must be encoded in UTC time | RFC 5280: 4.1.2.5

The following is a list of lints we plan to implement in the future.

**Lint Name** | **Description** | **Providence**
------------------- | :-------------------------------------------------------------------: | ---:
br_contains_ulabel | Warning thrown if commonNames in BR certificate contains U-labels | -
br_san_wildcard_unknown | Wildcard other than \*.<fqdn> in SAN | -
ca_valid_time_too_long | CA certificates SHOULD NOT have a validity of greater than 25 years | -
dir_name_empty | Simply checks if the DirectoryName field is empty. No further tests are performed | -
distng_name_bmp_string | SHOULD NOT use deprecated BMPString | -
distng_name_deprecated_attr | Names SHOULD NOT have deprecated attribute (includes attribute normally) | -
distng_name_dup_attr | Names SHOULD NOT have multiple (attribute) attributes | -
distng_name_general_string | SHOULD NOT use deprecated GeneralString | -
distng_name_graphic_string | SHOULD NOT use deprecated GraphicString | -
distng_name_invalid_label | Tests to make sure the name contains a valid label | -
distng_name_leading_whitespace | SHOULD NOT have leading whitespace in attribute | -
distng_name_multiple_attr_in_rdn | SHOULD NOT have multiple attributes in a single RDN in the subject Name | -
distng_name_not_ia5 | SHOULD be encoded as IA5String | -
distng_name_not_utf8 | SHOULD be encoded as UF8String | -
distng_name_teletex_string | SHOULD NOT be using deprecated TeletexString | -
distng_name_too_long | Verifies the attribute is not too long => how long is too long? | -
distng_name_trailing_whitespace | SHOULD NOT have trailing whitespace in attribute | -
distng_name_universal_string | SHOULD NOT be using deprecated UniversalString | -
distng_name_unknown_attr | Name SHOULD NOT has unknown attribute (includes attribute normally) | -
distng_name_unparsable | Error thrown if unparsable name | -
distng_name_videoex_string | SHOULD NOT be using deprecated VideoexString | -
dns_name_fqdn_too_long | FQDN in DNSName must be less than 253 (bits?) | -
dns_name_underscore_present | DNS names SHOULD NOT have an underscore | -
dns_name_wildcard_present | Wildcards must not be FQDN of DNSName => how this is this different from line 5 | -
dsa_param_not_empty | “DSA signatures MUST NOT have a parameter specified” | -
ec_infinite | EC Public key is not infinity | -
ec_no_parameters_included | “EC keys must have parameters” (no parameters were included) | -
ec_off_curve | EC Public key MUST be on curve | -
edi_part_name_missing | Simply checks if the EDIPartyName field is present. No further tests are performed. | -
eku_extra_usage | Extended Key Usage should pertain only to TLS Web Server Authentication, TLS Web Client Authentication, and E-mail Protection. Everything else is a warning | -
encode_not_der | The certificate SHOULD be encoded in the DER format. | -
ext_explicit_false | critical:FALSE explicitly encoded (X.690) | -
ext_san_placed_first | Some python versions will not see SAN extension if it is the first extension | -
ext_unknown | Any unknown extension causes an error => what's unknown? is there a list?? | -
ext_unknown_crit | Any unknown or opaque extension that is marked as critical causes an error | -
ext_unsupported | Any extension from Netscape, Microsoft, or just unsupported is marked as opaque | -
iana_base_domain_underscore | SHOULD NOT contain underscore in base domain | -
iana_domain_not_qualified | MUST have qualified domain name | -
iana_special | SHOULD NOT be special name => what's a special name? | -
iana_tld_type_unknown | MUST have known type of TLD | -
iana_tld_unknown | MUST have known tld from list | -
pre_br_valid_too_long | Pre-BR certificates should not be more than 120 months in validity | -
registered_id_empty | Simply checks if the RegisteredID field is filled. No further tests are performed. | -
rfc822_name_empty_value | Checks to verify the validity of a RFC822 name field. Checking for empty value | -
rfc822_name_invalid_padding | Checks to verify the validity of a RFC822 name field. Checking for invalid padding | -
rfc822_name_null_char | Checks to verify the validity of a RFC822 name field. Checking for null character | -
rsa_mod_is_negative | “RSA public key modulus must be positive” (negative modulus found) | -
rsa_no_null_param | RSA keys must have a null parameter, this was missing | -
rsa_no_param | “RSA keys must have a parameter specified” (no rsa parameter) | -
signature_algorithm_pss_not_supported | Checks if the signature algorithm is RSA Probabilistic Signature Scheme. This is not supported by most browsers | -
signature_algorithm_weak | Compares chosen signature algorithm to list containing algs and their corresponding strength. A ‘weak’ alg should have at least 64 bits on entropy. A ‘good’ alg should have at least 20 bits | -
string_contains_null | If a null byte was found in a String (UTF8, Teletex, Videotex, Graphic, General) in a certificate element this error is thrown. UTF8 currently returns a separate error for this same error, condensed here. | -
string_control_char_found | A control character found in a String (IA5, Visible) in a certificate element will throw this error | -
string_teletext_encoded_incorrectly | Teletex strings found in a certificate element should be correctly encoded. | -
string_utf8_encoded_incorrectly | If an incorrectly encoded UTF8String is found, fatal error. This means that an element in the certificate failed: value.force_encoding('UTF-8').valid_encoding | -
uid_client_not_empty | clientUniqueID must not be included | -
uid_issuer_not_empty | issuerUniqueID must not be included | -
uri_field_empty | Simply checks if the URI field is present. No further tests are performed. | -
x400_name_empty | Fails if a x400 address field is empty. No further tests are performed. | -
