/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "PKIX1Explicit88"
 * 	found in "asn1/rfc5280-PKIX1Explicit88.asn1"
 * 	`asn1c -S /home/fotisl/Projects/revlintsynt/asn1c/skeletons -pdu=all -fwide-types -fcompound-names`
 */

#ifndef	_TeletexOrganizationalUnitNames_H_
#define	_TeletexOrganizationalUnitNames_H_


#include <asn_application.h>

/* Including external dependencies */
#include "TeletexOrganizationalUnitName.h"
#include <asn_SEQUENCE_OF.h>
#include <constr_SEQUENCE_OF.h>

#ifdef __cplusplus
extern "C" {
#endif

/* TeletexOrganizationalUnitNames */
typedef struct TeletexOrganizationalUnitNames {
	A_SEQUENCE_OF(TeletexOrganizationalUnitName_t) list;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} TeletexOrganizationalUnitNames_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_TeletexOrganizationalUnitNames;

#ifdef __cplusplus
}
#endif

#endif	/* _TeletexOrganizationalUnitNames_H_ */
#include <asn_internal.h>
