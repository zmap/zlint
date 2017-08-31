package lints

import (
	"testing"
	"time"
)

func TestAllLintsHaveNameDescriptionSource(t *testing.T) {
	for name, lint := range Lints {
		if lint.Name == "" {
			t.Errorf("lint %s has empty name", name)
		}
		if lint.Description == "" {
			t.Errorf("lint %s has empty description", name)
		}
		if lint.Source == "" {
			t.Errorf("lint %s has empty source", name)
		}
	}
}

func TestLintCheckEffective(t *testing.T) {
	l := Lint{}
	c := ReadCertificate("../testlint/testCerts/caBasicConstCrit.pem")

	l.EffectiveDate = time.Time{}
	if l.CheckEffective(c) != true {
		t.Errorf("EffectiveDate of zero should always be true")
	}
	l.EffectiveDate = time.Unix(1, 0)
	if l.CheckEffective(c) != true {
		t.Errorf("EffectiveDate of 1970-01-01 should be true")
	}
	l.EffectiveDate = time.Unix(32503680000, 0) // 3000-01-01
	if l.CheckEffective(c) != false {
		t.Errorf("EffectiveDate of 3000 should be false")
	}
}

func TestLintExecute(t *testing.T) {
	c := ReadCertificate("../testlint/testCerts/goodRsaExp.pem")
	lint := Lint{}

	lint.Lint = &dsaParamsMissing{}
	res := lint.Execute(c)
	if res.Status != NA {
		t.Errorf("Expected NA, got %s", res.Status)
	}

	lint.Lint = &rsaParsedTestsExpBounds{}
	lint.EffectiveDate = time.Unix(32503680000, 0)
	res = lint.Execute(c)
	if res.Status != NE {
		t.Errorf("Expected NE, got %s", res.Status)
	}
}
