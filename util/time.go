package util

import (
	"encoding/asn1"
	"time"

	"github.com/zmap/zcrypto/x509"
)

var (
	ZeroDate                   = time.Date(0000, time.January, 1, 0, 0, 0, 0, time.UTC)
	RFC1035Date                = time.Date(1987, time.January, 1, 0, 0, 0, 0, time.UTC)
	RFC2459Date                = time.Date(1999, time.January, 1, 0, 0, 0, 0, time.UTC)
	RFC3280Date                = time.Date(2002, time.April, 1, 0, 0, 0, 0, time.UTC)
	RFC3490Date                = time.Date(2003, time.March, 1, 0, 0, 0, 0, time.UTC)
	RFC4325Date                = time.Date(2005, time.December, 1, 0, 0, 0, 0, time.UTC)
	RFC4630Date                = time.Date(2006, time.August, 1, 0, 0, 0, 0, time.UTC)
	RFC5280Date                = time.Date(2008, time.May, 1, 0, 0, 0, 0, time.UTC)
	RFC6818Date                = time.Date(2013, time.January, 1, 0, 0, 0, 0, time.UTC)
	CABEffectiveDate           = time.Date(2012, time.July, 1, 0, 0, 0, 0, time.UTC)
	CABSerialNumberEntropyDate = time.Date(2016, time.September, 30, 0, 0, 0, 0, time.UTC)
	CABV102Date                = time.Date(2012, time.June, 8, 0, 0, 0, 0, time.UTC)
	CABV113Date                = time.Date(2013, time.February, 21, 0, 0, 0, 0, time.UTC)
	CABV114Date                = time.Date(2013, time.May, 3, 0, 0, 0, 0, time.UTC)
	CABV116Date                = time.Date(2013, time.July, 29, 0, 0, 0, 0, time.UTC)
	CABV130Date                = time.Date(2015, time.April, 16, 0, 0, 0, 0, time.UTC)
	CABV131Date                = time.Date(2015, time.September, 28, 0, 0, 0, 0, time.UTC)
	NO_SHA1                    = time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC)
	NoRSA1024RootDate          = time.Date(2011, time.January, 1, 0, 0, 0, 0, time.UTC)
	NoRSA1024Date              = time.Date(2014, time.January, 1, 0, 0, 0, 0, time.UTC)
	GeneralizedDate            = time.Date(2050, time.January, 1, 0, 0, 0, 0, time.UTC)
	NoReservedIP               = time.Date(2015, time.November, 1, 0, 0, 0, 0, time.UTC)
	SubCert39Month             = time.Date(2016, time.June, 30, 0, 0, 0, 0, time.UTC)
)

func FindTimeType(firstDate, secondDate asn1.RawValue) (int, int) {
	return firstDate.Tag, secondDate.Tag
}

func GetTimes(cert *x509.Certificate) (asn1.RawValue, asn1.RawValue) {
	var outSeq, firstDate, secondDate asn1.RawValue
	// Unmarshal into the sequence
	rest, err := asn1.Unmarshal(cert.RawTBSCertificate, &outSeq)
	// Start unmarshalling the bytes
	rest, err = asn1.Unmarshal(outSeq.Bytes, &outSeq)
	// This is here to account for if version is not included
	if outSeq.Tag == 0 {
		rest, err = asn1.Unmarshal(rest, &outSeq)
	}
	rest, err = asn1.Unmarshal(rest, &outSeq)
	rest, err = asn1.Unmarshal(rest, &outSeq)
	rest, err = asn1.Unmarshal(rest, &outSeq)
	//Finally at the validity date, load them into a different RawValue
	rest, err = asn1.Unmarshal(outSeq.Bytes, &firstDate)
	rest, err = asn1.Unmarshal(rest, &secondDate)
	if err != nil {
		//fmt.Println(err)
		return asn1.RawValue{}, asn1.RawValue{}
	}
	return firstDate, secondDate
}
