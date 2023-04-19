ZLint
=====

[![CI Status](https://github.com/zmap/zlint/workflows/Go/badge.svg)](https://github.com/zmap/zlint/actions?query=workflow%3AGo)
[![Integration Tests](https://github.com/zmap/zlint/workflows/integration-test/badge.svg)](https://github.com/zmap/zlint/actions?query=workflow%3Aintegration-test)
[![Lint Status](https://github.com/zmap/zlint/workflows/golangci-lint/badge.svg)](https://github.com/zmap/zlint/actions?query=workflow%3Agolangci-lint)
[![Go Report Card](https://goreportcard.com/badge/github.com/zmap/zlint)](https://goreportcard.com/report/github.com/zmap/zlint)

ZLint is a X.509 certificate linter written in Go that checks for consistency
with standards (e.g. [RFC 5280]) and other relevant PKI requirements (e.g.
[CA/Browser Forum Baseline Requirements][BR v1.4.8]).

It can be used as a command line tool or as a library integrated into CA
software.

[RFC 5280]: https://www.ietf.org/rfc/rfc5280.txt
[BR v1.4.8]: https://cabforum.org/wp-content/uploads/CA-Browser-Forum-BR-1.4.8.pdf

Requirements
------------

ZLint requires [Go 1.16.x or newer](https://golang.org/doc/install) be
installed. The command line setup instructions assume the `go` command is in
your `$PATH`.

Lint Sources
------------

Historically ZLint was focused on only [RFC 5280] and [v1.4.8][BR v1.4.8] of the
[CA/Browser Forum baseline requirements][BRs]. A detailed list of the original
BR coverage can be found [in this spreadsheet][Coverage Spreadsheet].

More recently ZLint has been restructured to make it easier to add lints based
on other sources. While not complete, presently ZLint has lints sourced from:

* [CA/Browser Forum EV SSL Certificate Guidelines][CABF EV]
* [ETSI ESI]
* [Mozilla's PKI policy][MozPolicy]
* [Apple's CT policy][AppleCT]
* Various RFCs (e.g. [RFC 6818], [RFC 4055], [RFC 8399])

By default ZLint will apply applicable lints from all sources but consumers may
also customize which lints are used by including/exclduing specific sources.

[BRs]: https://cabforum.org/baseline-requirements-documents/
[Coverage Spreadsheet]: https://docs.google.com/spreadsheets/d/1ywp0op9mkTaggigpdF2YMTubepowJ50KQBhc_b00e-Y
[CABF EV]: https://cabforum.org/extended-validation/
[MozPolicy]: https://github.com/mozilla/pkipolicy
[ETSI ESI]: https://www.etsi.org/technologies/digital-signature
[AppleCT]: https://support.apple.com/en-us/HT205280
[RFC 6818]: https://www.ietf.org/rfc/rfc6818.txt
[RFC 4055]: https://www.ietf.org/rfc/rfc4055.txt
[RFC 8399]: https://www.ietf.org/rfc/rfc8399.txt


Versioning and Releases
-----------------------

ZLint aims to follow [semantic versioning](https://semver.org/). The addition of
new lints will generally result in a MINOR version revision. Since downstream
projects depend on lint results and names for policy decisions changes of this
nature will result in MAJOR version revision.

Where possible we will try to make available a release candidate (RC) a week
before finalizing a production ready release tag. We encourage users to test RC
releases to provide feedback early enough for bugs to be addressed before the
final release is made available.

Please subscribe to the [ZLint Announcements][zlint-announce] mailing list to
receive notifications of new releases/release candidates. This low-volumne
announcements mailing list is only used for new ZLint releases and major
project announcements, not questions/support/bug reports.

[zlint-announce]:  https://groups.google.com/forum/#!forum/zlint-announcements


Command Line Usage
------------------

ZLint can be used on the command-line through a simple bundled executable
_ZLint_ as well as through
[ZCertificate](https://github.com/zmap/zcertificate), a more full-fledged
command-line certificate parser that links against ZLint.

Example ZLint CLI usage:

	go get github.com/zmap/zlint/v3/cmd/zlint
	echo "Lint mycert.pem with all applicable lints"
	zlint mycert.pem

	echo "Lint mycert.pem with just the two named lints"
	zlint -includeNames=e_mp_exponent_cannot_be_one,e_mp_modulus_must_be_divisible_by_8 mycert.pem

	echo "List available lint sources"
	zlint -list-lints-source

	echo "Lint mycert.pem with all of the lints except for ETSI ESI sourced lints"
	zlint -excludeSources=ETSI_ESI mycert.pem

	echo "Receive a copy of the full (default) configuration for all configurable lints"
	zlint -exampleConfig

	echo "Lint mycert.pem using a custom configuration for any configurable lints"
	zlint -config configFile.toml mycert.pemr

	echo "List available lint profiles. A profile is a pre-defined collection of lints."
	zlint -list-profiles

See `zlint -h` for all available command line options.


Library Usage
-------------

ZLint can also be used as a library. To lint a certificate with all applicable
lints is as simple as using `zlint.LintCertificate` with a parsed certificate:

```go
import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3"
)

var certDER []byte = ...
parsed, err := x509.ParseCertificate(certDER)
if err != nil {
	// If x509.ParseCertificate fails, the certificate is too broken to lint.
	// This should be treated as ZLint rejecting the certificate
	log.Fatal("unable to parse certificate:", err)
}
zlintResultSet := zlint.LintCertificate(parsed)
```

To lint a certificate with a subset of lints (e.g. based on lint source, or
name) filter the global lint registry and use it with `zlint.LintCertificateEx`:

```go
import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3"
	"github.com/zmap/zlint/v3/lint"
)

var certDER []byte = ...
parsed, err := x509.ParseCertificate(certDER)
if err != nil {
	// If x509.ParseCertificate fails, the certificate is too broken to lint.
	// This should be treated as ZLint rejecting the certificate
	log.Fatal("unable to parse certificate:", err)
}

registry, err := lint.GlobalRegistry().Filter(lint.FilterOptions{
  ExcludeSources: []lint.LintSource{lint.EtsiEsi},
})
if err != nil {
	log.Fatal("lint registry filter failed to apply:", err)
}
zlintResultSet := zlint.LintCertificateEx(parsed, registry)
```

To lint a certificate in the presence of a particular configuration file, you must first construct the configuration and then make a call to `SetConfiguration` in the `Registry` interface.

A `Configuration` may be constructed using any of the following functions:

* `lint.NewConfig(r io.Reader) (Configuration, error)`
* `lint.NewConfigFromFile(path string) (Configuration, error)`
* `lint.NewConfigFromString(config string) (Configuration, error)`

The contents of the input to all three constructors must be a valid TOML document.

```go
import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3"
)

var certDER []byte = ...
parsed, err := x509.ParseCertificate(certDER)
if err != nil {
	// If x509.ParseCertificate fails, the certificate is too broken to lint.
	// This should be treated as ZLint rejecting the certificate
	log.Fatal("unable to parse certificate:", err)
}
configuration, err := lint.NewConfigFromString(`
        [some_configurable_lint]
        IsWebPki = true
        NumIterations = 42
        
        [some_configurable_lint.AnySubMapping]
        something = "else"
        anything = "at all"
`)
if err != nil {
	log.Fatal("unable to parse configuration:", err)
}
lint.GlobalRegistry().SetConfigutration(configuration)
zlintResultSet := zlint.LintCertificate(parsed)
```

See [the `zlint` command][zlint cmd]'s source code for an example.

[zlint cmd]: https://github.com/zmap/zlint/blob/master/v3/cmd/zlint/main.go


Extending ZLint
----------------

For information on extending ZLint with new lints see [CONTRIBUTING.md]

[CONTRIBUTING.md]: https://github.com/zmap/zlint/blob/master/CONTRIBUTING.md


Zlint Users/Integrations
-------------------------

Pre-issuance linting is **strongly recommended** by the [Mozilla root
program](https://wiki.allizom.org/CA/Required_or_Recommended_Practices#Pre-Issuance_Linting).
Here are some projects/CAs known to integrate with ZLint in some fashion:

* [Actalis](https://www.actalis.it/en/home.aspx)
* [ANF AC](https://www.anf.es/)
* [Camerfirma](https://www.camerfirma.com/)
* [CFSSL](https://github.com/cloudflare/cfssl)
* [Digicert](https://www.digicert.com/)
* [EJBCA](https://download.primekey.com/docs/EJBCA-Enterprise/6_11_1/adminguide.html#Post%20Processing%20Validators%20(Pre-Certificate%20or%20Certificate%20Validation))
* [Entrust](https://www.entrust.com/)
* [Globalsign](https://www.globalsign.com/en/)
* [GoDaddy](https://www.godaddy.com)
* [Google Trust Services](https://pki.goog/)
* [Government of Spain, FNMT](http://www.fnmt.es/)
* [Izenpe](https://www.izenpe.eus/)
* [Let's Encrypt](https://letsencrypt.org) and [Boulder](https://github.com/letsencrypt/boulder)
* [Microsec](https://www.microsec.com/)
* [Microsoft](https://www.microsoft.com)
* [Nexus Certificate Manager](https://doc.nexusgroup.com/display/PUB/Smart+ID+Certificate+Manager)
* [QuoVadis](https://www.quovadisglobal.com/)
* [Sectigo](https://sectigo.com/) and [crt.sh](https://crt.sh)
* [Siemens](https://siemens.com/pki)
* [SSL.com](https://www.ssl.com/)

Please submit a pull request to update the README if you are aware of
another CA/project that uses zlint.


License and Copyright
---------------------

ZMap Copyright 2021 Regents of the University of Michigan

Licensed under the Apache License, Version 2.0 (the "License"); you may not use
this file except in compliance with the License. You may obtain a copy of the
License at http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed
under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
CONDITIONS OF ANY KIND, either express or implied. See LICENSE for the specific
language governing permissions and limitations under the License.
