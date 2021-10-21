/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "PKIXqualified88"
 * 	found in "asn1/rfc3739-PKIXqualified88.asn1"
 * 	`asn1c -S /home/fotisl/Projects/revlintsynt/asn1c/skeletons -pdu=all -fwide-types -fcompound-names`
 */

#ifndef	_QCStatement_H_
#define	_QCStatement_H_


#include <asn_application.h>

/* Including external dependencies */
#include <OBJECT_IDENTIFIER.h>
#include <ANY.h>
#include <constr_SEQUENCE.h>

#ifdef __cplusplus
extern "C" {
#endif

/* QCStatement */
typedef struct QCStatement {
	OBJECT_IDENTIFIER_t	 statementId;
	ANY_t	*statementInfo	/* OPTIONAL */;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} QCStatement_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_QCStatement;
extern asn_SEQUENCE_specifics_t asn_SPC_QCStatement_specs_1;
extern asn_TYPE_member_t asn_MBR_QCStatement_1[2];

#ifdef __cplusplus
}
#endif

#endif	/* _QCStatement_H_ */
#include <asn_internal.h>
