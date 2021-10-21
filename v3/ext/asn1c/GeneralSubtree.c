/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "PKIX1Implicit88"
 * 	found in "asn1/rfc5280-PKIX1Implicit88.asn1"
 * 	`asn1c -S /home/fotisl/Projects/revlintsynt/asn1c/skeletons -pdu=all -fwide-types -fcompound-names`
 */

#include "GeneralSubtree.h"

static int asn_DFL_3_cmp_0(const void *sptr) {
	const BaseDistance_t *st = sptr;
	
	if(!st) {
		return -1; /* No value is not a default value */
	}
	
	/* Test default value 0 */
	long value;
	if(asn_INTEGER2long(st, &value))
		return -1;
	return (value != 0);
}
static int asn_DFL_3_set_0(void **sptr) {
	BaseDistance_t *st = *sptr;
	
	if(!st) {
		st = (*sptr = CALLOC(1, sizeof(*st)));
		if(!st) return -1;
	}
	
	/* Install default value 0 */
	return asn_long2INTEGER(st, 0);
}
asn_TYPE_member_t asn_MBR_GeneralSubtree_1[] = {
	{ ATF_NOFLAGS, 0, offsetof(struct GeneralSubtree, base),
		-1 /* Ambiguous tag (CHOICE?) */,
		0,
		&asn_DEF_GeneralName,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"base"
		},
	{ ATF_POINTER, 2, offsetof(struct GeneralSubtree, minimum),
		(ASN_TAG_CLASS_CONTEXT | (0 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_BaseDistance,
		0,
		{ 0, 0, 0 },
		&asn_DFL_3_cmp_0,	/* Compare DEFAULT 0 */
		&asn_DFL_3_set_0,	/* Set DEFAULT 0 */
		"minimum"
		},
	{ ATF_POINTER, 1, offsetof(struct GeneralSubtree, maximum),
		(ASN_TAG_CLASS_CONTEXT | (1 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_BaseDistance,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"maximum"
		},
};
static const int asn_MAP_GeneralSubtree_oms_1[] = { 1, 2 };
static const ber_tlv_tag_t asn_DEF_GeneralSubtree_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static const asn_TYPE_tag2member_t asn_MAP_GeneralSubtree_tag2el_1[] = {
    { (ASN_TAG_CLASS_CONTEXT | (0 << 2)), 0, 0, 1 }, /* otherName */
    { (ASN_TAG_CLASS_CONTEXT | (0 << 2)), 1, -1, 0 }, /* minimum */
    { (ASN_TAG_CLASS_CONTEXT | (1 << 2)), 0, 0, 1 }, /* rfc822Name */
    { (ASN_TAG_CLASS_CONTEXT | (1 << 2)), 2, -1, 0 }, /* maximum */
    { (ASN_TAG_CLASS_CONTEXT | (2 << 2)), 0, 0, 0 }, /* dNSName */
    { (ASN_TAG_CLASS_CONTEXT | (3 << 2)), 0, 0, 0 }, /* x400Address */
    { (ASN_TAG_CLASS_CONTEXT | (4 << 2)), 0, 0, 0 }, /* directoryName */
    { (ASN_TAG_CLASS_CONTEXT | (5 << 2)), 0, 0, 0 }, /* ediPartyName */
    { (ASN_TAG_CLASS_CONTEXT | (6 << 2)), 0, 0, 0 }, /* uniformResourceIdentifier */
    { (ASN_TAG_CLASS_CONTEXT | (7 << 2)), 0, 0, 0 }, /* iPAddress */
    { (ASN_TAG_CLASS_CONTEXT | (8 << 2)), 0, 0, 0 } /* registeredID */
};
asn_SEQUENCE_specifics_t asn_SPC_GeneralSubtree_specs_1 = {
	sizeof(struct GeneralSubtree),
	offsetof(struct GeneralSubtree, _asn_ctx),
	asn_MAP_GeneralSubtree_tag2el_1,
	11,	/* Count of tags in the map */
	asn_MAP_GeneralSubtree_oms_1,	/* Optional members */
	2, 0,	/* Root/Additions */
	-1,	/* First extension addition */
};
asn_TYPE_descriptor_t asn_DEF_GeneralSubtree = {
	"GeneralSubtree",
	"GeneralSubtree",
	&asn_OP_SEQUENCE,
	asn_DEF_GeneralSubtree_tags_1,
	sizeof(asn_DEF_GeneralSubtree_tags_1)
		/sizeof(asn_DEF_GeneralSubtree_tags_1[0]), /* 1 */
	asn_DEF_GeneralSubtree_tags_1,	/* Same as above */
	sizeof(asn_DEF_GeneralSubtree_tags_1)
		/sizeof(asn_DEF_GeneralSubtree_tags_1[0]), /* 1 */
	{ 0, 0, SEQUENCE_constraint },
	asn_MBR_GeneralSubtree_1,
	3,	/* Elements count */
	&asn_SPC_GeneralSubtree_specs_1	/* Additional specs */
};

