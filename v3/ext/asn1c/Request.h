/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "OCSP"
 * 	found in "asn1/rfc2560-OCSP.asn1"
 * 	`asn1c -S /home/fotisl/Projects/revlintsynt/asn1c/skeletons -pdu=all -fwide-types -fcompound-names`
 */

#ifndef	_Request_H_
#define	_Request_H_


#include <asn_application.h>

/* Including external dependencies */
#include "CertID.h"
#include <constr_SEQUENCE.h>

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct Extensions;

/* Request */
typedef struct Request {
	CertID_t	 reqCert;
	struct Extensions	*singleRequestExtensions	/* OPTIONAL */;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} Request_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_Request;
extern asn_SEQUENCE_specifics_t asn_SPC_Request_specs_1;
extern asn_TYPE_member_t asn_MBR_Request_1[2];

#ifdef __cplusplus
}
#endif

/* Referred external types */
#include "Extensions.h"

#endif	/* _Request_H_ */
#include <asn_internal.h>
