/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "PKIX1-PSS-OAEP-Algorithms"
 * 	found in "asn1/rfc4055-PKIX1-PSS-OAEP-Algorithms.asn1"
 * 	`asn1c -S /home/fotisl/Projects/revlintsynt/asn1c/skeletons -pdu=all -fwide-types -fcompound-names`
 */

#include "RSASSA-PSS-params.h"

static int asn_DFL_4_cmp_20(const void *sptr) {
	const INTEGER_t *st = sptr;
	
	if(!st) {
		return -1; /* No value is not a default value */
	}
	
	/* Test default value 20 */
	long value;
	if(asn_INTEGER2long(st, &value))
		return -1;
	return (value != 20);
}
static int asn_DFL_4_set_20(void **sptr) {
	INTEGER_t *st = *sptr;
	
	if(!st) {
		st = (*sptr = CALLOC(1, sizeof(*st)));
		if(!st) return -1;
	}
	
	/* Install default value 20 */
	return asn_long2INTEGER(st, 20);
}
static int asn_DFL_5_cmp_1(const void *sptr) {
	const INTEGER_t *st = sptr;
	
	if(!st) {
		return -1; /* No value is not a default value */
	}
	
	/* Test default value 1 */
	long value;
	if(asn_INTEGER2long(st, &value))
		return -1;
	return (value != 1);
}
static int asn_DFL_5_set_1(void **sptr) {
	INTEGER_t *st = *sptr;
	
	if(!st) {
		st = (*sptr = CALLOC(1, sizeof(*st)));
		if(!st) return -1;
	}
	
	/* Install default value 1 */
	return asn_long2INTEGER(st, 1);
}
asn_TYPE_member_t asn_MBR_RSASSA_PSS_params_1[] = {
	{ ATF_POINTER, 4, offsetof(struct RSASSA_PSS_params, hashAlgorithm),
		(ASN_TAG_CLASS_CONTEXT | (0 << 2)),
		+1,	/* EXPLICIT tag at current level */
		&asn_DEF_HashAlgorithm,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"hashAlgorithm"
		},
	{ ATF_POINTER, 3, offsetof(struct RSASSA_PSS_params, maskGenAlgorithm),
		(ASN_TAG_CLASS_CONTEXT | (1 << 2)),
		+1,	/* EXPLICIT tag at current level */
		&asn_DEF_MaskGenAlgorithm,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"maskGenAlgorithm"
		},
	{ ATF_POINTER, 2, offsetof(struct RSASSA_PSS_params, saltLength),
		(ASN_TAG_CLASS_CONTEXT | (2 << 2)),
		+1,	/* EXPLICIT tag at current level */
		&asn_DEF_INTEGER,
		0,
		{ 0, 0, 0 },
		&asn_DFL_4_cmp_20,	/* Compare DEFAULT 20 */
		&asn_DFL_4_set_20,	/* Set DEFAULT 20 */
		"saltLength"
		},
	{ ATF_POINTER, 1, offsetof(struct RSASSA_PSS_params, trailerField),
		(ASN_TAG_CLASS_CONTEXT | (3 << 2)),
		+1,	/* EXPLICIT tag at current level */
		&asn_DEF_INTEGER,
		0,
		{ 0, 0, 0 },
		&asn_DFL_5_cmp_1,	/* Compare DEFAULT 1 */
		&asn_DFL_5_set_1,	/* Set DEFAULT 1 */
		"trailerField"
		},
};
static const int asn_MAP_RSASSA_PSS_params_oms_1[] = { 0, 1, 2, 3 };
static const ber_tlv_tag_t asn_DEF_RSASSA_PSS_params_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static const asn_TYPE_tag2member_t asn_MAP_RSASSA_PSS_params_tag2el_1[] = {
    { (ASN_TAG_CLASS_CONTEXT | (0 << 2)), 0, 0, 0 }, /* hashAlgorithm */
    { (ASN_TAG_CLASS_CONTEXT | (1 << 2)), 1, 0, 0 }, /* maskGenAlgorithm */
    { (ASN_TAG_CLASS_CONTEXT | (2 << 2)), 2, 0, 0 }, /* saltLength */
    { (ASN_TAG_CLASS_CONTEXT | (3 << 2)), 3, 0, 0 } /* trailerField */
};
asn_SEQUENCE_specifics_t asn_SPC_RSASSA_PSS_params_specs_1 = {
	sizeof(struct RSASSA_PSS_params),
	offsetof(struct RSASSA_PSS_params, _asn_ctx),
	asn_MAP_RSASSA_PSS_params_tag2el_1,
	4,	/* Count of tags in the map */
	asn_MAP_RSASSA_PSS_params_oms_1,	/* Optional members */
	4, 0,	/* Root/Additions */
	-1,	/* First extension addition */
};
asn_TYPE_descriptor_t asn_DEF_RSASSA_PSS_params = {
	"RSASSA-PSS-params",
	"RSASSA-PSS-params",
	&asn_OP_SEQUENCE,
	asn_DEF_RSASSA_PSS_params_tags_1,
	sizeof(asn_DEF_RSASSA_PSS_params_tags_1)
		/sizeof(asn_DEF_RSASSA_PSS_params_tags_1[0]), /* 1 */
	asn_DEF_RSASSA_PSS_params_tags_1,	/* Same as above */
	sizeof(asn_DEF_RSASSA_PSS_params_tags_1)
		/sizeof(asn_DEF_RSASSA_PSS_params_tags_1[0]), /* 1 */
	{ 0, 0, SEQUENCE_constraint },
	asn_MBR_RSASSA_PSS_params_1,
	4,	/* Elements count */
	&asn_SPC_RSASSA_PSS_params_specs_1	/* Additional specs */
};

