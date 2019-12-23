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

func TestGivenNameCorrectPolicy(t *testing.T) {
	inputPath := "../testlint/testCerts/givenNameCorrectPolicy.pem"
	expected := Pass
	out := Lints["e_sub_cert_given_name_surname_contains_correct_policy"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSurnameCorrectPolicy(t *testing.T) {
	inputPath := "../testlint/testCerts/surnameCorrectPolicy.pem"
	expected := Pass
	out := Lints["e_sub_cert_given_name_surname_contains_correct_policy"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestGivenNameIncorrectPolicy(t *testing.T) {
	inputPath := "../testlint/testCerts/givenNameIncorrectPolicy.pem"
	expected := Error
	out := Lints["e_sub_cert_given_name_surname_contains_correct_policy"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSurnameIncorrectPolicy(t *testing.T) {
	inputPath := "../testlint/testCerts/surnameIncorrectPolicy.pem"
	expected := Error
	out := Lints["e_sub_cert_given_name_surname_contains_correct_policy"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
