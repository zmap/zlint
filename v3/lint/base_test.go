package lint

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
	"time"

	"github.com/zmap/zcrypto/x509"
)

// This test attempts to simplify the truth table by assigning dates to the
// single digit values 1 through 5, inclusive. As per the standard library,
// 0 is taken to be the null value.
//
// E.G.
//
// If a lint is effective between 2 and 5, then the certs {2, 3, 4} return true.
// If a lint is effective between 0 and 4, then the certs {0, 1, 2, 3} return true.
// If a lint is effective between 2 and 0, then the certs {2, 3, 4, 5} return true.
// if a lint is effective between 0 and 0, then the certs {0, 1, 2, 3, 4, 5} return true.
func TestLint_CheckEffective(t *testing.T) {
	zero := time.Time{}
	one := time.Unix(1, 0)
	two := time.Unix(2, 0)
	three := time.Unix(3, 0)
	four := time.Unix(4, 0)
	five := time.Unix(5, 0)

	lZeroZero := Lint{
		Description:   "ZeroZero",
		EffectiveDate: zero, IneffectiveDate: zero}
	lTwoZero := Lint{
		Description:   "TwoZero",
		EffectiveDate: two, IneffectiveDate: zero}
	lZeroFour := Lint{
		Description:   "ZeroFour",
		EffectiveDate: zero, IneffectiveDate: four}
	lTwoFour := Lint{
		Description:   "TwoFour",
		EffectiveDate: two, IneffectiveDate: four}

	type cert struct {
		Description string
		Certificate *x509.Certificate
	}

	cZero := cert{
		Description: "cZero",
		Certificate: &x509.Certificate{NotBefore: zero},
	}
	cOne := cert{
		Description: "cOne",
		Certificate: &x509.Certificate{NotBefore: one},
	}
	cTwo := cert{
		Description: "cTwo",
		Certificate: &x509.Certificate{NotBefore: two},
	}
	cThree := cert{
		Description: "cThree",
		Certificate: &x509.Certificate{NotBefore: three},
	}
	cFour := cert{
		Description: "cFour",
		Certificate: &x509.Certificate{NotBefore: four},
	}
	cFive := cert{
		Description: "cFive",
		Certificate: &x509.Certificate{NotBefore: five},
	}

	data := []struct {
		Lint        Lint
		Certificate cert
		Want        bool
	}{
		///////////////
		{
			Lint:        lZeroZero,
			Certificate: cZero,
			Want:        true,
		},
		{
			Lint:        lZeroZero,
			Certificate: cOne,
			Want:        true,
		},
		//////////
		{
			Lint:        lTwoZero,
			Certificate: cOne,
			Want:        false,
		},
		{
			Lint:        lTwoZero,
			Certificate: cTwo,
			Want:        true,
		},
		{
			Lint:        lTwoZero,
			Certificate: cThree,
			Want:        true,
		},
		///////////////
		{
			Lint:        lZeroFour,
			Certificate: cTwo,
			Want:        true,
		},
		{
			Lint:        lZeroFour,
			Certificate: cFour,
			Want:        false,
		},
		{
			Lint:        lZeroFour,
			Certificate: cFive,
			Want:        false,
		},
		////////////
		{
			Lint:        lTwoFour,
			Certificate: cOne,
			Want:        false,
		},
		{
			Lint:        lTwoFour,
			Certificate: cTwo,
			Want:        true,
		},
		{
			Lint:        lTwoFour,
			Certificate: cThree,
			Want:        true,
		},
		{
			Lint:        lTwoFour,
			Certificate: cFour,
			Want:        false,
		},
		{
			Lint:        lTwoFour,
			Certificate: cFive,
			Want:        false,
		},
	}
	for _, d := range data {
		got := d.Lint.CheckEffective(d.Certificate.Certificate)
		if got != d.Want {
			t.Errorf("Lint %s, cert %s, got %v want %v",
				d.Lint.Description, d.Certificate.Description, got, d.Want)
		}
	}
}
