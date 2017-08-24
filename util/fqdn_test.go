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

func TestIsFQDNQuestionMarkIncorrectPlaceFQDN(t *testing.T) {
	domain := "?.?.abc?.com"
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

func TestIsFQDNManyQuestionMarksFQDN(t *testing.T) {
	domain := "?.?.?.?.?.?.?.abc.com"
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

func TestIsFQDNIDNA(t *testing.T) {
	domain := "www.2.xn--80abixftle8gl7a.xn--p1ai"
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
