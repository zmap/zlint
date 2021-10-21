/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "PKIX1Implicit88"
 * 	found in "asn1/rfc5280-PKIX1Implicit88.asn1"
 * 	`asn1c -S /home/fotisl/Projects/revlintsynt/asn1c/skeletons -pdu=all -fwide-types -fcompound-names`
 */

#ifndef	_ReasonFlags_H_
#define	_ReasonFlags_H_


#include <asn_application.h>

/* Including external dependencies */
#include <BIT_STRING.h>

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum ReasonFlags {
	ReasonFlags_unused	= 0,
	ReasonFlags_keyCompromise	= 1,
	ReasonFlags_cACompromise	= 2,
	ReasonFlags_affiliationChanged	= 3,
	ReasonFlags_superseded	= 4,
	ReasonFlags_cessationOfOperation	= 5,
	ReasonFlags_certificateHold	= 6,
	ReasonFlags_privilegeWithdrawn	= 7,
	ReasonFlags_aACompromise	= 8
} e_ReasonFlags;

/* ReasonFlags */
typedef BIT_STRING_t	 ReasonFlags_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_ReasonFlags;
asn_struct_free_f ReasonFlags_free;
asn_struct_print_f ReasonFlags_print;
asn_constr_check_f ReasonFlags_constraint;
ber_type_decoder_f ReasonFlags_decode_ber;
der_type_encoder_f ReasonFlags_encode_der;
xer_type_decoder_f ReasonFlags_decode_xer;
xer_type_encoder_f ReasonFlags_encode_xer;
oer_type_decoder_f ReasonFlags_decode_oer;
oer_type_encoder_f ReasonFlags_encode_oer;
per_type_decoder_f ReasonFlags_decode_uper;
per_type_encoder_f ReasonFlags_encode_uper;

#ifdef __cplusplus
}
#endif

#endif	/* _ReasonFlags_H_ */
#include <asn_internal.h>
