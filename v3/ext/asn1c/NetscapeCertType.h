/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "NetscapeExtensions"
 * 	found in "asn1/netscape.asn1"
 * 	`asn1c -S /home/fotisl/Projects/revlintsynt/asn1c/skeletons -pdu=all -fwide-types -fcompound-names`
 */

#ifndef	_NetscapeCertType_H_
#define	_NetscapeCertType_H_


#include <asn_application.h>

/* Including external dependencies */
#include <BIT_STRING.h>

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum NetscapeCertType {
	NetscapeCertType_sSLClient	= 0,
	NetscapeCertType_sSLServer	= 1,
	NetscapeCertType_sMIME	= 2,
	NetscapeCertType_objectSigning	= 3,
	NetscapeCertType_reserved	= 4,
	NetscapeCertType_sSLCA	= 5,
	NetscapeCertType_sMIMECA	= 6,
	NetscapeCertType_objectSigningCA	= 7
} e_NetscapeCertType;

/* NetscapeCertType */
typedef BIT_STRING_t	 NetscapeCertType_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_NetscapeCertType;
asn_struct_free_f NetscapeCertType_free;
asn_struct_print_f NetscapeCertType_print;
asn_constr_check_f NetscapeCertType_constraint;
ber_type_decoder_f NetscapeCertType_decode_ber;
der_type_encoder_f NetscapeCertType_encode_der;
xer_type_decoder_f NetscapeCertType_decode_xer;
xer_type_encoder_f NetscapeCertType_encode_xer;
oer_type_decoder_f NetscapeCertType_decode_oer;
oer_type_encoder_f NetscapeCertType_encode_oer;
per_type_decoder_f NetscapeCertType_decode_uper;
per_type_encoder_f NetscapeCertType_encode_uper;

#ifdef __cplusplus
}
#endif

#endif	/* _NetscapeCertType_H_ */
#include <asn_internal.h>
