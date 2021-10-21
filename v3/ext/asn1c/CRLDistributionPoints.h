/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "PKIX1Implicit88"
 * 	found in "asn1/rfc5280-PKIX1Implicit88.asn1"
 * 	`asn1c -S /home/fotisl/Projects/revlintsynt/asn1c/skeletons -pdu=all -fwide-types -fcompound-names`
 */

#ifndef	_CRLDistributionPoints_H_
#define	_CRLDistributionPoints_H_


#include <asn_application.h>

/* Including external dependencies */
#include <asn_SEQUENCE_OF.h>
#include <constr_SEQUENCE_OF.h>

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct DistributionPoint;

/* CRLDistributionPoints */
typedef struct CRLDistributionPoints {
	A_SEQUENCE_OF(struct DistributionPoint) list;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} CRLDistributionPoints_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_CRLDistributionPoints;
extern asn_SET_OF_specifics_t asn_SPC_CRLDistributionPoints_specs_1;
extern asn_TYPE_member_t asn_MBR_CRLDistributionPoints_1[1];
extern asn_per_constraints_t asn_PER_type_CRLDistributionPoints_constr_1;

#ifdef __cplusplus
}
#endif

/* Referred external types */
#include "DistributionPoint.h"

#endif	/* _CRLDistributionPoints_H_ */
#include <asn_internal.h>
