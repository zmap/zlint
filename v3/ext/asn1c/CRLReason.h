/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "PKIX1Implicit88"
 * 	found in "asn1/rfc5280-PKIX1Implicit88.asn1"
 * 	`asn1c -S /home/fotisl/Projects/revlintsynt/asn1c/skeletons -pdu=all -fwide-types -fcompound-names`
 */

#ifndef	_CRLReason_H_
#define	_CRLReason_H_


#include <asn_application.h>

/* Including external dependencies */
#include <ENUMERATED.h>

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum CRLReason {
	CRLReason_unspecified	= 0,
	CRLReason_keyCompromise	= 1,
	CRLReason_cACompromise	= 2,
	CRLReason_affiliationChanged	= 3,
	CRLReason_superseded	= 4,
	CRLReason_cessationOfOperation	= 5,
	CRLReason_certificateHold	= 6,
	CRLReason_removeFromCRL	= 8,
	CRLReason_privilegeWithdrawn	= 9,
	CRLReason_aACompromise	= 10
} e_CRLReason;

/* CRLReason */
typedef ENUMERATED_t	 CRLReason_t;

/* Implementation */
extern asn_per_constraints_t asn_PER_type_CRLReason_constr_1;
extern asn_TYPE_descriptor_t asn_DEF_CRLReason;
extern const asn_INTEGER_specifics_t asn_SPC_CRLReason_specs_1;
asn_struct_free_f CRLReason_free;
asn_struct_print_f CRLReason_print;
asn_constr_check_f CRLReason_constraint;
ber_type_decoder_f CRLReason_decode_ber;
der_type_encoder_f CRLReason_encode_der;
xer_type_decoder_f CRLReason_decode_xer;
xer_type_encoder_f CRLReason_encode_xer;
oer_type_decoder_f CRLReason_decode_oer;
oer_type_encoder_f CRLReason_encode_oer;
per_type_decoder_f CRLReason_decode_uper;
per_type_encoder_f CRLReason_encode_uper;

#ifdef __cplusplus
}
#endif

#endif	/* _CRLReason_H_ */
#include <asn_internal.h>
