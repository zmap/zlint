/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "PKIX1Explicit88"
 * 	found in "asn1/rfc5280-PKIX1Explicit88.asn1"
 * 	`asn1c -S /home/fotisl/Projects/revlintsynt/asn1c/skeletons -pdu=all -fwide-types -fcompound-names`
 */

#include "StreetAddress2.h"

/*
 * This type is implemented using PDSParameter,
 * so here we adjust the DEF accordingly.
 */
static const ber_tlv_tag_t asn_DEF_StreetAddress2_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (17 << 2))
};
asn_TYPE_descriptor_t asn_DEF_StreetAddress2 = {
	"StreetAddress2",
	"StreetAddress2",
	&asn_OP_SET,
	asn_DEF_StreetAddress2_tags_1,
	sizeof(asn_DEF_StreetAddress2_tags_1)
		/sizeof(asn_DEF_StreetAddress2_tags_1[0]), /* 1 */
	asn_DEF_StreetAddress2_tags_1,	/* Same as above */
	sizeof(asn_DEF_StreetAddress2_tags_1)
		/sizeof(asn_DEF_StreetAddress2_tags_1[0]), /* 1 */
	{ 0, 0, SET_constraint },
	asn_MBR_PDSParameter_1,
	2,	/* Elements count */
	&asn_SPC_PDSParameter_specs_1	/* Additional specs */
};

