/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "PKIX1Algorithms2008"
 * 	found in "asn1/rfc5480-PKIX1Algorithms2008.asn1"
 * 	`asn1c -S /home/fotisl/Projects/revlintsynt/asn1c/skeletons -pdu=all -fwide-types -fcompound-names`
 */

#ifndef	_DSAPublicKey_H_
#define	_DSAPublicKey_H_


#include <asn_application.h>

/* Including external dependencies */
#include <INTEGER.h>

#ifdef __cplusplus
extern "C" {
#endif

/* DSAPublicKey */
typedef INTEGER_t	 DSAPublicKey_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_DSAPublicKey;
asn_struct_free_f DSAPublicKey_free;
asn_struct_print_f DSAPublicKey_print;
asn_constr_check_f DSAPublicKey_constraint;
ber_type_decoder_f DSAPublicKey_decode_ber;
der_type_encoder_f DSAPublicKey_encode_der;
xer_type_decoder_f DSAPublicKey_decode_xer;
xer_type_encoder_f DSAPublicKey_encode_xer;
oer_type_decoder_f DSAPublicKey_decode_oer;
oer_type_encoder_f DSAPublicKey_encode_oer;
per_type_decoder_f DSAPublicKey_decode_uper;
per_type_encoder_f DSAPublicKey_encode_uper;

#ifdef __cplusplus
}
#endif

#endif	/* _DSAPublicKey_H_ */
#include <asn_internal.h>
