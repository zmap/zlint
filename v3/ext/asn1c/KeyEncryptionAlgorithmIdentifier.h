/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "CryptographicMessageSyntax"
 * 	found in "asn1/rfc3369-CryptographicMessageSyntax.asn1"
 * 	`asn1c -S /home/fotisl/Projects/revlintsynt/asn1c/skeletons -pdu=all -fwide-types -fcompound-names`
 */

#ifndef	_KeyEncryptionAlgorithmIdentifier_H_
#define	_KeyEncryptionAlgorithmIdentifier_H_


#include <asn_application.h>

/* Including external dependencies */
#include "AlgorithmIdentifier.h"

#ifdef __cplusplus
extern "C" {
#endif

/* KeyEncryptionAlgorithmIdentifier */
typedef AlgorithmIdentifier_t	 KeyEncryptionAlgorithmIdentifier_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_KeyEncryptionAlgorithmIdentifier;
asn_struct_free_f KeyEncryptionAlgorithmIdentifier_free;
asn_struct_print_f KeyEncryptionAlgorithmIdentifier_print;
asn_constr_check_f KeyEncryptionAlgorithmIdentifier_constraint;
ber_type_decoder_f KeyEncryptionAlgorithmIdentifier_decode_ber;
der_type_encoder_f KeyEncryptionAlgorithmIdentifier_encode_der;
xer_type_decoder_f KeyEncryptionAlgorithmIdentifier_decode_xer;
xer_type_encoder_f KeyEncryptionAlgorithmIdentifier_encode_xer;
oer_type_decoder_f KeyEncryptionAlgorithmIdentifier_decode_oer;
oer_type_encoder_f KeyEncryptionAlgorithmIdentifier_encode_oer;
per_type_decoder_f KeyEncryptionAlgorithmIdentifier_decode_uper;
per_type_encoder_f KeyEncryptionAlgorithmIdentifier_encode_uper;

#ifdef __cplusplus
}
#endif

#endif	/* _KeyEncryptionAlgorithmIdentifier_H_ */
#include <asn_internal.h>
