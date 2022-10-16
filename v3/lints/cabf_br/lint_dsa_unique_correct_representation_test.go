package cabf_br

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
	"math/big"
	"testing"

	"github.com/zmap/zcrypto/dsa"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestDSAUniqueCorrectRepresentation(t *testing.T) {
	inputPath := "dsaUniqueRep.pem"
	expected := lint.Pass
	out := test.TestLint("e_dsa_unique_correct_representation", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestDSANotUniqueCorrectRepresentation(t *testing.T) {
	inputPath := "dsaUniqueRep.pem"
	c := test.ReadTestCert(inputPath)

	// Replace Y with P - 1
	dsaKey := c.PublicKey.(*dsa.PublicKey)
	pMinusOne := big.NewInt(0)
	pMinusOne.Sub(dsaKey.P, big.NewInt(1))
	dsaKey.Y = pMinusOne

	// Expect failure
	expected := lint.Error
	out := test.TestLintCert("e_dsa_unique_correct_representation", c, lint.NewEmptyConfig())
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
