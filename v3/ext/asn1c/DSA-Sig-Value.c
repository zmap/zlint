/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "PKIX1Algorithms2008"
 * 	found in "asn1/rfc5480-PKIX1Algorithms2008.asn1"
 * 	`asn1c -S /home/fotisl/Projects/revlintsynt/asn1c/skeletons -pdu=all -fwide-types -fcompound-names`
 */

#include "DSA-Sig-Value.h"

static asn_TYPE_member_t asn_MBR_DSA_Sig_Value_1[] = {
	{ ATF_NOFLAGS, 0, offsetof(struct DSA_Sig_Value, r),
		(ASN_TAG_CLASS_UNIVERSAL | (2 << 2)),
		0,
		&asn_DEF_INTEGER,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"r"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct DSA_Sig_Value, s),
		(ASN_TAG_CLASS_UNIVERSAL | (2 << 2)),
		0,
		&asn_DEF_INTEGER,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"s"
		},
};
static const ber_tlv_tag_t asn_DEF_DSA_Sig_Value_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static const asn_TYPE_tag2member_t asn_MAP_DSA_Sig_Value_tag2el_1[] = {
    { (ASN_TAG_CLASS_UNIVERSAL | (2 << 2)), 0, 0, 1 }, /* r */
    { (ASN_TAG_CLASS_UNIVERSAL | (2 << 2)), 1, -1, 0 } /* s */
};
static asn_SEQUENCE_specifics_t asn_SPC_DSA_Sig_Value_specs_1 = {
	sizeof(struct DSA_Sig_Value),
	offsetof(struct DSA_Sig_Value, _asn_ctx),
	asn_MAP_DSA_Sig_Value_tag2el_1,
	2,	/* Count of tags in the map */
	0, 0, 0,	/* Optional elements (not needed) */
	-1,	/* First extension addition */
};
asn_TYPE_descriptor_t asn_DEF_DSA_Sig_Value = {
	"DSA-Sig-Value",
	"DSA-Sig-Value",
	&asn_OP_SEQUENCE,
	asn_DEF_DSA_Sig_Value_tags_1,
	sizeof(asn_DEF_DSA_Sig_Value_tags_1)
		/sizeof(asn_DEF_DSA_Sig_Value_tags_1[0]), /* 1 */
	asn_DEF_DSA_Sig_Value_tags_1,	/* Same as above */
	sizeof(asn_DEF_DSA_Sig_Value_tags_1)
		/sizeof(asn_DEF_DSA_Sig_Value_tags_1[0]), /* 1 */
	{ 0, 0, SEQUENCE_constraint },
	asn_MBR_DSA_Sig_Value_1,
	2,	/* Elements count */
	&asn_SPC_DSA_Sig_Value_specs_1	/* Additional specs */
};

