/*
 * ZLint Copyright 2024 Regents of the University of Michigan
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

// Package cabf_bf implements the defintions from the 2.0.4 version of the
// baseline requirements as described at
// https://cabforum.org/working-groups/server/baseline-requirements/requirements/.
package cabf_bf

// 1.6.1 Definitions
//
// Affiliate: A corporation, partnership, joint venture or other entity
// controlling, controlled by, or under common control with another entity, or
// an agency, department, political subdivision, or any entity operating under
// the direct control of a Government Entity.
//

// Applicant: The natural person or Legal Entity that applies for (or seeks
// renewal of) a Certificate. Once the Certificate is issued, the Applicant is
// referred to as the Subscriber. For Certificates issued to devices, the
// Applicant is the entity that controls or operates the device named in the
// Certificate, even if the device is sending the actual certificate request.
//

// Applicant Representative: A natural person or human sponsor who is either the
// Applicant, employed by the Applicant, or an authorized agent who has express
// authority to represent the Applicant: who signs and submits, or approves a
// certificate request on behalf of the Applicant, and/or who signs and submits
// a Subscriber Agreement on behalf of the Applicant, and/or who acknowledges
// the Terms of Use on behalf of the Applicant when the Applicant is an
// Affiliate of the CA or is the CA.

// Application Software Supplier: A supplier of Internet browser software or
// other relying-party application software that displays or uses Certificates
// and incorporates Root Certificates.
//

// Attestation Letter: A letter attesting that Subject Information is correct
// written by an accountant, lawyer, government official, or other reliable
// third party customarily relied upon for such information.
//

// Audit Period: In a period-of-time audit, the period between the first day
// (start) and the last day of operations (end) covered by the auditors in their
// engagement. (This is not the same as the period of time when the auditors are
// on-site at the CA.) The coverage rules and maximum length of audit periods
// are defined in Section 8.1.
//

// Audit Report: A report from a Qualified Auditor stating the Qualified
// Auditor’s opinion on whether an entity’s processes and controls comply with
// the mandatory provisions of these Requirements.
//

// Authorization Domain Name: The FQDN used to obtain authorization for a given
// FQDN to be included in a Certificate. The CA may use the FQDN returned from a
// DNS CNAME lookup as the FQDN for the purposes of domain validation. If a
// Wildcard Domain Name is to be included in a Certificate, then the CA MUST
// remove “*.” from the left-most portion of the Wildcard Domain Name to yield
// the corresponding FQDN. The CA may prune zero or more Domain Labels of the
// FQDN from left to right until encountering a Base Domain Name and may use any
// one of the values that were yielded by pruning (including the Base Domain
// Name itself) for the purpose of domain validation.

// Authorized Ports: One of the following ports: 80 (http), 443 (https), 25
// (smtp), 22 (ssh).
//

// Base Domain Name: The portion of an applied-for FQDN that is the first Domain
// Name node left of a registry-controlled or public suffix plus the
// registry-controlled or public suffix (e.g. “example.co.uk” or “example.com”).
// For FQDNs where the right-most Domain Name node is a gTLD having ICANN
// Specification 13 in its registry agreement, the gTLD itself may be used as
// the Base Domain Name.
//

// CAA: From RFC 8659 (http://tools.ietf.org/html/rfc8659): “The Certification
// Authority Authorization (CAA) DNS Resource Record allows a DNS domain name
// holder to specify one or more Certification Authorities (CAs) authorized to
// issue certificates for that domain name. CAA Resource Records allow a public
// CA to implement additional controls to reduce the risk of unintended
// certificate mis-issue.”
//

// CA Key Pair: A Key Pair where the Public Key appears as the Subject Public
// Key Info in one or more Root CA Certificate(s) and/or Subordinate CA
// Certificate(s).
//

// Certificate: An electronic document that uses a digital signature to bind a
// public key and an identity.
//

// Certificate Data: Certificate requests and data related thereto (whether
// obtained from the Applicant or otherwise) in the CA’s possession or control
// or to which the CA has access.
//

// Certificate Management Process: Processes, practices, and procedures
// associated with the use of keys, software, and hardware, by which the CA
// verifies Certificate Data, issues Certificates, maintains a Repository, and
// revokes Certificates.
//

// Certificate Policy: A set of rules that indicates the applicability of a
// named Certificate to a particular community and/or PKI implementation with
// common security requirements.
//

// Certificate Problem Report: Complaint of suspected Key Compromise,
// Certificate misuse, or other types of fraud, compromise, misuse, or
// inappropriate conduct related to Certificates.
//

// Certificate Profile: A set of documents or files that defines requirements
// for Certificate content and Certificate extensions in accordance with Section
// 7, e.g. a Section in a CA’s CPS or a certificate template file used by CA
// software.
//

// Certificate Revocation List: A regularly updated time-stamped list of revoked
// Certificates that is created and digitally signed by the CA that issued the
// Certificates.
//

// Certification Authority: An organization that is responsible for the
// creation, issuance, revocation, and management of Certificates. The term
// applies equally to both Root CAs and Subordinate CAs.
//

// Certification Practice Statement: One of several documents forming the
// governance framework in which Certificates are created, issued, managed, and
// used.
//

// Control: “Control” (and its correlative meanings, “controlled by” and “under
// common control with”) means possession, directly or indirectly, of the power
// to: (1) direct the management, personnel, finances, or plans of such entity;
// (2) control the election of a majority of the directors ; or (3) vote that
// portion of voting shares required for “control” under the law of the entity’s
// Jurisdiction of Incorporation or Registration but in no case less than 10%.
//

// Country: Either a member of the United Nations OR a geographic region
// recognized as a Sovereign State by at least two UN member nations.
//

// Cross-Certified Subordinate CA Certificate: A certificate that is used to
// establish a trust relationship between two CAs.
//

// CSPRNG: A random number generator intended for use in a cryptographic system.
//

// Delegated Third Party: A natural person or Legal Entity that is not the CA
// but is authorized by the CA, and whose activities are not within the scope of
// the appropriate CA audits, to assist in the Certificate Management Process by
// performing or fulfilling one or more of the CA requirements found herein.
//

// DNS CAA Email Contact: The email address defined in Appendix A.1.1.
//

// DNS CAA Phone Contact: The phone number defined in Appendix A.1.2.
//

// DNS TXT Record Email Contact: The email address defined in Appendix A.2.1.
//

// DNS TXT Record Phone Contact: The phone number defined in Appendix A.2.2.
//

// Domain Contact: The Domain Name Registrant, technical contact, or
// administrative contact (or the equivalent under a ccTLD) as listed in the
// WHOIS record of the Base Domain Name or in a DNS SOA record, or as obtained
// through direct contact with the Domain Name Registrar.
//

// Domain Label: From RFC 8499 (http://tools.ietf.org/html/rfc8499): “An ordered
// list of zero or more octets that makes up a portion of a domain name. Using
// graph theory, a label identifies one node in a portion of the graph of all
// possible domain names.”
//
// TODO

// Domain Name: An ordered list of one or more Domain Labels assigned to a node
// in the Domain Name System.
//
// TODO

// Domain Namespace: The set of all possible Domain Names that are subordinate
// to a single node in the Domain Name System.
//

// Domain Name Registrant: Sometimes referred to as the “owner” of a Domain
// Name, but more properly the person(s) or entity(ies) registered with a Domain
// Name Registrar as having the right to control how a Domain Name is used, such
// as the natural person or Legal Entity that is listed as the “Registrant” by
// WHOIS or the Domain Name Registrar.
//

// Domain Name Registrar: A person or entity that registers Domain Names under
// the auspices of or by agreement with: the Internet Corporation for Assigned
// Names and Numbers (ICANN), a national Domain Name authority/registry, or a
// Network Information Center (including their affiliates, contractors,
// delegates, successors, or assignees).  Enterprise RA: An employee or agent of
// an organization unaffiliated with the CA who authorizes issuance of
// Certificates to that organization.
//

// Expiry Date: The “Not After” date in a Certificate that defines the end of a
// Certificate’s validity period.
//

// Fully-Qualified Domain Name: A Domain Name that includes the Domain Labels of
// all superior nodes in the Internet Domain Name System.
//

// Government Entity: A government-operated legal entity, agency, department, ministry, branch, or similar element of the government of a country, or political subdivision within such country (such as a state, province, city, county, etc.).
//
// High Risk Certificate Request: A Request that the CA flags for additional scrutiny by reference to internal criteria and databases maintained by the CA, which may include names at higher risk for phishing or other fraudulent usage, names contained in previously rejected certificate requests or revoked Certificates, names listed on the Miller Smiles phishing list or the Google Safe Browsing list, or names that the CA identifies using its own risk-mitigation criteria.
//
// Internal Name: A string of characters (not an IP address) in a Common Name or Subject Alternative Name field of a Certificate that cannot be verified as globally unique within the public DNS at the time of certificate issuance because it does not end with a Top Level Domain registered in IANA’s Root Zone Database.
//
// IP Address: A 32-bit or 128-bit number assigned to a device that uses the Internet Protocol for communication.
//
// IP Address Contact: The person(s) or entity(ies) registered with an IP Address Registration Authority as having the right to control how one or more IP Addresses are used.
//
// IP Address Registration Authority: The Internet Assigned Numbers Authority (IANA) or a Regional Internet Registry (RIPE, APNIC, ARIN, AfriNIC, LACNIC).
//
// Issuing CA: In relation to a particular Certificate, the CA that issued the Certificate. This could be either a Root CA or a Subordinate CA.
//
// Key Compromise: A Private Key is said to be compromised if its value has been disclosed to an unauthorized person, or an unauthorized person has had access to it.
//
// Key Generation Script: A documented plan of procedures for the generation of a CA Key Pair.
//
// Key Pair: The Private Key and its associated Public Key.
//
// LDH Label: From RFC 5890 (http://tools.ietf.org/html/rfc5890): “A string consisting of ASCII letters, digits, and the hyphen with the further restriction that the hyphen cannot appear at the beginning or end of the string. Like all DNS labels, its total length must not exceed 63 octets.”
//
// Legal Entity: An association, corporation, partnership, proprietorship, trust, government entity or other entity with legal standing in a country’s legal system.
//
// Non-Reserved LDH Label: From RFC 5890 (http://tools.ietf.org/html/rfc5890): “The set of valid LDH labels that do not have ‘--’ in the third and fourth positions.”
//
// Object Identifier: A unique alphanumeric or numeric identifier registered under the International Organization for Standardization’s applicable standard for a specific object or object class.
//
// OCSP Responder: An online server operated under the authority of the CA and connected to its Repository for processing Certificate status requests. See also, Online Certificate Status Protocol.
//
// Onion Domain Name: A Fully Qualified Domain Name ending with the RFC 7686 “.onion” Special-Use Domain Name. For example, 2gzyxa5ihm7nsggfxnu52rck2vv4rvmdlkiu3zzui5du4xyclen53wid.onion is an Onion Domain Name, whereas torproject.org is not an Onion Domain Name.
//
// Online Certificate Status Protocol: An online Certificate-checking protocol that enables relying-party application software to determine the status of an identified Certificate. See also OCSP Responder.
//
// Parent Company: A company that Controls a Subsidiary Company.
//
// Pending Prohibition​​: The use of a behavior described with this label is highly discouraged, as it is planned to be deprecated and will likely be designated as MUST NOT in the future.
//
// Private Key: The key of a Key Pair that is kept secret by the holder of the Key Pair, and that is used to create Digital Signatures and/or to decrypt electronic records or files that were encrypted with the corresponding Public Key.
//
// Public Key: The key of a Key Pair that may be publicly disclosed by the holder of the corresponding Private Key and that is used by a Relying Party to verify Digital Signatures created with the holder’s corresponding Private Key and/or to encrypt messages so that they can be decrypted only with the holder’s corresponding Private Key.
//
// Public Key Infrastructure: A set of hardware, software, people, procedures, rules, policies, and obligations used to facilitate the trustworthy creation, issuance, management, and use of Certificates and keys based on Public Key Cryptography.
//
// Publicly-Trusted Certificate: A Certificate that is trusted by virtue of the fact that its corresponding Root Certificate is distributed as a trust anchor in widely-available application software.
//
// P-Label: A XN-Label that contains valid output of the Punycode algorithm (as defined in RFC 3492, Section 6.3) from the fifth and subsequent positions.
//
// Qualified Auditor: A natural person or Legal Entity that meets the requirements of Section 8.2.
//
// Random Value: A value specified by a CA to the Applicant that exhibits at least 112 bits of entropy.
//
// Registered Domain Name: A Domain Name that has been registered with a Domain Name Registrar.
//
// Registration Authority (RA): Any Legal Entity that is responsible for identification and authentication of subjects of Certificates, but is not a CA, and hence does not sign or issue Certificates. An RA may assist in the certificate application process or revocation process or both. When “RA” is used as an adjective to describe a role or function, it does not necessarily imply a separate body, but can be part of the CA.
//
// Reliable Data Source: An identification document or source of data used to verify Subject Identity Information that is generally recognized among commercial enterprises and governments as reliable, and which was created by a third party for a purpose other than the Applicant obtaining a Certificate.
//
// Reliable Method of Communication: A method of communication, such as a postal/courier delivery address, telephone number, or email address, that was verified using a source other than the Applicant Representative.
//
// Relying Party: Any natural person or Legal Entity that relies on a Valid Certificate. An Application Software Supplier is not considered a Relying Party when software distributed by such Supplier merely displays information relating to a Certificate.
//
// Repository: An online database containing publicly-disclosed PKI governance documents (such as Certificate Policies and Certification Practice Statements) and Certificate status information, either in the form of a CRL or an OCSP response.
//
// Request Token: A value, derived in a method specified by the CA which binds this demonstration of control to the certificate request. The CA SHOULD define within its CPS (or a document clearly referenced by the CPS) the format and method of Request Tokens it accepts.
//
// The Request Token SHALL incorporate the key used in the certificate request.
//
// A Request Token MAY include a timestamp to indicate when it was created.
//
// A Request Token MAY include other information to ensure its uniqueness.
//
// A Request Token that includes a timestamp SHALL remain valid for no more than 30 days from the time of creation.
//
// A Request Token that includes a timestamp SHALL be treated as invalid if its timestamp is in the future.
//
// A Request Token that does not include a timestamp is valid for a single use and the CA SHALL NOT re-use it for a subsequent validation.
//
// The binding SHALL use a digital signature algorithm or a cryptographic hash algorithm at least as strong as that to be used in signing the certificate request.
//
// Note: Examples of Request Tokens include, but are not limited to:
//
// a hash of the public key; or
// a hash of the Subject Public Key Info [X.509]; or
// a hash of a PKCS#10 CSR.
// A Request Token may also be concatenated with a timestamp or other data. If a CA wanted to always use a hash of a PKCS#10 CSR as a Request Token and did not want to incorporate a timestamp and did want to allow certificate key re-use then the applicant might use the challenge password in the creation of a CSR with OpenSSL to ensure uniqueness even if the subject and key are identical between subsequent requests.
//
// Note: This simplistic shell command produces a Request Token which has a timestamp and a hash of a CSR. echo `date -u +%Y%m%d%H%M` `sha256sum <r2.csr` \| sed "s/[ -]//g" The script outputs: 201602251811c9c863405fe7675a3988b97664ea6baf442019e4e52fa335f406f7c5f26cf14f
//
// Required Website Content: Either a Random Value or a Request Token, together with additional information that uniquely identifies the Subscriber, as specified by the CA.
//
// Requirements: The Baseline Requirements found in this document.
//
// Reserved IP Address: An IPv4 or IPv6 address that is contained in the address block of any entry in either of the following IANA registries:
//
// https://www.iana.org/assignments/iana-ipv4-special-registry/iana-ipv4-special-registry.xhtml
//
// https://www.iana.org/assignments/iana-ipv6-special-registry/iana-ipv6-special-registry.xhtml
//
// Root CA: The top level Certification Authority whose Root Certificate is distributed by Application Software Suppliers and that issues Subordinate CA Certificates.
//
// Root Certificate: The self-signed Certificate issued by the Root CA to identify itself and to facilitate verification of Certificates issued to its Subordinate CAs.
//
// Short-lived Subscriber Certificate: For Certificates issued on or after 15 March 2024 and prior to 15 March 2026, a Subscriber Certificate with a Validity Period less than or equal to 10 days (864,000 seconds). For Certificates issued on or after 15 March 2026, a Subscriber Certificate with a Validity Period less than or equal to 7 days (604,800 seconds).
//
// Sovereign State: A state or country that administers its own government, and is not dependent upon, or subject to, another power.
//
// Subject: The natural person, device, system, unit, or Legal Entity identified in a Certificate as the Subject. The Subject is either the Subscriber or a device under the control and operation of the Subscriber.
//
// Subject Identity Information: Information that identifies the Certificate Subject. Subject Identity Information does not include a Domain Name listed in the subjectAltName extension or the Subject commonName field.
//
// Subordinate CA: A Certification Authority whose Certificate is signed by the Root CA, or another Subordinate CA.
//
// Subscriber: A natural person or Legal Entity to whom a Certificate is issued and who is legally bound by a Subscriber Agreement or Terms of Use.
//
// Subscriber Agreement: An agreement between the CA and the Applicant/Subscriber that specifies the rights and responsibilities of the parties.
//
// Subsidiary Company: A company that is controlled by a Parent Company.
//
// Technically Constrained Subordinate CA Certificate: A Subordinate CA certificate which uses a combination of Extended Key Usage and/or Name Constraint extensions, as defined within the relevant Certificate Profiles of this document, to limit the scope within which the Subordinate CA Certificate may issue Subscriber or additional Subordinate CA Certificates.
//
// Terms of Use: Provisions regarding the safekeeping and acceptable uses of a Certificate issued in accordance with these Requirements when the Applicant/Subscriber is an Affiliate of the CA or is the CA.
//
// Test Certificate: This term is no longer used in these Baseline Requirements.
//
// Trustworthy System: Computer hardware, software, and procedures that are: reasonably secure from intrusion and misuse; provide a reasonable level of availability, reliability, and correct operation; are reasonably suited to performing their intended functions; and enforce the applicable security policy.
//
// Unregistered Domain Name: A Domain Name that is not a Registered Domain Name.
//
// Valid Certificate: A Certificate that passes the validation procedure specified in RFC 5280.
//
// Validation Specialist: Someone who performs the information verification duties specified by these Requirements.
//
// Validity Period: From RFC 5280 (http://tools.ietf.org/html/rfc5280): “The period of time from notBefore through notAfter, inclusive.”
//
// WHOIS: Information retrieved directly from the Domain Name Registrar or registry operator via the protocol defined in RFC 3912, the Registry Data Access Protocol defined in RFC 7482, or an HTTPS website.
//
// Wildcard Certificate: A Certificate containing at least one Wildcard Domain Name in the Subject Alternative Names in the Certificate.
//
// Wildcard Domain Name: A string starting with “*.” (U+002A ASTERISK, U+002E FULL STOP) immediately followed by a Fully-Qualified Domain Name.
//
// XN-Label: From RFC 5890 (http://tools.ietf.org/html/rfc5890): “The class of labels that begin with the prefix "xn--" (case independent), but otherwise conform to the rules for LDH labels.”
