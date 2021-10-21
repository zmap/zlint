/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "OCSP"
 * 	found in "asn1/rfc2560-OCSP.asn1"
 * 	`asn1c -S /home/fotisl/Projects/revlintsynt/asn1c/skeletons -pdu=all -fwide-types -fcompound-names`
 */

#include "ArchiveCutoff.h"

/*
 * This type is implemented using GeneralizedTime,
 * so here we adjust the DEF accordingly.
 */
static const ber_tlv_tag_t asn_DEF_ArchiveCutoff_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (24 << 2))
};
asn_TYPE_descriptor_t asn_DEF_ArchiveCutoff = {
	"ArchiveCutoff",
	"ArchiveCutoff",
	&asn_OP_GeneralizedTime,
	asn_DEF_ArchiveCutoff_tags_1,
	sizeof(asn_DEF_ArchiveCutoff_tags_1)
		/sizeof(asn_DEF_ArchiveCutoff_tags_1[0]), /* 1 */
	asn_DEF_ArchiveCutoff_tags_1,	/* Same as above */
	sizeof(asn_DEF_ArchiveCutoff_tags_1)
		/sizeof(asn_DEF_ArchiveCutoff_tags_1[0]), /* 1 */
	{ 0, 0, GeneralizedTime_constraint },
	0, 0,	/* No members */
	0	/* No specifics */
};

