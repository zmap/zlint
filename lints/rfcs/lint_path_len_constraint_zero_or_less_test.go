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

func TestCaMaxLenNegative(t *testing.T) {
	inputPath := "../testlint/testCerts/caMaxPathNegative.pem"
	expected := Error
	out := Lints["e_path_len_constraint_zero_or_less"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCerMaxLenNegative(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertPathLenNegative.pem"
	expected := Error
	out := Lints["e_path_len_constraint_zero_or_less"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCaMaxLenPositive(t *testing.T) {
	inputPath := "../testlint/testCerts/caMaxPathLenPositive.pem"
	expected := Pass
	out := Lints["e_path_len_constraint_zero_or_less"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCertMaxLenPositive(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertPathLenPositive.pem"
	expected := Pass
	out := Lints["e_path_len_constraint_zero_or_less"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCertMaxLenMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/caBasicConstMissing.pem"
	expected := NA
	out := Lints["e_path_len_constraint_zero_or_less"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCAMaxLenNone(t *testing.T) {
	inputPath := "../testlint/testCerts/caMaxPathLenMissing.pem"
	expected := Pass
	out := Lints["e_path_len_constraint_zero_or_less"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
