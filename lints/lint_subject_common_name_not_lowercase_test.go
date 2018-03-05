package lints

import "testing"

func TestCnNotCaplitalised(t *testing.T) {
	inputPath := "../testlint/testCerts/SANRegisteredIdBeginning.pem"
	expected := Pass
	out := Lints["n_lint_subject_common_name_not_lowercase"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCnCapitalised(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectDNCommonNameCapitalised.pem"
	expected := Notice
	out := Lints["n_lint_subject_common_name_not_lowercase"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
