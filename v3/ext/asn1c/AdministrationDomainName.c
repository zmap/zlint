/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "PKIX1Explicit88"
 * 	found in "asn1/rfc5280-PKIX1Explicit88.asn1"
 * 	`asn1c -S /home/fotisl/Projects/revlintsynt/asn1c/skeletons -pdu=all -fwide-types -fcompound-names`
 */

#include "AdministrationDomainName.h"

static const int permitted_alphabet_table_2[256] = {
 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,	/*                  */
 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,	/*                  */
 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,	/* .                */
 2, 3, 4, 5, 6, 7, 8, 9,10,11, 0, 0, 0, 0, 0, 0,	/* 0123456789       */
};
static const int permitted_alphabet_code2value_2[11] = {
32,48,49,50,51,52,53,54,55,56,57,};


static int check_permitted_alphabet_2(const void *sptr) {
	const int *table = permitted_alphabet_table_2;
	/* The underlying type is NumericString */
	const NumericString_t *st = (const NumericString_t *)sptr;
	const uint8_t *ch = st->buf;
	const uint8_t *end = ch + st->size;
	
	for(; ch < end; ch++) {
		uint8_t cv = *ch;
		if(!table[cv]) return -1;
	}
	return 0;
}

static const int permitted_alphabet_table_3[256] = {
 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,	/*                  */
 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,	/*                  */
 1, 0, 0, 0, 0, 0, 0, 2, 3, 4, 0, 5, 6, 7, 8, 9,	/* .      '() +,-./ */
10,11,12,13,14,15,16,17,18,19,20, 0, 0,21, 0,22,	/* 0123456789:  = ? */
 0,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,	/*  ABCDEFGHIJKLMNO */
38,39,40,41,42,43,44,45,46,47,48, 0, 0, 0, 0, 0,	/* PQRSTUVWXYZ      */
 0,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,	/*  abcdefghijklmno */
64,65,66,67,68,69,70,71,72,73,74, 0, 0, 0, 0, 0,	/* pqrstuvwxyz      */
};
static const int permitted_alphabet_code2value_3[74] = {
32,39,40,41,43,44,45,46,47,48,49,50,51,52,53,54,
55,56,57,58,61,63,65,66,67,68,69,70,71,72,73,74,
75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,
97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,
113,114,115,116,117,118,119,120,121,122,};


static int check_permitted_alphabet_3(const void *sptr) {
	const int *table = permitted_alphabet_table_3;
	/* The underlying type is PrintableString */
	const PrintableString_t *st = (const PrintableString_t *)sptr;
	const uint8_t *ch = st->buf;
	const uint8_t *end = ch + st->size;
	
	for(; ch < end; ch++) {
		uint8_t cv = *ch;
		if(!table[cv]) return -1;
	}
	return 0;
}

static int
memb_numeric_constraint_1(const asn_TYPE_descriptor_t *td, const void *sptr,
			asn_app_constraint_failed_f *ctfailcb, void *app_key) {
	const NumericString_t *st = (const NumericString_t *)sptr;
	size_t size;
	
	if(!sptr) {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: value not given (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
	
	size = st->size;
	
	if((size <= 16)
		 && !check_permitted_alphabet_2(st)) {
		/* Constraint check succeeded */
		return 0;
	} else {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: constraint failed (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
}

static int asn_PER_MAP_numeric_2_v2c(unsigned int value) {
	if(value >= sizeof(permitted_alphabet_table_2)/sizeof(permitted_alphabet_table_2[0]))
		return -1;
	return permitted_alphabet_table_2[value] - 1;
}
static int asn_PER_MAP_numeric_2_c2v(unsigned int code) {
	if(code >= sizeof(permitted_alphabet_code2value_2)/sizeof(permitted_alphabet_code2value_2[0]))
		return -1;
	return permitted_alphabet_code2value_2[code];
}
static int
memb_printable_constraint_1(const asn_TYPE_descriptor_t *td, const void *sptr,
			asn_app_constraint_failed_f *ctfailcb, void *app_key) {
	const PrintableString_t *st = (const PrintableString_t *)sptr;
	size_t size;
	
	if(!sptr) {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: value not given (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
	
	size = st->size;
	
	if((size <= 16)
		 && !check_permitted_alphabet_3(st)) {
		/* Constraint check succeeded */
		return 0;
	} else {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: constraint failed (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
}

static int asn_PER_MAP_printable_3_v2c(unsigned int value) {
	if(value >= sizeof(permitted_alphabet_table_3)/sizeof(permitted_alphabet_table_3[0]))
		return -1;
	return permitted_alphabet_table_3[value] - 1;
}
static int asn_PER_MAP_printable_3_c2v(unsigned int code) {
	if(code >= sizeof(permitted_alphabet_code2value_3)/sizeof(permitted_alphabet_code2value_3[0]))
		return -1;
	return permitted_alphabet_code2value_3[code];
}
static asn_oer_constraints_t asn_OER_memb_numeric_constr_2 CC_NOTUSED = {
	{ 0, 0 },
	-1	/* (SIZE(0..16)) */};
static asn_per_constraints_t asn_PER_memb_numeric_constr_2 CC_NOTUSED = {
	{ APC_CONSTRAINED,	 4,  4,  32,  57 }	/* (32..57) */,
	{ APC_CONSTRAINED,	 5,  5,  0,  16 }	/* (SIZE(0..16)) */,
	asn_PER_MAP_numeric_2_v2c,	/* Value to PER code map */
	asn_PER_MAP_numeric_2_c2v	/* PER code to value map */
};
static asn_oer_constraints_t asn_OER_memb_printable_constr_3 CC_NOTUSED = {
	{ 0, 0 },
	-1	/* (SIZE(0..16)) */};
static asn_per_constraints_t asn_PER_memb_printable_constr_3 CC_NOTUSED = {
	{ APC_CONSTRAINED,	 7,  7,  32,  122 }	/* (32..122) */,
	{ APC_CONSTRAINED,	 5,  5,  0,  16 }	/* (SIZE(0..16)) */,
	asn_PER_MAP_printable_3_v2c,	/* Value to PER code map */
	asn_PER_MAP_printable_3_c2v	/* PER code to value map */
};
static asn_oer_constraints_t asn_OER_type_AdministrationDomainName_constr_1 CC_NOTUSED = {
	{ 0, 0 },
	-1};
asn_per_constraints_t asn_PER_type_AdministrationDomainName_constr_1 CC_NOTUSED = {
	{ APC_CONSTRAINED,	 1,  1,  0,  1 }	/* (0..1) */,
	{ APC_UNCONSTRAINED,	-1, -1,  0,  0 },
	0, 0	/* No PER value map */
};
asn_TYPE_member_t asn_MBR_AdministrationDomainName_1[] = {
	{ ATF_NOFLAGS, 0, offsetof(struct AdministrationDomainName, choice.numeric),
		(ASN_TAG_CLASS_UNIVERSAL | (18 << 2)),
		0,
		&asn_DEF_NumericString,
		0,
		{ &asn_OER_memb_numeric_constr_2, &asn_PER_memb_numeric_constr_2,  memb_numeric_constraint_1 },
		0, 0, /* No default value */
		"numeric"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct AdministrationDomainName, choice.printable),
		(ASN_TAG_CLASS_UNIVERSAL | (19 << 2)),
		0,
		&asn_DEF_PrintableString,
		0,
		{ &asn_OER_memb_printable_constr_3, &asn_PER_memb_printable_constr_3,  memb_printable_constraint_1 },
		0, 0, /* No default value */
		"printable"
		},
};
static const ber_tlv_tag_t asn_DEF_AdministrationDomainName_tags_1[] = {
	(ASN_TAG_CLASS_APPLICATION | (2 << 2))
};
static const asn_TYPE_tag2member_t asn_MAP_AdministrationDomainName_tag2el_1[] = {
    { (ASN_TAG_CLASS_UNIVERSAL | (18 << 2)), 0, 0, 0 }, /* numeric */
    { (ASN_TAG_CLASS_UNIVERSAL | (19 << 2)), 1, 0, 0 } /* printable */
};
asn_CHOICE_specifics_t asn_SPC_AdministrationDomainName_specs_1 = {
	sizeof(struct AdministrationDomainName),
	offsetof(struct AdministrationDomainName, _asn_ctx),
	offsetof(struct AdministrationDomainName, present),
	sizeof(((struct AdministrationDomainName *)0)->present),
	asn_MAP_AdministrationDomainName_tag2el_1,
	2,	/* Count of tags in the map */
	0, 0,
	-1	/* Extensions start */
};
asn_TYPE_descriptor_t asn_DEF_AdministrationDomainName = {
	"AdministrationDomainName",
	"AdministrationDomainName",
	&asn_OP_CHOICE,
	asn_DEF_AdministrationDomainName_tags_1,
	sizeof(asn_DEF_AdministrationDomainName_tags_1)
		/sizeof(asn_DEF_AdministrationDomainName_tags_1[0]), /* 1 */
	asn_DEF_AdministrationDomainName_tags_1,	/* Same as above */
	sizeof(asn_DEF_AdministrationDomainName_tags_1)
		/sizeof(asn_DEF_AdministrationDomainName_tags_1[0]), /* 1 */
	{ &asn_OER_type_AdministrationDomainName_constr_1, &asn_PER_type_AdministrationDomainName_constr_1, CHOICE_constraint },
	asn_MBR_AdministrationDomainName_1,
	2,	/* Elements count */
	&asn_SPC_AdministrationDomainName_specs_1	/* Additional specs */
};

