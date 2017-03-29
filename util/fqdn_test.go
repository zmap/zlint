package util

import (
	"testing"
)

func TestIsFQDNCorrectFQDN(t *testing.T) {
	domain := "google.com"
	expected := true
	actual := IsFQDN(domain)
	if expected != actual {
		t.Error(
			"For", domain,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestIsFQDNQuestionMarkFQDN(t *testing.T) {
	domain := "?.?.abc.com"
	expected := true
	actual := IsFQDN(domain)
	if expected != actual {
		t.Error(
			"For", domain,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestIsFQDNWildcardFQDN(t *testing.T) {
	domain := "*.abc.com"
	expected := true
	actual := IsFQDN(domain)
	if expected != actual {
		t.Error(
			"For", domain,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestIsFQDNNotFQDN(t *testing.T) {
	domain := "abc"
	expected := false
	actual := IsFQDN(domain)
	if expected != actual {
		t.Error(
			"For", domain,
			"expected", expected,
			"got", actual,
		)
	}
}