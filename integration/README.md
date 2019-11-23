ZLint Integration Tests
=======================

Overview
--------

Integration tests are run during Travis with the `make integration` target of
the Zlint makefile. This uses the default configuration located in
`integration/config.json`.

At a high level the integration test process involves fetching configured CSV
data files, parsing certificates from the data file rows, linting the
certificates, and finally comparing the results to the expected values from the
configuration file. Any differences between the results and the expected values
will fail the integration test.

The ZLint integration tests are intended to make it easier to develop and test
new lints against representative data as well as to catch regressions and bugs
with existing lints.

Running the integration tests
-----------------------------

To run the integration tests with the default configuration use the
`integration` make target:

```
make integration
```

To increase the number of linting Go routines set the `PARALLELISM` variable:

```
make integration PARALLELISM=10
```

To pass other integration test command line parameters use the `INT_TEST`
variable:

```
make integration INT_FLAGS="-lintSummary -fingerprintSummary -lintFilter='^e_' -config small.config.json"
```

Config options
--------------

* `-parallelism` - number of linting Go routines to spawn (_Default: 5_)

* `-configFile` - integration test config file (_Default `integration/config.json`_)

* `-forceDownload` - ignore cached data files on disk forcing it to be downloaded fresh (_Default false_)

* `-overwriteExpected` - overwrite the expected results map in the `-configFile` with the results of the test run. This is useful when new lints or bugfixes are added and the changes in the results map have been vetted and are ready to be committed to the repository. (_Default false_)

* `-fingerprintSummarize` - print a summary of all certificate fingerprints that had lint findings. Can be quite spammy with the default data set. (_Default false_)

* `-fingerprintFilterString` - only lint certificates with hex encoded fingerprints that match the provided regular expression (_Default none_)

* `-lintSummarize` - print a summary of result type counts by lint name. (_Default false_)

* `-lintFilterString` - only lint certificates with lints that have a name that matches the provided regular expression (_Default: none_)

* `-outputTick` - number of certificates to lint before printing a "." marker to output (_Default 1000_)

Data
----

The certificate data used by the integration tests were collected from
[Censys](https://censys.io/) using [a
query](https://github.com/zmap/zlint-test-corpus/blob/847bdf990a0f1ca4f709457d235c850a7a891b73/query.sql)
intended to select random samples of certificates that chain to a Mozilla
trusted root

The exported CSV data files created by this query live in a separate Github
repository to avoid bloating the ZLint repo:
[zmap/zlint-test-corpus](https://github.com/zmap/zlint-test-corpus).

The default configuration uses 60 CSV files from the `zlint-test-corpus`
repository. This represents just shy of 600,000 certificates.

Care is taken by the integration test tooling to download the data only once.
Cached copies on-disk are used for subsequent runs unless the `-forceDownload`
flag is provided.


Adding a new lint
-----------------

TODO(@cpu)

Investigating failures
----------------------

TODO(@cpu)
