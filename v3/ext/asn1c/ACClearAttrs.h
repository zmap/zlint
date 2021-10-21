/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "PKIXAttributeCertificate"
 * 	found in "asn1/rfc3281-PKIXAttributeCertificate.asn1"
 * 	`asn1c -S /home/fotisl/Projects/revlintsynt/asn1c/skeletons -pdu=all -fwide-types -fcompound-names`
 */

#ifndef	_ACClearAttrs_H_
#define	_ACClearAttrs_H_


#include <asn_application.h>

/* Including external dependencies */
#include "GeneralName.h"
#include <INTEGER.h>
#include <asn_SEQUENCE_OF.h>
#include <constr_SEQUENCE_OF.h>
#include <constr_SEQUENCE.h>

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct Attribute;

/* ACClearAttrs */
typedef struct ACClearAttrs {
	GeneralName_t	 acIssuer;
	INTEGER_t	 acSerial;
	struct ACClearAttrs__attrs {
		A_SEQUENCE_OF(struct Attribute) list;
		
		/* Context for parsing across buffer boundaries */
		asn_struct_ctx_t _asn_ctx;
	} attrs;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} ACClearAttrs_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_ACClearAttrs;

#ifdef __cplusplus
}
#endif

/* Referred external types */
#include "Attribute.h"

#endif	/* _ACClearAttrs_H_ */
#include <asn_internal.h>
