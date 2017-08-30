// lint_distribution_point_incomplete.go
/********************************************************************
The cRLDistributionPoints extension is a SEQUENCE of
DistributionPoint.  A DistributionPoint consists of three fields,
each of which is optional: distributionPoint, reasons, and cRLIssuer.
While each of these fields is optional, a DistributionPoint MUST NOT
consist of only the reasons field; either distributionPoint or
cRLIssuer MUST be present.  If the certificate issuer is not the CRL
issuer, then the cRLIssuer field MUST be present and contain the Name
of the CRL issuer.  If the certificate issuer is also the CRL issuer,
then conforming CAs MUST omit the cRLIssuer field and MUST include
the distributionPoint field.
********************************************************************/

package lints

import (
	"encoding/asn1"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zcrypto/x509/pkix"
	"github.com/zmap/zlint/util"
)

type distributionPoint struct {
	DistributionPoint distributionPointName `asn1:"optional,tag:0"`
	Reason            asn1.BitString        `asn1:"optional,tag:1"`
	CRLIssuer         asn1.RawValue         `asn1:"optional,tag:2"`
}

type distributionPointName struct {
	FullName     asn1.RawValue    `asn1:"optional,tag:0"`
	RelativeName pkix.RDNSequence `asn1:"optional,tag:1"`
}

type dpIncomplete struct {
	// Internal data here
}

func (l *dpIncomplete) Initialize() error {
	return nil
}

func (l *dpIncomplete) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.CrlDistOID)
}

func (l *dpIncomplete) RunTest(c *x509.Certificate) (ResultStruct, error) {
	dp := util.GetExtFromCert(c, util.CrlDistOID)
	var cdp []distributionPoint
	_, err := asn1.Unmarshal(dp.Value, &cdp)
	if err != nil {
		return ResultStruct{Result: Fatal}, nil
	}
	for _, dp := range cdp {
		if dp.Reason.BitLength != 0 && len(dp.DistributionPoint.FullName.Bytes) == 0 &&
			dp.DistributionPoint.RelativeName == nil && len(dp.CRLIssuer.Bytes) == 0 {
			return ResultStruct{Result: Error}, nil
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_distribution_point_incomplete",
		Description:   "A DistributionPoint from the CRLDistributionPoints extension MUST NOT consist of only the reasons field; either distributionPoint or CRLIssuer must be present",
		Source:        "RFC 5280: 4.2.1.13",
		EffectiveDate: util.RFC3280Date,
		Test:          &dpIncomplete{},
	})
}
