Contributing Code
-----------------

**Submitting Code for Review.** We strongly prefer multiple small pull
requests (PR), each of which contain a single lint or a small handful of lints,
over a single large PR. This allows for better code review, faster turnaround
times on comments and merging, as well as for contributors to learn from any
requested changes in the initial round of review. We are happy to wait to cut
new a version of ZLint until a set of PRs have been approved and merged.

Adding New Lints
----------------

**Generating Lint Scaffolding.** The scaffolding for a new lints can be created
by running `./newLint.sh <path_name> <lint_name> <structName>`. Path name may be
one of the existing folders under `lints` (for example `apple`, `cabf_br`, `rfc`
etc) and the choice depends on who authors/suggests the lint specification. Lint
names are generally of the form `e_subject_common_name_not_from_san` where the
first letter is one of: `e`, `w`, or `n` (error, warning, or notice respectively).
Struct names following Go conventions, e.g., `subjectCommonNameNotFromSAN`. Example:
`./newLint.sh rfc e_subject_common_name_not_from_san subjectCommonNameNotFromSAN`.
This will generate a new lint in the `lints/rfc` directory with the necessary
fields filled out.

**Choosing a Lint Result Level.** When choosing what `lints.LintStatus` your new
lint should return (e.g. `Notice`,`Warn`, `Error`, or `Fatal`) the following
general guidance may help. `Error` should be used for clear violations of RFC/BR
`MUST` or `MUST NOT` requirements and include strong citations. `Warn` should be
used for violations of RFC/BR `SHOULD` or `SHOULD NOT` requirements and again
should include strong citations. `Notice` should be used for more general "FYI"
statements that violate non-codified community standards or for cases where
citations are unclear. Lastly `Fatal` should be used when there is an
unresolvable error in `zlint`, `zcrypto` or some other part of the certificate
processing.

**Scoping a Lint.** Lints are executed in three steps. First, the ZLint
framework determines whether a certificate falls within the scope of a given
lint by calling `CheckApplies`. This is often used to scope lints to only check
subscriber, intermediate CA, or root CAs. This function commonly calls one of a
select number of helper functions: `IsCA`, `IsSubscriber`, `IsExtInCert`, or
`DNSNamesExist`. Example:

```go
func (l *caCRLSignNotSet) CheckApplies(c *x509.Certificate) bool {
	return c.IsCA && util.IsExtInCert(c, util.KeyUsageOID)
}
```

Next, the framework determines whether the certificate was issued after the
effective date of a Lint by checking whether the certificate was issued prior
to the lint's `EffectiveDate`. You'll also need to fill out the source and
description of what the lint is checking. We encourage you to copy text
directly from the BR or RFC here. Example:

```go
func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_ca_country_name_missing",
		Description:   "Root and Subordinate CA certificates MUST have a countryName present in subject information",
		Citation:      "BRs: 7.1.2.1",
		Source:        lint.CABFBaselineRequirements,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &caCountryNameMissing{},
	})
}
```

The meat of the lint is contained within the `Execute` function, which is
passed a `x509.Certificate` instance. **Note:** This is an X.509 object from
[ZCrypto](https://github.com/zmap/zcrypto) not the Go standard library. 

Lints should perform their described test and then return a `*LintResult` that
contains a `Status` and optionally a `Details` string, e.g.,
`&LintResult{Status: Pass}`. If you encounter a situation in which you
typically would return a Go `error` object, instead return
`&LintResult{Status: Fatal}`.

Example:

```go
func (l *caCRLSignNotSet) Execute(c *x509.Certificate) *lint.LintResult {
	if c.KeyUsage&x509.KeyUsageCRLSign != 0 {
		return &lint.LintResult{Result: Pass}
	}
	return &lint.LintResult{Result: Error}
}
```

Testing Lints
-------------

**Creating Unit Tests.** Every lint should also have corresponding unit
tests (generally at least one for a success and one for afailure condition). We
have typically generated test certificates using Go (see
[documentation][CreateCertificates] for details), but OpenSSL
could also be used. Test certificates should be placed in `testdata/` and called
from the test file created by `newLint.sh`. You may want to prepend the PEM with
the output of `openssl x509 -text`. You can run your lint against a test
certificate from a unit test using the `test.TestLint` helper function.

[CreateCertificates]: https://golang.org/pkg/crypto/x509/#CreateCertificate 

Example:

```go
func TestBasicConstNotCritical(t *testing.T) {
	inputPath := "caBlankCountry.pem"
	expected := Error
	out := test.TestLint("e_basic_constraints_not_critical", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

```

**Integration Tests.** ZLint's [continuous integration][CI] includes an
integration test phase where all lints are run against a large corpus of
certificates. The number of notice, warning, error and fatal results for each
lint are captured and compared to a set of expected values in a configuration
file. You may need to update these expected values when you add/change lints.
Please see the [integration tests README] for more information.

[CI]: https://travis-ci.org/zmap/zlint
[integration tests README]: https://github.com/zmap/zlint/blob/master/integration/README.md


Updating the TLD Map
--------------------

ZLint maintains [a map of top-level-domains][TLD Map] and their validity periods
that is referenced by linters. This data is updated periodically by a bot
integration using the `zlint-gltd-update` command.

To update the data manually ensure the `zlint-gtld-update` command is installed
and in your `$PATH` and run `go generate`:

	go get github.com/zmap/zlint/cmd/zlint-gtld-update
	go generate github.com/zmap/zlint/...

[TLD Map]: https://github.com/zmap/zlint/blob/master/util/gtld_map.go
