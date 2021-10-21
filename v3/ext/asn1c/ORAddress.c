/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "PKIX1Explicit88"
 * 	found in "asn1/rfc5280-PKIX1Explicit88.asn1"
 * 	`asn1c -S /home/fotisl/Projects/revlintsynt/asn1c/skeletons -pdu=all -fwide-types -fcompound-names`
 */

#include "ORAddress.h"

asn_TYPE_member_t asn_MBR_ORAddress_1[] = {
	{ ATF_NOFLAGS, 0, offsetof(struct ORAddress, built_in_standard_attributes),
		(ASN_TAG_CLASS_UNIVERSAL | (16 << 2)),
		0,
		&asn_DEF_BuiltInStandardAttributes,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"built-in-standard-attributes"
		},
	{ ATF_POINTER, 2, offsetof(struct ORAddress, built_in_domain_defined_attributes),
		(ASN_TAG_CLASS_UNIVERSAL | (16 << 2)),
		0,
		&asn_DEF_BuiltInDomainDefinedAttributes,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"built-in-domain-defined-attributes"
		},
	{ ATF_POINTER, 1, offsetof(struct ORAddress, extension_attributes),
		(ASN_TAG_CLASS_UNIVERSAL | (17 << 2)),
		0,
		&asn_DEF_ExtensionAttributes,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"extension-attributes"
		},
};
static const int asn_MAP_ORAddress_oms_1[] = { 1, 2 };
static const ber_tlv_tag_t asn_DEF_ORAddress_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static const asn_TYPE_tag2member_t asn_MAP_ORAddress_tag2el_1[] = {
    { (ASN_TAG_CLASS_UNIVERSAL | (16 << 2)), 0, 0, 1 }, /* built-in-standard-attributes */
    { (ASN_TAG_CLASS_UNIVERSAL | (16 << 2)), 1, -1, 0 }, /* built-in-domain-defined-attributes */
    { (ASN_TAG_CLASS_UNIVERSAL | (17 << 2)), 2, 0, 0 } /* extension-attributes */
};
asn_SEQUENCE_specifics_t asn_SPC_ORAddress_specs_1 = {
	sizeof(struct ORAddress),
	offsetof(struct ORAddress, _asn_ctx),
	asn_MAP_ORAddress_tag2el_1,
	3,	/* Count of tags in the map */
	asn_MAP_ORAddress_oms_1,	/* Optional members */
	2, 0,	/* Root/Additions */
	-1,	/* First extension addition */
};
asn_TYPE_descriptor_t asn_DEF_ORAddress = {
	"ORAddress",
	"ORAddress",
	&asn_OP_SEQUENCE,
	asn_DEF_ORAddress_tags_1,
	sizeof(asn_DEF_ORAddress_tags_1)
		/sizeof(asn_DEF_ORAddress_tags_1[0]), /* 1 */
	asn_DEF_ORAddress_tags_1,	/* Same as above */
	sizeof(asn_DEF_ORAddress_tags_1)
		/sizeof(asn_DEF_ORAddress_tags_1[0]), /* 1 */
	{ 0, 0, SEQUENCE_constraint },
	asn_MBR_ORAddress_1,
	3,	/* Elements count */
	&asn_SPC_ORAddress_specs_1	/* Additional specs */
};

