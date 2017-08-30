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


See https://github.com/zmap/zlint/blob/master/cmd/zlint/main.go for an example.

Contributing
-------------

If you would like to add a new x509 Lint:

0. Make sure the lint doesn't already exist.
1. Come up with a name for your lint. If your lint returns an error (i.e., an RFC or the BRs use a MUST 
clause), prepend your lint name with `e_`. If your lint returns a warning (i.e., an RFC or the BRs use a SHOULD
clause), prepend your lint name with `w_`. For example, `e_subject_common_name_not_from_san`. 
2. Come up with a struct name for your lint. Typically just camelCase the name of your lint. From the previous example, a suitable struct name would be `subjectCommonNameNotFromSAN`. 
3. Run the following command:
`./newLint.sh <lint_name> <structName>`
This will generate a new lint, in the `lints` directory, with the necessary fields filled out.
4. Determine what prerequisites are necessary for your lint, and add code to 
the `CheckApplies` function that ensures the prerequisites are met. For example, 
if your lint only applies to subscriber certificates, you would add 
`return util.IsSubscriberCert(c)` 
in the `CheckApplies` function.
5. Fill out the `Description` of the Lint, as well as the `Provenance` of 
the Lint (where did the lint come from?), as well as the earliest date 
that the lint was effective in the `EffectiveDate` field.
6. Write the logic of your lint in the `RunTest` function.
7. Create a test file for your lint by creating a file in the lints directory called `<lint_name>_test.go`.
8. Create test certifiates that test your lint. You can do this via `openssl`
configs, or writing pure Golang (https://golang.org/pkg/crypto/x509/#CreateCertificate)
9. For each new test certificate, run the following command.
`openssl x509 -in <testCert> -text -noout | cat - <testCert> > /tmp/out && mv /tmp/out <testCert>`
10. Place these test certificates in the `testlint/testCerts` directory.
11. Run 
`go test ./...` 
in the top level directory to ensure that your tests pass.
12. Send a PR.

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

