package lints

/*
 * ZLint Copyright 2018 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

import (
	"testing"
)

func TestSubCertKeyUsageNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/keyUsageNotCriticalSubCert.pem"
	expected := Warn
	out := Lints["w_ext_key_usage_not_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCaKeyUsageNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/caKeyUsageNotCrit.pem"
	expected := Warn
	out := Lints["w_ext_key_usage_not_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCertKeyUsageCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/domainValGoodSubject.pem"
	expected := Pass
	out := Lints["w_ext_key_usage_not_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCaKeyUsageCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/caKeyUsageCrit.pem"
	expected := Pass
	out := Lints["w_ext_key_usage_not_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCertKeyUsageNotIncludedCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/caKeyUsageMissing.pem"
	expected := NA
	out := Lints["e_ext_key_usage_without_bits"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
