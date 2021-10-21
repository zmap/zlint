include Makefile.am.libasncodec

ASN1C=asn1c

ASN_MODULE_SRCS+=pdu_collection.c

OBJS := $(ASN_MODULE_SRCS:%.c=%.o)

all: libpdu.a

%.o: %.c
	gcc -c -o $@ $< -I.

libpdu.a: $(OBJS)
	rm -f $@
	ar rcs $@ $^
	ranlib $@

regenerate:
	$(ASN1C) -pdu=all -fwide-types -fcompound-names asn1/rfc2560-OCSP.asn1 asn1/rfc3281-PKIXAttributeCertificate.asn1 asn1/rfc3369-AttributeCertificateVersion1.asn1 asn1/rfc3369-CryptographicMessageSyntax.asn1 asn1/rfc3709-LogotypeCertExtn.asn1 asn1/rfc3739-PKIXqualified88.asn1 asn1/rfc3851-SecureMimeMessageV3dot1.asn1 asn1/rfc4055-PKIX1-PSS-OAEP-Algorithms.asn1 asn1/rfc5280-PKIX1Explicit88.asn1 asn1/rfc5280-PKIX1Implicit88.asn1 asn1/rfc5480-PKIX1Algorithms2008.asn1 asn1/rfc7633-TLS-Feature-Module-2015.asn1 asn1/cabfev.asn1 asn1/ct.asn1 asn1/otherattrs.asn1 asn1/sxg.asn1
	rm converter-example.c converter-example.mk

clean:
	rm -f $(OBJS) libpdu.a
