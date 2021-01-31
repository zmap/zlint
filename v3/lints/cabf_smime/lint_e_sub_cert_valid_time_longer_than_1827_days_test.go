package cabf_smime

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

<<<<<<< HEAD
func TestSubCertValidTimeLongerThan1827Days(t *testing.T) {
=======
func TestSubCertValidTimeLongerThan825Days(t *testing.T) {
>>>>>>> b046d8654b4cab9c71aafc84b73c765c793560be
	inputPath := "subCertOver825DaysBad.pem"
	expected := lint.Error
	out := test.TestLint("e_sub_cert_valid_time_longer_than_825_days", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

<<<<<<< HEAD
func TestSubCertValidTimeLongerThan1827DaysBeforeCutoff(t *testing.T) {
=======
func TestSubCertValidTimeLongerThan825DaysBeforeCutoff(t *testing.T) {
>>>>>>> b046d8654b4cab9c71aafc84b73c765c793560be
	inputPath := "subCertOver825DaysOK.pem"
	expected := lint.NE
	out := test.TestLint("e_sub_cert_valid_time_longer_than_825_days", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

<<<<<<< HEAD
func TestSubCertValidTime1827Days(t *testing.T) {
=======
func TestSubCertValidTime825Days(t *testing.T) {
>>>>>>> b046d8654b4cab9c71aafc84b73c765c793560be
	inputPath := "subCert825DaysOK.pem"
	expected := lint.Pass
	out := test.TestLint("e_sub_cert_valid_time_longer_than_825_days", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
