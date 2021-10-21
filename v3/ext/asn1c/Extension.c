/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "PKIX1Explicit88"
 * 	found in "asn1/rfc5280-PKIX1Explicit88.asn1"
 * 	`asn1c -S /home/fotisl/Projects/revlintsynt/asn1c/skeletons -pdu=all -fwide-types -fcompound-names`
 */

#include "Extension.h"

static int asn_DFL_3_cmp_0(const void *sptr) {
	const BOOLEAN_t *st = sptr;
	
	if(!st) {
		return -1; /* No value is not a default value */
	}
	
	/* Test default value 0 */
	return (*st != 0);
}
static int asn_DFL_3_set_0(void **sptr) {
	BOOLEAN_t *st = *sptr;
	
	if(!st) {
		st = (*sptr = CALLOC(1, sizeof(*st)));
		if(!st) return -1;
	}
	
	/* Install default value 0 */
	*st = 0;
	return 0;
}
asn_TYPE_member_t asn_MBR_Extension_1[] = {
	{ ATF_NOFLAGS, 0, offsetof(struct Extension, extnID),
		(ASN_TAG_CLASS_UNIVERSAL | (6 << 2)),
		0,
		&asn_DEF_OBJECT_IDENTIFIER,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"extnID"
		},
	{ ATF_NOFLAGS, 1, offsetof(struct Extension, critical),
		(ASN_TAG_CLASS_UNIVERSAL | (1 << 2)),
		0,
		&asn_DEF_BOOLEAN,
		0,
		{ 0, 0, 0 },
		&asn_DFL_3_cmp_0,	/* Compare DEFAULT 0 */
		&asn_DFL_3_set_0,	/* Set DEFAULT 0 */
		"critical"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct Extension, extnValue),
		(ASN_TAG_CLASS_UNIVERSAL | (4 << 2)),
		0,
		&asn_DEF_OCTET_STRING,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"extnValue"
		},
};
static const int asn_MAP_Extension_oms_1[] = { 1 };
static const ber_tlv_tag_t asn_DEF_Extension_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static const asn_TYPE_tag2member_t asn_MAP_Extension_tag2el_1[] = {
    { (ASN_TAG_CLASS_UNIVERSAL | (1 << 2)), 1, 0, 0 }, /* critical */
    { (ASN_TAG_CLASS_UNIVERSAL | (4 << 2)), 2, 0, 0 }, /* extnValue */
    { (ASN_TAG_CLASS_UNIVERSAL | (6 << 2)), 0, 0, 0 } /* extnID */
};
asn_SEQUENCE_specifics_t asn_SPC_Extension_specs_1 = {
	sizeof(struct Extension),
	offsetof(struct Extension, _asn_ctx),
	asn_MAP_Extension_tag2el_1,
	3,	/* Count of tags in the map */
	asn_MAP_Extension_oms_1,	/* Optional members */
	1, 0,	/* Root/Additions */
	-1,	/* First extension addition */
};
asn_TYPE_descriptor_t asn_DEF_Extension = {
	"Extension",
	"Extension",
	&asn_OP_SEQUENCE,
	asn_DEF_Extension_tags_1,
	sizeof(asn_DEF_Extension_tags_1)
		/sizeof(asn_DEF_Extension_tags_1[0]), /* 1 */
	asn_DEF_Extension_tags_1,	/* Same as above */
	sizeof(asn_DEF_Extension_tags_1)
		/sizeof(asn_DEF_Extension_tags_1[0]), /* 1 */
	{ 0, 0, SEQUENCE_constraint },
	asn_MBR_Extension_1,
	3,	/* Elements count */
	&asn_SPC_Extension_specs_1	/* Additional specs */
};

