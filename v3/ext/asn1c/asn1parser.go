// Package asn1parser provides an ASN.1 parser generated using asn1c from the
// relevant ASN.1 modules.
package asn1parser

// #cgo LDFLAGS: -L. -lpdu
/*
#include <stdlib.h>
#include <stdbool.h>
#include "asn_application.h"
#include "Certificate.h"

struct asn1vretval {
	bool valid;
	char *errbuf;
	// This is messy, but the easiest way to do it.
	void *pdu;
};

extern asn_TYPE_descriptor_t *asn_pdu_collection[];

char *err_invalid_pdu = "invalid pdu";
char *err_cannot_decode = "cannot decode pdu";
char *err_check_constraints = "cannot validate constraints";

asn_TYPE_descriptor_t *lookup_type(char *name) {
	asn_TYPE_descriptor_t **p = asn_pdu_collection;

	while(*p) {
		if(!strcmp((*p)->name, name))
			return *p;
		p++;
	}

	return NULL;
}

struct asn1vretval decode_pdu(char *pdu_name, void *data, size_t len) {
	asn_dec_rval_t rv;
	asn_TYPE_descriptor_t *pdutype;
	struct asn1vretval retval;
	void *pdustruct = NULL;

	retval.valid = false;

	if((pdutype = lookup_type(pdu_name)) == NULL) {
		retval.errbuf = strdup(err_invalid_pdu);
		return retval;
	}

	rv = ber_decode(0, pdutype, &pdustruct, data, len);
	if(rv.code != RC_OK) {
		retval.errbuf = strdup(err_cannot_decode);
		return retval;
	}

	// TODO: remove 1 line
	//xer_fprint(stdout, pdutype, pdustruct);

	retval.valid = true;
	retval.pdu = pdustruct;

	return retval;
}

struct asn1vretval check_constraints(char *pdu_name, void *pdustruct) {
	int ret;
	char errbuf[128];
	size_t errlen = sizeof(errbuf);
	struct asn1vretval retval;
	asn_TYPE_descriptor_t *pdutype;

	retval.valid = false;

	if((pdutype = lookup_type(pdu_name)) == NULL) {
		retval.errbuf = strdup(err_invalid_pdu);
		return retval;
	}

	//TODO: remove 1 line
	//errbuf[0] = 0;

	ret = asn_check_constraints(pdutype, pdustruct, errbuf, &errlen);

	// TODO: remove 1 line
	//printf("error=%s\n", errbuf);

	if(ret) {
		retval.errbuf = strdup(err_check_constraints);
		return retval;
	}

	retval.valid = true;
	return retval;
}

void free_pdu(char *pdu_name, void *pdustruct) {
	asn_TYPE_descriptor_t *pdutype;

	if((pdutype = lookup_type(pdu_name)) == NULL) {
		return;
	}

	ASN_STRUCT_FREE(*pdutype, pdustruct);
}

*/
import "C"
import (
	"crypto/x509/pkix"
	"encoding/asn1"
	"errors"
	"unsafe"
)

// DecodePdu decodes a PDU named `name` stored in `data`. Note that an
// unsafe.Pointer is returned, which should only be used with the rest of the
// functions in this package. In addition, since the memory is unmanaged, it
// must be freed with FreePdu.
func DecodePdu(name string, data []byte) (unsafe.Pointer, error) {
	cname := C.CString(name)
	cdata := C.CBytes(data)

	r := C.decode_pdu(cname, cdata, C.size_t(len(data)))
	C.free(unsafe.Pointer(cname))
	C.free(cdata)

	if !r.valid {
		errstr := C.GoString(r.errbuf)
		C.free(unsafe.Pointer(r.errbuf))
		return nil, errors.New(errstr)
	}

	return r.pdu, nil
}

// CheckConstraints verifies the constrains of a decoded PDU based on the ASN.1
// module.
func CheckConstraints(name string, pdu unsafe.Pointer) error {
	cname := C.CString(name)

	r := C.check_constraints(cname, pdu)
	C.free(unsafe.Pointer(cname))

	if !r.valid {
		errstr := C.GoString(r.errbuf)
		C.free(unsafe.Pointer(r.errbuf))
		return errors.New(errstr)
	}

	return nil
}

// FreePdu releases all memory held by a decoded PDU.
func FreePdu(name string, pdu unsafe.Pointer) {
	cname := C.CString(name)

	C.free_pdu(cname, pdu)
	C.free(unsafe.Pointer(cname))
}

func parseAsn1cOID(oid C.struct_ASN__PRIMITIVE_TYPE_s) (asn1.ObjectIdentifier, error) {
	oidv := C.GoBytes((unsafe.Pointer)(oid.buf), C.int(oid.size))
	oidf := []byte{0x06, byte(oid.size)}
	oidf = append(oidf, oidv...)

	var ret asn1.ObjectIdentifier
	_, err := asn1.Unmarshal(oidf, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetCertExtensions returns all the extensions from a Certificate PDU. Since
// the PDU is an unsafe.Pointer, you need to ensure that it is a Certificate.
func GetCertExtensions(pdu unsafe.Pointer) ([]pkix.Extension, error) {
	var exts []pkix.Extension

	cert := (*C.struct_Certificate)(pdu)

	if cert.tbsCertificate.extensions == nil || cert.tbsCertificate.extensions.list.count == 0 {
		return nil, nil
	}

	cexts := (*[1 << 30]*C.struct_Extension)(unsafe.Pointer(cert.tbsCertificate.extensions.list.array))[:cert.tbsCertificate.extensions.list.count:cert.tbsCertificate.extensions.list.count]
	for _, cext := range cexts {
		var ext pkix.Extension

		oid, err := parseAsn1cOID(cext.extnID)
		if err != nil {
			return nil, err
		}
		ext.Id = oid

		if int(cext.critical) > 0 {
			ext.Critical = true
		}

		ext.Value = C.GoBytes((unsafe.Pointer)(cext.extnValue.buf), C.int(cext.extnValue.size))

		exts = append(exts, ext)
	}

	return exts, nil
}

func getNameATV(name C.struct_Name) ([]pkix.AttributeTypeAndValue, error) {
	var atvs []pkix.AttributeTypeAndValue

	rdnSeq := (*C.struct_RelativeDistinguishedName)(unsafe.Pointer(&name.choice[0]))

	if rdnSeq.list.count == 0 {
		return nil, nil
	}

	rdnSeqSlice := (*[1 << 30]*C.struct_RelativeDistinguishedName)(unsafe.Pointer(rdnSeq.list.array))[:rdnSeq.list.count:rdnSeq.list.count]

	for _, rdn := range rdnSeqSlice {
		atvSlice := (*[1 << 30]*C.struct_AttributeTypeAndValue)(unsafe.Pointer(rdn.list.array))[:rdn.list.count:rdn.list.count]
		for _, atv := range atvSlice {
			var gatv pkix.AttributeTypeAndValue

			oid, err := parseAsn1cOID(atv._type)
			if err != nil {
				return nil, err
			}

			gatv.Type = oid
			gatv.Value = C.GoBytes((unsafe.Pointer)(atv.value.buf), C.int(atv.value.size))
			atvs = append(atvs, gatv)
		}
	}

	return atvs, nil
}

// GetSubjectATV returns all attribute value pairs from the subject of a
// Certificate PDU.
func GetSubjectATV(pdu unsafe.Pointer) ([]pkix.AttributeTypeAndValue, error) {
	cert := (*C.struct_Certificate)(pdu)
	return getNameATV(cert.tbsCertificate.subject)
}

// GetIssuerATV returns all attribute value pairs from the issuer of a
// Certificate PDU.
func GetIssuerATV(pdu unsafe.Pointer) ([]pkix.AttributeTypeAndValue, error) {
	cert := (*C.struct_Certificate)(pdu)
	return getNameATV(cert.tbsCertificate.issuer)
}

// GetSubjectPublicKeyInfo returns the algorithm, the params and the public key
// of a Certificate PDU.
func GetSubjectPublicKeyInfo(pdu unsafe.Pointer) (asn1.ObjectIdentifier, []byte, []byte, error) {
	cert := (*C.struct_Certificate)(pdu)
	spki := cert.tbsCertificate.subjectPublicKeyInfo

	algoid, err := parseAsn1cOID(spki.algorithm.algorithm)
	if err != nil {
		return nil, nil, nil, err
	}

	var params, key []byte
	if spki.algorithm.parameters != nil && spki.algorithm.parameters.buf != nil {
		params = C.GoBytes((unsafe.Pointer)(spki.algorithm.parameters.buf), spki.algorithm.parameters.size)
	}
	if spki.subjectPublicKey.buf != nil {
		key = C.GoBytes((unsafe.Pointer)(spki.subjectPublicKey.buf), C.int(spki.subjectPublicKey.size))
	}

	return algoid, params, key, nil
}
