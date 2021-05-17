package rfc

/*
 * ZLint Copyright 2021 Regents of the University of Michigan
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

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestCaMaxLenNegative(t *testing.T) {
	inputPath := "caMaxPathNegative.pem"
	expected := lint.Error
	out := test.TestLint("e_path_len_constraint_zero_or_less", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCerMaxLenNegative(t *testing.T) {
	inputPath := "subCertPathLenNegative.pem"
	expected := lint.Error
	out := test.TestLint("e_path_len_constraint_zero_or_less", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCaMaxLenPositive(t *testing.T) {
	inputPath := "caMaxPathLenPositive.pem"
	expected := lint.Pass
	out := test.TestLint("e_path_len_constraint_zero_or_less", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCertMaxLenPositive(t *testing.T) {
	inputPath := "subCertPathLenPositive.pem"
	expected := lint.Pass
	out := test.TestLint("e_path_len_constraint_zero_or_less", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCertMaxLenMissing(t *testing.T) {
	inputPath := "caBasicConstMissing.pem"
	expected := lint.NA
	out := test.TestLint("e_path_len_constraint_zero_or_less", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCAMaxLenNone(t *testing.T) {
	inputPath := "caMaxPathLenMissing.pem"
	expected := lint.Pass
	out := test.TestLint("e_path_len_constraint_zero_or_less", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
