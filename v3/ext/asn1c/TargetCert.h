/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "PKIXAttributeCertificate"
 * 	found in "asn1/rfc3281-PKIXAttributeCertificate.asn1"
 * 	`asn1c -S /home/fotisl/Projects/revlintsynt/asn1c/skeletons -pdu=all -fwide-types -fcompound-names`
 */

#ifndef	_TargetCert_H_
#define	_TargetCert_H_


#include <asn_application.h>

/* Including external dependencies */
#include "IssuerSerial.h"
#include <constr_SEQUENCE.h>

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct GeneralName;
struct ObjectDigestInfo;

/* TargetCert */
typedef struct TargetCert {
	IssuerSerial_t	 targetCertificate;
	struct GeneralName	*targetName	/* OPTIONAL */;
	struct ObjectDigestInfo	*certDigestInfo	/* OPTIONAL */;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} TargetCert_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_TargetCert;
extern asn_SEQUENCE_specifics_t asn_SPC_TargetCert_specs_1;
extern asn_TYPE_member_t asn_MBR_TargetCert_1[3];

#ifdef __cplusplus
}
#endif

/* Referred external types */
#include "GeneralName.h"
#include "ObjectDigestInfo.h"

#endif	/* _TargetCert_H_ */
#include <asn_internal.h>
