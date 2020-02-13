package lint

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
	"time"

	"github.com/zmap/zcrypto/x509"
)

func TestLintCheckEffective(t *testing.T) {
	c := &x509.Certificate{
		NotBefore: time.Now(),
	}
	l := Lint{}

	l.EffectiveDate = time.Time{}
	if l.CheckEffective(c) != true {
		t.Errorf("EffectiveDate of zero should always be true")
	}
	l.EffectiveDate = time.Unix(1, 0)
	if l.CheckEffective(c) != true {
		t.Errorf("EffectiveDate of 1970-01-01 should be true")
	}
	l.EffectiveDate = time.Unix(32503680000, 0) // 3000-01-01
	if l.CheckEffective(c) != false {
		t.Errorf("EffectiveDate of 3000 should be false")
	}
}
