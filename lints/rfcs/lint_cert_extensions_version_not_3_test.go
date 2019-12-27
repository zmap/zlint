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

func TestExtsV2(t *testing.T) {
	inputPath := "../testlint/testCerts/certVersion2WithExtension.pem"
	expected := Error
	out := Lints["e_cert_extensions_version_not_3"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestExtsV3(t *testing.T) {
	inputPath := "../testlint/testCerts/caBasicConstCrit.pem"
	expected := Pass
	out := Lints["e_cert_extensions_version_not_3"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestNoExtsV2(t *testing.T) {
	inputPath := "../testlint/testCerts/certVersion2NoExtensions.pem"
	expected := Pass
	out := Lints["e_cert_extensions_version_not_3"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
