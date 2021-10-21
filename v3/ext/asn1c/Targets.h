/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "PKIXAttributeCertificate"
 * 	found in "asn1/rfc3281-PKIXAttributeCertificate.asn1"
 * 	`asn1c -S /home/fotisl/Projects/revlintsynt/asn1c/skeletons -pdu=all -fwide-types -fcompound-names`
 */

#ifndef	_Targets_H_
#define	_Targets_H_


#include <asn_application.h>

/* Including external dependencies */
#include <asn_SEQUENCE_OF.h>
#include <constr_SEQUENCE_OF.h>

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct Target;

/* Targets */
typedef struct Targets {
	A_SEQUENCE_OF(struct Target) list;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} Targets_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_Targets;
extern asn_SET_OF_specifics_t asn_SPC_Targets_specs_1;
extern asn_TYPE_member_t asn_MBR_Targets_1[1];

#ifdef __cplusplus
}
#endif

/* Referred external types */
#include "Target.h"

#endif	/* _Targets_H_ */
#include <asn_internal.h>
