ZLint
=====

[![Build Status](https://travis-ci.org/zmap/zlint.svg?branch=master)](https://travis-ci.org/zmap/zlint)
[![Go Report Card](https://goreportcard.com/badge/github.com/zmap/zlint)](https://goreportcard.com/report/github.com/zmap/zlint)

ZLint is a golang-based X.509 certificate linter that checks for consistency
with [RFC 5280](https://www.ietf.org/rfc/rfc5280.txt) and the CA/Browser Forum
Baseline Requirements
([v.1.4.8](https://cabforum.org/wp-content/uploads/CA-Browser-Forum-BR-1.4.8.pdf)).

A detailed list of BR coverage can be found here:
https://docs.google.com/spreadsheets/d/1ywp0op9mkTaggigpdF2YMTubepowJ50KQBhc_b00e-Y.

Command Line Usage
------------------

ZLint can be used on the command-line through a simple bundled executable
_ZLint_ as well as through
[ZCertificate](https://github.com/zmap/zcertificate), a more full-fledged
command-line certificate parser that links against ZLint.

Example ZLint CLI usage:

	go get github.com/zmap/zlint/cmd/zlint
	zlint mycert.pem


Library Usage
-------------

ZLint can also be used as a library:

```golang
import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint"
)

parsed, err := x509.ParseCertificate(raw)
if err != nil {
	// The certificate could not be parsed. Either error or halt.
	log.Errorf("could not parse certificate: %s", err)
}
zlintResult := zlint.LintCertificate(parsed)
```


See https://github.com/zmap/zlint/blob/master/cmd/zlint/main.go for an example.


Adding New Lints
----------------

**Generating Lint Scaffolding.** The scaffolding for a new lints can be created by running `./newLint.sh <lint_name> <structName>`. Lint names are generally of the form `e_subject_common_name_not_from_san` where the first letter is one of: `e`, `w`, or `n` (error, warning, or notice respectively). Struct names following golang conventions, e.g., `subjectCommonNameNotFromSAN`. Example: `./newLint.sh e_subject_common_name_not_from_san subjectCommonNameNotFromSAN`. This will generate a new lint in the `lints` directory with the necessary fields filled out.

**Scoping a Lint.** Lints are executed in three steps. First, the ZLint framework determines whether a certificate falls within the scope of a given lint by calling `CheckApplies`. This is often used to scope lints to only check subscriber, intermediate CA, or root CAs. This function commonly calls one of a select number of helper functions: `IsCA`, `IsSubscriber`, `IsExtInCert`, or `DNSNamesExist`. Example:

```golang
func (l *caCRLSignNotSet) CheckApplies(c *x509.Certificate) bool {
	return c.IsCA && util.IsExtInCert(c, util.KeyUsageOID)
}
```

Next, the framework determines whether the certificate was issued after the effective date of a Lint by checking whether the certificate was issued prior to the lint's `EffectiveDate`. You'll also need to fill out the source and description of what the lint is checking. We encourage you to copy text directly from the BR or RFC here. Example:

```golang
func init() {
	RegisterLint(&Lint{
		Name:          "e_ca_country_name_missing",
		Description:   "Root and Subordinate CA certificates MUST have a countryName present in subject information",
		Source:        "BRs: 7.1.2.1",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &caCountryNameMissing{},
	})
}
```

The meat of the lint is contained within the `RunTest` function, which is passed `x509.Certificate`. **Note:** This is an X.509 object from [ZCrypto](https://github.com/zmap/zcrypto) not Golang stdlib. Lints should perform their described test and then return a `ResultStruct` that contains a Result and optionally a `Details` string, e.g., `ResultStruct{Result: Pass}`.

Example:

```golang
func (l *caCRLSignNotSet) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if c.KeyUsage&x509.KeyUsageCRLSign != 0 {
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: Error}, nil
	}
}
```

**Creating Tests.** Every lint should also have two corresponding tests for a success and failure condition. We have typically generated test certificates using Golang (see https://golang.org/pkg/crypto/x509/#CreateCertificate for details), but OpenSSL could also be used. Test certificates should be placed in `testlint/testCerts` and called from the test file created by `newLint.sh`. Example:

```golang
func TestBasicConstNotCrit(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/caBasicConstNotCrit.pem"
	desEnum := Error
	out, _ := Lints["e_basic_constraints_not_critical"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

```


License and Copyright
---------------------

ZMap Copyright 2017 Regents of the University of Michigan

Licensed under the Apache License, Version 2.0 (the "License"); you may not use
this file except in compliance with the License. You may obtain a copy of the
License at http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed
under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
CONDITIONS OF ANY KIND, either express or implied. See LICENSE for the specific
language governing permissions and limitations under the License.

