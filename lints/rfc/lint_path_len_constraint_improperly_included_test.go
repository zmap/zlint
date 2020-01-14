package rfc

/*
 * ZLint Copyright 2020 Regents of the University of Michigan
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

	"github.com/zmap/zlint/lint"
	"github.com/zmap/zlint/util"
)

func TestCaMaxLenPresentNoCertSign(t *testing.T) {
	inputPath := "../../testlint/testCerts/caMaxPathLenPresentNoCertSign.pem"
	expected := lint.Error
	out := lint.Lints["e_path_len_constraint_improperly_included"].Execute(util.ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCaMaxLenPresentGood(t *testing.T) {
	inputPath := "../../testlint/testCerts/caMaxPathLenPositive.pem"
	expected := lint.Pass
	out := lint.Lints["e_path_len_constraint_improperly_included"].Execute(util.ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCaMaxLenMissing(t *testing.T) {
	inputPath := "../../testlint/testCerts/caMaxPathLenMissing.pem"
	expected := lint.Pass
	out := lint.Lints["e_path_len_constraint_improperly_included"].Execute(util.ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCertMaxLenPresent(t *testing.T) {
	inputPath := "../../testlint/testCerts/subCertPathLenPositive.pem"
	expected := lint.Error
	out := lint.Lints["e_path_len_constraint_improperly_included"].Execute(util.ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCertMaxLenNone(t *testing.T) {
	inputPath := "../../testlint/testCerts/orgValGoodAllFields.pem"
	expected := lint.Pass
	out := lint.Lints["e_path_len_constraint_improperly_included"].Execute(util.ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
