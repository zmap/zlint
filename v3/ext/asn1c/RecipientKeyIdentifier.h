/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "CryptographicMessageSyntax"
 * 	found in "asn1/rfc3369-CryptographicMessageSyntax.asn1"
 * 	`asn1c -S /home/fotisl/Projects/revlintsynt/asn1c/skeletons -pdu=all -fwide-types -fcompound-names`
 */

#ifndef	_RecipientKeyIdentifier_H_
#define	_RecipientKeyIdentifier_H_


#include <asn_application.h>

/* Including external dependencies */
#include "SubjectKeyIdentifier.h"
#include <GeneralizedTime.h>
#include <constr_SEQUENCE.h>

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct OtherKeyAttribute;

/* RecipientKeyIdentifier */
typedef struct RecipientKeyIdentifier {
	SubjectKeyIdentifier_t	 subjectKeyIdentifier;
	GeneralizedTime_t	*date	/* OPTIONAL */;
	struct OtherKeyAttribute	*other	/* OPTIONAL */;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} RecipientKeyIdentifier_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_RecipientKeyIdentifier;
extern asn_SEQUENCE_specifics_t asn_SPC_RecipientKeyIdentifier_specs_1;
extern asn_TYPE_member_t asn_MBR_RecipientKeyIdentifier_1[3];

#ifdef __cplusplus
}
#endif

/* Referred external types */
#include "OtherKeyAttribute.h"

#endif	/* _RecipientKeyIdentifier_H_ */
#include <asn_internal.h>
