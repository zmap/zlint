/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "PKIX1Explicit88"
 * 	found in "asn1/rfc5280-PKIX1Explicit88.asn1"
 * 	`asn1c -S /home/fotisl/Projects/revlintsynt/asn1c/skeletons -pdu=all -fwide-types -fcompound-names`
 */

#ifndef	_TeletexOrganizationalUnitName_H_
#define	_TeletexOrganizationalUnitName_H_


#include <asn_application.h>

/* Including external dependencies */
#include <TeletexString.h>

#ifdef __cplusplus
extern "C" {
#endif

/* TeletexOrganizationalUnitName */
typedef TeletexString_t	 TeletexOrganizationalUnitName_t;

/* Implementation */
extern asn_per_constraints_t asn_PER_type_TeletexOrganizationalUnitName_constr_1;
extern asn_TYPE_descriptor_t asn_DEF_TeletexOrganizationalUnitName;
asn_struct_free_f TeletexOrganizationalUnitName_free;
asn_struct_print_f TeletexOrganizationalUnitName_print;
asn_constr_check_f TeletexOrganizationalUnitName_constraint;
ber_type_decoder_f TeletexOrganizationalUnitName_decode_ber;
der_type_encoder_f TeletexOrganizationalUnitName_encode_der;
xer_type_decoder_f TeletexOrganizationalUnitName_decode_xer;
xer_type_encoder_f TeletexOrganizationalUnitName_encode_xer;
oer_type_decoder_f TeletexOrganizationalUnitName_decode_oer;
oer_type_encoder_f TeletexOrganizationalUnitName_encode_oer;
per_type_decoder_f TeletexOrganizationalUnitName_decode_uper;
per_type_encoder_f TeletexOrganizationalUnitName_encode_uper;

#ifdef __cplusplus
}
#endif

#endif	/* _TeletexOrganizationalUnitName_H_ */
#include <asn_internal.h>
