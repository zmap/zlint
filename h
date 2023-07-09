[33mcommit 92902fc7d9ae7ad9f221235c74b992be6f101812[m[33m ([m[1;36mHEAD -> [m[1;32mmaster[m[33m, [m[1;31morigin/master[m[33m, [m[1;31morigin/HEAD[m[33m)[m
Merge: 526f9be 8c46bdf
Author: mtgag <githreg@mtg.de>
Date:   Sat Jul 1 09:28:04 2023 +0200

    Merge https://github.com/zmap/zlint

[33mcommit 8c46bdf0e6c8f3ccab7d3101cbf56eea9b7a856a[m
Author: Aaron Gable <aaron@letsencrypt.org>
Date:   Fri Jun 30 12:56:49 2023 -0700

    Fix typo in LintRevocationListEx comment (#730)

[33mcommit 7ef1f8451ba9894bb27645321618de2bf9a158be[m
Author: github-actions[bot] <41898282+github-actions[bot]@users.noreply.github.com>
Date:   Sun Jun 25 16:11:22 2023 -0700

    util: gtld_map autopull updates for 2023-06-14T22:18:50 UTC (#727)
    
    Co-authored-by: GitHub <noreply@github.com>
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit 5e0219d2a818f0d8c71f20191d79e010890c2269[m
Author: MTG <36234449+mtgag@users.noreply.github.com>
Date:   Mon Jun 26 01:02:29 2023 +0200

    Bc critical (#722)
    
    * lint about the encoding of qcstatements for PSD2
    
    * Revert "lint about the encoding of qcstatements for PSD2"
    
    This reverts commit 6c2367080d148f4b8c01f96a4c80e3ac55d1ef26.
    
    * util: gtld_map autopull updates for 2021-10-21T07:25:20 UTC
    
    * always check and perform the operation in the execution
    
    * returning fatal rather than na
    
    * Update v3/lints/rfc/lint_basic_constraints_not_critical.go
    
    Error instead of fatal
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>
    
    * adding error description.
    
    ---------
    
    Co-authored-by: mtg <git@mtg.de>
    Co-authored-by: GitHub <noreply@github.com>
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit 3746088f87cde72a751b8f8a68c9b0a9e9a6a8b0[m
Author: github-actions[bot] <41898282+github-actions[bot]@users.noreply.github.com>
Date:   Sun Jun 11 12:21:00 2023 -0700

    util: gtld_map autopull updates for 2023-06-06T18:20:14 UTC (#698)
    
    Co-authored-by: GitHub <noreply@github.com>
    Co-authored-by: Zakir Durumeric <zakird@gmail.com>

[33mcommit 9b18bdcd8fedb5013bda10ba13de27e3bf4ed908[m
Author: MTG <36234449+mtgag@users.noreply.github.com>
Date:   Sun Jun 11 21:13:48 2023 +0200

    Ca field empty description (#723)
    
    * lint about the encoding of qcstatements for PSD2
    
    * Revert "lint about the encoding of qcstatements for PSD2"
    
    This reverts commit 6c2367080d148f4b8c01f96a4c80e3ac55d1ef26.
    
    * util: gtld_map autopull updates for 2021-10-21T07:25:20 UTC
    
    * always check and perform the operation in the execution
    
    * simply must not have a non-empty distinguished name should suffice. The field is always present, the lints tests if the Sequence is empty.
    
    ---------
    
    Co-authored-by: mtg <git@mtg.de>
    Co-authored-by: GitHub <noreply@github.com>
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit 59a91a2b1b7562e80894103cf8f8e03319b82a92[m
Author: MTG <36234449+mtgag@users.noreply.github.com>
Date:   Sun Jun 11 21:02:42 2023 +0200

    Max length check applies (#724)
    
    * lint about the encoding of qcstatements for PSD2
    
    * Revert "lint about the encoding of qcstatements for PSD2"
    
    This reverts commit 6c2367080d148f4b8c01f96a4c80e3ac55d1ef26.
    
    * util: gtld_map autopull updates for 2021-10-21T07:25:20 UTC
    
    * always check and perform the operation in the execution
    
    * max length check only if component is present.
    
    ---------
    
    Co-authored-by: mtg <git@mtg.de>
    Co-authored-by: GitHub <noreply@github.com>
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit 526f9be2c26b63477a2d03d8a6a2736e2fe89b72[m
Merge: b52111b 45e8dff
Author: mtgag <githreg@mtg.de>
Date:   Fri Jun 9 06:52:40 2023 +0200

    Merge https://github.com/zmap/zlint

[33mcommit 45e8dff6fe0d2a6989366a3dbd44713c360afc8f[m
Author: mwahaj <mwahaj@mail.com>
Date:   Sun Jun 4 23:13:06 2023 +0500

    Update README.md (#719)
    
    Added PKI Insights which also used zlint for X.509 Certificate verification against the PKI and Industry standards
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit af903824a31385208566fa640cc13036a0e4d8e4[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sun Jun 4 11:02:45 2023 -0700

    Enable accepting a PEM encoded CRL via the command line interface (#721)
    
    * dispatching CRLs to the CRL linting infra
    
    * fixing typo in README

[33mcommit 1d8591cffbd9513c7302ef8187297e7463358291[m
Author: toddgaunt-gs <107932811+toddgaunt-gs@users.noreply.github.com>
Date:   Mon May 29 12:05:30 2023 -0400

    Remove references in comments to Initialize() method of lints (#718)
    
    Some comments still refer to lints having an Initialize method. This
    appears to no longer be the case but a warning in the comments for
    RegisterLint, RegisterCertificateLint, and RegisterRevocationListLint
    was still referencing lints having such a method.

[33mcommit b52111baec7700cadeafd21ca74e448cec162483[m
Merge: 351a379 2438596
Author: mtgag <githreg@mtg.de>
Date:   Tue May 16 08:44:04 2023 +0200

    Merge https://github.com/zmap/zlint

[33mcommit 24385962110d84a33e403ae611169297e8d205c1[m
Author: MTG <36234449+mtgag@users.noreply.github.com>
Date:   Sun May 14 20:16:08 2023 +0200

    Always perform e_cert_unique_identifier_version_not_2_or_3  (#711)
    
    * lint about the encoding of qcstatements for PSD2
    
    * Revert "lint about the encoding of qcstatements for PSD2"
    
    This reverts commit 6c2367080d148f4b8c01f96a4c80e3ac55d1ef26.
    
    * util: gtld_map autopull updates for 2021-10-21T07:25:20 UTC
    
    * always check and perform the operation in the execution
    
    ---------
    
    Co-authored-by: mtg <git@mtg.de>
    Co-authored-by: GitHub <noreply@github.com>
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit 351a37987e16c681f69725836a73dc888179d2be[m
Merge: 92e659c a5c869f
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sun May 14 11:06:52 2023 -0700

    Merge branch 'master' into master

[33mcommit a5c869f807cbfce8a689aeba5682eb8f326845ea[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sat May 13 09:23:45 2023 -0700

    Update copyright text to 2023 (#716)
    
    * Updating copyright headers to 2023

[33mcommit 92e659c5aefeeea3afd8a32cc768b112a9355218[m
Author: mtgag <githreg@mtg.de>
Date:   Thu Apr 27 08:55:54 2023 +0200

    always check and perform the operation in the execution

[33mcommit 30b096ee5b613af5eff751d9c5b878e8d07f529e[m
Merge: 8600050 997ad51
Author: mtgag <githreg@mtg.de>
Date:   Wed Apr 19 08:41:37 2023 +0200

    Merge https://github.com/zmap/zlint

[33mcommit 997ad5143216f4a3f461545f277be7e20bdcb557[m
Author: Amir Omidi <amir@aaomidi.com>
Date:   Sun Mar 26 14:02:27 2023 -0400

    Add CRL linting infrastructure (#699)
    
    * Add the skeleton around linting CRLs
    
    * Change the entrypoint of zlint
    
    * Add tests for the new skeleton
    
    * Address reviews
    
    * starting my own suggestions to work coopertaively on he change
    
    * Take out generics from the registration struct (#3)
    
    * Update to use Zcrypto instead of stdlib crypto for RevocationList (#4)
    
    * Take out generics from the registration struct (#3)
    
    * updating to use zcrypto
    
    * pointing zcrypto back to master
    
    * go tidy up
    
    ---------
    
    Co-authored-by: Amir Omidi <amir@aaomidi.com>
    
    * Tidy go mod
    
    * Update zcrypto
    
    * go mod tidy one more time
    
    * Bypass lint for Registry
    
    * Add NextUpdate CRL lint (#5)
    
    ---------
    
    Co-authored-by: christopher-henderson <chris@chenderson.org>

[33mcommit 64ae4e500e020b535a475a6c99007f77b917e1e9[m
Author: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>
Date:   Sun Mar 12 13:06:18 2023 -0700

    build(deps): bump golang.org/x/net in /v3/cmd/genTestCerts (#704)
    
    Bumps [golang.org/x/net](https://github.com/golang/net) from 0.0.0-20220412020605-290c469a71a5 to 0.7.0.
    - [Release notes](https://github.com/golang/net/releases)
    - [Commits](https://github.com/golang/net/commits/v0.7.0)
    
    ---
    updated-dependencies:
    - dependency-name: golang.org/x/net
      dependency-type: indirect
    ...
    
    Signed-off-by: dependabot[bot] <support@github.com>
    Co-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit 68901ea435cd9be1c5f37765ed178120c3f570f9[m
Author: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>
Date:   Sun Mar 12 12:58:25 2023 -0700

    build(deps): bump golang.org/x/net in /v3 (#702)
    
    Bumps [golang.org/x/net](https://github.com/golang/net) from 0.0.0-20220412020605-290c469a71a5 to 0.7.0.
    - [Release notes](https://github.com/golang/net/releases)
    - [Commits](https://github.com/golang/net/commits/v0.7.0)
    
    ---
    updated-dependencies:
    - dependency-name: golang.org/x/net
      dependency-type: direct:production
    ...
    
    Signed-off-by: dependabot[bot] <support@github.com>
    Co-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit 5ed8e34fe97edb3fedd7f1fb5cbc48a1444ea195[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sun Mar 12 12:48:34 2023 -0700

    asserting human readable strings is error prone (#707)

[33mcommit c7740fad1793b30df07212f9297066363efb19ce[m
Author: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>
Date:   Sun Mar 12 12:32:52 2023 -0700

    build(deps): bump golang.org/x/text in /v3/cmd/genTestCerts (#701)
    
    Bumps [golang.org/x/text](https://github.com/golang/text) from 0.3.7 to 0.3.8.
    - [Release notes](https://github.com/golang/text/releases)
    - [Commits](https://github.com/golang/text/compare/v0.3.7...v0.3.8)
    
    ---
    updated-dependencies:
    - dependency-name: golang.org/x/text
      dependency-type: indirect
    ...
    
    Signed-off-by: dependabot[bot] <support@github.com>
    Co-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit a476724019152fa17e7ebb3c0bba6b896aecf89d[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sun Mar 12 10:55:47 2023 -0700

    Upgrading golangci-lint to v1.51.2 (#705)

[33mcommit 46f7185e35ed0a7af55db60004a66ac4f15520fa[m
Author: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>
Date:   Sun Mar 5 09:18:23 2023 -0800

    build(deps): bump golang.org/x/text from 0.3.7 to 0.3.8 in /v3 (#700)
    
    Bumps [golang.org/x/text](https://github.com/golang/text) from 0.3.7 to 0.3.8.
    - [Release notes](https://github.com/golang/text/releases)
    - [Commits](https://github.com/golang/text/compare/v0.3.7...v0.3.8)
    
    ---
    updated-dependencies:
    - dependency-name: golang.org/x/text
      dependency-type: direct:production
    ...
    
    Signed-off-by: dependabot[bot] <support@github.com>
    Co-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>

[33mcommit 8a9f61eb9d9b2ee4b14519573ee2f0d09474c316[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Thu Nov 3 09:18:18 2022 -0700

    test.ReadTestCert breaks for downstream consumers dependent on the previous relative certificate path building behavior (#695)
    
    * util: gtld_map autopull updates for 2022-10-06T19:22:06 UTC
    
    * Trigger GHA
    
    * revert change
    
    * fixing our own tests
    
    Co-authored-by: GitHub <noreply@github.com>

[33mcommit 6292ca4c07afed0c9e4f43470126901161fd0c2c[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sun Oct 16 11:41:20 2022 -0700

    Adding support for linting profiles (#595)
    
    * adding support for linting profiles
    
    * at least tests running
    
    * Update v3/lint/profile.go
    
    Absolutely
    
    Co-authored-by: Daniel McCarney <daniel@binaryparadox.net>
    
    * Update v3/newProfile.sh
    
    * adding godoc to AllProfiles
    
    * util: gtld_map autopull updates for 2022-10-06T19:22:06 UTC
    
    * Trigger GHA
    
    * fixing linter
    
    Co-authored-by: Daniel McCarney <daniel@binaryparadox.net>
    Co-authored-by: GitHub <noreply@github.com>

[33mcommit c6273337f37bce57a42c61f61566465ba81a8f4d[m
Author: github-actions[bot] <41898282+github-actions[bot]@users.noreply.github.com>
Date:   Sun Oct 16 10:20:03 2022 -0700

    util: gtld_map autopull updates for 2022-10-10T19:22:35 UTC (#694)
    
    Co-authored-by: GitHub <noreply@github.com>

[33mcommit 13fcc6ff15096c615205e0073681d571227522f9[m
Author: github-actions[bot] <41898282+github-actions[bot]@users.noreply.github.com>
Date:   Sun Oct 9 07:06:19 2022 -0700

    util: gtld_map autopull updates for 2022-10-06T19:22:06 UTC (#693)
    
    Co-authored-by: GitHub <noreply@github.com>

[33mcommit 137e46e0ca400af8c38465773a9d9ef8dc044b62[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sun Sep 18 11:18:06 2022 -0700

    Lint to check for invalid KU lengths (#686)
    
    * lint for incorrecty KU length
    
    * better code comment
    
    * correcting linter
    
    * fixing lint to check for combinations with nine possible flags
    
    * fixing comments
    
    * using cryptobyte
    
    * accounting for jumbo sized KUs

[33mcommit 1209017ea441820ff41f4ef6b05e946ed53efcda[m
Author: Rob <3725956+robplee@users.noreply.github.com>
Date:   Sun Sep 18 19:08:44 2022 +0100

    Prevent OU lint from applying to CA certificates.  Add unit test to confirm change of behaviour (#691)

[33mcommit 44e12c12ca43a4af86f0dc2da4a71493ac9f8345[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sun Aug 28 07:33:00 2022 -0700

    Add lint to check for incorrect 'unused' bit encoding in KeyUsages (#684)
    
    * Add lint to check for incorrect 'unused' bit encoding
    
    * using real life test data as a failure case

[33mcommit 3f5e40d69c7dd1ed2049051f00dba88e97794ef0[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sun Jul 31 11:02:44 2022 -0700

    Lint for RSA close prime Fermat factorization susceptibility (#674)
    
    * lint for close prime factorization with a default round setting of 100

[33mcommit e5ee614b989dca0615c7fdb9cb6d621f281c5a20[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sat Jul 23 11:55:36 2022 -0700

    Support for Configurable Lints (#648)
    
    * Support for configurable lints

[33mcommit ed9a20f851f487d6d280b72dc9db232779fc11e3[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sun Jul 17 13:06:32 2022 -0700

    Added lint to check for superfluous zero byte on KU (#682)

[33mcommit d8b86f771ea068173826b2088f0c502c17eaaa8d[m
Author: MTG <36234449+mtgag@users.noreply.github.com>
Date:   Sun Jun 19 19:58:35 2022 +0200

    Lints for allowable key usages as per RFC 8813 Section 3 and RFC 3279 Section 2.3.1 (#678)
    
    * lint about the encoding of qcstatements for PSD2
    
    * Revert "lint about the encoding of qcstatements for PSD2"
    
    This reverts commit 6c2367080d148f4b8c01f96a4c80e3ac55d1ef26.
    
    * util: gtld_map autopull updates for 2021-10-21T07:25:20 UTC
    
    * added lints that adress issues about correct key usage values for a certain public key type
    
    * adjustments in config.json
    
    * adjustments after code review
    
    * adjustments after code review
    
    * warnings are turned to errors
    
    * fixed error count
    
    Co-authored-by: mtg <git@mtg.de>
    Co-authored-by: GitHub <noreply@github.com>

[33mcommit c7955ed482857439faa68dfdfb67b94a1510bce1[m
Author: MTG <36234449+mtgag@users.noreply.github.com>
Date:   Mon Jun 13 16:19:30 2022 +0200

    Sunset subject:organizationalUnitName (Section 7.1.4.2.2.i, CAB-Forum BR) (#643)
    
    * lint about the encoding of qcstatements for PSD2
    
    * Revert "lint about the encoding of qcstatements for PSD2"
    
    This reverts commit 6c2367080d148f4b8c01f96a4c80e3ac55d1ef26.
    
    * util: gtld_map autopull updates for 2021-10-21T07:25:20 UTC
    
    * added lint for presence of OU in subject
    
    * Update v3/lints/cabf_br/lint_subject_contains_organizational_unit_name.go
    
    Co-authored-by: Ryan Sleevi <ryan.sleevi@gmail.com>
    
    * separated lints to adress two requirements
    
    * separated lints to adress two requirements
    
    * reverted change proposed by IDE
    
    * aligning to #644
    
    * Update v3/util/time.go
    
    * Update v3/util/time.go
    
    * Update v3/util/time.go
    
    * addressed requested changes, removing lint that is implemented in 675
    
    Co-authored-by: mtg <git@mtg.de>
    Co-authored-by: GitHub <noreply@github.com>
    Co-authored-by: Ryan Sleevi <ryan.sleevi@gmail.com>
    Co-authored-by: Christopher Henderson <chris@chenderson.org>
