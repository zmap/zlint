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
by running `./newLint.sh -r <requirement> -n <lint_name> -s <structName>`. Path name may
be one of the existing folders under `lints` (for example `apple`, `cabf_br`, `rfc`
etc) and the choice depends on who authors/suggests the lint specification. Lint
names are generally of the form `e_subject_common_name_not_from_san` where the
first letter is one of: `e`, `w`, or `n` (error, warning, or notice respectively).
Struct names following Go conventions, e.g., `subjectCommonNameNotFromSAN`. Example:
`./newLint.sh -r rfc -n e_subject_common_name_not_from_san -s subjectCommonNameNotFromSAN`.
This will generate a new lint in the `lints/rfc` directory with the necessary
fields filled out.

**Choosing Result Level.**  Lints return a single type of status:

 * **Error:** `Error` can only be used for clear violations of `MUST` or `MUST
   NOT` requirements and must include a specific citation.

 * **Warning:** `Warn` can only be used for violations of `SHOULD` or `SHOULD
   NOT` requirements and again should include strong citations. Many
   certificate authorities block on both Error and Warning lints, and Warning
   lints should not be used for non-deterministic errors (e.g., calculating
   whether a serial number has sufficient entropy based on high-order bits.)

 * **Notice:** `Notice` should be used for more general "FYI" statements that
   indicate there may be a problem. Non-deterministic lints are OK.

Lints only return one non-success or non-fatal status, which must also match
their name prefix. For example, `e_ian_wildcard_not_first` can only return a
`SUCCESS`, `ERROR`, or `FATAL` status.  It cannot return a `NOTICE` or
`WARNING` status. Any lint can return a `FATAL` error, but `FATAL` should only
be used when there is an unresolvable error in `zlint`, `zcrypto` or some other
part of the certificate processing.

**Lint Source:** Typically Lint Source is straightfoward since every lint needs
a citation. However, sometimes the community has lints that aren't codified in
a formal document. In these situations, do not create a `NOTICE` lint under a
common source (e.g,. RFC or Baseline Requirements). Instead, create a lint
using the `ZLint` source. Lints in this source are included at the maintainers'
discretion, though we typically shy away from lints with significant
controversy.  We encourage certificate authorities and other users to
participate in the ZLint review process and to express their opinions on
community lints during the Pull Request review period.

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
		Lint:          NewCaCountryNameMissing,
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

Making your Lint Configurable
-------------
Lints may implement an optional interface - `Configurable`...

```go
type Configurable interface {
    Configure() interface{}
}
```

...where the returned `interface{}` is a pointer to the target struct to deserialize your configuration into.

This struct may encode any arbitrary data that may be deserialized from [TOML](https://toml.io/en/). Examples may include:

* PEM encoded certificates or certificate chains
* File paths
* Resolvable DNS entries or URIs
* Dates or Unix timestamps

...and so on. How stable and/or appropriate a given configuration field is is left as a code review exercise on a per-lint basis.

If a lint is `Configurable` then a new step is injected at the beginning of its lifecycle.

---
##### Non-Configurable Lifecycle
> * CheckApplies
> * CheckEffective
> * Execute

##### Configurable Lifecycle
> * Configure
> * CheckApplies
> * CheckEffective
> * Execute

### Higher Scoped Configurations

Lints may embed within theselves either pointers or structs to the following definitions within the `lint` package.

```go
type Global struct {}
type RFC5280Config struct{}
type RFC5480Config struct{}
type RFC5891Config struct{}
type CABFBaselineRequirementsConfig struct {}
type CABFEVGuidelinesConfig struct{}
type MozillaRootStorePolicyConfig struct{}
type AppleRootStorePolicyConfig struct{}
type CommunityConfig struct{}
type EtsiEsiConfig struct{}
```

Doing so will enable receiving a _copy_ of any such defintions from a higher scope within the configuration.

```toml
# Top level (non-scoped) fields will be copied into any Global struct that you declare within your lint.
something_global = 5
something_else_global = "The funniest joke in the world."

[RFC5280]
# Top level (non-scoped) fields will be copied into any RFC5280Config struct that you declare within your lint.
wildcard_allowed = true

[MyLint]
# You can also embed comments!
my_config = "Some arbitrary data."
```

An example of the above might be...

```go
type MyLint struct {
	Global      lint.Global
	RFC         lint.RFC5280Config
	MyConfig    string `toml:"my_config",comment:"You can also embed comments!"`
}
```

Testing Lints
-------------

**Creating Unit Tests.** Every lint should also have corresponding unit tests
(generally at least one for a success and one for a failure condition). There
are various ways for generating test certificates. The following options have
been used by contributers successfully:

* Create new certificates using [Go][CreateCertificate] (compare [this
  article on SO][certGenerator] as starting point)
* Modify existing certificates using [der-ascii][DERASCII] (compare [this
  documentation][resign] how to re-sign the modified certificate)
* Using OpenSSL

Test certificates should be placed in `testdata/` and called from the test file
created by `newLint.sh`. All test certificates must have the textual description
from `openssl x509 -text` added before the PEM header or CI will flag them as a
build error. You can add the text decoding to all of the test certs missing it
by running `test/prepend_testcerts_openssl.sh`.

[CreateCertificate]: https://golang.org/pkg/crypto/x509/#CreateCertificate
[certGenerator]: https://stackoverflow.com/q/26441547/1426535
[DERASCII]:https://github.com/google/der-ascii
[resign]:https://github.com/google/der-ascii/blob/master/samples/certificates.md

If you only have one or two test cases separate unit test functions are
acceptable, example:

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

If you have more than two or three test cases we prefer new unit tests to be
written in a [table driven style][table-tests]. Each testcase should be invoked
as a [subtest][subtests] so that it's easy to figure out which subtest failed
and to allow control over which subtests are run.

Example: see [`lint_ct_sct_policy_count_unsatisfied_test.go`][sct_test_eg]

[table-tests]: https://github.com/golang/go/wiki/TableDrivenTests
[subtests]: https://golang.org/pkg/testing/#hdr-Subtests_and_Sub_benchmarks
[sct_test_eg]: https://github.com/zmap/zlint/blob/master/v3/lints/apple/lint_ct_sct_policy_count_unsatisfied_test.go

**Integration Tests.** ZLint's [continuous integration][CI] includes an
integration test phase where all lints are run against a large corpus of
certificates. The number of notice, warning, error and fatal results for each
lint are captured and compared to a set of expected values in a configuration
file. You may need to update these expected values when you add/change lints.
Please see the [integration tests README] for more information.

[CI]: https://travis-ci.org/zmap/zlint
[integration tests README]: https://github.com/zmap/zlint/blob/master/v3/integration/README.md

### Testing Configurable Lints

Testing a lint that is configurable is much the same as testing one that is not. However, if you wish to exercise
various configurations then you may do so by utilizing the `test.TestLintWithConfig` function which takes in an extra
string which is the raw TOML of your target test configuration.

```go
func TestCaCommonNameNotMissing2(t *testing.T) {
	inputPath := "caCommonNameNotMissing.pem"
	expected := lint.Pass
	config := `
            [e_ca_common_name_missing2]
            BeerHall = "liedershousen"
        `
	out := test.TestLintWithConfig("e_ca_common_name_missing2", inputPath, config)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
```

Adding New Profiles
----------------
**Generating Profile Scaffolding.** The scaffolding for a new profiles can be created
by running `./newProfile.sh <profile_name>`.

An example is:

```bash
$ ./newProfile.sh my_new_profile
```

This will generate a new file in the `profiles` directory by the name `profile_my_new_profile.go` for you.

Updating the TLD Map
--------------------

ZLint maintains [a map of top-level-domains][TLD Map] and their validity periods
that is referenced by linters. This data is updated periodically by a bot
integration using the `zlint-gltd-update` command.

To update the data manually ensure the `zlint-gtld-update` command is installed
and in your `$PATH` and run `go generate`:

	go get github.com/zmap/zlint/v3/cmd/zlint-gtld-update
	go generate github.com/zmap/v3/zlint/...

[TLD Map]: https://github.com/zmap/zlint/blob/master/v3/util/gtld_map.go


Publishing a Release
--------------------

ZLint releases are published via Github Actions using Goreleaser. Most of the
release process is automated but there is still some manual effort involved in
creating good release notes & communicating news of the release.

At a high level the release process requires:

1. Preparing release notes.
1. Choosing an appropriate new version per semver.
1. Pushing an annotated release candidate tag.
1. Monitoring CI for successful completion.
1. Editing & Publishing the Github release candidate created by CI.
1. Creating a call-for-testing announcement in Github issues.
1. Emailing the announcement list.
1. Waiting a week.
1. Pushing a final release tag.
1. Editing & Publishing the Github release created by CI.
1. Closing the release announcement Github issue.
1. Emailing the announcement list.

To prepare the release notes examine the diff between `HEAD` and the previous
release tag. E.g. if `v2.0.0` is the latest release, use:

```bash
git log v2.0.0..HEAD --oneline
```

Try to pull out the commits of importance, following the format of [previous
release notes](https://github.com/zmap/zlint/releases/tag/v2.2.0-rc1). E.g.
pulling out new lints, updated lints, bug fixes, etc. Remember that you don't
need to mention every commit because the release tooling will include a full
change-log of commits. Your job is to emphasize the highlights.

When choosing a new version tag you should reference [the semver
philosophy](http://semver.org/) and the commitments made in the [ZLint
README](https://github.com/zmap/zlint#versioning-and-releases).

Release tags should be annotated with the release notes you prepared so use `-a`
when creating the new tag. You may want to GPG sign the tag, if so add `-s`.
Lastly remember to obey the expected format for the tag name. For final versions
`'v$MAJOR.$MINOR.$PATCH'` and for release candidates
`'v$MAJOR.$MINOR.$PATCH-rc$NUMBER'`. See `git tag` for previous examples to
match.

As an example to create a tag for a first v2.2.0 release candidate run:
```bash
git tag -s -a v2.2.0-rc1
git push origin v2.2.0-rc1
```

After pushing a tag with the expected release format the deploy job
configured in the `.github/workflows/release.yml` workflow will kick in and
invoke [Goreleaser](https://goreleaser.com/).

Once the build completes Goreleaser and Github actions will have created
a **draft** release in [the project release section of
Github](https://github.com/zmap/zlint/releases). You will need to edit this
release to add your release notes in front of the full change-log of commits. The
release will not be visible until you explicitly publish it. The Goreleaser
automation will attach binary artifacts to the release as they are available.

Now is a good time to create a call-for-testing issue. You can copy a [previous
example](https://github.com/zmap/zlint/issues/466) to create a new one. It
should reference the Github release you just published and is a central place
for folks to report issues with a release candidate.

Next, post to the [ZLint Announcements Mailing
List](https://groups.google.com/forum/#!forum/zlint-announcements). You should
copy the release notes in, link to the Github release, and also reference the
call-for-testing issue.

Assuming the release candidate has no issues that need to be addressed with bug
fixes & a new release candidate tag you can "finalize" the release by pushing
a new tag with the `-rc$NUMBER` portion removed. Repeat the process of editing
the draft Github release to add notes, publishing it, and posting to the mailing
list.

You're done!

For more detail consult the [Goreleaser
docs](https://goreleaser.com/quick-start/), the release workflow configuration in
[`release.yml`](https://github.com/zmap/zlint/blob/master/.github/workflows/release.yml),
and the
[`.goreleaser.yml`](https://github.com/zmap/zlint/blob/master/v3/.goreleaser.yml)
project configuration.

Generating Test Certificates
-----------------
At times, it may be difficult to generate examples, or counter examples, for a particular lint.
To that end, we have `genTestCerts.go` - a playground script that is intended for contributors
to edit (but not commit) to their heart's content in order to generate the oddly specific
certificates that one may need in order to sufficiently exercise one's lint.

Of course, generating x509 certificates is a _highly_ configurable procedure which is why this script
is intended to be edited and ran locally rather than as an extremely complex command line tool or service
(that project already exists - openssl).

The intent of the script is that authors can modify and run in it any way they see fit in order
to get themselves off the ground, but to ultimately not submit any local changes made to the script.
In that regard, please feel free to whack this file around to your heart's content in order to accomplish
your goals. If you think that you have improved upon the contents of the script itself, then please do
open a pull against the script itself (however, please refrain from bundling it with anything else such as
a new lint).

This script has a facility for generating a self signed trust anchor to act as a CA, a facility
for generating intermediate certificates, and a facility for generating a leaf certificate.
The certificates generated by each are NOT healthy nor acceptable to any reasonable PKI system.
However, being a complete and usable certificate is not necessarily required when you are writing
a lint for, say, checking that a certificate does not expire on Valentine's Day (because no certificate
should be alone on Valentine's Day).

In general, you should generate whatever certificate/s you need in order to pass the CheckApplies method for your
particular lint and modify the one (hopefully) field that you are checking. For the sake of coverage it may also
be a good idea to generate a certificate for which CheckApplies returns false.
