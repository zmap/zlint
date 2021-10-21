/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "PKIXAttributeCertificate"
 * 	found in "asn1/rfc3281-PKIXAttributeCertificate.asn1"
 * 	`asn1c -S /home/fotisl/Projects/revlintsynt/asn1c/skeletons -pdu=all -fwide-types -fcompound-names`
 */

#ifndef	_IetfAttrSyntax_H_
#define	_IetfAttrSyntax_H_


#include <asn_application.h>

/* Including external dependencies */
#include <asn_SEQUENCE_OF.h>
#include <OCTET_STRING.h>
#include <OBJECT_IDENTIFIER.h>
#include <UTF8String.h>
#include <constr_CHOICE.h>
#include <constr_SEQUENCE_OF.h>
#include <constr_SEQUENCE.h>

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum IetfAttrSyntax__values__Member_PR {
	IetfAttrSyntax__values__Member_PR_NOTHING,	/* No components present */
	IetfAttrSyntax__values__Member_PR_octets,
	IetfAttrSyntax__values__Member_PR_oid,
	IetfAttrSyntax__values__Member_PR_string
} IetfAttrSyntax__values__Member_PR;

/* Forward declarations */
struct GeneralNames;

/* Forward definitions */
typedef struct IetfAttrSyntax__values__Member {
	IetfAttrSyntax__values__Member_PR present;
	union IetfAttrSyntax__values__Member_u {
		OCTET_STRING_t	 octets;
		OBJECT_IDENTIFIER_t	 oid;
		UTF8String_t	 string;
	} choice;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} IetfAttrSyntax__values__Member;

/* IetfAttrSyntax */
typedef struct IetfAttrSyntax {
	struct GeneralNames	*policyAuthority	/* OPTIONAL */;
	struct IetfAttrSyntax__values {
		A_SEQUENCE_OF(IetfAttrSyntax__values__Member) list;
		
		/* Context for parsing across buffer boundaries */
		asn_struct_ctx_t _asn_ctx;
	} values;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} IetfAttrSyntax_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_IetfAttrSyntax;

#ifdef __cplusplus
}
#endif

/* Referred external types */
#include "GeneralNames.h"

#endif	/* _IetfAttrSyntax_H_ */
#include <asn_internal.h>
