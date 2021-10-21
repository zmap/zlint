/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "PKIXAttributeCertificate"
 * 	found in "asn1/rfc3281-PKIXAttributeCertificate.asn1"
 * 	`asn1c -S /home/fotisl/Projects/revlintsynt/asn1c/skeletons -pdu=all -fwide-types -fcompound-names`
 */

#include "AttributeCertificateInfo.h"

static asn_TYPE_member_t asn_MBR_attributes_8[] = {
	{ ATF_POINTER, 0, 0,
		(ASN_TAG_CLASS_UNIVERSAL | (16 << 2)),
		0,
		&asn_DEF_Attribute,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		""
		},
};
static const ber_tlv_tag_t asn_DEF_attributes_tags_8[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static asn_SET_OF_specifics_t asn_SPC_attributes_specs_8 = {
	sizeof(struct AttributeCertificateInfo__attributes),
	offsetof(struct AttributeCertificateInfo__attributes, _asn_ctx),
	0,	/* XER encoding is XMLDelimitedItemList */
};
static /* Use -fall-defs-global to expose */
asn_TYPE_descriptor_t asn_DEF_attributes_8 = {
	"attributes",
	"attributes",
	&asn_OP_SEQUENCE_OF,
	asn_DEF_attributes_tags_8,
	sizeof(asn_DEF_attributes_tags_8)
		/sizeof(asn_DEF_attributes_tags_8[0]), /* 1 */
	asn_DEF_attributes_tags_8,	/* Same as above */
	sizeof(asn_DEF_attributes_tags_8)
		/sizeof(asn_DEF_attributes_tags_8[0]), /* 1 */
	{ 0, 0, SEQUENCE_OF_constraint },
	asn_MBR_attributes_8,
	1,	/* Single element */
	&asn_SPC_attributes_specs_8	/* Additional specs */
};

asn_TYPE_member_t asn_MBR_AttributeCertificateInfo_1[] = {
	{ ATF_NOFLAGS, 0, offsetof(struct AttributeCertificateInfo, version),
		(ASN_TAG_CLASS_UNIVERSAL | (2 << 2)),
		0,
		&asn_DEF_AttCertVersion,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"version"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct AttributeCertificateInfo, holder),
		(ASN_TAG_CLASS_UNIVERSAL | (16 << 2)),
		0,
		&asn_DEF_Holder,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"holder"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct AttributeCertificateInfo, issuer),
		-1 /* Ambiguous tag (CHOICE?) */,
		0,
		&asn_DEF_AttCertIssuer,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"issuer"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct AttributeCertificateInfo, signature),
		(ASN_TAG_CLASS_UNIVERSAL | (16 << 2)),
		0,
		&asn_DEF_AlgorithmIdentifier,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"signature"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct AttributeCertificateInfo, serialNumber),
		(ASN_TAG_CLASS_UNIVERSAL | (2 << 2)),
		0,
		&asn_DEF_CertificateSerialNumber,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"serialNumber"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct AttributeCertificateInfo, attrCertValidityPeriod),
		(ASN_TAG_CLASS_UNIVERSAL | (16 << 2)),
		0,
		&asn_DEF_AttCertValidityPeriod,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"attrCertValidityPeriod"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct AttributeCertificateInfo, attributes),
		(ASN_TAG_CLASS_UNIVERSAL | (16 << 2)),
		0,
		&asn_DEF_attributes_8,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"attributes"
		},
	{ ATF_POINTER, 2, offsetof(struct AttributeCertificateInfo, issuerUniqueID),
		(ASN_TAG_CLASS_UNIVERSAL | (3 << 2)),
		0,
		&asn_DEF_UniqueIdentifier,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"issuerUniqueID"
		},
	{ ATF_POINTER, 1, offsetof(struct AttributeCertificateInfo, extensions),
		(ASN_TAG_CLASS_UNIVERSAL | (16 << 2)),
		0,
		&asn_DEF_Extensions,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"extensions"
		},
};
static const int asn_MAP_AttributeCertificateInfo_oms_1[] = { 7, 8 };
static const ber_tlv_tag_t asn_DEF_AttributeCertificateInfo_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static const asn_TYPE_tag2member_t asn_MAP_AttributeCertificateInfo_tag2el_1[] = {
    { (ASN_TAG_CLASS_UNIVERSAL | (2 << 2)), 0, 0, 1 }, /* version */
    { (ASN_TAG_CLASS_UNIVERSAL | (2 << 2)), 4, -1, 0 }, /* serialNumber */
    { (ASN_TAG_CLASS_UNIVERSAL | (3 << 2)), 7, 0, 0 }, /* issuerUniqueID */
    { (ASN_TAG_CLASS_UNIVERSAL | (16 << 2)), 1, 0, 5 }, /* holder */
    { (ASN_TAG_CLASS_UNIVERSAL | (16 << 2)), 2, -1, 4 }, /* v1Form */
    { (ASN_TAG_CLASS_UNIVERSAL | (16 << 2)), 3, -2, 3 }, /* signature */
    { (ASN_TAG_CLASS_UNIVERSAL | (16 << 2)), 5, -3, 2 }, /* attrCertValidityPeriod */
    { (ASN_TAG_CLASS_UNIVERSAL | (16 << 2)), 6, -4, 1 }, /* attributes */
    { (ASN_TAG_CLASS_UNIVERSAL | (16 << 2)), 8, -5, 0 }, /* extensions */
    { (ASN_TAG_CLASS_CONTEXT | (0 << 2)), 2, 0, 0 } /* v2Form */
};
asn_SEQUENCE_specifics_t asn_SPC_AttributeCertificateInfo_specs_1 = {
	sizeof(struct AttributeCertificateInfo),
	offsetof(struct AttributeCertificateInfo, _asn_ctx),
	asn_MAP_AttributeCertificateInfo_tag2el_1,
	10,	/* Count of tags in the map */
	asn_MAP_AttributeCertificateInfo_oms_1,	/* Optional members */
	2, 0,	/* Root/Additions */
	-1,	/* First extension addition */
};
asn_TYPE_descriptor_t asn_DEF_AttributeCertificateInfo = {
	"AttributeCertificateInfo",
	"AttributeCertificateInfo",
	&asn_OP_SEQUENCE,
	asn_DEF_AttributeCertificateInfo_tags_1,
	sizeof(asn_DEF_AttributeCertificateInfo_tags_1)
		/sizeof(asn_DEF_AttributeCertificateInfo_tags_1[0]), /* 1 */
	asn_DEF_AttributeCertificateInfo_tags_1,	/* Same as above */
	sizeof(asn_DEF_AttributeCertificateInfo_tags_1)
		/sizeof(asn_DEF_AttributeCertificateInfo_tags_1[0]), /* 1 */
	{ 0, 0, SEQUENCE_constraint },
	asn_MBR_AttributeCertificateInfo_1,
	9,	/* Elements count */
	&asn_SPC_AttributeCertificateInfo_specs_1	/* Additional specs */
};

