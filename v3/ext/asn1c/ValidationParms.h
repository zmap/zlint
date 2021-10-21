/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "PKIX1Algorithms2008"
 * 	found in "asn1/rfc5480-PKIX1Algorithms2008.asn1"
 * 	`asn1c -S /home/fotisl/Projects/revlintsynt/asn1c/skeletons -pdu=all -fwide-types -fcompound-names`
 */

#ifndef	_ValidationParms_H_
#define	_ValidationParms_H_


#include <asn_application.h>

/* Including external dependencies */
#include <BIT_STRING.h>
#include <INTEGER.h>
#include <constr_SEQUENCE.h>

#ifdef __cplusplus
extern "C" {
#endif

/* ValidationParms */
typedef struct ValidationParms {
	BIT_STRING_t	 seed;
	INTEGER_t	 pgenCounter;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} ValidationParms_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_ValidationParms;
extern asn_SEQUENCE_specifics_t asn_SPC_ValidationParms_specs_1;
extern asn_TYPE_member_t asn_MBR_ValidationParms_1[2];

#ifdef __cplusplus
}
#endif

#endif	/* _ValidationParms_H_ */
#include <asn_internal.h>
