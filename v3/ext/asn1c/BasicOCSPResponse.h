/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "OCSP"
 * 	found in "asn1/rfc2560-OCSP.asn1"
 * 	`asn1c -S /home/fotisl/Projects/revlintsynt/asn1c/skeletons -pdu=all -fwide-types -fcompound-names`
 */

#ifndef	_BasicOCSPResponse_H_
#define	_BasicOCSPResponse_H_


#include <asn_application.h>

/* Including external dependencies */
#include "ResponseData.h"
#include "AlgorithmIdentifier.h"
#include <BIT_STRING.h>
#include <asn_SEQUENCE_OF.h>
#include <constr_SEQUENCE_OF.h>
#include <constr_SEQUENCE.h>

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct Certificate;

/* BasicOCSPResponse */
typedef struct BasicOCSPResponse {
	ResponseData_t	 tbsResponseData;
	AlgorithmIdentifier_t	 signatureAlgorithm;
	BIT_STRING_t	 signature;
	struct BasicOCSPResponse__certs {
		A_SEQUENCE_OF(struct Certificate) list;
		
		/* Context for parsing across buffer boundaries */
		asn_struct_ctx_t _asn_ctx;
	} *certs;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} BasicOCSPResponse_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_BasicOCSPResponse;

#ifdef __cplusplus
}
#endif

/* Referred external types */
#include "Certificate.h"

#endif	/* _BasicOCSPResponse_H_ */
#include <asn_internal.h>
