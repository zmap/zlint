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

The certificate data used by the integration tests was collected from
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

Example failure investigation
-----------------------------

Here's an example of using the integration test tooling to investigate a linter
bug.

First, let's revert [a
bugfix](https://github.com/cpu/zlint/commit/5dcecad773158b82b5e52064ee2782d1b8a79314)
for the `e_subject_printable_string_badalpha` lint so we can see what happens
when there's a difference between the test results and the expected results.

* `git revert 5dcecad773158b82b5e52064ee2782d1b8a79314`

Now let's run the integration tests. We'll use a higher than default
parallelism value since our dev machine probably has a few cores laying around.

This will take approximately ~15 minutes (Longer if you haven't downloaded the
integration test data in previous runs). If you want to tighten the iteration
time (e.g. while you're developing a new lint vs chasing a bug) try specifying
a `-config` file that has fewer data files than the default one.

* `make integration PARALLELISM=6`

As we'd expect after reverting a bugfix the integration tests fail.

```
--- FAIL: TestCorpus (448.05s)
    corpus_test.go:139: linted 599997 certificates
    corpus_test.go:163: expected lint "e_subject_printable_string_badalpha" to have result fatals: 0    errs: 7    warns: 0    infos: 0    got fatals: 0    errs: 221  warns: 0    infos: 0   
FAIL
FAIL	github.com/zmap/zlint/integration	448.244s
FAIL
make: *** [makefile:33: integration] Error 1
```

The `e_subject_printable_string_badalpha` lint was expected to find only 7
certificates with errors and it found 221!

The next step is to find out which certificates in the integration test data
are failing. To do that we'll re-run the integration tests specifying a
`-lintFilter` flag so that only the `e_subject_printable_string_badalpha` is
run and a `-fingerprintSummary` flag so the certificate fingerprints that have
a non-pass result from this lint are printed.

* `make integration PARALLELISM=6 INT_FLAGS="-fingerprintSummary -lintFilter='e_subject_printable_string_badalpha'"`

Once that completes (which should be faster than before now that we're only running one lint per certificate) the 221 certificate fingerprints that failed the lint are printed:

```
2019/11/23 18:52:43 Finished reading data from 60 CSV files. Closing work channel

summary of result type by certificate fingerprint:
0037ae7546555efca0935dfedf3cef79b1a0301b18bb6a86382becf6aa53f1c4  fatals: 0    errs: 1    warns: 0    infos: 0
004e38dd0ae5410010a0ebfc6afddeed2020008b146908fd635dc725960fad53  fatals: 0    errs: 1    warns: 0    infos: 0
0066f781f91c6e694e7ad98babc89c9f96cf1087005e8f713559b1ceb16d417b  fatals: 0    errs: 1    warns: 0    infos: 0
008bedb904a6c7a8219c14da91d433863d9d27fbb225c12bfcc7dc3a59657999  fatals: 0    errs: 1    warns: 0    infos: 0
00b308aafa26b3315a9c7371c5ff14807fcd567ea4f543a70dabfa873502d3fb  fatals: 0    errs: 1    warns: 0    infos: 0
00b579f8b86ddca8e2a9d2d610f91786db1bace28327ee9d6c2d7099df78d3f8  fatals: 0    errs: 1    warns: 0    infos: 0
<snipped>
ffe2f3264d9b41980c8c1ebae0f69533b4ed6486e45827447e98ac27c3ddb791  fatals: 0    errs: 1    warns: 0    infos: 0
fff61b942a56b87c5d5dd3725f43d3708bc646df87adb5db1792bbf61ad6875c  fatals: 0    errs: 1    warns: 0    infos: 0
fffd96497d21df4d55fa5e8883645325e1b9472db99e1b1a322d4df8f5b0bd3a  fatals: 0    errs: 1    warns: 0    infos: 0

--- FAIL: TestCorpus (126.13s)
    corpus_test.go:139: linted 599997 certificates
    corpus_test.go:163: expected lint "e_subject_printable_string_badalpha" to have result fatals: 0    errs: 7    warns: 0    infos: 0    got fatals: 0    errs: 221  warns: 0    infos: 0
FAIL
FAIL  github.com/zmap/zlint/integration 126.143s
FAIL

```

The next step is to look at some of the certificates corresponding to the
fingerprints shown. Since the full certificate data is already present on disk
we can do this easily with a small utility script (`integrate/certByFP.sh`)
included with ZLint. 

To check out the first fingerprint from the summary output
(`0037ae7546555efca0935dfedf3cef79b1a0301b18bb6a86382becf6aa53f1c4`) we can run:

```
./integration/certByFP.sh 0037ae7546555efca0935dfedf3cef79b1a0301b18bb6a86382becf6aa53f1c4
```

This will find the matching certificate in the cached integration test data
directory, parse it with OpenSSL, print the text version and the PEM version,
and finally show a Censys.io URL:

```
./integration/certByFP.sh 0037ae7546555efca0935dfedf3cef79b1a0301b18bb6a86382becf6aa53f1c4

Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            3f:3d:fc:65:2d:d6:bc:ea:dc:70:4f:df
        Signature Algorithm: sha256WithRSAEncryption
        Issuer: C = BE, O = GlobalSign nv-sa, CN = GlobalSign RSA OV SSL CA 2018
        Validity
            Not Before: Jun 19 08:54:52 2019 GMT
            Not After : Jun 19 08:54:52 2021 GMT
        Subject: C = CH, ST = Vaud, L = Lausanne, O = FONDATION ECOLE D'ETUDES SOCIALES ET PEDAGOGIQUES, CN = cuc01-ms.eesp.ch
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                RSA Public-Key: (2048 bit)
                Modulus:
                    00:b6:1b:b9:6a:7f:99:18:a8:1e:8b:43:ff:c4:81:
                    90:9f:e3:42:7a:2f:53:39:bd:e9:6a:d3:7b:24:1c:
                    6b:4f:65:61:35:03:c3:9a:7b:c7:6a:5f:a9:39:7f:
                    0d:82:36:30:ac:03:4b:61:4c:bc:be:33:4c:e4:bb:
                    aa:f9:4b:a6:1b:ef:d8:4d:e1:77:88:89:ad:16:db:
                    7c:0e:fd:b1:de:07:7b:a5:78:a7:a0:9d:4d:55:18:
                    ed:6c:9d:db:a6:c3:01:24:c7:5d:31:0c:93:86:e5:
                    f3:f7:37:f2:31:04:3d:b5:7f:35:6c:bb:17:30:bb:
                    8c:ae:24:6a:b9:57:12:71:97:a9:04:94:fd:8b:b5:
                    06:07:eb:e6:c2:06:c3:73:47:89:6e:a6:42:44:fe:
                    36:4b:fa:76:6d:4c:c7:78:1b:b9:98:75:d4:81:1c:
                    d0:af:57:dd:14:ed:bb:b0:96:10:ff:85:67:e1:c0:
                    e0:d4:b4:34:b1:ef:6f:d9:05:13:ce:71:99:8c:51:
                    12:92:88:60:d5:ee:7d:9c:1b:69:c8:b0:e0:7d:43:
                    05:d8:76:2e:fe:13:8f:46:e5:45:9b:a3:fe:98:af:
                    8e:2d:3d:5b:8a:e1:1e:11:42:92:0e:f6:1f:7a:e3:
                    c9:f5:5c:58:97:b0:10:fb:cd:e8:b6:f3:55:38:ea:
                    8e:29
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Key Usage: critical
                Digital Signature, Key Encipherment
            Authority Information Access: 
                CA Issuers - URI:http://secure.globalsign.com/cacert/gsrsaovsslca2018.crt
                OCSP - URI:http://ocsp.globalsign.com/gsrsaovsslca2018

            X509v3 Certificate Policies: 
                Policy: 1.3.6.1.4.1.4146.1.20
                  CPS: https://www.globalsign.com/repository/
                Policy: 2.23.140.1.2.2

            X509v3 Basic Constraints: 
                CA:FALSE
            X509v3 CRL Distribution Points: 

                Full Name:
                  URI:http://crl.globalsign.com/gsrsaovsslca2018.crl

            X509v3 Subject Alternative Name: 
                DNS:cuc01-ms.eesp.ch, DNS:eesp.ch, DNS:cuc01.eesp.ch, DNS:cuc02.eesp.ch
            X509v3 Extended Key Usage: 
                TLS Web Server Authentication, TLS Web Client Authentication
            X509v3 Authority Key Identifier: 
                keyid:F8:EF:7F:F2:CD:78:67:A8:DE:6F:8F:24:8D:88:F1:87:03:02:B3:EB

            X509v3 Subject Key Identifier: 
                4A:23:C8:49:41:68:67:21:B8:C9:91:D2:3C:7B:F9:E6:2B:76:34:37
            CT Precertificate Poison: critical
                NULL
    Signature Algorithm: sha256WithRSAEncryption
         03:68:b9:11:c0:b9:43:a7:0b:17:55:95:83:30:40:a4:74:31:
         ad:5b:8d:17:8b:26:ee:c3:a0:ce:a8:5f:53:55:34:75:11:33:
         b1:25:58:33:6c:a8:db:e5:7a:40:da:c4:47:a0:3e:77:41:0f:
         7b:29:7c:5d:54:cd:ac:98:f7:e2:7c:9c:f5:92:0f:da:bc:26:
         ad:a7:44:26:b1:93:89:69:01:d8:18:a1:a1:bc:c2:9d:84:27:
         45:c4:01:96:c1:b6:86:95:fe:82:01:75:a5:d0:e4:6e:6b:bb:
         6b:22:15:83:71:67:dc:f2:54:30:90:4d:7b:be:6e:30:11:50:
         3e:9d:94:eb:75:4a:7c:67:ee:d5:bd:3b:8a:db:58:c1:42:1e:
         aa:5c:65:96:5e:83:b6:29:e2:5f:f4:4d:a5:2a:4f:19:01:e8:
         2b:d8:14:16:da:c9:a1:68:15:d5:34:24:b9:4f:eb:d3:6c:1d:
         26:d2:50:3a:0d:b4:f3:fd:cf:ce:91:2e:c4:4c:95:95:0c:3f:
         2b:62:b4:97:8a:41:96:97:97:6a:4c:c0:12:20:9f:ac:87:9c:
         f1:f7:09:f0:f0:43:72:e2:42:f4:ab:5e:33:9c:ec:14:8a:5f:
         e9:3d:8e:f4:aa:dc:5e:b7:41:62:cd:ea:fb:08:1a:c2:01:e5:
         f0:c3:c8:b0
-----BEGIN CERTIFICATE-----
MIIFYDCCBEigAwIBAgIMPz38ZS3WvOrccE/fMA0GCSqGSIb3DQEBCwUAMFAxCzAJ
BgNVBAYTAkJFMRkwFwYDVQQKExBHbG9iYWxTaWduIG52LXNhMSYwJAYDVQQDEx1H
bG9iYWxTaWduIFJTQSBPViBTU0wgQ0EgMjAxODAeFw0xOTA2MTkwODU0NTJaFw0y
MTA2MTkwODU0NTJaMIGGMQswCQYDVQQGEwJDSDENMAsGA1UECBMEVmF1ZDERMA8G
A1UEBxMITGF1c2FubmUxOjA4BgNVBAoTMUZPTkRBVElPTiBFQ09MRSBEJ0VUVURF
UyBTT0NJQUxFUyBFVCBQRURBR09HSVFVRVMxGTAXBgNVBAMTEGN1YzAxLW1zLmVl
c3AuY2gwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQC2G7lqf5kYqB6L
Q//EgZCf40J6L1M5velq03skHGtPZWE1A8Oae8dqX6k5fw2CNjCsA0thTLy+M0zk
u6r5S6Yb79hN4XeIia0W23wO/bHeB3uleKegnU1VGO1sndumwwEkx10xDJOG5fP3
N/IxBD21fzVsuxcwu4yuJGq5VxJxl6kElP2LtQYH6+bCBsNzR4lupkJE/jZL+nZt
TMd4G7mYddSBHNCvV90U7buwlhD/hWfhwODUtDSx72/ZBRPOcZmMURKSiGDV7n2c
G2nIsOB9QwXYdi7+E49G5UWbo/6Yr44tPVuK4R4RQpIO9h9648n1XFiXsBD7zei2
81U46o4pAgMBAAGjggIBMIIB/TAOBgNVHQ8BAf8EBAMCBaAwgY4GCCsGAQUFBwEB
BIGBMH8wRAYIKwYBBQUHMAKGOGh0dHA6Ly9zZWN1cmUuZ2xvYmFsc2lnbi5jb20v
Y2FjZXJ0L2dzcnNhb3Zzc2xjYTIwMTguY3J0MDcGCCsGAQUFBzABhitodHRwOi8v
b2NzcC5nbG9iYWxzaWduLmNvbS9nc3JzYW92c3NsY2EyMDE4MFYGA1UdIARPME0w
QQYJKwYBBAGgMgEUMDQwMgYIKwYBBQUHAgEWJmh0dHBzOi8vd3d3Lmdsb2JhbHNp
Z24uY29tL3JlcG9zaXRvcnkvMAgGBmeBDAECAjAJBgNVHRMEAjAAMD8GA1UdHwQ4
MDYwNKAyoDCGLmh0dHA6Ly9jcmwuZ2xvYmFsc2lnbi5jb20vZ3Nyc2FvdnNzbGNh
MjAxOC5jcmwwQgYDVR0RBDswOYIQY3VjMDEtbXMuZWVzcC5jaIIHZWVzcC5jaIIN
Y3VjMDEuZWVzcC5jaIINY3VjMDIuZWVzcC5jaDAdBgNVHSUEFjAUBggrBgEFBQcD
AQYIKwYBBQUHAwIwHwYDVR0jBBgwFoAU+O9/8s14Z6jeb48kjYjxhwMCs+swHQYD
VR0OBBYEFEojyElBaGchuMmR0jx7+eYrdjQ3MBMGCisGAQQB1nkCBAMBAf8EAgUA
MA0GCSqGSIb3DQEBCwUAA4IBAQADaLkRwLlDpwsXVZWDMECkdDGtW40Xiybuw6DO
qF9TVTR1ETOxJVgzbKjb5XpA2sRHoD53QQ97KXxdVM2smPfifJz1kg/avCatp0Qm
sZOJaQHYGKGhvMKdhCdFxAGWwbaGlf6CAXWl0ORua7trIhWDcWfc8lQwkE17vm4w
EVA+nZTrdUp8Z+7VvTuK21jBQh6qXGWWXoO2KeJf9E2lKk8ZAegr2BQW2smhaBXV
NCS5T+vTbB0m0lA6DbTz/c/OkS7ETJWVDD8rYrSXikGWl5dqTMASIJ+sh5zx9wnw
8ENy4kL0q14znOwUil/pPY70qtxet0Fizer7CBrCAeXww8iw
-----END CERTIFICATE-----

+ View on Censys: https://censys.io/certificates/0037ae7546555efca0935dfedf3cef79b1a0301b18bb6a86382becf6aa53f1c4

```

If we wanted to step through the linter in question in a debugger when it's
linting this certificate we could run the integration tests again specifying a
`-fingerprintFilter` that limits linting to the certificate we're interested
in:

* `make integration PARALLELISM=6 INT_FLAGS="-fingerprintSummary -lintFilter='e_subject_printable_string_badalpha' -fingerprintFilter='0037ae7546555efca0935dfedf3cef79b1a0301b18bb6a86382becf6aa53f1c4'"`

By spot-checking a few of the new 221 certificate fingerprints with
`certByFP.sh` and with `-lintFilter/-fingerprintFilter` we're likely to notice
that all of the certificates causing new error results have a `'` character in
their PrintableString encoded Subjects, which should be allowed.

The `'` character being omitted from the regexp used by the
`e_subject_printable_string_badalpha` lint was the root cause of the bugfix we
reverted and so the integration tests have done the right thing and flagged an
unintended regression.

Adding a new lint
-----------------

Adding a new lint is very similar to the process undertaken above while
debugging an integration test failure.

After adding your lint the integration tests can be run to see which of the
existing test corpus certificates are flagged by the new linter. Because there
is no expected data for the new lint, the integration tests will fail unless
there are no info level or higher findings from your new lint across the whole
test corpus.

If your lint has findings in the corpus you can see which certificates
fingerprints tripped the new lint by using the `-serialSummary` flag with
a `-lintFilter`. Spot check the flagged certificates with `certByFP.sh` and any
other other required techniques until you're certain the new lint is operating
correctly.

Once you're confident the observed results match expectations you can add the
new lint results to the expected data by running the integration tests with
`-overwriteExpected` and committing the updated config file along with your new
lint. Nice work!
