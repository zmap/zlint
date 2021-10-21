/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "OtherAttributesModule"
 * 	found in "asn1/otherattrs.asn1"
 * 	`asn1c -S /home/fotisl/Projects/revlintsynt/asn1c/skeletons -pdu=all -fwide-types -fcompound-names`
 */

#ifndef	_Description_H_
#define	_Description_H_


#include <asn_application.h>

/* Including external dependencies */
#include <TeletexString.h>
#include <PrintableString.h>
#include <UniversalString.h>
#include <UTF8String.h>
#include <BMPString.h>
#include <constr_CHOICE.h>

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum Description_PR {
	Description_PR_NOTHING,	/* No components present */
	Description_PR_teletexString,
	Description_PR_printableString,
	Description_PR_universalString,
	Description_PR_utf8String,
	Description_PR_bmpString
} Description_PR;

/* Description */
typedef struct Description {
	Description_PR present;
	union Description_u {
		TeletexString_t	 teletexString;
		PrintableString_t	 printableString;
		UniversalString_t	 universalString;
		UTF8String_t	 utf8String;
		BMPString_t	 bmpString;
	} choice;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} Description_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_Description;

#ifdef __cplusplus
}
#endif

#endif	/* _Description_H_ */
#include <asn_internal.h>
