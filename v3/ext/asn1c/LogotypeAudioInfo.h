/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "LogotypeCertExtn"
 * 	found in "asn1/rfc3709-LogotypeCertExtn.asn1"
 * 	`asn1c -S /home/fotisl/Projects/revlintsynt/asn1c/skeletons -pdu=all -fwide-types -fcompound-names`
 */

#ifndef	_LogotypeAudioInfo_H_
#define	_LogotypeAudioInfo_H_


#include <asn_application.h>

/* Including external dependencies */
#include <INTEGER.h>
#include <IA5String.h>
#include <constr_SEQUENCE.h>

#ifdef __cplusplus
extern "C" {
#endif

/* LogotypeAudioInfo */
typedef struct LogotypeAudioInfo {
	INTEGER_t	 fileSize;
	INTEGER_t	 playTime;
	INTEGER_t	 channels;
	INTEGER_t	*sampleRate	/* OPTIONAL */;
	IA5String_t	*language	/* OPTIONAL */;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} LogotypeAudioInfo_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_LogotypeAudioInfo;
extern asn_SEQUENCE_specifics_t asn_SPC_LogotypeAudioInfo_specs_1;
extern asn_TYPE_member_t asn_MBR_LogotypeAudioInfo_1[5];

#ifdef __cplusplus
}
#endif

#endif	/* _LogotypeAudioInfo_H_ */
#include <asn_internal.h>
