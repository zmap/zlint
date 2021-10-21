/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "PKIXAttributeCertificate"
 * 	found in "asn1/rfc3281-PKIXAttributeCertificate.asn1"
 * 	`asn1c -S /home/fotisl/Projects/revlintsynt/asn1c/skeletons -pdu=all -fwide-types -fcompound-names`
 */

#include "SecurityCategory.h"

asn_TYPE_member_t asn_MBR_SecurityCategory_1[] = {
	{ ATF_NOFLAGS, 0, offsetof(struct SecurityCategory, type),
		(ASN_TAG_CLASS_CONTEXT | (0 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_OBJECT_IDENTIFIER,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"type"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct SecurityCategory, value),
		(ASN_TAG_CLASS_CONTEXT | (1 << 2)),
		+1,	/* EXPLICIT tag at current level */
		&asn_DEF_ANY,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"value"
		},
};
static const ber_tlv_tag_t asn_DEF_SecurityCategory_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static const asn_TYPE_tag2member_t asn_MAP_SecurityCategory_tag2el_1[] = {
    { (ASN_TAG_CLASS_CONTEXT | (0 << 2)), 0, 0, 0 }, /* type */
    { (ASN_TAG_CLASS_CONTEXT | (1 << 2)), 1, 0, 0 } /* value */
};
asn_SEQUENCE_specifics_t asn_SPC_SecurityCategory_specs_1 = {
	sizeof(struct SecurityCategory),
	offsetof(struct SecurityCategory, _asn_ctx),
	asn_MAP_SecurityCategory_tag2el_1,
	2,	/* Count of tags in the map */
	0, 0, 0,	/* Optional elements (not needed) */
	-1,	/* First extension addition */
};
asn_TYPE_descriptor_t asn_DEF_SecurityCategory = {
	"SecurityCategory",
	"SecurityCategory",
	&asn_OP_SEQUENCE,
	asn_DEF_SecurityCategory_tags_1,
	sizeof(asn_DEF_SecurityCategory_tags_1)
		/sizeof(asn_DEF_SecurityCategory_tags_1[0]), /* 1 */
	asn_DEF_SecurityCategory_tags_1,	/* Same as above */
	sizeof(asn_DEF_SecurityCategory_tags_1)
		/sizeof(asn_DEF_SecurityCategory_tags_1[0]), /* 1 */
	{ 0, 0, SEQUENCE_constraint },
	asn_MBR_SecurityCategory_1,
	2,	/* Elements count */
	&asn_SPC_SecurityCategory_specs_1	/* Additional specs */
};

