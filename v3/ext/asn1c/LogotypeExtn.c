/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "LogotypeCertExtn"
 * 	found in "asn1/rfc3709-LogotypeCertExtn.asn1"
 * 	`asn1c -S /home/fotisl/Projects/revlintsynt/asn1c/skeletons -pdu=all -fwide-types -fcompound-names`
 */

#include "LogotypeExtn.h"

static asn_TYPE_member_t asn_MBR_communityLogos_2[] = {
	{ ATF_POINTER, 0, 0,
		-1 /* Ambiguous tag (CHOICE?) */,
		0,
		&asn_DEF_LogotypeInfo,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		""
		},
};
static const ber_tlv_tag_t asn_DEF_communityLogos_tags_2[] = {
	(ASN_TAG_CLASS_CONTEXT | (0 << 2)),
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static asn_SET_OF_specifics_t asn_SPC_communityLogos_specs_2 = {
	sizeof(struct LogotypeExtn__communityLogos),
	offsetof(struct LogotypeExtn__communityLogos, _asn_ctx),
	2,	/* XER encoding is XMLValueList */
};
static /* Use -fall-defs-global to expose */
asn_TYPE_descriptor_t asn_DEF_communityLogos_2 = {
	"communityLogos",
	"communityLogos",
	&asn_OP_SEQUENCE_OF,
	asn_DEF_communityLogos_tags_2,
	sizeof(asn_DEF_communityLogos_tags_2)
		/sizeof(asn_DEF_communityLogos_tags_2[0]), /* 2 */
	asn_DEF_communityLogos_tags_2,	/* Same as above */
	sizeof(asn_DEF_communityLogos_tags_2)
		/sizeof(asn_DEF_communityLogos_tags_2[0]), /* 2 */
	{ 0, 0, SEQUENCE_OF_constraint },
	asn_MBR_communityLogos_2,
	1,	/* Single element */
	&asn_SPC_communityLogos_specs_2	/* Additional specs */
};

static asn_TYPE_member_t asn_MBR_otherLogos_6[] = {
	{ ATF_POINTER, 0, 0,
		(ASN_TAG_CLASS_UNIVERSAL | (16 << 2)),
		0,
		&asn_DEF_OtherLogotypeInfo,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		""
		},
};
static const ber_tlv_tag_t asn_DEF_otherLogos_tags_6[] = {
	(ASN_TAG_CLASS_CONTEXT | (3 << 2)),
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static asn_SET_OF_specifics_t asn_SPC_otherLogos_specs_6 = {
	sizeof(struct LogotypeExtn__otherLogos),
	offsetof(struct LogotypeExtn__otherLogos, _asn_ctx),
	0,	/* XER encoding is XMLDelimitedItemList */
};
static /* Use -fall-defs-global to expose */
asn_TYPE_descriptor_t asn_DEF_otherLogos_6 = {
	"otherLogos",
	"otherLogos",
	&asn_OP_SEQUENCE_OF,
	asn_DEF_otherLogos_tags_6,
	sizeof(asn_DEF_otherLogos_tags_6)
		/sizeof(asn_DEF_otherLogos_tags_6[0]), /* 2 */
	asn_DEF_otherLogos_tags_6,	/* Same as above */
	sizeof(asn_DEF_otherLogos_tags_6)
		/sizeof(asn_DEF_otherLogos_tags_6[0]), /* 2 */
	{ 0, 0, SEQUENCE_OF_constraint },
	asn_MBR_otherLogos_6,
	1,	/* Single element */
	&asn_SPC_otherLogos_specs_6	/* Additional specs */
};

static asn_TYPE_member_t asn_MBR_LogotypeExtn_1[] = {
	{ ATF_POINTER, 4, offsetof(struct LogotypeExtn, communityLogos),
		(ASN_TAG_CLASS_CONTEXT | (0 << 2)),
		0,
		&asn_DEF_communityLogos_2,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"communityLogos"
		},
	{ ATF_POINTER, 3, offsetof(struct LogotypeExtn, issuerLogo),
		(ASN_TAG_CLASS_CONTEXT | (1 << 2)),
		+1,	/* EXPLICIT tag at current level */
		&asn_DEF_LogotypeInfo,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"issuerLogo"
		},
	{ ATF_POINTER, 2, offsetof(struct LogotypeExtn, subjectLogo),
		(ASN_TAG_CLASS_CONTEXT | (2 << 2)),
		+1,	/* EXPLICIT tag at current level */
		&asn_DEF_LogotypeInfo,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"subjectLogo"
		},
	{ ATF_POINTER, 1, offsetof(struct LogotypeExtn, otherLogos),
		(ASN_TAG_CLASS_CONTEXT | (3 << 2)),
		0,
		&asn_DEF_otherLogos_6,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"otherLogos"
		},
};
static const int asn_MAP_LogotypeExtn_oms_1[] = { 0, 1, 2, 3 };
static const ber_tlv_tag_t asn_DEF_LogotypeExtn_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static const asn_TYPE_tag2member_t asn_MAP_LogotypeExtn_tag2el_1[] = {
    { (ASN_TAG_CLASS_CONTEXT | (0 << 2)), 0, 0, 0 }, /* communityLogos */
    { (ASN_TAG_CLASS_CONTEXT | (1 << 2)), 1, 0, 0 }, /* issuerLogo */
    { (ASN_TAG_CLASS_CONTEXT | (2 << 2)), 2, 0, 0 }, /* subjectLogo */
    { (ASN_TAG_CLASS_CONTEXT | (3 << 2)), 3, 0, 0 } /* otherLogos */
};
static asn_SEQUENCE_specifics_t asn_SPC_LogotypeExtn_specs_1 = {
	sizeof(struct LogotypeExtn),
	offsetof(struct LogotypeExtn, _asn_ctx),
	asn_MAP_LogotypeExtn_tag2el_1,
	4,	/* Count of tags in the map */
	asn_MAP_LogotypeExtn_oms_1,	/* Optional members */
	4, 0,	/* Root/Additions */
	-1,	/* First extension addition */
};
asn_TYPE_descriptor_t asn_DEF_LogotypeExtn = {
	"LogotypeExtn",
	"LogotypeExtn",
	&asn_OP_SEQUENCE,
	asn_DEF_LogotypeExtn_tags_1,
	sizeof(asn_DEF_LogotypeExtn_tags_1)
		/sizeof(asn_DEF_LogotypeExtn_tags_1[0]), /* 1 */
	asn_DEF_LogotypeExtn_tags_1,	/* Same as above */
	sizeof(asn_DEF_LogotypeExtn_tags_1)
		/sizeof(asn_DEF_LogotypeExtn_tags_1[0]), /* 1 */
	{ 0, 0, SEQUENCE_constraint },
	asn_MBR_LogotypeExtn_1,
	4,	/* Elements count */
	&asn_SPC_LogotypeExtn_specs_1	/* Additional specs */
};

