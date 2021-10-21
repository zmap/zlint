/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "OtherAttributesModule"
 * 	found in "asn1/otherattrs.asn1"
 * 	`asn1c -S /home/fotisl/Projects/revlintsynt/asn1c/skeletons -pdu=all -fwide-types -fcompound-names`
 */

#ifndef	_PostalAddress_H_
#define	_PostalAddress_H_


#include <asn_application.h>

/* Including external dependencies */
#include <asn_SEQUENCE_OF.h>
#include <constr_SEQUENCE_OF.h>

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct PostalString;

/* PostalAddress */
typedef struct PostalAddress {
	A_SEQUENCE_OF(struct PostalString) list;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} PostalAddress_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_PostalAddress;

#ifdef __cplusplus
}
#endif

/* Referred external types */
#include "PostalString.h"

#endif	/* _PostalAddress_H_ */
#include <asn_internal.h>
