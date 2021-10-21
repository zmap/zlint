/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "PKIXAttributeCertificate"
 * 	found in "asn1/rfc3281-PKIXAttributeCertificate.asn1"
 * 	`asn1c -S /home/fotisl/Projects/revlintsynt/asn1c/skeletons -pdu=all -fwide-types -fcompound-names`
 */

#include "ProxyInfo.h"

static asn_TYPE_member_t asn_MBR_ProxyInfo_1[] = {
	{ ATF_POINTER, 0, 0,
		(ASN_TAG_CLASS_UNIVERSAL | (16 << 2)),
		0,
		&asn_DEF_Targets,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		""
		},
};
static const ber_tlv_tag_t asn_DEF_ProxyInfo_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static asn_SET_OF_specifics_t asn_SPC_ProxyInfo_specs_1 = {
	sizeof(struct ProxyInfo),
	offsetof(struct ProxyInfo, _asn_ctx),
	0,	/* XER encoding is XMLDelimitedItemList */
};
asn_TYPE_descriptor_t asn_DEF_ProxyInfo = {
	"ProxyInfo",
	"ProxyInfo",
	&asn_OP_SEQUENCE_OF,
	asn_DEF_ProxyInfo_tags_1,
	sizeof(asn_DEF_ProxyInfo_tags_1)
		/sizeof(asn_DEF_ProxyInfo_tags_1[0]), /* 1 */
	asn_DEF_ProxyInfo_tags_1,	/* Same as above */
	sizeof(asn_DEF_ProxyInfo_tags_1)
		/sizeof(asn_DEF_ProxyInfo_tags_1[0]), /* 1 */
	{ 0, 0, SEQUENCE_OF_constraint },
	asn_MBR_ProxyInfo_1,
	1,	/* Single element */
	&asn_SPC_ProxyInfo_specs_1	/* Additional specs */
};

