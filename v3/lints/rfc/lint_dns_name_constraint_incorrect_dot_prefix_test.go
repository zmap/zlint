package rfc

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestDotPrefixedDNSNameConstraints(t *testing.T) {
	data := []struct {
		file string
		want lint.LintStatus
	}{
		{
			"dnsNameConstraintExcludedWithDotPrefix.pem", lint.Error,
		},
		{
			"dnsNameConstraintPermittedWithDotPrefix.pem", lint.Error,
		},
		{
			"dnsNameConstraintExcludedAndPermittedWithDotPrefix.pem", lint.Error,
		},
		{
			"dnsNameConstraintExcludedAndPermittedWithoutDotPrefix.pem", lint.Pass,
		},
		{
			"dnsNameConstraintNotApplicableDotPrefix.pem", lint.NA,
		},
	}
	for _, d := range data {
		out := test.TestLint("e_dns_name_constraint_incorrect_dot_prefix", d.file)
		if out.Status != d.want {
			t.Errorf("%s: expected %s, got %s", d.file, d.want, out.Status)
		}
	}
}
