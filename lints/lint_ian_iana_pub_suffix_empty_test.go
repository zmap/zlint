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

package lints

import (
	"testing"
)

func TestIANBarePubSuffix(t *testing.T) {
	inputPath := "../testlint/testCerts/IANBareSuffix.pem"
	expected := Warn
	out := Lints["w_ian_iana_pub_suffix_empty"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestIANGoodPubSuffix(t *testing.T) {
	inputPath := "../testlint/testCerts/IANGoodSuffix.pem"
	expected := Pass
	out := Lints["w_ian_iana_pub_suffix_empty"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
