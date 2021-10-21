/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "PKIXAttributeCertificate"
 * 	found in "asn1/rfc3281-PKIXAttributeCertificate.asn1"
 * 	`asn1c -S /home/fotisl/Projects/revlintsynt/asn1c/skeletons -pdu=all -fwide-types -fcompound-names`
 */

#include "SvceAuthInfo.h"

static asn_TYPE_member_t asn_MBR_SvceAuthInfo_1[] = {
	{ ATF_NOFLAGS, 0, offsetof(struct SvceAuthInfo, service),
		-1 /* Ambiguous tag (CHOICE?) */,
		0,
		&asn_DEF_GeneralName,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"service"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct SvceAuthInfo, ident),
		-1 /* Ambiguous tag (CHOICE?) */,
		0,
		&asn_DEF_GeneralName,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"ident"
		},
	{ ATF_POINTER, 1, offsetof(struct SvceAuthInfo, authInfo),
		(ASN_TAG_CLASS_UNIVERSAL | (4 << 2)),
		0,
		&asn_DEF_OCTET_STRING,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"authInfo"
		},
};
static const int asn_MAP_SvceAuthInfo_oms_1[] = { 2 };
static const ber_tlv_tag_t asn_DEF_SvceAuthInfo_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static const asn_TYPE_tag2member_t asn_MAP_SvceAuthInfo_tag2el_1[] = {
    { (ASN_TAG_CLASS_UNIVERSAL | (4 << 2)), 2, 0, 0 }, /* authInfo */
    { (ASN_TAG_CLASS_CONTEXT | (0 << 2)), 0, 0, 1 }, /* otherName */
    { (ASN_TAG_CLASS_CONTEXT | (0 << 2)), 1, -1, 0 }, /* otherName */
    { (ASN_TAG_CLASS_CONTEXT | (1 << 2)), 0, 0, 1 }, /* rfc822Name */
    { (ASN_TAG_CLASS_CONTEXT | (1 << 2)), 1, -1, 0 }, /* rfc822Name */
    { (ASN_TAG_CLASS_CONTEXT | (2 << 2)), 0, 0, 1 }, /* dNSName */
    { (ASN_TAG_CLASS_CONTEXT | (2 << 2)), 1, -1, 0 }, /* dNSName */
    { (ASN_TAG_CLASS_CONTEXT | (3 << 2)), 0, 0, 1 }, /* x400Address */
    { (ASN_TAG_CLASS_CONTEXT | (3 << 2)), 1, -1, 0 }, /* x400Address */
    { (ASN_TAG_CLASS_CONTEXT | (4 << 2)), 0, 0, 1 }, /* directoryName */
    { (ASN_TAG_CLASS_CONTEXT | (4 << 2)), 1, -1, 0 }, /* directoryName */
    { (ASN_TAG_CLASS_CONTEXT | (5 << 2)), 0, 0, 1 }, /* ediPartyName */
    { (ASN_TAG_CLASS_CONTEXT | (5 << 2)), 1, -1, 0 }, /* ediPartyName */
    { (ASN_TAG_CLASS_CONTEXT | (6 << 2)), 0, 0, 1 }, /* uniformResourceIdentifier */
    { (ASN_TAG_CLASS_CONTEXT | (6 << 2)), 1, -1, 0 }, /* uniformResourceIdentifier */
    { (ASN_TAG_CLASS_CONTEXT | (7 << 2)), 0, 0, 1 }, /* iPAddress */
    { (ASN_TAG_CLASS_CONTEXT | (7 << 2)), 1, -1, 0 }, /* iPAddress */
    { (ASN_TAG_CLASS_CONTEXT | (8 << 2)), 0, 0, 1 }, /* registeredID */
    { (ASN_TAG_CLASS_CONTEXT | (8 << 2)), 1, -1, 0 } /* registeredID */
};
static asn_SEQUENCE_specifics_t asn_SPC_SvceAuthInfo_specs_1 = {
	sizeof(struct SvceAuthInfo),
	offsetof(struct SvceAuthInfo, _asn_ctx),
	asn_MAP_SvceAuthInfo_tag2el_1,
	19,	/* Count of tags in the map */
	asn_MAP_SvceAuthInfo_oms_1,	/* Optional members */
	1, 0,	/* Root/Additions */
	-1,	/* First extension addition */
};
asn_TYPE_descriptor_t asn_DEF_SvceAuthInfo = {
	"SvceAuthInfo",
	"SvceAuthInfo",
	&asn_OP_SEQUENCE,
	asn_DEF_SvceAuthInfo_tags_1,
	sizeof(asn_DEF_SvceAuthInfo_tags_1)
		/sizeof(asn_DEF_SvceAuthInfo_tags_1[0]), /* 1 */
	asn_DEF_SvceAuthInfo_tags_1,	/* Same as above */
	sizeof(asn_DEF_SvceAuthInfo_tags_1)
		/sizeof(asn_DEF_SvceAuthInfo_tags_1[0]), /* 1 */
	{ 0, 0, SEQUENCE_constraint },
	asn_MBR_SvceAuthInfo_1,
	3,	/* Elements count */
	&asn_SPC_SvceAuthInfo_specs_1	/* Additional specs */
};

