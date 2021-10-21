/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "PKIX1Implicit88"
 * 	found in "asn1/rfc5280-PKIX1Implicit88.asn1"
 * 	`asn1c -S /home/fotisl/Projects/revlintsynt/asn1c/skeletons -pdu=all -fwide-types -fcompound-names`
 */

#include "UserNotice.h"

static asn_TYPE_member_t asn_MBR_UserNotice_1[] = {
	{ ATF_POINTER, 2, offsetof(struct UserNotice, noticeRef),
		(ASN_TAG_CLASS_UNIVERSAL | (16 << 2)),
		0,
		&asn_DEF_NoticeReference,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"noticeRef"
		},
	{ ATF_POINTER, 1, offsetof(struct UserNotice, explicitText),
		-1 /* Ambiguous tag (CHOICE?) */,
		0,
		&asn_DEF_DisplayText,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"explicitText"
		},
};
static const int asn_MAP_UserNotice_oms_1[] = { 0, 1 };
static const ber_tlv_tag_t asn_DEF_UserNotice_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static const asn_TYPE_tag2member_t asn_MAP_UserNotice_tag2el_1[] = {
    { (ASN_TAG_CLASS_UNIVERSAL | (12 << 2)), 1, 0, 0 }, /* utf8String */
    { (ASN_TAG_CLASS_UNIVERSAL | (16 << 2)), 0, 0, 0 }, /* noticeRef */
    { (ASN_TAG_CLASS_UNIVERSAL | (22 << 2)), 1, 0, 0 }, /* ia5String */
    { (ASN_TAG_CLASS_UNIVERSAL | (26 << 2)), 1, 0, 0 }, /* visibleString */
    { (ASN_TAG_CLASS_UNIVERSAL | (30 << 2)), 1, 0, 0 } /* bmpString */
};
static asn_SEQUENCE_specifics_t asn_SPC_UserNotice_specs_1 = {
	sizeof(struct UserNotice),
	offsetof(struct UserNotice, _asn_ctx),
	asn_MAP_UserNotice_tag2el_1,
	5,	/* Count of tags in the map */
	asn_MAP_UserNotice_oms_1,	/* Optional members */
	2, 0,	/* Root/Additions */
	-1,	/* First extension addition */
};
asn_TYPE_descriptor_t asn_DEF_UserNotice = {
	"UserNotice",
	"UserNotice",
	&asn_OP_SEQUENCE,
	asn_DEF_UserNotice_tags_1,
	sizeof(asn_DEF_UserNotice_tags_1)
		/sizeof(asn_DEF_UserNotice_tags_1[0]), /* 1 */
	asn_DEF_UserNotice_tags_1,	/* Same as above */
	sizeof(asn_DEF_UserNotice_tags_1)
		/sizeof(asn_DEF_UserNotice_tags_1[0]), /* 1 */
	{ 0, 0, SEQUENCE_constraint },
	asn_MBR_UserNotice_1,
	2,	/* Elements count */
	&asn_SPC_UserNotice_specs_1	/* Additional specs */
};

