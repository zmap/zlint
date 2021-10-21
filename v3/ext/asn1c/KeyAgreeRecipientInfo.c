/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "CryptographicMessageSyntax"
 * 	found in "asn1/rfc3369-CryptographicMessageSyntax.asn1"
 * 	`asn1c -S /home/fotisl/Projects/revlintsynt/asn1c/skeletons -pdu=all -fwide-types -fcompound-names`
 */

#include "KeyAgreeRecipientInfo.h"

asn_TYPE_member_t asn_MBR_KeyAgreeRecipientInfo_1[] = {
	{ ATF_NOFLAGS, 0, offsetof(struct KeyAgreeRecipientInfo, version),
		(ASN_TAG_CLASS_UNIVERSAL | (2 << 2)),
		0,
		&asn_DEF_CMSVersion,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"version"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct KeyAgreeRecipientInfo, originator),
		(ASN_TAG_CLASS_CONTEXT | (0 << 2)),
		+1,	/* EXPLICIT tag at current level */
		&asn_DEF_OriginatorIdentifierOrKey,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"originator"
		},
	{ ATF_POINTER, 1, offsetof(struct KeyAgreeRecipientInfo, ukm),
		(ASN_TAG_CLASS_CONTEXT | (1 << 2)),
		+1,	/* EXPLICIT tag at current level */
		&asn_DEF_UserKeyingMaterial,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"ukm"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct KeyAgreeRecipientInfo, keyEncryptionAlgorithm),
		(ASN_TAG_CLASS_UNIVERSAL | (16 << 2)),
		0,
		&asn_DEF_KeyEncryptionAlgorithmIdentifier,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"keyEncryptionAlgorithm"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct KeyAgreeRecipientInfo, recipientEncryptedKeys),
		(ASN_TAG_CLASS_UNIVERSAL | (16 << 2)),
		0,
		&asn_DEF_RecipientEncryptedKeys,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"recipientEncryptedKeys"
		},
};
static const int asn_MAP_KeyAgreeRecipientInfo_oms_1[] = { 2 };
static const ber_tlv_tag_t asn_DEF_KeyAgreeRecipientInfo_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static const asn_TYPE_tag2member_t asn_MAP_KeyAgreeRecipientInfo_tag2el_1[] = {
    { (ASN_TAG_CLASS_UNIVERSAL | (2 << 2)), 0, 0, 0 }, /* version */
    { (ASN_TAG_CLASS_UNIVERSAL | (16 << 2)), 3, 0, 1 }, /* keyEncryptionAlgorithm */
    { (ASN_TAG_CLASS_UNIVERSAL | (16 << 2)), 4, -1, 0 }, /* recipientEncryptedKeys */
    { (ASN_TAG_CLASS_CONTEXT | (0 << 2)), 1, 0, 0 }, /* originator */
    { (ASN_TAG_CLASS_CONTEXT | (1 << 2)), 2, 0, 0 } /* ukm */
};
asn_SEQUENCE_specifics_t asn_SPC_KeyAgreeRecipientInfo_specs_1 = {
	sizeof(struct KeyAgreeRecipientInfo),
	offsetof(struct KeyAgreeRecipientInfo, _asn_ctx),
	asn_MAP_KeyAgreeRecipientInfo_tag2el_1,
	5,	/* Count of tags in the map */
	asn_MAP_KeyAgreeRecipientInfo_oms_1,	/* Optional members */
	1, 0,	/* Root/Additions */
	-1,	/* First extension addition */
};
asn_TYPE_descriptor_t asn_DEF_KeyAgreeRecipientInfo = {
	"KeyAgreeRecipientInfo",
	"KeyAgreeRecipientInfo",
	&asn_OP_SEQUENCE,
	asn_DEF_KeyAgreeRecipientInfo_tags_1,
	sizeof(asn_DEF_KeyAgreeRecipientInfo_tags_1)
		/sizeof(asn_DEF_KeyAgreeRecipientInfo_tags_1[0]), /* 1 */
	asn_DEF_KeyAgreeRecipientInfo_tags_1,	/* Same as above */
	sizeof(asn_DEF_KeyAgreeRecipientInfo_tags_1)
		/sizeof(asn_DEF_KeyAgreeRecipientInfo_tags_1[0]), /* 1 */
	{ 0, 0, SEQUENCE_constraint },
	asn_MBR_KeyAgreeRecipientInfo_1,
	5,	/* Elements count */
	&asn_SPC_KeyAgreeRecipientInfo_specs_1	/* Additional specs */
};

