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

Example ZLint CLI usage::

	go get github.com/zmap/zlint/cmd/zlint
	zlint mycert.pem


Library Usage
-------------

ZLint can also be used as a library::


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

