[33mcommit 9bd48d488919a55a0e2228677ef57b3e1f273ba9[m[33m ([m[1;36mHEAD -> [m[1;32mqc_type_web_also_smime[m[33m)[m
Author: mtgag <githreg@mtg.de>
Date:   Thu Feb 27 09:35:06 2025 +0100

    updated test config for new lint

[33mcommit 1b69c403f1e1cde9bcb193d043f6da64706d1556[m
Author: mtgag <githreg@mtg.de>
Date:   Thu Feb 27 09:26:39 2025 +0100

    considering SMIME certificates

[33mcommit d5aec9bd93fa361e4a276a4c0ece53d286303193[m[33m ([m[1;31morigin/master[m[33m, [m[1;31morigin/HEAD[m[33m, [m[1;32mmaster[m[33m)[m
Author: mtgag <githreg@mtg.de>
Date:   Tue Feb 18 15:33:26 2025 +0100

    synchronised with project

[33mcommit 6662edf0e1159a5c025e284c47ae58e123302e71[m
Merge: f0991f98 04d863f7
Author: mtgag <githreg@mtg.de>
Date:   Tue Jun 18 05:55:48 2024 +0200

    Merge https://github.com/zmap/zlint

[33mcommit 04d863f7660edfe0498162334524742397226fb2[m
Author: Martijn Katerbarg <martijn.katerbarg@sectigo.com>
Date:   Mon Jun 17 16:17:27 2024 +0200

    cabfOrganizationIdentifier extension for VAT and PSD based organizationIdentifiers cannot have referenceStateOrProvince (#848)
    
    * cabfOrganizationIdentifier lint for PSD based QWAC certificates
    
    * cabfOrganizationIdentifier referenceStateOrProvince lint for PSD and VAT based QWAC certificates
    
    * Provide Bad Test Cert
    
    * Add "e_" to lint name
    
    * Also add "e_" to test case
    
    * Update lint_cabf_org_identifier_psd_vat_has_state.go
    
    * Reference v1.7.0 section 9.2.8
    
    ---------
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit e5da476b15be77968e50510f819a54ab1fa3b952[m
Author: Adriano Santoni <asantoni64@gmail.com>
Date:   Sun Jun 16 21:23:03 2024 +0200

    Improve the util.IsServerAuthCert() function (#856)
    
    * Add files via upload
    
    * Add files via upload
    
    * Add files via upload
    
    * Add files via upload
    
    * Update lint_invalid_subject_rdn_order_test.go
    
    Added //nolint:all to comment block to avoid golangci-lint to complain about duplicate words in comment
    
    * Update lint_invalid_subject_rdn_order.go
    
    Fixed import block
    
    * Update v3/lints/cabf_br/lint_invalid_subject_rdn_order.go
    
    Fine to me.
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>
    
    * Update lint_invalid_subject_rdn_order.go
    
    As per Chris Henderson's suggestion, to "improve readability".
    
    * Update lint_invalid_subject_rdn_order_test.go
    
    As per Chris Henderson's suggestion.
    
    * Update time.go
    
    Added CABFEV_Sec9_2_8_Date
    
    * Add files via upload
    
    * Add files via upload
    
    * Revised according to Chris and Corey suggestions
    
    * Add files via upload
    
    * Add files via upload
    
    * Delete v3/lints/cabf_br/lint_e_invalid_cps_uri.go
    
    * Delete v3/lints/cabf_br/lint_e_invalid_cps_uri_test.go
    
    * Delete v3/testdata/invalid_cps_uri_ko_01.pem
    
    * Delete v3/testdata/invalid_cps_uri_ko_02.pem
    
    * Delete v3/testdata/invalid_cps_uri_ko_03.pem
    
    * Delete v3/testdata/invalid_cps_uri_ok_01.pem
    
    * Delete v3/testdata/invalid_cps_uri_ok_02.pem
    
    * Delete v3/testdata/invalid_cps_uri_ok_03.pem
    
    * Update ca.go
    
    * Update config.json
    
    * Update config.json
    
    ---------
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit 5b73e7b8fdcbabe138c745f1e6151fb18737f3c6[m
Author: Mathew Hodson <mathew.hodson@gmail.com>
Date:   Sun Jun 16 15:12:28 2024 -0400

    Fix ExpectedDetails of passing invalid subject test (#846)
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit 899709e95046383a8f6bdd52bd61c45b9eab279e[m
Author: MTG <36234449+mtgag@users.noreply.github.com>
Date:   Sun Jun 16 20:22:02 2024 +0200

    Aia ca issuers must have http only (#852)
    
    * lint about the encoding of qcstatements for PSD2
    
    * Revert "lint about the encoding of qcstatements for PSD2"
    
    This reverts commit 6c2367080d148f4b8c01f96a4c80e3ac55d1ef26.
    
    * util: gtld_map autopull updates for 2021-10-21T07:25:20 UTC
    
    * always check and perform the operation in the execution
    
    * synchronised with project
    
    * synchronised with project
    
    * synchronised with project
    
    * synchronised with project
    
    * fixed merge error
    
    * synchronised with project
    
    * synchronised with project
    
    * Revert "synchronised with project"
    
    This reverts commit bad73ee2d5669394cde3053d300f285a91f75fd6.
    
    * Revert "synchronised with project"
    
    This reverts commit 2cd7d087f4a812d4ef3640560edf1d07cce2ea56.
    
    * new lint; The id-ad-caIssuers accessMethod must contain an HTTP URL of the Issuing CA’s certificate; removed unnecessary assertion from older lint.
    
    * update to consider HTTPS (not only HTTP) URLs also.
    
    * this is already covered by PR #846
    
    * addressing issues in PR discussion
    
    ---------
    
    Co-authored-by: mtg <git@mtg.de>
    Co-authored-by: GitHub <noreply@github.com>
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit ae8d59405f1926eb418d496cd0415b8a4fa88e04[m
Author: github-actions[bot] <41898282+github-actions[bot]@users.noreply.github.com>
Date:   Wed Jun 12 23:52:44 2024 -0400

    util: gtld_map autopull updates for 2024-06-12T22:19:30 UTC (#854)
    
    Co-authored-by: GitHub <noreply@github.com>

[33mcommit b14a83bb192056a51b26cb9d66370fa7d978f373[m
Author: Martijn Katerbarg <martijn.katerbarg@sectigo.com>
Date:   Sun Jun 9 19:30:35 2024 +0200

    fix: Only apply CN check for Subscriber certificates (#851)
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit bf3764c2b225b5942bea7e96de3d429e1a4cb093[m
Author: Phil Porada <pgporada@users.noreply.github.com>
Date:   Sun Jun 9 13:21:44 2024 -0400

    Cleanup some unnecessary allocations (#849)
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit f0991f985f5ceb4797b7de39c236bdfe62c47ed3[m
Merge: 4d467299 26ca0f3b
Author: mtgag <githreg@mtg.de>
Date:   Thu Jun 6 07:11:46 2024 +0200

    Merge https://github.com/zmap/zlint

[33mcommit 26ca0f3bed092ef6e6b34f546f68068ae9d808a1[m
Author: Adriano Santoni <asantoni64@gmail.com>
Date:   Sun Jun 2 22:07:35 2024 +0200

    Add lint to check for duplicate subject attributes (ATVs) (#850)
    
    * Add files via upload
    
    * Add files via upload
    
    * Add files via upload
    
    * Add files via upload
    
    * Update lint_invalid_subject_rdn_order_test.go
    
    Added //nolint:all to comment block to avoid golangci-lint to complain about duplicate words in comment
    
    * Update lint_invalid_subject_rdn_order.go
    
    Fixed import block
    
    * Update v3/lints/cabf_br/lint_invalid_subject_rdn_order.go
    
    Fine to me.
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>
    
    * Update lint_invalid_subject_rdn_order.go
    
    As per Chris Henderson's suggestion, to "improve readability".
    
    * Update lint_invalid_subject_rdn_order_test.go
    
    As per Chris Henderson's suggestion.
    
    * Update time.go
    
    Added CABFEV_Sec9_2_8_Date
    
    * Add files via upload
    
    * Add files via upload
    
    * Revised according to Chris and Corey suggestions
    
    * Add files via upload
    
    * Add files via upload
    
    * Delete v3/lints/cabf_br/lint_e_invalid_cps_uri.go
    
    * Delete v3/lints/cabf_br/lint_e_invalid_cps_uri_test.go
    
    * Delete v3/testdata/invalid_cps_uri_ko_01.pem
    
    * Delete v3/testdata/invalid_cps_uri_ko_02.pem
    
    * Delete v3/testdata/invalid_cps_uri_ko_03.pem
    
    * Delete v3/testdata/invalid_cps_uri_ok_01.pem
    
    * Delete v3/testdata/invalid_cps_uri_ok_02.pem
    
    * Delete v3/testdata/invalid_cps_uri_ok_03.pem
    
    * Add files via upload
    
    * Add files via upload
    
    ---------
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit c8164d8a086ff6e3dd419b2ace95784d32f49c57[m
Author: Adriano Santoni <asantoni64@gmail.com>
Date:   Sun May 26 18:32:55 2024 +0200

    Add lint to check that SubCA certificates do not have illegal values in their EKU extension (#840)
    
    * Add files via upload
    
    * Add files via upload
    
    * Add files via upload
    
    * Add files via upload
    
    * Update lint_invalid_subject_rdn_order_test.go
    
    Added //nolint:all to comment block to avoid golangci-lint to complain about duplicate words in comment
    
    * Update lint_invalid_subject_rdn_order.go
    
    Fixed import block
    
    * Update v3/lints/cabf_br/lint_invalid_subject_rdn_order.go
    
    Fine to me.
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>
    
    * Update lint_invalid_subject_rdn_order.go
    
    As per Chris Henderson's suggestion, to "improve readability".
    
    * Update lint_invalid_subject_rdn_order_test.go
    
    As per Chris Henderson's suggestion.
    
    * Update time.go
    
    Added CABFEV_Sec9_2_8_Date
    
    * Add files via upload
    
    * Add files via upload
    
    * Revised according to Chris and Corey suggestions
    
    * Add files via upload
    
    * Add files via upload
    
    * Delete v3/lints/cabf_br/lint_e_invalid_cps_uri.go
    
    * Delete v3/lints/cabf_br/lint_e_invalid_cps_uri_test.go
    
    * Delete v3/testdata/invalid_cps_uri_ko_01.pem
    
    * Delete v3/testdata/invalid_cps_uri_ko_02.pem
    
    * Delete v3/testdata/invalid_cps_uri_ko_03.pem
    
    * Delete v3/testdata/invalid_cps_uri_ok_01.pem
    
    * Delete v3/testdata/invalid_cps_uri_ok_02.pem
    
    * Delete v3/testdata/invalid_cps_uri_ok_03.pem
    
    * Add files via upload
    
    * Add files via upload
    
    * Update config.json
    
    * Update lint_ca_invalid_eku.go
    
    ---------
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit 068ae82324696a6f484be9baa6085318e7851112[m
Author: MTG <36234449+mtgag@users.noreply.github.com>
Date:   Sat May 25 17:12:33 2024 +0200

    Avoid warning dv cn (#843)
    
    * lint about the encoding of qcstatements for PSD2
    
    * Revert "lint about the encoding of qcstatements for PSD2"
    
    This reverts commit 6c2367080d148f4b8c01f96a4c80e3ac55d1ef26.
    
    * util: gtld_map autopull updates for 2021-10-21T07:25:20 UTC
    
    * always check and perform the operation in the execution
    
    * synchronised with project
    
    * synchronised with project
    
    * synchronised with project
    
    * synchronised with project
    
    * fixed merge error
    
    * synchronised with project
    
    * synchronised with project
    
    * Revert "synchronised with project"
    
    This reverts commit bad73ee2d5669394cde3053d300f285a91f75fd6.
    
    * Revert "synchronised with project"
    
    This reverts commit 2cd7d087f4a812d4ef3640560edf1d07cce2ea56.
    
    * avoiding warning when CN is present.
    
    ---------
    
    Co-authored-by: mtg <git@mtg.de>
    Co-authored-by: GitHub <noreply@github.com>
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit 8523152e2c47c83321a145b1e777a9996bd714dd[m
Author: Rob Stradling <rob@sectigo.com>
Date:   Fri May 24 22:58:46 2024 +0100

    Fix handling of Subject:commonName not present in lint for BR 7.1.4.2.2a mailbox-validated (#845)
    
    * Fix handling of Subject:commonName not present in lint for BR 7.1.4.2.2a mailbox-validated
    
    * Add test case for no commonName attribute present

[33mcommit 456dc01dad591ddaaf005f6a955fbca032379c0f[m
Author: Adriano Santoni <asantoni64@gmail.com>
Date:   Sun May 19 20:09:35 2024 +0200

    Add lint to check that an SCT list is not empty  (#837)
    
    * Add files via upload
    
    * Add files via upload
    
    * Add files via upload
    
    * Add files via upload
    
    * Update lint_invalid_subject_rdn_order_test.go
    
    Added //nolint:all to comment block to avoid golangci-lint to complain about duplicate words in comment
    
    * Update lint_invalid_subject_rdn_order.go
    
    Fixed import block
    
    * Update v3/lints/cabf_br/lint_invalid_subject_rdn_order.go
    
    Fine to me.
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>
    
    * Update lint_invalid_subject_rdn_order.go
    
    As per Chris Henderson's suggestion, to "improve readability".
    
    * Update lint_invalid_subject_rdn_order_test.go
    
    As per Chris Henderson's suggestion.
    
    * Update time.go
    
    Added CABFEV_Sec9_2_8_Date
    
    * Add files via upload
    
    * Add files via upload
    
    * Revised according to Chris and Corey suggestions
    
    * Add files via upload
    
    * Add files via upload
    
    * Delete v3/lints/cabf_br/lint_e_invalid_cps_uri.go
    
    * Delete v3/lints/cabf_br/lint_e_invalid_cps_uri_test.go
    
    * Delete v3/testdata/invalid_cps_uri_ko_01.pem
    
    * Delete v3/testdata/invalid_cps_uri_ko_02.pem
    
    * Delete v3/testdata/invalid_cps_uri_ko_03.pem
    
    * Delete v3/testdata/invalid_cps_uri_ok_01.pem
    
    * Delete v3/testdata/invalid_cps_uri_ok_02.pem
    
    * Delete v3/testdata/invalid_cps_uri_ok_03.pem
    
    * Add files via upload
    
    * Add files via upload
    
    * Add files via upload
    
    * Add files via upload
    
    * Add files via upload
    
    * Add files via upload
    
    * Add files via upload
    
    ---------
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit c73f78bfa648887dffe592f02fd6519b514fbb36[m
Author: Adriano Santoni <asantoni64@gmail.com>
Date:   Sun May 19 19:09:17 2024 +0200

    Add lint to check that precertificates do not contain an SCT list (#841)
    
    * Add files via upload
    
    * Add files via upload
    
    * Add files via upload
    
    * Add files via upload
    
    * Update lint_invalid_subject_rdn_order_test.go
    
    Added //nolint:all to comment block to avoid golangci-lint to complain about duplicate words in comment
    
    * Update lint_invalid_subject_rdn_order.go
    
    Fixed import block
    
    * Update v3/lints/cabf_br/lint_invalid_subject_rdn_order.go
    
    Fine to me.
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>
    
    * Update lint_invalid_subject_rdn_order.go
    
    As per Chris Henderson's suggestion, to "improve readability".
    
    * Update lint_invalid_subject_rdn_order_test.go
    
    As per Chris Henderson's suggestion.
    
    * Update time.go
    
    Added CABFEV_Sec9_2_8_Date
    
    * Add files via upload
    
    * Add files via upload
    
    * Revised according to Chris and Corey suggestions
    
    * Add files via upload
    
    * Add files via upload
    
    * Delete v3/lints/cabf_br/lint_e_invalid_cps_uri.go
    
    * Delete v3/lints/cabf_br/lint_e_invalid_cps_uri_test.go
    
    * Delete v3/testdata/invalid_cps_uri_ko_01.pem
    
    * Delete v3/testdata/invalid_cps_uri_ko_02.pem
    
    * Delete v3/testdata/invalid_cps_uri_ko_03.pem
    
    * Delete v3/testdata/invalid_cps_uri_ok_01.pem
    
    * Delete v3/testdata/invalid_cps_uri_ok_02.pem
    
    * Delete v3/testdata/invalid_cps_uri_ok_03.pem
    
    * Add files via upload
    
    * Add files via upload
    
    * Add files via upload
    
    * Add files via upload
    
    * Update lint_precert_with_sct_list.go
    
    * Update source.go
    
    As per Chris' request
    
    * Update source.go
    
    * Update registration_test.go
    
    * Update registration_test.go
    
    ---------
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit 26ab5b0a05d2a70c1a5e98c38fc8a08794fcf950[m
Author: Adriano Santoni <asantoni64@gmail.com>
Date:   Sat May 11 20:04:08 2024 +0200

    Add lint for checking that the 'critical' field is properly DER-encoded in extensions (#839)
    
    * Add files via upload
    
    * Add files via upload
    
    * Add files via upload
    
    * Add files via upload
    
    * Update lint_invalid_subject_rdn_order_test.go
    
    Added //nolint:all to comment block to avoid golangci-lint to complain about duplicate words in comment
    
    * Update lint_invalid_subject_rdn_order.go
    
    Fixed import block
    
    * Update v3/lints/cabf_br/lint_invalid_subject_rdn_order.go
    
    Fine to me.
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>
    
    * Update lint_invalid_subject_rdn_order.go
    
    As per Chris Henderson's suggestion, to "improve readability".
    
    * Update lint_invalid_subject_rdn_order_test.go
    
    As per Chris Henderson's suggestion.
    
    * Update time.go
    
    Added CABFEV_Sec9_2_8_Date
    
    * Add files via upload
    
    * Add files via upload
    
    * Revised according to Chris and Corey suggestions
    
    * Add files via upload
    
    * Add files via upload
    
    * Delete v3/lints/cabf_br/lint_e_invalid_cps_uri.go
    
    * Delete v3/lints/cabf_br/lint_e_invalid_cps_uri_test.go
    
    * Delete v3/testdata/invalid_cps_uri_ko_01.pem
    
    * Delete v3/testdata/invalid_cps_uri_ko_02.pem
    
    * Delete v3/testdata/invalid_cps_uri_ko_03.pem
    
    * Delete v3/testdata/invalid_cps_uri_ok_01.pem
    
    * Delete v3/testdata/invalid_cps_uri_ok_02.pem
    
    * Delete v3/testdata/invalid_cps_uri_ok_03.pem
    
    * Add files via upload
    
    * Add files via upload
    
    * Add files via upload
    
    * Add files via upload
    
    * Add files via upload
    
    * Add files via upload
    
    * Add files via upload
    
    * Add files via upload
    
    * Add files via upload
    
    * Delete v3/lints/rfc/lint_empty_sct_list.go
    
    * Delete v3/lints/rfc/lint_empty_sct_list_test.go
    
    * Delete v3/testdata/empty_sct_list_ko_01.pem
    
    * Delete v3/testdata/empty_sct_list_na_01.pem
    
    * Delete v3/testdata/empty_sct_list_na_02.pem
    
    * Delete v3/testdata/empty_sct_list_ok_01.pem
    
    * Delete v3/testdata/empty_sct_list_ok_02.pem
    
    * Update source.go
    
    * Update time.go
    
    ---------
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit 208af03b5346578ba252fed89c93ceda0d6dc62e[m
Author: Adriano Santoni <asantoni64@gmail.com>
Date:   Sun Apr 28 20:33:13 2024 +0200

    Add lint for checking that a CRL contains the CRL Number extension (#834)
    
    * Add files via upload
    
    * Add files via upload
    
    * Add files via upload
    
    * Add files via upload
    
    * Update lint_invalid_subject_rdn_order_test.go
    
    Added //nolint:all to comment block to avoid golangci-lint to complain about duplicate words in comment
    
    * Update lint_invalid_subject_rdn_order.go
    
    Fixed import block
    
    * Update v3/lints/cabf_br/lint_invalid_subject_rdn_order.go
    
    Fine to me.
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>
    
    * Update lint_invalid_subject_rdn_order.go
    
    As per Chris Henderson's suggestion, to "improve readability".
    
    * Update lint_invalid_subject_rdn_order_test.go
    
    As per Chris Henderson's suggestion.
    
    * Update time.go
    
    Added CABFEV_Sec9_2_8_Date
    
    * Add files via upload
    
    * Add files via upload
    
    * Revised according to Chris and Corey suggestions
    
    * Add files via upload
    
    * Add files via upload
    
    * Delete v3/lints/cabf_br/lint_e_invalid_cps_uri.go
    
    * Delete v3/lints/cabf_br/lint_e_invalid_cps_uri_test.go
    
    * Delete v3/testdata/invalid_cps_uri_ko_01.pem
    
    * Delete v3/testdata/invalid_cps_uri_ko_02.pem
    
    * Delete v3/testdata/invalid_cps_uri_ko_03.pem
    
    * Delete v3/testdata/invalid_cps_uri_ok_01.pem
    
    * Delete v3/testdata/invalid_cps_uri_ok_02.pem
    
    * Delete v3/testdata/invalid_cps_uri_ok_03.pem
    
    * Add files via upload
    
    * Add files via upload
    
    * Add files via upload
    
    * Add files via upload
    
    * Update oid.go
    
    Add OID for CRL Number
    
    * Update v3/util/oid.go
    
    ---------
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit d5a09f841725281fd65d0003dea004fd75e31d8c[m
Author: Paul van Brouwershaven <vanbroup@users.noreply.github.com>
Date:   Sun Apr 28 21:14:06 2024 +0300

    Add lint to cover TLS BR v2 EKU checks (#833)
    
    * Add EV policy and Pre Certiicate Signing Certificate EKU
    
    * Apply serverAuth to certificates with CA/B TLS policy OID
    
    * lint subscriber EKU according to TLS BR v2
    
    * Make lint ineffective since TLS BR v2
    
    These lints are covered by the new `e_sub_cert_eku_check` lint that will lint according to the TLS BR v2 language.
    
    * Correct expected result
    
    * Correct numbers as result of CA/B policy inclusion in additon to serverAuth
    
    The `util.IsServerAuthCert` did not consider certificates that attest the CA/Browser Forum policy OIDs but do not include the `serverAuth` EKU. This has been addressed and caused some mintor changes in the test corpus.
    
    * Check if subscriber certificate with EKU extension
    
    * Pass certificate in subscriber certificate check
    
    * Remove unnecessary len check
    
    * Format
    
    ---------
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit 63e3f8654d742ba9e7b36881b1f8c003a426f201[m
Author: Adriano Santoni <asantoni64@gmail.com>
Date:   Sun Apr 28 17:02:34 2024 +0200

    Add lint to detect invalid cps uri (#828)
    
    * Add files via upload
    
    * Add files via upload
    
    * Add files via upload
    
    * Add files via upload
    
    * Update lint_invalid_subject_rdn_order_test.go
    
    Added //nolint:all to comment block to avoid golangci-lint to complain about duplicate words in comment
    
    * Update lint_invalid_subject_rdn_order.go
    
    Fixed import block
    
    * Update v3/lints/cabf_br/lint_invalid_subject_rdn_order.go
    
    Fine to me.
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>
    
    * Update lint_invalid_subject_rdn_order.go
    
    As per Chris Henderson's suggestion, to "improve readability".
    
    * Update lint_invalid_subject_rdn_order_test.go
    
    As per Chris Henderson's suggestion.
    
    * Update time.go
    
    Added CABFEV_Sec9_2_8_Date
    
    * Add files via upload
    
    * Add files via upload
    
    * Revised according to Chris and Corey suggestions
    
    * Add files via upload
    
    * Add files via upload
    
    * Delete v3/lints/cabf_ev/lint_ev_orgid_inconsistent_subj_and_ext.go
    
    * Delete v3/lints/cabf_ev/lint_ev_orgid_inconsistent_subj_and_ext_test.go
    
    * Delete v3/testdata/orgid_subj_and_ext_ko_01.pem
    
    * Delete v3/testdata/orgid_subj_and_ext_ko_02.pem
    
    * Delete v3/testdata/orgid_subj_and_ext_ko_03.pem
    
    * Delete v3/testdata/orgid_subj_and_ext_ok_01.pem
    
    * Delete v3/testdata/orgid_subj_and_ext_ok_02.pem
    
    * Delete v3/testdata/orgid_subj_and_ext_ok_03.pem
    
    * Delete v3/testdata/orgid_subj_and_ext_ok_04.pem
    
    * Delete v3/testdata/orgid_subj_and_ext_ok_05.pem
    
    * Update time.go
    
    ---------
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit 2988620fc3db96938dbfb71ca2afe8f5b2010920[m
Author: Adriano Santoni <asantoni64@gmail.com>
Date:   Sun Apr 28 16:09:22 2024 +0200

    Add lint to check that a CRL does not contain an empty revokedCertificates element (#831)
    
    * Add files via upload
    
    * Add files via upload
    
    * Add files via upload
    
    * Add files via upload
    
    * Update lint_invalid_subject_rdn_order_test.go
    
    Added //nolint:all to comment block to avoid golangci-lint to complain about duplicate words in comment
    
    * Update lint_invalid_subject_rdn_order.go
    
    Fixed import block
    
    * Update v3/lints/cabf_br/lint_invalid_subject_rdn_order.go
    
    Fine to me.
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>
    
    * Update lint_invalid_subject_rdn_order.go
    
    As per Chris Henderson's suggestion, to "improve readability".
    
    * Update lint_invalid_subject_rdn_order_test.go
    
    As per Chris Henderson's suggestion.
    
    * Update time.go
    
    Added CABFEV_Sec9_2_8_Date
    
    * Add files via upload
    
    * Add files via upload
    
    * Revised according to Chris and Corey suggestions
    
    * Add files via upload
    
    * Add files via upload
    
    * Delete v3/lints/cabf_br/lint_e_invalid_cps_uri.go
    
    * Delete v3/lints/cabf_br/lint_e_invalid_cps_uri_test.go
    
    * Delete v3/testdata/invalid_cps_uri_ko_01.pem
    
    * Delete v3/testdata/invalid_cps_uri_ko_02.pem
    
    * Delete v3/testdata/invalid_cps_uri_ko_03.pem
    
    * Delete v3/testdata/invalid_cps_uri_ok_01.pem
    
    * Delete v3/testdata/invalid_cps_uri_ok_02.pem
    
    * Delete v3/testdata/invalid_cps_uri_ok_03.pem
    
    * Delete v3/lints/cabf_ev/lint_ev_orgid_inconsistent_subj_and_ext.go
    
    * Delete v3/lints/cabf_ev/lint_ev_orgid_inconsistent_subj_and_ext_test.go
    
    * Delete v3/testdata/orgid_subj_and_ext_ko_01.pem
    
    * Delete v3/testdata/orgid_subj_and_ext_ko_02.pem
    
    * Delete v3/testdata/orgid_subj_and_ext_ko_03.pem
    
    * Delete v3/testdata/orgid_subj_and_ext_ok_01.pem
    
    * Delete v3/testdata/orgid_subj_and_ext_ok_02.pem
    
    * Delete v3/testdata/orgid_subj_and_ext_ok_03.pem
    
    * Delete v3/testdata/orgid_subj_and_ext_ok_04.pem
    
    * Delete v3/testdata/orgid_subj_and_ext_ok_05.pem
    
    * Add files via upload
    
    * Add files via upload
    
    * Update time.go
    
    * Add files via upload
    
    * Add files via upload
    
    * Delete v3/lints/cabf_br/crl_empty_revoked_certificates_ko.pem
    
    * Delete v3/lints/cabf_br/crl_empty_revoked_certificates_ok.pem
    
    * Update lint_crl_revoked_certificates_field_empty.go
    
    ---------
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>
    Co-authored-by: Zakir Durumeric <zakird@gmail.com>

[33mcommit 61c73edc6b2a2cf3e6eae6a5fb5f67dd334829ee[m
Author: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>
Date:   Sun Apr 21 10:10:35 2024 -0700

    build(deps): bump golang.org/x/net from 0.17.0 to 0.23.0 in /v3 (#835)
    
    Bumps [golang.org/x/net](https://github.com/golang/net) from 0.17.0 to 0.23.0.
    - [Commits](https://github.com/golang/net/compare/v0.17.0...v0.23.0)
    
    ---
    updated-dependencies:
    - dependency-name: golang.org/x/net
      dependency-type: direct:production
    ...
    
    Signed-off-by: dependabot[bot] <support@github.com>
    Co-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit a0112345dbd9d39f2f637edcecb2f313d56b7a35[m
Author: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>
Date:   Sun Apr 21 09:47:49 2024 -0700

    build(deps): bump golang.org/x/net in /v3/cmd/genTestCerts (#836)
    
    Bumps [golang.org/x/net](https://github.com/golang/net) from 0.17.0 to 0.23.0.
    - [Commits](https://github.com/golang/net/compare/v0.17.0...v0.23.0)
    
    ---
    updated-dependencies:
    - dependency-name: golang.org/x/net
      dependency-type: indirect
    ...
    
    Signed-off-by: dependabot[bot] <support@github.com>
    Co-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>

[33mcommit 6c7d024812dfb2f25143f31e4655240d66e4058a[m
Author: Phil Porada <pgporada@users.noreply.github.com>
Date:   Thu Apr 18 13:27:59 2024 -0400

    Add lint to verify CRL TBSCertList.revokedCertificates field is absent when there are no revoked certificates (#832)
    
    * Working lint and tests
    
    * Add negative test
    
    * Rename test crl
    
    * DER, PEM, vim smuggled inside testdata just like xz, you pick
    
    * Add more negative test cases and run through all of the files
    
    ---------
    
    Co-authored-by: Zakir Durumeric <zakird@gmail.com>

[33mcommit 4b2f38b56132eda5017d637ed07ef9be59ab6976[m
Author: Adriano Santoni <asantoni64@gmail.com>
Date:   Sun Apr 14 21:49:41 2024 +0200

    Lint for checking that organizationIdentifier Subject attribute and CABFOrganizationIdentifier extension are consistent as per EVG 9.2.8 (#820)
    
    * Add files via upload
    
    * Add files via upload
    
    * Add files via upload
    
    * Add files via upload
    
    * Update lint_invalid_subject_rdn_order_test.go
    
    Added //nolint:all to comment block to avoid golangci-lint to complain about duplicate words in comment
    
    * Update lint_invalid_subject_rdn_order.go
    
    Fixed import block
    
    * Update v3/lints/cabf_br/lint_invalid_subject_rdn_order.go
    
    Fine to me.
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>
    
    * Update lint_invalid_subject_rdn_order.go
    
    As per Chris Henderson's suggestion, to "improve readability".
    
    * Update lint_invalid_subject_rdn_order_test.go
    
    As per Chris Henderson's suggestion.
    
    * Update time.go
    
    Added CABFEV_Sec9_2_8_Date
    
    * Add files via upload
    
    * Add files via upload
    
    * Revised according to Chris and Corey suggestions
    
    * Add files via upload
    
    * Add files via upload
    
    * Delete v3/lints/cabf_br/lint_e_invalid_cps_uri.go
    
    * Delete v3/lints/cabf_br/lint_e_invalid_cps_uri_test.go
    
    * Delete v3/testdata/invalid_cps_uri_ko_01.pem
    
    * Delete v3/testdata/invalid_cps_uri_ko_02.pem
    
    * Delete v3/testdata/invalid_cps_uri_ko_03.pem
    
    * Delete v3/testdata/invalid_cps_uri_ok_01.pem
    
    * Delete v3/testdata/invalid_cps_uri_ok_02.pem
    
    * Delete v3/testdata/invalid_cps_uri_ok_03.pem
    
    ---------
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit 5de620c50c0621fffce102d391475f78e0fe3e89[m
Author: MTG <36234449+mtgag@users.noreply.github.com>
Date:   Sun Apr 14 19:58:33 2024 +0200

    Subject rdns correct encoding (#824)
    
    * lint about the encoding of qcstatements for PSD2
    
    * Revert "lint about the encoding of qcstatements for PSD2"
    
    This reverts commit 6c2367080d148f4b8c01f96a4c80e3ac55d1ef26.
    
    * util: gtld_map autopull updates for 2021-10-21T07:25:20 UTC
    
    * always check and perform the operation in the execution
    
    * synchronised with project
    
    * synchronised with project
    
    * synchronised with project
    
    * synchronised with project
    
    * fixed merge error
    
    * synchronised with project
    
    * goimports
    
    * trying to decrease cyclomatic complexity
    
    ---------
    
    Co-authored-by: mtg <git@mtg.de>
    Co-authored-by: GitHub <noreply@github.com>
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit ae3b1f32c23bdbb29998329b7e2fb13f0d00a015[m
Author: MTG <36234449+mtgag@users.noreply.github.com>
Date:   Tue Apr 9 21:24:34 2024 +0200

    Correct test descriptions (#829)
    
    Mark lint.NA expected results as lint.NA, not pass.
    
    https://github.com/zmap/zlint/pull/829

[33mcommit 4d4672997c53cdfe417d85a0abd6091d59deeb89[m
Merge: b3a86b3c 308a138e
Author: mtgag <githreg@mtg.de>
Date:   Tue Apr 9 11:48:05 2024 +0200

    Merge https://github.com/zmap/zlint

[33mcommit b3a86b3c0f6658a402f3a81dfb32b534e1abba3e[m
Author: mtgag <githreg@mtg.de>
Date:   Tue Apr 9 11:45:06 2024 +0200

    Revert "synchronised with project"
    
    This reverts commit 2cd7d087f4a812d4ef3640560edf1d07cce2ea56.

[33mcommit 63cf8e862a490ebd8769ffbc516d882449d67741[m
Author: mtgag <githreg@mtg.de>
Date:   Tue Apr 9 11:44:43 2024 +0200

    Revert "synchronised with project"
    
    This reverts commit bad73ee2d5669394cde3053d300f285a91f75fd6.

[33mcommit 2cd7d087f4a812d4ef3640560edf1d07cce2ea56[m
Author: mtgag <githreg@mtg.de>
Date:   Tue Apr 9 11:40:00 2024 +0200

    synchronised with project

[33mcommit 308a138ee20193335072c10b9b6ce7dec3d950c9[m
Author: MTG <36234449+mtgag@users.noreply.github.com>
Date:   Sun Apr 7 15:04:05 2024 +0200

    Limit scope for cn checking in SAN (#825)
    
    * lint about the encoding of qcstatements for PSD2
    
    * Revert "lint about the encoding of qcstatements for PSD2"
    
    This reverts commit 6c2367080d148f4b8c01f96a4c80e3ac55d1ef26.
    
    * util: gtld_map autopull updates for 2021-10-21T07:25:20 UTC
    
    * always check and perform the operation in the execution
    
    * synchronised with project
    
    * synchronised with project
    
    * synchronised with project
    
    * synchronised with project
    
    * fixed merge error
    
    * synchronised with project
    
    * address comments of PR #809
    
    * trying to decrease cyclomatic complexity
    
    * reverted commit in this branch
    
    ---------
    
    Co-authored-by: mtg <git@mtg.de>
    Co-authored-by: GitHub <noreply@github.com>
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit 2980c72629aaf86b440644d13a9d7e2f36c0f350[m
Author: David Adrian <davidcadrian@gmail.com>
Date:   Sat Apr 6 18:16:01 2024 -0400

    Add ineffective date to DSA lints. (#827)
    
    DSA is prohibited, so we can't maintain an up-to-date reference for how
    a DSA key should be structured. Instead of checking prohibited DSA certs
    against the old requirements, rely on lint_prohibit_dsa_usage.go

[33mcommit bad73ee2d5669394cde3053d300f285a91f75fd6[m
Author: mtgag <githreg@mtg.de>
Date:   Fri Apr 5 07:40:36 2024 +0200

    synchronised with project

[33mcommit 795d2068ca83cd2f08bb866ee4c367b09e444489[m
Merge: f1a66db9 f9496fad
Author: mtgag <githreg@mtg.de>
Date:   Fri Apr 5 07:40:35 2024 +0200

    Merge https://github.com/zmap/zlint

[33mcommit f9496fada52af23e776f898362b4074cb082f44b[m
Author: MTG <36234449+mtgag@users.noreply.github.com>
Date:   Thu Mar 28 18:17:32 2024 +0100

    Use help Method beforeoron instead of  (#717)
    
    * lint about the encoding of qcstatements for PSD2
    
    * Revert "lint about the encoding of qcstatements for PSD2"
    
    This reverts commit 6c2367080d148f4b8c01f96a4c80e3ac55d1ef26.
    
    * util: gtld_map autopull updates for 2021-10-21T07:25:20 UTC
    
    * always check and perform the operation in the execution
    
    * using the help method BeforeOrOn instead of simple Before, added certificates that cover the edge cases
    
    * update in integration data
    
    * reverted commit, kept certificates, changed assertion, after discussion in the pull request
    
    ---------
    
    Co-authored-by: mtg <git@mtg.de>
    Co-authored-by: GitHub <noreply@github.com>
    Co-authored-by: Christopher Henderson <chris@chenderson.org>
    Co-authored-by: Zakir Durumeric <zakird@gmail.com>

[33mcommit 92917299fd81f3247a1bbc69643d31a3a3e1552c[m
Author: github-actions[bot] <41898282+github-actions[bot]@users.noreply.github.com>
Date:   Thu Mar 28 09:55:29 2024 -0700

    util: gtld_map autopull updates for 2024-03-27T22:19:31 UTC (#817)
    
    Co-authored-by: GitHub <noreply@github.com>

[33mcommit e99e725e9aff1a9e9af427a1bf288389f693751b[m
Author: Martijn Katerbarg <martijn.katerbarg@sectigo.com>
Date:   Wed Mar 27 23:14:26 2024 +0100

    feat: Test EKU Criticality (#816)
    
    * feat: Test EKU Criticality
    
    * Correct capitalization
    
    * Correct capitalization

[33mcommit 38cfd72bd88b8688173ac63d408cfdfefb46801a[m
Author: Martijn Katerbarg <martijn.katerbarg@sectigo.com>
Date:   Sun Mar 24 20:21:39 2024 +0100

    cRLIssuer MUST NOT be present (#814)
    
    * cRLIssuer MUST NOT be present lint
    
    * Also cover Reason
    
    ---------
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit 990a074c32c8899e935552724fc773de8765ceef[m
Author: Adam <bitlux@users.noreply.github.com>
Date:   Sun Mar 24 08:44:09 2024 -0700

    Add lints for S/MIME BR 7.1.2.3l (#805)
    
    * Add lints for S/MIME BR 7.1.2.3l
    
    * Save results of util functions as variables to make logic clearer.
    
    ---------
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit 32bba7aeb74f82f604f99ee78d08aae1cb7e4985[m
Author: MTG <36234449+mtgag@users.noreply.github.com>
Date:   Sun Mar 17 18:10:55 2024 +0100

    Update single email if present (#808)
    
    * lint about the encoding of qcstatements for PSD2
    
    * Revert "lint about the encoding of qcstatements for PSD2"
    
    This reverts commit 6c2367080d148f4b8c01f96a4c80e3ac55d1ef26.
    
    * util: gtld_map autopull updates for 2021-10-21T07:25:20 UTC
    
    * always check and perform the operation in the execution
    
    * synchronised with project
    
    * synchronised with project
    
    * synchronised with project
    
    * synchronised with project
    
    * added same lint for subject values instead of SAN values
    
    * resolved conflict issue
    
    * addressed review comments and hint to citation from #795
    
    * addressing issue #795 and review comments of PR #802
    
    ---------
    
    Co-authored-by: mtg <git@mtg.de>
    Co-authored-by: GitHub <noreply@github.com>
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit e33bae9c194cff0acf80026363f7f36c45d42fd7[m
Author: MTG <36234449+mtgag@users.noreply.github.com>
Date:   Sun Mar 17 18:02:26 2024 +0100

    Update single email subject if present (#802)
    
    * lint about the encoding of qcstatements for PSD2
    
    * Revert "lint about the encoding of qcstatements for PSD2"
    
    This reverts commit 6c2367080d148f4b8c01f96a4c80e3ac55d1ef26.
    
    * util: gtld_map autopull updates for 2021-10-21T07:25:20 UTC
    
    * always check and perform the operation in the execution
    
    * synchronised with project
    
    * synchronised with project
    
    * synchronised with project
    
    * synchronised with project
    
    * added same lint for subject values instead of SAN values
    
    * resolved conflict issue
    
    * addressed review comments and hint to citation from #795
    
    ---------
    
    Co-authored-by: mtg <git@mtg.de>
    Co-authored-by: GitHub <noreply@github.com>
    Co-authored-by: Christopher Henderson <chris@chenderson.org>
    Co-authored-by: Zakir Durumeric <zakird@gmail.com>

[33mcommit 7c899eaaaa534b10489f457ffbea808235d4fc71[m
Author: Adam <bitlux@users.noreply.github.com>
Date:   Mon Mar 11 15:04:47 2024 -0700

    Add lint for BR 7.1.4.2.2a mailbox-validated (#806)
    
    * Add lint for BR 7.1.4.2.2a mailbox-validated
    
    * Remove test code
    
    * Update citation description
    
    ---------
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit e6650ebd433bea8cbe73b96c9d0d66015c0cd7e2[m
Author: Adam <bitlux@users.noreply.github.com>
Date:   Mon Mar 11 14:39:27 2024 -0700

    Add lints for S/MIME BR 7.1.4.2.2n country name (#807)
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit 8d2c57948e697330fc81c865f449e6922b7bc0bb[m
Author: MTG <36234449+mtgag@users.noreply.github.com>
Date:   Sun Mar 10 18:59:15 2024 +0100

    Lint for 7.1.2.7.2 BR (#810)
    
    * lint about the encoding of qcstatements for PSD2
    
    * Revert "lint about the encoding of qcstatements for PSD2"
    
    This reverts commit 6c2367080d148f4b8c01f96a4c80e3ac55d1ef26.
    
    * util: gtld_map autopull updates for 2021-10-21T07:25:20 UTC
    
    * always check and perform the operation in the execution
    
    * synchronised with project
    
    * synchronised with project
    
    * synchronised with project
    
    * synchronised with project
    
    * added lint to check values of subjectDN in DV certificates
    
    * fixed errors
    
    * fixed merge error
    
    * addressing review comment
    
    ---------
    
    Co-authored-by: mtg <git@mtg.de>
    Co-authored-by: GitHub <noreply@github.com>
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit e76cc77296612b97bb8df7a525b7cec68f77070f[m
Author: Adriano Santoni <asantoni64@gmail.com>
Date:   Sun Mar 10 18:36:37 2024 +0100

    Add lint for checking that Subject attributes (RDNs) appear in the order prescribed by CABF BR 7.1.4.2 (#813)
    
    * Add files via upload
    
    * Add files via upload
    
    * Add files via upload
    
    * Add files via upload
    
    * Update lint_invalid_subject_rdn_order_test.go
    
    Added //nolint:all to comment block to avoid golangci-lint to complain about duplicate words in comment
    
    * Update lint_invalid_subject_rdn_order.go
    
    Fixed import block
    
    * Update v3/lints/cabf_br/lint_invalid_subject_rdn_order.go
    
    Fine to me.
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>
    
    * Update lint_invalid_subject_rdn_order.go
    
    As per Chris Henderson's suggestion, to "improve readability".
    
    * Update lint_invalid_subject_rdn_order_test.go
    
    As per Chris Henderson's suggestion.
    
    ---------
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit f1a66db99213a6363d024c0108bb592d49094032[m
Merge: 53b911ef a063d317
Author: mtgag <githreg@mtg.de>
Date:   Sun Mar 10 10:17:34 2024 +0100

    Merge https://github.com/zmap/zlint

[33mcommit a063d317122dce598d26ac03b4975cbb6469f8e4[m
Author: Adam <bitlux@users.noreply.github.com>
Date:   Sat Mar 9 10:55:49 2024 -0800

    Add lints for S/MIME BR 7.1.2.3.b (#779)
    
    * Add lints for S/MIME BR 7.1.2.3.b
    
    * remove logging
    
    * Update logic to include legacy certs
    
    * Add test for legacy certs
    
    * add test
    
    * Add tests with mixed HTTP and non-HTTP
    
    * URL -> URI
    
    * Fix text
    
    * UseCertificateLint
    
    * Rename testdata files to reflect their type
    
    ---------
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit a72ff4ec44ebff74248e6631940c1b44f9bbffda[m
Author: github-actions[bot] <41898282+github-actions[bot]@users.noreply.github.com>
Date:   Sat Mar 9 10:30:35 2024 -0800

    util: gtld_map autopull updates for 2024-03-09T18:19:57 UTC (#811)
    
    Co-authored-by: GitHub <noreply@github.com>

[33mcommit 5501be19e0a5da8b260cc06474f09961d1423eb3[m
Author: MTG <36234449+mtgag@users.noreply.github.com>
Date:   Sat Mar 9 18:59:39 2024 +0100

    Mailbox addresses from san for all br (#809)
    
    * lint about the encoding of qcstatements for PSD2
    
    * Revert "lint about the encoding of qcstatements for PSD2"
    
    This reverts commit 6c2367080d148f4b8c01f96a4c80e3ac55d1ef26.
    
    * util: gtld_map autopull updates for 2021-10-21T07:25:20 UTC
    
    * always check and perform the operation in the execution
    
    * synchronised with project
    
    * synchronised with project
    
    * synchronised with project
    
    * synchronised with project
    
    * refactored lint to cover all SMIME BR certificates
    
    * fixed git merge issue
    
    ---------
    
    Co-authored-by: mtg <git@mtg.de>
    Co-authored-by: GitHub <noreply@github.com>
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit 9c67bdb4bde793753f7d98b78d089a49ddf83b7b[m
Author: Adam <bitlux@users.noreply.github.com>
Date:   Sat Mar 9 08:58:05 2024 -0800

    Fix typo (#804)
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit 53b911ef750ff906e9749252b599e8253fa84594[m
Author: mtgag <githreg@mtg.de>
Date:   Tue Mar 5 11:05:05 2024 +0100

    fixed merge error

[33mcommit d10444e4b15dcd252a2f18a583a7d4348e8ba659[m
Merge: 31e18450 83b5f8d6
Author: mtgag <githreg@mtg.de>
Date:   Mon Mar 4 07:51:14 2024 +0100

    Merge https://github.com/zmap/zlint

[33mcommit 83b5f8d6b7c243880a3dab6c95d954d681bf2e3f[m
Author: Adam <bitlux@users.noreply.github.com>
Date:   Sun Mar 3 07:51:11 2024 -0800

    Add lint for S/MIME BR 7.1.2.3 (k) (#799)
    
    * Add line for S/MIME BR 7.1.2.3.k.
    
    * Add tests generated by christopher-henderson
    
    ---------
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit b9ff71f1ec90d82ee01aabf5272b40b6ce1cf154[m
Author: toddgaunt-gs <todd.gaunt@globalsign.com>
Date:   Sun Mar 3 10:32:24 2024 -0500

    Add lint to enforce SMIME BRs: 7.1.4.2.1 requirement for mailbox addr… (#800)
    
    * Add lint to enforce SMIME BRs: 7.1.4.2.1 requirement for mailbox addresses
    
    All mailbox addresses appearing in subjectDN or dirName must be repeated
    in san:rfc822Name or san:otherName. This lint does its best to detect
    mailbox address values in the subjectDN or dirName and if any are
    detected ensures they are repeated.
    
    * Add expected integration failures for new lint e_mailbox_address_shall_contain_an_rfc822_name
    
    The failures all have email addresses that don't have an **exact** match
    in the SAN.
    
    How the integration tests were run:
    `make integration INT_FLAGS="-lintSummary -fingerprintSummary -lintFilter='e_mailbox_address_shall_contain_an_rfc822_name'"`
    
    Fingerprints of the relevant certificates:
    3087f97b6cff020b5320e18d3e326074cbaa128142660f2debe4564ab1ab0179
    5f3fcccca91a7b39e8995f79c35cb5e604d4ee0487ea1a41993c84304c0a5c99
    63d23132c2511f33bb947f27c398bb824109ccf2d6a2037e3713fe9f7a43b15d
    b034fa1aa9e501dc14b43d43dfe2210de3e5551744494b55d5f0abd865c67efc
    c6ac841c78191101725ca7d5ed499be47c15ebeece7d74e6d095e2925e7bb404
    e4dbfc94e616ffb59904e394d9dcdd3ab55c26c5586440f37c058eecb907a344
    
    * Revert "Add expected integration failures for new lint e_mailbox_address_shall_contain_an_rfc822_name"
    
    This reverts commit 037b5ec8918805bdb989726750c00d7d74e0d66a.
    
    * Add expected integration failures for new lint e_mailbox_address_shall_contain_an_rfc822_name
    
    This commit is a proper version of the previously reverted one. It was
    reverted because I accidently ran the script to update the config only
    for the failing lint, rather than lints.
    
    The failures all have email addresses that don't have an **exact** match
    in the SAN.
    
    How the integration tests were run:
    `make integration INT_FLAGS="-lintSummary -fingerprintSummary -lintFilter='e_mailbox_address_shall_contain_an_rfc822_name'"`
    
    Fingerprints of the relevant certificates:
    3087f97b6cff020b5320e18d3e326074cbaa128142660f2debe4564ab1ab0179
    5f3fcccca91a7b39e8995f79c35cb5e604d4ee0487ea1a41993c84304c0a5c99
    63d23132c2511f33bb947f27c398bb824109ccf2d6a2037e3713fe9f7a43b15d
    b034fa1aa9e501dc14b43d43dfe2210de3e5551744494b55d5f0abd865c67efc
    c6ac841c78191101725ca7d5ed499be47c15ebeece7d74e6d095e2925e7bb404
    e4dbfc94e616ffb59904e394d9dcdd3ab55c26c5586440f37c058eecb907a344
    
    * Use effective date from SMIME BR for mailbox_address_from_san lint
    
    * Address code style to fit with established conventions
    
    * Revert accidental changes to genTestCerts
    
    * Apply DeMorgan's law to fix incorrect code simplification
    
    * Remove redundant function literal
    
    * Run gofmt
    
    ---------
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit a23de3d51cddfc5e355bd6231f83103e86d936da[m
Author: github-actions[bot] <41898282+github-actions[bot]@users.noreply.github.com>
Date:   Sun Feb 25 08:55:00 2024 -0800

    util: gtld_map autopull updates for 2024-02-20T21:17:08 UTC (#794)
    
    Co-authored-by: GitHub <noreply@github.com>

[33mcommit bf84ed888ec0a3f0ac99220e14b12c31b53ccb94[m
Author: MTG <36234449+mtgag@users.noreply.github.com>
Date:   Sun Feb 25 17:24:52 2024 +0100

    Add test case for smime ext subject directory attr (#801)
    
    * lint about the encoding of qcstatements for PSD2
    
    * Revert "lint about the encoding of qcstatements for PSD2"
    
    This reverts commit 6c2367080d148f4b8c01f96a4c80e3ac55d1ef26.
    
    * util: gtld_map autopull updates for 2021-10-21T07:25:20 UTC
    
    * always check and perform the operation in the execution
    
    * synchronised with project
    
    * synchronised with project
    
    * synchronised with project
    
    * synchronised with project
    
    * added test case
    
    * resolved conflict issue
    
    ---------
    
    Co-authored-by: mtg <git@mtg.de>
    Co-authored-by: GitHub <noreply@github.com>
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit 31e18450ec45e7aaa559d20d470450957092fcd7[m
Merge: 51d498f8 060b3850
Author: mtgag <githreg@mtg.de>
Date:   Sun Feb 25 11:01:49 2024 +0100

    Merge https://github.com/zmap/zlint

[33mcommit 060b3850760d832415fc7b82113f1117ef93e285[m
Author: Adam <bitlux@users.noreply.github.com>
Date:   Tue Feb 20 12:58:54 2024 -0800

    Lint for S/MIME BR 7.1.2.3.g (#797)
    
    * Lint for S/MIME BR 7.1.2.3.g
    
    * Remove printf
    
    * Addresss rewview comments. Check for presence of AuthkeyOID extension. Use error details.
    
    ---------
    
    Co-authored-by: Zakir Durumeric <zakird@gmail.com>
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit a4b46ef6a8969ffbc0c41e72f8f2b21294a1cccd[m
Author: Adam <bitlux@users.noreply.github.com>
Date:   Mon Feb 19 09:35:52 2024 -0800

    Add lint for subject directory attributes extension (#798)

[33mcommit 1baec6eef208984b13b55d0f9545afab9c2315e8[m
Author: Adam <bitlux@users.noreply.github.com>
Date:   Wed Feb 14 07:03:31 2024 -0800

    Fix copy/paste error (#796)
    
    Co-authored-by: Zakir Durumeric <zakird@gmail.com>

[33mcommit 51d498f89ff4d5e8877465a6f43f85d6879d2529[m
Merge: e77fae15 8deb02ba
Author: mtgag <githreg@mtg.de>
Date:   Tue Feb 13 08:05:33 2024 +0100

    synchronised with project

[33mcommit 8deb02ba189e5f89d7f1e8a5bf2f75e86f81690e[m
Author: Arthur Gautier <superbaloo+registrations.github@superbaloo.net>
Date:   Sun Feb 11 07:29:20 2024 -0800

    Subject Key Identifier is not recommended by CABF BR v2 (#790)
    
    * Subject Key Identifier is not recommended by CABF BR v2
    
    With SC62, the CABF BR now lists SKI as not recommended.
    
    Per discussion in #762, zlint should provide two lints, one for rfc5280
    behavior and one for CABF BR.
    
    Both lint will conflict with each other, users are expected to select
    (or ignore) which behavior they mean to follow.
    
    Fixes #749
    
    * Test data for SKI not recommnended
    
    Co-Authored-By: Christopher Henderson <chris@chenderson.org>
    
    ---------
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit fa85598bd69a2c2aa0238a394cd74542bf3c6691[m
Author: MTG <36234449+mtgag@users.noreply.github.com>
Date:   Sat Feb 10 19:08:40 2024 +0100

    Handle ips in aia internal names (#791)
    
    * lint about the encoding of qcstatements for PSD2
    
    * Revert "lint about the encoding of qcstatements for PSD2"
    
    This reverts commit 6c2367080d148f4b8c01f96a4c80e3ac55d1ef26.
    
    * util: gtld_map autopull updates for 2021-10-21T07:25:20 UTC
    
    * always check and perform the operation in the execution
    
    * synchronised with project
    
    * synchronised with project
    
    * synchronised with project
    
    * synchronised with project
    
    * if the AIA contains an IP then pass instead of warn
    
    * fixed merge message
    
    * trying to resolve conflicts
    
    * enhancement; lint only if extension is present otherwise not applicable
    
    ---------
    
    Co-authored-by: mtg <git@mtg.de>
    Co-authored-by: GitHub <noreply@github.com>
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit 82d733e4dceb5e69296c9ac9dd4d1747182ebe26[m
Author: BJ Cardon <cardonator@users.noreply.github.com>
Date:   Fri Feb 9 14:32:23 2024 -0700

    Fix a bug in the check for 7.1.4.2.h - single email address in subject:emailAddress (#792)
    
    * fix bug in the email address checking in the smime package to allow multiple email address subject fields, but dsisallow multiple values in a single email address field
    
    fixes a comment on #753
    
    * fix typo

[33mcommit 5501b4fcf4f9891a1eaf463fb72b8d582d2684d2[m
Author: github-actions[bot] <41898282+github-actions[bot]@users.noreply.github.com>
Date:   Sat Jan 27 11:11:24 2024 -0800

    util: gtld_map autopull updates for 2024-01-22T23:19:16 UTC (#789)
    
    Co-authored-by: GitHub <noreply@github.com>

[33mcommit e77fae15e50dcbfc6c214557cac40c94ddd465c1[m
Author: mtgag <githreg@mtg.de>
Date:   Wed Jan 24 07:11:55 2024 +0100

    synchronised with project

[33mcommit ddd1a81ca77fbcb75a5c787d41b8b307762f7246[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sat Jan 20 09:57:02 2024 -0800

    Update copyright notices to 2024 (#787)
    
    * Update copyright notices to 2024
    
    * touched the gen test cert script and need to update the test file template

[33mcommit 8a61dfa6b62eb8caee5daa9cab97b4c1e6757f21[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sat Jan 20 09:16:10 2024 -0800

    Refactor and improve the new lint creation bash script (#786)
    
    * Improving the new lint creation bash scripts
    
    * fixed typo

[33mcommit be8dd6a629e36c9a9a34aeb7b34ed06327151ce3[m
Author: Rob <3725956+robplee@users.noreply.github.com>
Date:   Mon Jan 1 18:15:24 2024 +0000

    Limit e_registration_scheme_id_matches_subject_country to no longer apply to LEI or INT organizationIdentifiers (#781)
    
    * fix issue where e_registration_scheme_id_matches_subject_country was applying to LEI and INT certs where not required by the SMIME BRs
    
    * fix execution of e_registration_scheme_id_matches_subject_country lint in case where some organizationIdentifiers are subject to the check and others are not
    
    ---------
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit dfb985b620b4fcdc536885fbb12a6d99b582604d[m
Author: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>
Date:   Mon Jan 1 10:01:24 2024 -0800

    build(deps): bump golang.org/x/crypto from 0.14.0 to 0.17.0 in /v3 (#784)
    
    Bumps [golang.org/x/crypto](https://github.com/golang/crypto) from 0.14.0 to 0.17.0.
    - [Commits](https://github.com/golang/crypto/compare/v0.14.0...v0.17.0)
    
    ---
    updated-dependencies:
    - dependency-name: golang.org/x/crypto
      dependency-type: direct:production
    ...
    
    Signed-off-by: dependabot[bot] <support@github.com>
    Co-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit 832a1eae0de3256f99472d85f752fdcc9a4f024f[m
Author: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>
Date:   Mon Jan 1 09:44:24 2024 -0800

    build(deps): bump golang.org/x/crypto in /v3/cmd/genTestCerts (#785)
    
    Bumps [golang.org/x/crypto](https://github.com/golang/crypto) from 0.14.0 to 0.17.0.
    - [Commits](https://github.com/golang/crypto/compare/v0.14.0...v0.17.0)
    
    ---
    updated-dependencies:
    - dependency-name: golang.org/x/crypto
      dependency-type: indirect
    ...
    
    Signed-off-by: dependabot[bot] <support@github.com>
    Co-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>

[33mcommit d4e2de02f88f8ce38b06f1835b8bcbda72bb2ca9[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sat Dec 16 06:49:58 2023 -0800

    Fix goreleaser deprecation (#783)
    
    * Fix goreleaser deprecation
    
    * correction example syntax

[33mcommit f830602323170ee78c35dee0c6fb5218a667a247[m
Author: MTG <36234449+mtgag@users.noreply.github.com>
Date:   Sat Dec 16 15:07:29 2023 +0100

    Added IsSMIMEBRCertificate in checkApplies where missing (#780)
    
    * lint about the encoding of qcstatements for PSD2
    
    * Revert "lint about the encoding of qcstatements for PSD2"
    
    This reverts commit 6c2367080d148f4b8c01f96a4c80e3ac55d1ef26.
    
    * util: gtld_map autopull updates for 2021-10-21T07:25:20 UTC
    
    * always check and perform the operation in the execution
    
    * synchronised with project
    
    * synchronised with project
    
    * synchronised with project
    
    * added util.IsSMIMEBRCertificate(c) where missing, updated test data
    
    * removed GIT merge hints
    
    ---------
    
    Co-authored-by: mtg <git@mtg.de>
    Co-authored-by: GitHub <noreply@github.com>
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit c1aacb0afe4c5dd97d1542c27e9d4f2cfc21ecbf[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sat Dec 16 05:56:03 2023 -0800

    golangci-lint update and fixes (#782)
    
    * Code Linter Update
    
    * linter suggestions
    
    * fixing code lints

[33mcommit f90a51ecb3d36a190d0ca90ed3e5c5d80203ac72[m
Author: github-actions[bot] <41898282+github-actions[bot]@users.noreply.github.com>
Date:   Sat Dec 16 04:54:37 2023 -0800

    util: gtld_map autopull updates for 2023-12-16T12:21:31 UTC (#778)
    
    * util: gtld_map autopull updates for 2023-12-12T16:20:34 UTC
    
    * Triggering CICD
    
    ---------
    
    Co-authored-by: GitHub <noreply@github.com>
    Co-authored-by: christopher-henderson <chris@chenderson.org>

[33mcommit 67537e945e1e8008157589a05a03f94ec57f031b[m
Merge: 24085437 45de8804
Author: mtgag <githreg@mtg.de>
Date:   Thu Dec 14 07:04:42 2023 +0100

    synchronised with project

[33mcommit 24085437aa4e9a39b1f3ac86350774d68432055a[m
Author: mtgag <githreg@mtg.de>
Date:   Thu Dec 14 07:02:35 2023 +0100

    synchronised with project

[33mcommit 45de88040a22e2db4d962de9ec3847dcac59be92[m
Author: MTG <36234449+mtgag@users.noreply.github.com>
Date:   Tue Dec 12 17:10:21 2023 +0100

    refactor of SMIME aia contains (#777)
    
    * lint about the encoding of qcstatements for PSD2
    
    * Revert "lint about the encoding of qcstatements for PSD2"
    
    This reverts commit 6c2367080d148f4b8c01f96a4c80e3ac55d1ef26.
    
    * util: gtld_map autopull updates for 2021-10-21T07:25:20 UTC
    
    * always check and perform the operation in the execution
    
    * synchronised with project
    
    * synchronised with project
    
    * changed date, added check for existent extension
    
    * updates in config after tests
    
    * removed accidentally commited file
    
    * removed internal names part, kept only has http only
    
    * changes addressing discussion in PR. Internal names are checked, IP addresses are skipped.
    
    * the check for HTTP scheme is not needed here. This is covered by the other lint
    
    * fixed test
    
    * renamed file
    
    * one lint for internal names in AIA covers all S/MIME generations, legacy AIA has one HTTP moved to a new lint, added isSubscriberCert for all checkApplies
    
    ---------
    
    Co-authored-by: mtg <git@mtg.de>
    Co-authored-by: GitHub <noreply@github.com>
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit bc2c0fda0533178d96e8be0c73b01168a47e6304[m
Author: Eliot <145681652+eliot-gs@users.noreply.github.com>
Date:   Mon Dec 4 15:01:51 2023 +0000

     CABF SMIME BR Appendix A.1 - countryName matches registration scheme id (#768)
    
    * lint and unit test subject country in organization id
    
    * add lints and unit test for matching country in format id
    
    * deletes accidential workflow additions
    
    * updates according to PR comments
    
    * fixes indentation
    
    * updates following PR  comments
    
    * updates comment and formatting
    
    ---------
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit 7f6ef92e44f595d537ca8f0df6a7770090c11a50[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sun Dec 3 10:58:28 2023 -0800

    Metalint for checking against the deprecaetd lint.RegisterLint function (#775)
    
    * Metalint for checking against the deprecaetd lint.RegisterLint function
    
    * go imports

[33mcommit ebf2071ba0d7adb820a50b52fce8ea42df6b8e0b[m
Author: github-actions[bot] <41898282+github-actions[bot]@users.noreply.github.com>
Date:   Sun Dec 3 07:47:25 2023 -0800

    util: gtld_map autopull updates for 2023-11-27T16:20:42 UTC (#773)
    
    Co-authored-by: GitHub <noreply@github.com>

[33mcommit cee805f2b497508dfbf118a3d654072273b62bdb[m
Merge: 88c933e1 c35c9b9a
Author: mtgag <githreg@mtg.de>
Date:   Sun Dec 3 10:25:36 2023 +0100

    Merge https://github.com/zmap/zlint

[33mcommit c35c9b9a6aefebe6dcc4b1f003820776637561be[m
Author: Martijn Katerbarg <martijn.katerbarg@sectigo.com>
Date:   Mon Nov 27 17:04:44 2023 +0100

    Policy Qualifiers other than id-qt-cps are no longer allowed as per CABF BRs (#774)
    
    * feat: User Notice is no longer allowed as per CABF BRs
    
    * fix: Set proper title and description
    
    * fix: Rename files and align function names

[33mcommit 1bb58f0cc7b6b31cda2e93b6ffdbd038866ea136[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sun Nov 19 13:10:22 2023 -0800

    Updating certificate lint template to use the new certificate specific interface (#772)
    
    * Updating certificate lint template to use new interface
    
    * use tabs instead of space

[33mcommit 96a479935c3e90699e3f5e1f96a3386488af856b[m
Author: github-actions[bot] <41898282+github-actions[bot]@users.noreply.github.com>
Date:   Sun Nov 19 12:48:38 2023 -0800

    util: gtld_map autopull updates for 2023-11-17T20:19:40 UTC (#771)
    
    Co-authored-by: GitHub <noreply@github.com>
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit a08efa8121dc8e72c23bed0cf9dc473b9dfa32b7[m
Author: mara-soldan <87363716+mara-soldan@users.noreply.github.com>
Date:   Sun Nov 19 20:40:15 2023 +0000

    CABF SMIME BR 7.1.2.3.m - Adobe Extensions (#763)
    
    * add lints for adobe extensions presence and criticality in smime certs
    
    * move adobe extensions to preserve alphabetical order
    
    * update timestamp references and use new CertificateLint type
    
    * Update v3/lints/cabf_smime_br/lint_adobe_extensions_legacy_multipurpose_criticality.go
    
    Co-authored-by: Rob <3725956+robplee@users.noreply.github.com>
    
    * Update v3/lints/cabf_smime_br/lint_adobe_extensions_strict_presence.go
    
    Co-authored-by: Rob <3725956+robplee@users.noreply.github.com>
    
    * Update v3/lints/cabf_smime_br/lint_adobe_extensions_legacy_multipurpose_criticality.go
    
    Co-authored-by: Rob <3725956+robplee@users.noreply.github.com>
    
    * update comments
    
    ---------
    
    Co-authored-by: marahrehorciuc <mara.hrehorciuc@globalsign.com>
    Co-authored-by: Rob <3725956+robplee@users.noreply.github.com>
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit 45e62047222f1ac864c5ad5c2afbb8d0c1bcd10e[m
Author: Amir Omidi <amir@aaomidi.com>
Date:   Sun Nov 19 15:20:10 2023 -0500

    Convert all Lints to CertificateLints (#767)

[33mcommit 43b6954c46a5e9475c827a837d67932a31ed0b0e[m
Author: Rob <3725956+robplee@users.noreply.github.com>
Date:   Sun Nov 12 22:41:06 2023 +0000

    address smime lint applicability issue.  regenerate test certificates to fix unit tests broken by change (#764)
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit e8c0c248cc6815a9b69b2c5cfe8eb7377392d57c[m
Author: github-actions[bot] <41898282+github-actions[bot]@users.noreply.github.com>
Date:   Fri Nov 10 06:23:02 2023 -0500

    util: gtld_map autopull updates for 2023-11-06T23:18:29 UTC (#756)
    
    Co-authored-by: GitHub <noreply@github.com>

[33mcommit 64533b5c98de1b52db665bbaee9cd4992d4c3519[m
Author: BJ Cardon <cardonator@users.noreply.github.com>
Date:   Mon Nov 6 15:49:28 2023 -0700

    Ensure AIA URLs point to public paths (#760)
    
    * added lints to check if the aia has likely internal names
    
    * add tests for all aia path combinations
    
    * use Hostname instead of Host to account for ports, triage integration test results and update integration config
    
    * address code review feedback (Fatal->Error, handling for http schemes)
    
    * handle https as well
    
    * enforce http scheme, fix test data
    
    * don't require any OCSPServer to exist
    
    * also don't require IssuingCertificateURLs

[33mcommit 89231704f987ed5002618333c59d8004882aa1d6[m
Author: Rob <3725956+robplee@users.noreply.github.com>
Date:   Sun Nov 5 15:40:31 2023 +0000

    CABF SMIME BR 7.1.2.3.e - KeyUsages (#757)
    
    * add lints for smime ku presence and criticality, rsa KUs and ECC KUs
    
    * Finish lint for ECDSA key usages.  Add lint for edwards curve key usages
    
    * strict rsa ku lint unit tests
    
    * rename rsa strict ku lint test data to reflect strictness of SMIME policy oid
    
    * add unit tests to smime rsa legacy/multipurpose ku lint
    
    * add unit tests to key usage presence lint.  Fix present/presence typos
    
    * rename key usage critical lint to key usage criticality. unit tests for same
    
    * add unit tests to smime ecdsa key usage lint.  Fix issue in check applies
    
    * add unit tests for smime ed25519 ku lint
    
    * use iota constants for signing, key management and dual use to make rsa and ec ku lints clearer to read
    
    * replace bit mask checks with util.HasKeyUsage calls in smime KU lints
    
    * Refactor RSA and EC SMIME KU lints to cover other KUs without digitalSignature and/or keyAgreement/Encipherment with separate lints.
    
    ---------
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit f9f30bcd3fe1718c3022c7c6e45709ef7f9f0b60[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sun Nov 5 01:10:00 2023 -0700

    Fixing lint registration for CABF SMIME (#761)

[33mcommit 1c307f4b9ef04348f621ce0a03e2bf0d5e471fbc[m
Author: Rob <3725956+robplee@users.noreply.github.com>
Date:   Sun Oct 22 19:36:04 2023 +0100

    Lints for CABF SMIME BRs 7.1.2.3.f - EKUs (#747)
    
    * Add lints to enforce SMIME BR EKU restrictions
    
    * Tidy up smime_policies util file by removing some unused code.
    
    * Address issue raised with mailbox validated field restrictions lint checkApplies
    
    * Add subscriber certificate requirement to EKU lint CheckApplies functions
    
    ---------
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit 553276dabd988e4c1645c0bed62bed516e480cd8[m
Author: github-actions[bot] <41898282+github-actions[bot]@users.noreply.github.com>
Date:   Thu Oct 19 10:39:48 2023 -0700

    util: gtld_map autopull updates for 2023-10-19T17:18:28 UTC (#755)
    
    Co-authored-by: GitHub <noreply@github.com>

[33mcommit 2f544868a7a6d13a2dcc0b9dc42db7290dfaff7d[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sun Oct 15 18:14:19 2023 -0700

    CABF SMIME 7.1.4.2.h If present, the subject:emailAddress SHALL contain a single Mailbox Address (#752)
    
    * CABF SMIMS 4.1.4.2.h If present, the subject:emailAddress SHALL contain a single Mailbox Address
    
    * go imports the files

[33mcommit 2f0f4b8a071d4fb0699aba6e5f255336567f73c0[m
Author: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>
Date:   Sun Oct 15 09:30:12 2023 -0700

    build(deps): bump golang.org/x/net in /v3/cmd/genTestCerts (#751)
    
    Bumps [golang.org/x/net](https://github.com/golang/net) from 0.7.0 to 0.17.0.
    - [Commits](https://github.com/golang/net/compare/v0.7.0...v0.17.0)
    
    ---
    updated-dependencies:
    - dependency-name: golang.org/x/net
      dependency-type: indirect
    ...
    
    Signed-off-by: dependabot[bot] <support@github.com>
    Co-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>

[33mcommit 378c09f71e33d16f8d0d87b6bd78911902444817[m
Author: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>
Date:   Sun Oct 15 09:14:13 2023 -0700

    build(deps): bump golang.org/x/net from 0.8.0 to 0.17.0 in /v3 (#750)
    
    Bumps [golang.org/x/net](https://github.com/golang/net) from 0.8.0 to 0.17.0.
    - [Commits](https://github.com/golang/net/compare/v0.8.0...v0.17.0)
    
    ---
    updated-dependencies:
    - dependency-name: golang.org/x/net
      dependency-type: direct:production
    ...
    
    Signed-off-by: dependabot[bot] <support@github.com>
    Co-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit 88e01adce26bde53086eff8b3ec14f7207e18e62[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sun Oct 15 08:56:47 2023 -0700

    Lint for CABF SMIME 7.1.2.3.h - subjectAlternativeName SHOULD NOT be marked critical unless the subject field is an empty sequence (#746)
    
    * Lint for CABF SMIME 7.1.2.3.h - subjectAlternativeName SHOULD NOT be marked critical unless the subject field is an empty sequence.
    
    * removing implied warning interpretation

[33mcommit 08a9354f3002c9eb0b8823d663cd71e7d5a14aa3[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sun Oct 15 08:28:31 2023 -0700

    Lint for CABF SMIME 7.1.2.3.h - subjectAlternativeName, all: SHALL be present (7.1.2.3.h) (#744)
    
    * Lint for CABF SMIME 7.1.2.3.h - subjectAlternativeName, all: SHALL be present (7.1.2.3.h)
    
    * not exporting

[33mcommit 386a8dc413add9bb92d80badbb4d86f833f6a4e5[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sun Oct 8 08:33:38 2023 -0700

    Lint for CABF SMIME 7.1.2.3b - cRLDistributionPoints SHALL be present (#742)
    
    * Lint for CABF SMIME 7.1.2.3b - cRLDistributionPoints SHALL be present
    
    * adressing linter
    
    * correcting copying error
    
    * fixing typo in filename

[33mcommit 48baa89c49e477e223ac2c4644046530869d8af7[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Wed Sep 27 16:16:51 2023 -0700

    Permit underscores in DNSNames if-and-only-if replacing all underscores results in valid LDH labels during BR 1.6.2's permissibility period (#661)
    
    Co-authored-by: David Adrian <davidcadrian@gmail.com>

[33mcommit ba30b3b851aa56363f426ec4dc2377c0e43314b3[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Wed Sep 27 14:11:50 2023 -0700

    Permit underscores in DNSNames if-and-only-if those certificates are valid for less than 30 days and during BR 1.6.2's permissibility period (#660)
    
    Co-authored-by: Ryan Sleevi <ryan.sleevi@gmail.com>
    Co-authored-by: Zakir Durumeric <zakird@gmail.com>
    Co-authored-by: David Adrian <davidcadrian@gmail.com>

[33mcommit 1fd1c0d6964020e389955fb582c32ab603ceed28[m
Author: BJ Cardon <cardonator@users.noreply.github.com>
Date:   Sun Sep 17 17:53:44 2023 -0600

    Part 1 of SC-62 related updates to zlint (#739)
    
    * Updated lint for common name handling. The definition for the CN field has switched from deprecated to NOT RECOMMENDED (essentially SHOULD NOT). An IneffectiveDate was added to the original lint.
    
    Added a new lint for subscriber cert basic constraints checking. Post-SC62, basicConstraint MAY be included but MUST be critical if present.
    
    Added a date for SC62 Effective
    
    * fix CheckApplies
    
    * edited the wrong file, reverted and edited the right file.
    
    * add PEMs that exercise the tests properly
    
    * Update v3/lints/cabf_br/lint_sub_cert_basic_constraints_not_critical.go
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>
    
    * Update v3/lints/cabf_br/lint_sub_cert_basic_constraints_not_critical.go
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>
    
    * fix missing import
    
    ---------
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit 5c4e05fe5f30e1e0b68434c077009c3f15974f77[m
Author: github-actions[bot] <41898282+github-actions[bot]@users.noreply.github.com>
Date:   Sun Sep 17 09:40:14 2023 -0700

    util: gtld_map autopull updates for 2023-08-27T22:18:12 UTC (#737)
    
    Co-authored-by: GitHub <noreply@github.com>

[33mcommit 88c933e135487a4f39a75767e9207e880d4c040b[m
Merge: d4f2f9f2 71d5e4b1
Author: mtgag <githreg@mtg.de>
Date:   Wed Aug 30 10:04:10 2023 +0200

    Merge https://github.com/zmap/zlint

[33mcommit d4f2f9f20715c9d7f4c617254749917cce4834be[m
Author: mtgag <githreg@mtg.de>
Date:   Wed Aug 30 09:58:56 2023 +0200

    synchronised with project

[33mcommit 71d5e4b1cbf7636331bbf1c839d24abb4a945633[m
Author: Paul van Brouwershaven <vanbroup@users.noreply.github.com>
Date:   Sun Aug 27 23:26:07 2023 +0200

    Reintroduce lint for inconsistent KU and EKU (#708)
    
    * Add function to get human friendly KeyUsage names
    
    * Add lint to check for KU and EKU inconsistency
    
    * Add func to get EKU strings
    
    * Sort KeyUsage strings for consistency in messages
    
    * Consider multiple purposes
    
    * Update result for integration test
    
    * Fix formatting
    
    * Add KU/EKU inconsistent test cases
    
    * No error on undefined extended key usage
    
    * Move sort from util to lint and include comment
    
    * Add some comments around the cyclomatic complexity
    
    * Update count for test corpus incl email certs

[33mcommit 59d4dd332041087118bbbe86f7c5870c8bad980d[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sun Aug 20 09:11:06 2023 -0700

    Inclusion of approximately 190000 email protection certificates into the test corpus (#738)

[33mcommit d959c8318c817be31cf3e9823bfd146d7c218675[m
Author: Rob <3725956+robplee@users.noreply.github.com>
Date:   Sun Aug 13 16:27:31 2023 +0100

    Add lint enforcing the restrictions on subject DN fields for mailbox validated SMIME certificates (#713)
    
    * Add lint enforcing the restrictions on subject DN fields for mailbox validated SMIME certificates
    
    * Add zlint copyright text to new files.
    
    * Add cabf_smime_br lint source to TestNotMissingAnyLintSources
    
    * refactor lint to add lists of allowed and forbidden fields into the lint struct
    
    * rename mailboxValidatedEnforceSubjectFieldRestrictions lint to no longer export the underlying struct as per other lints in zlint
    
    * Update mailbox lint to use new certificatelint interface
    
    * fix mailbox validated field lint unit tests, reorganise smime testdata, remove unused test certificates
    
    * Update v3/lints/cabf_smime_br/mailbox_validated_enforce_subject_field_restrictions.go comment to list relevant policy OIDs
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>
    
    * attempt to address lint complaint with comment describing CheckApplies of mailbox field presence lint
    
    * Add explanatory comment to IsEmailProtectionCert
    
    * Fix styling in time.go
    
    ---------
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit 624744d33e68fb899766d58299f84cd1df2d680a[m
Author: Amir Omidi <amir@aaomidi.com>
Date:   Tue Aug 1 18:58:56 2023 -0400

    Include LintMetadata in the LintResult (#729)
    
    * Include LintMetadata in the LintResult
    
    * Don't include LintMetadata in LintResult's JSON output

[33mcommit 38b74849c31c105e7d8469b6efa2d9f9f45281f5[m
Author: Amir Omidi <amir@aaomidi.com>
Date:   Tue Aug 1 16:31:49 2023 -0400

    Add CRL Lints for the ReasonCode extension from the baseline requirements and RFC 5280 (#715)
    
    Add CRL Lints for the ReasonCode extension from the baseline requirements and RFC 5280.
    
    https://github.com/zmap/zlint/pull/715
    
    Co-authored-by: Rob <3725956+robplee@users.noreply.github.com>
    Co-authored-by: David Adrian <davidcadrian@gmail.com>

[33mcommit 1e3cf0111c7f97688d9037d9e58423883aeb9723[m
Author: github-actions[bot] <41898282+github-actions[bot]@users.noreply.github.com>
Date:   Sat Jul 29 09:50:44 2023 -0700

    util: gtld_map autopull updates for 2023-07-25T22:18:37 UTC (#736)
    
    Co-authored-by: GitHub <noreply@github.com>

[33mcommit b492fe7cd7618e7c1bf81217f3a9b42a6c391652[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Thu Jul 20 13:09:40 2023 -0400

    tidy: delete 'h' gitlog fragment from proj. root. (#735)
    
    In 4d38bfea a hunk of ANSI decorated `git log` output was committed to
    the root of the repository. This commit deletes it.

[33mcommit 4d38bfea8756d7b4fd4ebfb36aee8675a9eeeed4[m
Author: MTG <36234449+mtgag@users.noreply.github.com>
Date:   Sun Jul 9 20:50:43 2023 +0200

    E ext cert policy disallowed any policy qualifier refactor (#732)
    
    * lint about the encoding of qcstatements for PSD2
    
    * Revert "lint about the encoding of qcstatements for PSD2"
    
    This reverts commit 6c2367080d148f4b8c01f96a4c80e3ac55d1ef26.
    
    * util: gtld_map autopull updates for 2021-10-21T07:25:20 UTC
    
    * always check and perform the operation in the execution
    
    * synchronised with project
    
    * refactored implementation, tests, and testdata
    
    * refactored implementation
    
    * addressing high cyclomatic complexity
    
    * code format
    
    * code format
    
    ---------
    
    Co-authored-by: mtg <git@mtg.de>
    Co-authored-by: GitHub <noreply@github.com>
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit 7602109a26ca74845409fd61c18f698fe01930b0[m
Author: github-actions[bot] <41898282+github-actions[bot]@users.noreply.github.com>
Date:   Sun Jul 9 11:20:36 2023 -0700

    util: gtld_map autopull updates for 2023-07-08T13:20:31 UTC (#733)
    
    Co-authored-by: GitHub <noreply@github.com>
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit 40f2b32c4a866076818267d9ffbf4b9baa925e74[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sun Jul 9 11:18:59 2023 -0700

    Duplicate lints about keyIdentifier in certificates (#726)
    
    * Duplicate lints about keyIdentifier in certificates
    
    * fixed go imports styling
    
    * breaking up code comments to match conditional blocks
    
    * typo
    
    * simplifying check
    
    * Triggering GHA with empty commit
    
    * adding one more error cert to the corpus
    
    ---------
    
    Co-authored-by: Zakir Durumeric <zakird@gmail.com>

[33mcommit 3f1605e8704ade3a3f95c4b9a1392cdcf88fe3f9[m
Author: MTG <36234449+mtgag@users.noreply.github.com>
Date:   Sun Jul 9 20:05:31 2023 +0200

    Ecdsa ee invalid ku check applies (#731)
    
    * lint about the encoding of qcstatements for PSD2
    
    * Revert "lint about the encoding of qcstatements for PSD2"
    
    This reverts commit 6c2367080d148f4b8c01f96a4c80e3ac55d1ef26.
    
    * util: gtld_map autopull updates for 2021-10-21T07:25:20 UTC
    
    * always check and perform the operation in the execution
    
    * check applies could also check if the extension is present
    
    ---------
    
    Co-authored-by: mtg <git@mtg.de>
    Co-authored-by: GitHub <noreply@github.com>
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit 1652cfa597d7c4c37991484d35e4a6da57a06580[m
Author: mtgag <githreg@mtg.de>
Date:   Wed Jul 5 07:03:20 2023 +0200

    synchronised with project

[33mcommit 92902fc7d9ae7ad9f221235c74b992be6f101812[m
Merge: 526f9be2 8c46bdf0
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
Merge: b52111ba 45e8dff6
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
Merge: 351a3798 24385962
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
Merge: 92e659c5 a5c869f8
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
Merge: 8600050f 997ad514
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

[33mcommit b7abf25bdffae0b85a5c1058ac0dbf9775675803[m
Author: Rob <3725956+robplee@users.noreply.github.com>
Date:   Sun Jun 12 19:53:47 2022 +0100

    Add new lint to block organisational unit names as of 1st September 2022 (#675)
    
    * Add new lint to block organisational unit names as of 1st September 2022
    
    * update copyright year in all files changed by this PR
    
    * update name of date variable for ou prohibition
    
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit c32f6d3f7bdfa4d6773b1bc6bc60c36c93d6843e[m
Author: James Kasten <jdkasten@umich.edu>
Date:   Thu Jun 9 21:50:05 2022 -0700

    Fix SPKI Encoding Lint's RSA BR Section (#679)
    
    The RSA AlgorithmIdentifier is specified in 7.1.3.1.1. ECDSA is referenced in 7.1.3.1.2
    https://github.com/cabforum/servercert/blob/main/docs/BR.md#71311-rsa

[33mcommit ed6287a54ce1e6297e2576730e65b5c2f4faddcb[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sun Jun 5 11:16:44 2022 -0700

    Zlint incorrectly requires TorServiceDescriptors if V3 onion and DNS name (#677)
    
    * Correct false negative in the presence of a DNS name

[33mcommit 74f454196357f798ca087df6d43e80d9a7a4debd[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sat Apr 16 11:28:06 2022 -0700

    Update to Go 1.18 and update GolangCI Linter (#672)
    
    * upgrading the repo to Go 1.18

[33mcommit a34c016cb0f6d4e79fe584939e2a52fc68fd68a7[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Fri Apr 15 10:38:52 2022 -0700

    QoL changes to genTestCert.go (#664)
    
    Co-authored-by: Zakir Durumeric <zakird@gmail.com>

[33mcommit 20aeab4d82749f573c3a85dc48ad862f8a2c111c[m
Author: github-actions[bot] <41898282+github-actions[bot]@users.noreply.github.com>
Date:   Fri Apr 15 09:47:47 2022 -0700

    util: gtld_map autopull updates for 2022-04-15T16:45:51 UTC (#671)
    
    Co-authored-by: GitHub <noreply@github.com>
    Co-authored-by: Zakir Durumeric <zakird@gmail.com>

[33mcommit 6d874e67f06bf5e62ad2471c3d9627094e67fd2d[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Fri Apr 15 09:45:38 2022 -0700

    updating to zcrypto 599ec18ecbac (#670)

[33mcommit b3be71cf576a4f17272f675f5ee9cda66ccbabd5[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sun Mar 27 10:06:05 2022 -0700

    Skip checking for a Tor Descriptor Hash if the provided cert contains a V3 Onion address. (#669)
    
    * check for v3 addresses before asserting presence of tor descriptor hash
    
    * fixing linter

[33mcommit 3be391b56b004fbc501f4fc0c1edfb53b71a2937[m
Author: Pablo Díaz <59196303+Pabloanf@users.noreply.github.com>
Date:   Fri Mar 4 17:57:47 2022 +0100

    Update README.md (#666)
    
    * Update README.md
    
    Added ANF AC to the bullet list of CAs that integrate with this linter.
    
    * Alphabetizing.
    
    * Update README.md
    
    Co-authored-by: Zakir Durumeric <zakird@gmail.com>

[33mcommit b1bd967fe787933fdbb3be70377ab2508045c401[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sun Feb 20 10:22:25 2022 -0800

    No underscores are allowed in DNSNames before BR 1.6.2's permissibility period (#659)
    
    * no underscores before BR 1.6.2
    
    Co-authored-by: Ryan Sleevi <ryan.sleevi@gmail.com>

[33mcommit 6badb89602ca1102e34a0c1fe8b2c7b96ead639d[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sun Feb 20 10:10:57 2022 -0800

    No underscores are allowed in DNSNames after BR 1.6.2's permissibility period (#662)
    
    * underscores not permissible after hard enforcement

[33mcommit 4ab856795cdd8185c273ae873c52e37bea8535c4[m
Author: github-actions[bot] <41898282+github-actions[bot]@users.noreply.github.com>
Date:   Thu Feb 17 18:09:19 2022 -0500

    util: gtld_map autopull updates for 2022-02-17T22:26:31 UTC (#658)
    
    Co-authored-by: GitHub <noreply@github.com>

[33mcommit 7fc9fbd55c2f6bdca4cd5dee825da903b00844db[m
Author: Ryan Sleevi <ryan.sleevi@gmail.com>
Date:   Tue Jan 25 21:09:47 2022 -0500

    Add Microsoft to the known-ZLint users (#655)

[33mcommit b4a225e88a0a05bd9aed3bde83725a5168331b82[m
Author: MTG <36234449+mtgag@users.noreply.github.com>
Date:   Sat Jan 15 21:37:08 2022 +0100

    AlgorithmIdentifier encoding (Section 7.1.3.1, CAB-Forum BR) (#642)
    
    * lint about the encoding of qcstatements for PSD2
    
    * Revert "lint about the encoding of qcstatements for PSD2"
    
    This reverts commit 6c2367080d148f4b8c01f96a4c80e3ac55d1ef26.
    
    * util: gtld_map autopull updates for 2021-10-21T07:25:20 UTC
    
    * added lint for proper encoding of public key accoring to cab_br
    
    * fixed prefix error
    
    * Update v3/lints/cabf_br/lint_subject_public_key_info_improper_algorithm_object_identifier_encoding.go
    
    Co-authored-by: Ryan Sleevi <ryan.sleevi@gmail.com>
    
    * refactored lint after review
    
    * solving review issue
    
    Co-authored-by: mtg <git@mtg.de>
    Co-authored-by: GitHub <noreply@github.com>
    Co-authored-by: Ryan Sleevi <ryan.sleevi@gmail.com>
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit da67a2330fc2d880b0c751b55d2de5dddd8c6b86[m
Author: github-actions[bot] <41898282+github-actions[bot]@users.noreply.github.com>
Date:   Sun Jan 9 10:57:07 2022 -0800

    util: gtld_map autopull updates for 2021-12-30T02:43:35 UTC (#654)
    
    Co-authored-by: GitHub <noreply@github.com>

[33mcommit 3f7cf6cbc2f56d3fc1e06437b6d0fc69089abdc0[m
Author: Leo Grove <leo@ssl.com>
Date:   Sun Dec 12 21:47:56 2021 -0600

    Update README.md (#653)
    
    * Update README.md
    
    Adding SSL.com to the list of CAs that integrate with this linter.
    
    * Alphabetizing
    
    Co-authored-by: Zakir Durumeric <zakird@gmail.com>

[33mcommit 9199b6d9326f7be600458d3ac7de892ade1467db[m
Author: github-actions[bot] <41898282+github-actions[bot]@users.noreply.github.com>
Date:   Sat Dec 11 10:00:37 2021 -0800

    util: gtld_map autopull updates for 2021-12-09T20:29:24 UTC (#649)
    
    Co-authored-by: GitHub <noreply@github.com>
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit 0d7125864bc4f6baa17426aff2f884e8901fde50[m
Author: Paul van Brouwershaven <vanbroup@users.noreply.github.com>
Date:   Sat Dec 11 18:52:05 2021 +0100

    Entrust Datacard rebranded to Entrust (#652)

[33mcommit bbc7e360e6f8ae223ddd4b1f3e4c460294f2e7f2[m
Author: Paul van Brouwershaven <vanbroup@users.noreply.github.com>
Date:   Thu Dec 9 20:21:51 2021 +0100

    Add lint to detect IP addresses in EV certs (#650)

[33mcommit cb3e7e86e1cf73c82622de4768f15e69599d0751[m
Author: Paul van Brouwershaven <vanbroup@users.noreply.github.com>
Date:   Thu Dec 9 20:05:53 2021 +0100

    Mark CA/Browser Forum EV Policy OID as EV (#651)
    
    CAs are required to use this OID after 2020-09-30 (per CA/B Forum Ballot SC31 - https://cabforum.org/2020/07/16/ballot-sc31-browser-alignment/ ), so all new EV certs since then can be detected just by looking for this OID.

[33mcommit da4e374e427291aa8d3acaf860623ea83b14d915[m
Author: Eng Zer Jun <engzerjun@gmail.com>
Date:   Sun Nov 14 06:00:35 2021 +0800

    refactor: move from io/ioutil to io and os packages (#647)
    
    The io/ioutil package has been deprecated as of Go 1.16, see
    https://golang.org/doc/go1.16#ioutil. This commit replaces the existing
    io/ioutil functions with their new definitions in io and os packages.
    
    Signed-off-by: Eng Zer Jun <engzerjun@gmail.com>

[33mcommit 3a3de3c3cc9d3b1d9c1c191f395777dd7a4d5d0d[m
Author: github-actions[bot] <41898282+github-actions[bot]@users.noreply.github.com>
Date:   Sat Nov 13 10:12:36 2021 -0800

    util: gtld_map autopull updates for 2021-10-30T04:36:00 UTC (#637)
    
    Co-authored-by: GitHub <noreply@github.com>
    Co-authored-by: Zakir Durumeric <zakird@gmail.com>

[33mcommit 2ff21301bb5f5f5d41a9999705e2bf60b88ebb6b[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sat Nov 13 09:49:53 2021 -0800

    cleaning up some datetime logic (#644)

[33mcommit 8600050f905393376bd091ded9da59a205fde045[m
Merge: 749d8960 e56e2a09
Author: MTG <36234449+mtgag@users.noreply.github.com>
Date:   Thu Oct 21 09:42:15 2021 +0200

    Merge pull request #1 from mtgag/zlint-gtld-update
    
    util: gtld_map autopull updates for 2021-10-21T07:25:20 UTC

[33mcommit e56e2a09361056ae4f3d9ed9e03624bfbe2fb0cb[m
Author: GitHub <noreply@github.com>
Date:   Thu Oct 21 07:26:00 2021 +0000

    util: gtld_map autopull updates for 2021-10-21T07:25:20 UTC

[33mcommit 749d89604a42279f37efdc7f65a16a8814fc532a[m
Merge: 28481cc7 cb17369b
Author: mtg <git@mtg.de>
Date:   Thu Oct 21 09:13:49 2021 +0200

    Merge https://github.com/zmap/zlint

[33mcommit cb17369b4628c684ac68c1fc169ff2a38c00cfdf[m
Author: Corey Bonnell <corey.j.bonnell@outlook.com>
Date:   Tue Oct 19 13:35:30 2021 -0400

    Lint for Non-XN Reserved Labels (#635)
    
    * Lint for Non-XN Reserved Labels
    
    * Refactor to use idna functions
    
    Co-authored-by: Corey Bonnell <corey.bonnell@digicert.com>
    Co-authored-by: Zakir Durumeric <zakird@gmail.com>
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit 9113ed8c1f1dd14ca6e19e2d6096fdda8885dd09[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sun Oct 17 10:50:24 2021 -0700

    Forbid wildcard certs for non .onion EVs (#641)
    
    * adding lint to forbid wildcard certs for non .onion EVs

[33mcommit 0508b86cf4c558ad17daf7d4d3438dadaaf33376[m
Author: Corey Bonnell <corey.j.bonnell@outlook.com>
Date:   Sat Oct 16 13:18:59 2021 -0400

    Detect XN-Labels case-insensitively (#636)
    
    * Detect XN-Labels case-insensitively
    
    * Incorporate Chris's refactoring suggestion to create idna functions
    
    Co-authored-by: Corey Bonnell <corey.bonnell@digicert.com>
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit b6ec3270b8ff9c141e335a41e414beaaee0f1485[m
Author: github-actions[bot] <41898282+github-actions[bot]@users.noreply.github.com>
Date:   Tue Oct 5 18:42:05 2021 -0400

    util: gtld_map autopull updates for 2021-10-05T22:26:49 UTC (#633)
    
    Co-authored-by: GitHub <noreply@github.com>

[33mcommit b4060ec70d7ec1d6203a837dafa859710e3e3a9c[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sun Sep 26 11:04:29 2021 -0700

    Correct lint attribution for dnsname_etc lints and limit scope to just DNS SAN entries (#609)
    
    Adding an RFC5280 specific version that does not check for CommonName of the dnsname lints that are already present in CABF lints

[33mcommit 74dfff29b023931eed2372d248035c2d2e62d394[m
Author: Attila Rozgonyi <81579568+attilarozgonyi@users.noreply.github.com>
Date:   Wed Sep 8 16:43:59 2021 +0200

    Update README.md (#631)
    
    Added Microsec to the bullet list "Zlint Users/Integrations".

[33mcommit 0944e91628ccd5b175f4468ab0cde0b7753c1fcd[m
Author: Rob <3725956+robplee@users.noreply.github.com>
Date:   Mon Sep 6 19:35:20 2021 +0100

    e_subject_common_name_not_from_san is no longer sufficient for enforcing CABF BRs (#627)
    
    * Add new subject common name not exactly from san lint.  Ineffective from the previous subject common name not from san lint.  Tests for both
    
    * Fix typo in subject_common
    _name_not_from_san_test
    
    * Update integration test config.json
    
    * Update CABF SC48 deffective date to CABF_1_8_0_Date
    
    * Add test cases covering IP address common name and SAN IP Addresses
    
    * Add tests for extra IPv6 scenarios
    
    * Remove commented out experimental code
    
    * Rename SANWithoutCNSeptember2021 to CNWithoutSANSeptember2021 test certificate to describe its contents correctly
    
    * Extend common name in SAN check to verify all provided CNs are present in SAN fields.  Add tests
    
    * Update v3/lints/cabf_br/lint_subject_common_name_not_exactly_from_san.go
    
    Additional detail information about which CN was missing.
    
    * Update v3/lints/cabf_br/lint_subject_common_name_not_exactly_from_san.go
    
    Adding fmt as import
    
    Co-authored-by: Zakir Durumeric <zakird@gmail.com>
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit 1b894052a035ea9cca7c82326e0d07cef61bcae4[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sat Sep 4 13:16:42 2021 -0700

    bump zcrypto to v0.0.0-20210811211718-6f9bc4aff20f (#629)

[33mcommit 28481cc7ccdd1381f4917ab6094fbf4b3f3bf493[m
Merge: 01996c6f 9da3c9fa
Author: mtg <git@mtg.de>
Date:   Wed Sep 1 12:30:17 2021 +0200

    Merge https://github.com/zmap/zlint

[33mcommit 9da3c9fa110527f4b11902970a4c11f3a4d80d3d[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sat Jul 24 11:33:34 2021 -0700

    disallow duplicate entries in config.json (#616)

[33mcommit 4940d55870f19d5c2ca9447cc2ee0af3cba84138[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sat Jul 24 11:18:31 2021 -0700

    Test certificate generator doesn't create a certificate chain (#622)
    
    * making generated certs chain correctly with correct attributes

[33mcommit e2742152692085cf1c4c6898774c938b7775c9db[m
Author: Rob <3725956+robplee@users.noreply.github.com>
Date:   Sun Jul 18 15:45:37 2021 +0100

    split lint_sub_ca_aia_missing lint into an error lint for before CABF_BR 1.7.1 and a warning for after.  Add test data (#613)

[33mcommit 7bba3627145fc5b3f4618862f9868bef0af0d718[m
Author: Jaime Hablutzel <hablutzel1@gmail.com>
Date:   Sun Jul 18 09:33:05 2021 -0500

    Code clarification to match BRs wording. (#621)

[33mcommit 48b300e877f2ed3e6b7e75426a273f84455f17a6[m
Author: Adriano Santoni <asantoni64@gmail.com>
Date:   Sat Jun 26 14:39:47 2021 +0200

    Update README.md (#614)

[33mcommit 8e8930ee04b5a10052ae29f8b37e8344f0be1d34[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sun Jun 20 14:01:34 2021 -0700

    subsitute the initialize method for a constructor in the Lint struct (#607)
    
    Co-authored-by: Zakir Durumeric <zakird@gmail.com>

[33mcommit dbd9bfd21e5cb6c78c4ba16c87da21492a999c93[m
Author: Denis Issoupov <denis@ekspand.com>
Date:   Sun Jun 20 12:56:09 2021 -0700

    dep: upgrade to latest ZCrypto with permissive asn1 parsing (#611)
    
    * upgrade to ZCrypto with permissive asn1 parsing (#596)
    
    * Merge from master (#610)
    
    * cmd: add `-version` to `zlint`, `zlint-gtld-update`. (#598)
    
    Our GoReleaser configuration already populates a `version` variable in
    LDFLAGS at build time. Prior to this commit we only included the dynamic
    version var in the usage output from `--help`. This made it easy to
    overlook. We also didn't set the dynamic var from the makefile, leaving
    all src builds as the static version "dev".
    
    In this commit we add a `--version` flag to the `zlint` and
    `zlint-gtld-update` commands that prints the dynamic version and
    exits. We also update the `makefile` so that both binaries get built
    with a version that includes the latest tag, and the SHA of the local
    git checkout, e.g. `v3.1.0-19-g0807bf95`. This should better match user
    expectation for CLI tools.
    
    * lints: fix anyKeyUsage typo in `n_mp_allowed_eku`. (#600)
    
    * deps: update zcrypto to ea3fdbd (#604)
    
    * upgrade to ZCrypto with permissive asn1 parsing
    
    Co-authored-by: Daniel McCarney <daniel@binaryparadox.net>
    Co-authored-by: Rob Stradling <rob@sectigo.com>
    
    * deps: update zcrypto to ea3fdbd (#604)
    
    * upgrade to ZCrypto with permissive asn1 parsing
    
    Co-authored-by: Daniel McCarney <daniel@binaryparadox.net>
    Co-authored-by: Rob Stradling <rob@sectigo.com>

[33mcommit 7e75dc35e04f682f0d0eb3de6d4af49ccd5db5af[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Mon May 17 21:59:59 2021 -0400

    deps: update zcrypto to ea3fdbd (#604)

[33mcommit d5d0ed9565c2b2284d0f4eddf8aa83ca7a735bf1[m
Author: Rob Stradling <rob@sectigo.com>
Date:   Thu May 13 13:32:40 2021 +0100

    lints: fix anyKeyUsage typo in `n_mp_allowed_eku`. (#600)

[33mcommit c47eab4fe42cda40cc4c56117869bac4c3850037[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Wed May 12 08:43:29 2021 -0400

    cmd: add `-version` to `zlint`, `zlint-gtld-update`. (#598)
    
    Our GoReleaser configuration already populates a `version` variable in
    LDFLAGS at build time. Prior to this commit we only included the dynamic
    version var in the usage output from `--help`. This made it easy to
    overlook. We also didn't set the dynamic var from the makefile, leaving
    all src builds as the static version "dev".
    
    In this commit we add a `--version` flag to the `zlint` and
    `zlint-gtld-update` commands that prints the dynamic version and
    exits. We also update the `makefile` so that both binaries get built
    with a version that includes the latest tag, and the SHA of the local
    git checkout, e.g. `v3.1.0-19-g0807bf95`. This should better match user
    expectation for CLI tools.

[33mcommit 0807bf95d58b4f0c35831674caf02d40a6972303[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sat Apr 24 11:18:45 2021 -0700

    Updating RFC surname and givenname character limits (#586)
    
    * updating RFC surname and givenname character limits

[33mcommit 3de0a7c3319280bd56a9230f07c70c1526cfda60[m
Author: github-actions[bot] <41898282+github-actions[bot]@users.noreply.github.com>
Date:   Thu Apr 22 08:21:00 2021 -0400

    util: gtld_map autopull updates for 2021-04-22T03:40:32 UTC (#590)
    
    Co-authored-by: GitHub <noreply@github.com>

[33mcommit 5ca3470ab97282d85196552dd0872ea57ed84e16[m
Author: github-actions[bot] <41898282+github-actions[bot]@users.noreply.github.com>
Date:   Wed Apr 21 17:50:42 2021 -0400

    util: gtld_map autopull updates for 2021-04-21T21:31:31 UTC (#589)
    
    Co-authored-by: GitHub <noreply@github.com>

[33mcommit 740b212a296bc321966cfbffd2b58868edcfc217[m
Author: github-actions[bot] <41898282+github-actions[bot]@users.noreply.github.com>
Date:   Sat Apr 17 11:37:47 2021 -0400

    util: gtld_map autopull updates for 2021-04-17T02:48:14 UTC (#588)
    
    Co-authored-by: GitHub <noreply@github.com>

[33mcommit d5ab97e9ca1cef2c7b594672c0265cef3b703637[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Thu Apr 1 07:53:26 2021 -0700

    Make zero an invalid serial number for RFC lints (#584)
    
    * making zero an invalid serial number

[33mcommit 2cac1fd10fa9ad71e692c2ca9aea0f341c64055a[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sun Mar 28 10:02:41 2021 -0700

    Lint that DSA is not used - BR  (#577)
    
    * prohibit DSA in BR and sunset dh_params_missing

[33mcommit 30c55c549f8ea880b958888c8eeb7ea62577b28a[m
Author: Mathew Hodson <mathew.hodson@gmail.com>
Date:   Sun Mar 28 10:52:38 2021 -0400

    lints: fix typo in e_ext_name_constraints_not_critical description (#579)

[33mcommit a6348f94131181245c3cb33f6d06bbe7a6adef85[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sun Mar 28 07:34:19 2021 -0700

    Update zcrypto for vendored crypto/dsa package (#578)
    
    * update to zcrypto@6b615bf2dd2e for vendored crypto/dsa package

[33mcommit 35273f10a0b56408d64b783d07b4a97d11bc1e7d[m
Author: github-actions[bot] <41898282+github-actions[bot]@users.noreply.github.com>
Date:   Fri Mar 26 21:09:00 2021 -0400

    util: gtld_map autopull updates for 2021-03-26T21:30:44 UTC (#580)
    
    Co-authored-by: GitHub <noreply@github.com>

[33mcommit b313d9f438043ac53924afb3e27d38ef17be7780[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sun Mar 14 12:59:50 2021 -0700

    Introduce an upper bounds to effective dates  (#576)
    
    * adding a field for declaring a lint's ineffective date

[33mcommit 3223b2a6047ecf0a9883bd3bc7a40d6418de3ade[m
Author: Rob <3725956+robplee@users.noreply.github.com>
Date:   Wed Mar 3 22:15:03 2021 +0000

    Add a new lint to prohibit using DSA (#572)
    
    * Create a lint to prevent DSA usage

[33mcommit 3615e0fedae577abb7b8cf7837eeac5c29009057[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Wed Mar 3 08:35:14 2021 -0800

    Include a playground script for generating one off certificates and certificate chains (#569)
    
    *Add a playground script for generating test data certificates

[33mcommit 7fcf0da63f3e7467d818c66cd99d54fc1189ba1b[m
Author: github-actions[bot] <41898282+github-actions[bot]@users.noreply.github.com>
Date:   Fri Feb 19 17:43:09 2021 -0500

    util: gtld_map autopull updates for 2021-02-19T22:31:45 UTC (#571)
    
    Co-authored-by: GitHub <noreply@github.com>

[33mcommit 2aa588fc12da8f02a5f0fdb7f6971ba83030a2c5[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Wed Feb 17 11:26:01 2021 -0500

    project: switch to go 1.16. (#570)
    
    Go 1.16 is the latest stable release[0]. This commit switches the
    `go.mod` go version, as well as the version used in Github workflows for
    CI.
    
    [0]: https://golang.org/doc/go1.16

[33mcommit 1f157ab6c8ab2a591daefeaad74bd061b13b4d3f[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Mon Feb 15 11:00:41 2021 -0800

    Lint template produces a file with an `init` function that is not at the top of the new lint (#565)
    
    * update template
    
    * and v2 to v3
    
    Co-authored-by: Zakir Durumeric <zakird@gmail.com>

[33mcommit 835500b231bc6214ae534cd8e703ca54e8ead959[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Mon Feb 15 10:58:49 2021 -0800

    Custom static analysis tooling for CI/CD (#551)
    
    * adding custom linters to the code base for static analysis
    
    * renaming main test directory to make more consistent
    
    * Update v3/integration/lints/main.go
    
    Co-authored-by: Daniel McCarney <daniel@binaryparadox.net>
    
    * Update v3/integration/lints/lint/lint.go
    
    Co-authored-by: Daniel McCarney <daniel@binaryparadox.net>
    
    * Update v3/integration/lints/lint/lint.go
    
    Co-authored-by: Daniel McCarney <daniel@binaryparadox.net>
    
    * Update v3/integration/lints/lint/lint.go
    
    Co-authored-by: Daniel McCarney <daniel@binaryparadox.net>
    
    * Update v3/integration/lints/main.go
    
    * Update v3/integration/lints/filters/nodes.go
    
    Co-authored-by: Daniel McCarney <daniel@binaryparadox.net>
    
    * Update v3/integration/lints/filters/nodes.go
    
    Co-authored-by: Daniel McCarney <daniel@binaryparadox.net>
    
    * Update makefile
    
    * Update makefile
    
    * Update v3/integration/lints/lint/lint.go
    
    Co-authored-by: Daniel McCarney <daniel@binaryparadox.net>
    
    Co-authored-by: Daniel McCarney <daniel@binaryparadox.net>

[33mcommit 1cbdd0cc99a976580a3e7968e4abe079ee5c5fa3[m
Author: Rufus Buschart <rufus.buschart@siemens.com>
Date:   Fri Feb 12 01:43:15 2021 +0100

    docs: update CONTRIBUTING.md with cert generation resources (#560)

[33mcommit 59e0d7802dcc38059ed17e961800b74c06b97806[m
Author: github-actions[bot] <41898282+github-actions[bot]@users.noreply.github.com>
Date:   Thu Feb 11 08:53:50 2021 -0500

    util: gtld_map autopull updates for 2021-02-11T11:26:01 UTC (#563)
    
    Co-authored-by: GitHub <noreply@github.com>

[33mcommit f091dd34980fda129fadf2b4a4a0fc9104e11a85[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Mon Feb 8 17:44:59 2021 -0500

    deps: update zcrypto to 2a2d9c3 (#562)

[33mcommit 848c50b07013a773f2d98620400a6be900c88d2c[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Fri Jan 29 08:47:25 2021 -0500

    integration: fix resultCount types to not overflow, update expected vals. (#557)
    
    * integration: fix resultCount types to not overflow.
    
    Using `uint8` as the type for the count fields in the `resultCount` type
    produces overflows if more than 255 certificates with a given result
    level are linted.
    
    Our integration test corpus is just shy of 600,000 certificates so
    `uint32` should be more than sufficient.
    
    * integration: update expected values to correct overflows.
    
    Any lints that had more than 255 results at a given level will have
    overflowed, meaning expected counts were not correct.

[33mcommit 12bb0ed27a9859a48a91831fdb6a5f49bd9ef88b[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Thu Jan 28 20:09:24 2021 -0500

    lints: revert e_key_usage_and_extended_key_usage_inconsistent. (#556)
    
    The `e_key_usage_and_extended_key_usage_inconsistent` lint's
    interpretation of RFC 5280 is under question (see zmap#553).
    
    We also had an integration test bug that resulted in massively
    under-estimating it's impact on our integration test corpus (see zmap#555).
    
    Let's remove this lint while we sort out the correct logic.

[33mcommit c1c6681339ef2247dca5ff81162ab80c4811b154[m
Author: Mathew Hodson <mathew.hodson@gmail.com>
Date:   Wed Jan 27 19:54:49 2021 -0500

    lints: fix description of e_ext_ian_uri_not_ia5 (#554)

[33mcommit 2549ed3615faeef82f2c3a297da65f605260aea0[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Sun Jan 24 13:21:15 2021 -0500

    lints: return detail for e_ext_duplicate_extension. (#550)
    
    Previously the `e_ext_duplicate_extension` lint from the `lint.RFC5280`
    source only returned a `lint.Error` result as soon as one duplicate
    extension was found in a certificate. It did not indicate which
    extension OID was duplicated, or if there was more than one duplicated
    extensions.
    
    This commit reworks the lint to do both of these things. The detail
    string now indicates all of the extension OIDs that were present more
    than once.

[33mcommit 6dde095f6f92a1a58b0cb4abee9a4dfe8cde1bb7[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Sat Jan 23 14:16:32 2021 -0500

    deps: update zcrypto to 9cf5bea (#548)

[33mcommit 30943995122c332fa0c505d8e6f7e3802543a7c3[m
Author: Zsófia Tomicskó <72446712+ZsofiaTomicsko@users.noreply.github.com>
Date:   Wed Jan 20 18:02:32 2021 +0100

    tests: coverage for e_name_constraint_not_fqdn detail msgs (#547)
    
    This is a follow-up PR to #533 that adds test coverage for the details
    msg constructed for a non-pass result.

[33mcommit ea233116759a73b38900d929e89427cf8f3a9413[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Mon Jan 18 12:03:02 2021 -0500

    lints: move init to start of lint_name_constraint_not_fqdn.go (#544)
    
    This matches the work done in https://github.com/zmap/zlint/pull/536

[33mcommit 6d643b9bddd500d5585876859efa09403a1f0f42[m
Author: Rob <3725956+robplee@users.noreply.github.com>
Date:   Mon Jan 18 16:43:46 2021 +0000

    project: re-order lint init functions (#536)

[33mcommit edd0d0c0474304375e7a783784fdcde720dc7c13[m
Author: Zsófia Tomicskó <72446712+ZsofiaTomicsko@users.noreply.github.com>
Date:   Mon Jan 18 17:34:40 2021 +0100

    lints: adds `e_name_constraint_not_fqdn` lint (RFC5280 4.2.1.10) (#533)
    
    This commit adds a new `e_name_constraint_not_fqdn` lint to the
    lint.RFC5280 source that enforces RFC 5280, Section 4.2.1.10
    requirement that name constraints are fully qualified domain names.

[33mcommit 186e2c116bfb533babda67e1033ac6831f6dcf53[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sat Jan 9 13:20:44 2021 -0800

    project: update copyright year to 2021 (#543)

[33mcommit 5316fa5a79e7c3be7cf3a1032396fe52d24a51b0[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Sat Jan 9 12:51:16 2021 -0800

    lints: adds e_ev_organization_id_missing lint (CABF EVG 1.7.0 Section 9.8.2) (#532)
    
    This commit adds a new `e_ev_organization_id_missing` lint to the
    lint.CABFEVGuidelines source that enforces
    CA-Browser-Forum-EV-Guidelines-v1.7.0 Section 9.8.2 and the
    presence of the CabfExtensionOrganizationIdentifier:
    
    > Effective January 31, 2020, if the subject:organizationIdentifier
    > field is present, this field MUST be present.

[33mcommit b0e20c85df2f646252d3bf568fad98bb87166139[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Thu Jan 7 17:43:45 2021 -0500

    docs: CONTRIBUTING.md updates, couple copyright year tweaks. (#535)
    
    * docs: update CONTRIBUTING.md to fix typo & rec. test-driven subtests.
    * docs: update copyright year in template and README

[33mcommit 747b41fe8a541f5dd39e5a323a81729ccfe8629e[m
Author: Christopher Henderson <chris@chenderson.org>
Date:   Wed Jan 6 11:36:12 2021 -0800

    lints: fix boundary condition in `e_serial_number_longer_than_20_octets` lint (#527)
    
    Previously the `e_serial_number_longer_than_20_octets` lint would mistakenly
    pass certificates that had a DER encoded serial number that was exactly 21
    octets long. This case typically arises when a serial number is 20 octets long
    with an MSB of 1 since the encoded form will be prefixed with 0x00 to remain
    a positive DER encoded integer, thus bumping the encoded length to 21 octets.
    This commit fixes the calculation to correctly return an error finding
    for this class of certificates/encoded serial numbers.

[33mcommit 30424383782f921d502e8aa388e0cc9901195c3e[m
Author: Zsófia Tomicskó <72446712+ZsofiaTomicsko@users.noreply.github.com>
Date:   Sat Jan 2 20:45:07 2021 +0100

    KU and EKU Inconsistent lint correction (#528)
    
    * Added lint and tests for KU&EKU consistency check
    
    * Added errors to config.json
    
    * update to v3
    
    * removal of merge artifacts
    
    * corrected name of eku bits inside comments
    
    * no need for helper function anymore
    
    * replaced function with mapping
    
    * Changed comments
    
    * Added truth table tests
    
    * removed empty lines
    
    Co-authored-by: Rufus Buschart <rufus.buschart@siemens.com>
    Co-authored-by: Zakir Durumeric <zakird@gmail.com>

[33mcommit 4d0ac7ae1afc5f04e5111cb0b8eb7ebf536de550[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Sat Jan 2 14:31:26 2021 -0500

    deps: update zmap/zcrypto to 1eef276 (#529)

[33mcommit b691fe912db35cf719005041a5c24ba69096ce27[m
Author: Zsófia Tomicskó <72446712+ZsofiaTomicsko@users.noreply.github.com>
Date:   Wed Dec 23 22:24:01 2020 +0100

    Added a new lint and tests for correlation between KU&EKU (#497)
    
    * Added lint and tests for KU&EKU consistency check
    
    * Added errors to config.json
    
    * update to v3
    
    * removal of merge artifacts
    
    * corrected name of eku bits inside comments
    
    * no need for helper function anymore
    
    * replaced function with mapping
    
    Co-authored-by: Rufus Buschart <rufus.buschart@siemens.com>
    Co-authored-by: Zakir Durumeric <zakird@gmail.com>
    Co-authored-by: Christopher Henderson <chris@chenderson.org>

[33mcommit a1b837ad58380c74b25c4c92f856fe2ebc4dce90[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Mon Dec 21 20:51:54 2020 -0500

    deps: update zmap/zcrypto to deeac00. (#526)
    
    Most notably this includes an updated publicsuffix-go dependency with
    fresher PSL data.

[33mcommit 9e16bfc62c32ffa8f29d6ad671780334a31901f5[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Mon Dec 21 12:30:42 2020 -0500

    util: remove unused `ICANNPublicSuffixParse` helper. (#525)
    
    The `util.ICANNPublicSuffixParse` helper function is the only place in
    ZLint that directly impots `publicsuffix-go`, and it isn't called
    anywhere.
    
    By removing this unused helper we can remove the direct dependency on
    `weppos/publicsuffix-go` from ZLint. It remains a transitive dependency
    via ZCrypto.

[33mcommit f47c9d6ebeb7e2ea62b63c34b39fc03638f1d36c[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Mon Dec 21 12:12:50 2020 -0500

    CI: Cleanup hacky tld-update workflow env var use. (#524)
    
    The `tld-update.yml` workflow needs the current date in a human readable
    form to include in commit messages, PR titles, and the PR body.
    Previously this was achieved in a clunky way with an env var echoed into
    `$GITHUB_ENV`.
    
    This commit replaces that logic with a cleaner way to achieve the same
    thing: setting the date string as the output of a step and then
    referencing it later in the workflow.

[33mcommit d8314a3cab4a3b971016681608f217b9a2655126[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Sat Dec 12 20:30:51 2020 -0500

    CI: Have tld-update workflow build & test pre-PR. (#521)
    
    Once the `tld-update.yml` github workflow job has run `go generate
    ./...` and potentially created a local diff with updated gTLD data it
    should also build & test. This helps ensure that if something goes
    haywire and the local diff won't build/run it won't have a PR opened.

[33mcommit 83f15ca06fc8224d088fd2ffccb8e89db9735045[m
Author: github-actions[bot] <41898282+github-actions[bot]@users.noreply.github.com>
Date:   Thu Dec 10 21:34:14 2020 -0800

    util: gtld_map autopull updates for 2020-12-11T05:27:56 UTC (#520)
    
    Co-authored-by: GitHub <noreply@github.com>

[33mcommit b6e5ba7064ae73cf9aa561b8dba59bd171fda062[m
Author: github-actions[bot] <41898282+github-actions[bot]@users.noreply.github.com>
Date:   Tue Dec 8 13:35:51 2020 -0500

    util: gtld_map autopull updates for 2020-12-08T18:31:14 UTC (#518)
    
    Co-authored-by: GitHub <noreply@github.com>

[33mcommit 1eb11ce3552704904404d56a65f4be3511828dbf[m
Author: Rufus Buschart <rufus.buschart@siemens.com>
Date:   Mon Dec 7 17:33:25 2020 +0100

    Ocsp eku check for tls certificates (#490)
    
    * included error overview
    
    * fix: see https://github.com/zmap/zlint/pull/488#discussion_r508524594
    
    * add a check for OCSP responder certs
    
    * add lint, test and update integration config
    
    * added test data
    
    * fixes for v3
    
    Co-authored-by: Daniel McCarney <daniel@binaryparadox.net>

[33mcommit 662504d527366b54ce9dd58a9fb96bec1b6c8f98[m
Author: Zakir Durumeric <zakird@gmail.com>
Date:   Mon Nov 30 16:15:17 2020 -0800

    change tld updator to not be me (#516)

[33mcommit 931c5d4d22f946f33d19a64e7d9c3668196abef1[m
Author: github-actions[bot] <41898282+github-actions[bot]@users.noreply.github.com>
Date:   Mon Nov 30 18:47:34 2020 -0500

    util: gtld_map autopull updates for 2020-11-30T23:23:57 UTC (#514)
    
    Co-authored-by: zakird <zakird@users.noreply.github.com>

[33mcommit 12dfc1846136a481fa889340d591701a3dfaa516[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Mon Nov 30 18:05:31 2020 -0500

    CI: Add cron workflow for gTLD update PRs. (#513)
    
    Previously we used a separate repo with a bash script[0] hooked up to
    a bot github user[1] in a Travis CI cron build to automatically create
    PRs updating ZLint TLD data on a periodic basis.
    
    Now that we're using Github Actions we can make things much simpler and
    self-contained.
    
    This commit adds a `tld-update.yml` workflow that uses
    a create-pull-request Github action to replace the separate repo/bash
    script/bot user approach.
    
    Not only does this let us delete the bot user's write access to the
    ZLint repo but it's also a smarter integration overall and won't
    recreate the same PR over and over if it isn't merged right away. Lastly
    the Github Actions cron schedule is more flexible so we can run the new
    action once every hour instead of just once a day like the Travis
    version.
    
    [0]: https://github.com/cpu/zlint-autopull/blob/master/autopull
    [1]: https://github.com/tld-update-bot
    [2]: https://github.com/peter-evans/create-pull-request

[33mcommit fe65bae28cbc112082e21449a373f7ccf059422f[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Sun Nov 29 15:34:33 2020 -0500

    project: bump major version to 3.0.0 (#510)
    
    * project: move v2 subdir to v3
    * project: update references from v2 to v3
    * project: bump module version

[33mcommit 0d48ea15507cb16a08f694b8d9b64f91b601d3e7[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Sun Nov 29 14:59:18 2020 -0500

    lint: combine ZLint and AWSLabs Sources into Community. (#509)
    
    This commit is a breaking API change that combines the `lint.ZLint` and
    `lint.AWSLabs` `lint.LintSource`'s into one: `lint.Community`. This
    better matches the directory structure we store the lints under and is
    more indicative of the fact that we don't intend to have perfect
    matching coverage of all awslabs certlint lints.

[33mcommit 8dc66d05deb6ef92eb21c815e151a5c805730018[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Sun Nov 29 13:49:50 2020 -0500

    Update to Go 1.15, latest , fix n_san_iana_pub_suffix_empty. (#508)
    
    * CI: Switch from Travis to Github Actions.
    
    This commit replaces the existing Travis CI integration with Github
    Actions equivalents.
    
    Notable changes:
    
    * Golanci-lint had to be updated to a newer version to support the
    action. I fixed a few of the new linter findings and added ignores for
    others.
    
    * Integration tests - The Github Action cache support is much more
    generous (5GB!). This lets us cache the integration test corpus data
    between runs saving ~1-2m each run.
    
    * chore: update to Go 1.15
    
    This switches CI and `go.mod` to use the latest major Go release stream,
    1.15.x.
    
    * lints: drop one ETSI test case.
    
    Updates in Go 1.15 mean that an error test case that was used by
    `v2/lints/etsi/lint_qcstatem_qccompliance_valid_test.go` now panics
    during parse.
    
    Since we expect the standard library to throw an error for this case now
    instead of the ZLint lint we can remove the test case.
    
    * deps: update ZCrypto to latest.
    
    * lints: rewrite n_san_iana_pub_suffix_empty.
    
    Reworked `n_san_iana_pub_suffix_empty`:
    
    While looking at this lint I noticed it has some weird behaviour:
    
    * It could only flag at most one info result per cert.
    * It didn't include the name that tripped the result in the details.
    * It would return NA for a whole cert if one name failed to parse but
    with a different error than the one used to determine it's a suffix.
    
    I've updated the implementation to address all of the above. It collects
    up a list of bad names per-cert, includes them in a details message, and
    doesn't return NA for parse errors unrelated to suffixes. The unit tests
    were also not table driven so I rewrote those quickly and included
    a test case based on my integration test update and a test case with
    multiple bare public suffixes.
    
    Updated integration test data:
    
    Here's the story about the expected integration test data change for
    `n_san_iana_pub_suffix_empty` with lots of detail on the debugging process.
    
    First to find out what certificate changed from an info result to a pass result
    I ran the integration tests with `-fingerprintSummary` for just that one lint
    with both the ZCrypto version in this branch and the one on master.
    
    ```
    make integration INT_FLAGS="-fingerprintSummary -lintFilter='n_san_iana_pub_suffix_empty'"
    ```
    
    I cut out the fingerprint summary and diff'd the output between versions to spot
    the one certificate FP that disappeared with the updated ZCrypto. It was
    `d570517b96eb7e3db7c6986f421e988fdae8f417295baade0dfc9e6edf8d12cc`.
    
    Next I ran the `certByFP.sh` script to pull out that cert from the corpus and
    get a link to it on Censys.
    
    ```
    ./integration/certByFP.sh d570517b96eb7e3db7c6986f421e988fdae8f417295baade0dfc9e6edf8d12cc
    <snipped>
    https://censys.io/certificates/d570517b96eb7e3db7c6986f421e988fdae8f417295baade0dfc9e6edf8d12cc
    ```
    
    Then I ran the lint with the old ZCrypto in dlv and set a breakpoint in the
    lint.
    
    ```
    dlv debug ./cmd/zlint -- -includeNames='n_san_iana_pub_suffix_empty' -pretty ~/Downloads/UnsortedChrome/d570517b96eb7e3db7c6986f421e988fdae8f417295baade0dfc9e6edf8d12cc.pem
    break pubsuffix.Execute
    ```
    
    I stepped through until there was a parse error and printed it:
    
    ```
    (dlv) p parsedName
    github.com/zmap/zcrypto/x509.ParsedDomainName {
            DomainString: "www.theaterpreise.ch\n\n",
            ParsedDomain: *github.com/weppos/publicsuffix-go/publicsuffix.DomainName nil,
            ParseError: error(*errors.errorString) *{
                    s: "www.theaterpreise.ch\n\n is a suffix",},}
    ```
    
    A borked DNS SAN with two trailing newlines was incorrectly deemed a suffix by publicsuffix-go.
    
    Repeating the process with the updated ZCrypto/publicsuffix-go:
    
    ```
    (dlv) p parsedName
    github.com/zmap/zcrypto/x509.ParsedDomainName {
            DomainString: "www.theaterpreise.ch\n\n",
            ParsedDomain: *github.com/weppos/publicsuffix-go/publicsuffix.DomainName {
                    TLD: "ch\n\n",
                    SLD: "theaterpreise",
                    TRD: "www",
                    Rule: *(*"github.com/weppos/publicsuffix-go/publicsuffix.Rule")(0xc000108e40),},
            ParseError: error nil,}
    ```
    
    No more parse error, just a whacky TLD with two newlines. This shouldn't return
    an info result for `n_san_iana_pub_suffix_empty` and now with updated deps it
    doesn't. TL;DR - old behaviour was a bug we accidentally fixed.

[33mcommit da00f3f052ac0c0b9362486d89908755d18c748f[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Sat Nov 28 17:01:32 2020 -0500

    CI: Switch from Travis to Github Actions. (#505)
    
    This commit replaces the existing Travis CI integration with Github
    Actions equivalents.
    
    Notable changes:
    
    * Golanci-lint had to be updated to a newer version to support the
    action. I fixed a few of the new linter findings and added ignores for
    others.
    
    * Integration tests - The Github Action cache support is much more
    generous (5GB!). This lets us cache the integration test corpus data
    between runs saving ~1-2m each run.

[33mcommit 7f7ef1f90617b717b1e0c80ff5c8ca588c6f84c0[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Sat Nov 28 16:47:02 2020 -0500

    lints: split Apple cert lifetime lint per-result. (#506)
    
    Previously the `e_tls_server_cert_valid_time_longer_than_398_days` lint
    could return either an error result or a warning result.  We prefer
    having lints return one status level only.
    
    This commit breaks the lint up so that
    `e_tls_server_cert_valid_time_longer_than_398_days` only handles the
    error case and a new `w_tls_server_cert_valid_time_longer_than_397_days`
    lint handles the warning case.

[33mcommit c42a35826efbac5643ab1b93709380efc6b08a5f[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Mon Nov 23 15:01:34 2020 -0500

    lint: rename Source AppleCTPolicy -> AppleRootProgramPolicy (#501)

[33mcommit 71e2966bbbbd8a3ae6e7ff5f07502f6f59926757[m
Author: TLD Update Robot <47792085+tld-update-bot@users.noreply.github.com>
Date:   Sun Nov 22 17:36:53 2020 -0500

    gTLD autopull: 2020-11-21T16:05:09Z (#498)
    
    Co-authored-by: tld-update-bot <cpu+tldbot@letsencrypt.org>

[33mcommit 29b3fa90013c35c61f635f58f0d306a23a169a86[m
Author: Zakir Durumeric <zakird@gmail.com>
Date:   Tue Nov 10 13:22:33 2020 -0800

    Update Contributing Guidelines (#495)
    
    * contributing guidelines
    
    * Update CONTRIBUTING.md
    
    * Update CONTRIBUTING.md
    
    * Update CONTRIBUTING.md
    
    * Update CONTRIBUTING.md
    
    Co-authored-by: Daniel McCarney <daniel@binaryparadox.net>
    
    * Update CONTRIBUTING.md
    
    * Update CONTRIBUTING.md
    
    Co-authored-by: Daniel McCarney <daniel@binaryparadox.net>

[33mcommit e2b36583cd24a86a63ac4d7bcd3da204b42428b8[m
Author: BJ Cardon <cardonator@users.noreply.github.com>
Date:   Tue Nov 10 13:08:45 2020 -0700

    make two lints notice instead of warn,  (#493)
    
    * make two lints notice instead of warn, closes #492
    
    * update lint results in tests
    
    Co-authored-by: Zakir Durumeric <zakird@gmail.com>

[33mcommit 7b54a38ec9d5701d3d65294605c0eaa6299bf213[m
Author: Rufus Buschart <rufus.buschart@siemens.com>
Date:   Mon Nov 9 21:55:08 2020 +0100

    Improve readability of "EKU" abbreviation (#489)
    
    * included error overview
    
    * fix: see https://github.com/zmap/zlint/pull/488#discussion_r508524594
    
    * improve readability

[33mcommit f46d09c4ec0f3e1036731915a4d61938004b2d76[m
Author: Rufus Buschart <rufus.buschart@siemens.com>
Date:   Tue Oct 20 23:44:45 2020 +0200

    tests: include error/warning/info overview for integration test failures (#488)
    
    This commit introduces an overview of how many lints fail when the complete
    integration test is being executed. This is information can be helpful when making
    a change that affected multiple lints.

[33mcommit cca4a6b67f9e32a0659b60f44378660c7cd2a1f5[m
Author: TLD Update Robot <47792085+tld-update-bot@users.noreply.github.com>
Date:   Mon Oct 19 12:25:28 2020 -0400

    gTLD autopull: 2020-10-19T15:48:38Z (#487)
    
    Co-authored-by: tld-update-bot <cpu+tldbot@letsencrypt.org>

[33mcommit def029d0380875e92d4dbb518875a195fad76509[m
Author: Rufus Buschart <rufus.buschart@siemens.com>
Date:   Fri Oct 9 15:40:23 2020 +0200

    misc: gitignore Visual Studio Code configuration files (#485)

[33mcommit 1fd478276e8f96630ed888780cbbf1ac58e5a97d[m
Author: Rufus Buschart <rufus.buschart@siemens.com>
Date:   Fri Oct 9 15:39:26 2020 +0200

    README: Correction of link to Siemens PKI (#486)
    
    Correction of the generic link to Siemens into the link to the Siemens PKI web site

[33mcommit 5ed7e1316baa5eb001b0222645c90b76c140b195[m
Author: TLD Update Robot <47792085+tld-update-bot@users.noreply.github.com>
Date:   Thu Oct 8 14:03:39 2020 -0400

    gTLD autopull: 2020-10-08T15:44:26Z (#484)
    
    Co-authored-by: tld-update-bot <cpu+tldbot@letsencrypt.org>

[33mcommit 6b73243356213d5ff575a547a6131c4d6cc05646[m
Author: Hugo Stijns <hugo@boosboos.net>
Date:   Tue Oct 6 19:08:57 2020 +0200

    deps: bump golang.org/x/text to 0.3.3 to fix CVE-2020-14040 (#481)
    
    Package golang.org/x/text has a vulnerability which is fixed in 0.3.3.
    
    See: https://nvd.nist.gov/vuln/detail/CVE-2020-14040

[33mcommit f7543c75b181fe680a544549ad5dacf151060966[m
Author: James Kasten <jdkasten@umich.edu>
Date:   Fri Sep 25 16:55:16 2020 -0400

    Improve error message of ReadTestCert panic (#478)
    
    The error from x509.ParseCertificate was not being included within the panic.Debugging is easier if this information is retained.
    
    Co-authored-by: Zakir Durumeric <zakird@gmail.com>

[33mcommit c16b5bd64d638ed745f1be545bf3e2ee16989eb6[m
Author: Simon Edänge <70440015+OathMeadow@users.noreply.github.com>
Date:   Thu Sep 24 21:26:09 2020 +0200

    README: Add Nexus CM to list of users/integrations (#477)
    
    * README: Add Nexus CM to list of users/integrations
    
    * Update README.md
    
    Co-authored-by: Simon Edänge <simon.edange@nexusgroup.com>
    Co-authored-by: Zakir Durumeric <zakird@gmail.com>

[33mcommit aa4e2619db00dbdea768e9afa70262f0d6af3417[m
Author: TLD Update Robot <47792085+tld-update-bot@users.noreply.github.com>
Date:   Tue Sep 8 11:55:19 2020 -0400

    autopull: 2020-09-08T15:28:12Z (#470)
    
    Co-authored-by: tld-update-bot <cpu+tldbot@letsencrypt.org>

[33mcommit 2b994a741bfde638a648a759c16644724d07e708[m
Author: Corey Bonnell <corey.j.bonnell@outlook.com>
Date:   Mon Sep 7 20:37:09 2020 -0400

    Align Validity Period definition with RFC 5280 (#469)
    
    Co-authored-by: Corey Bonnell <cbonnell@outlook.com>

[33mcommit f20a717a628edc3181b9802b0879951808e02803[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Wed Sep 2 20:15:53 2020 -0400

    CONTRIBUTING: Add notes on publishing a release. (#468)
    
    The `CONTRIBUTING.md` docs now describe the ZLint release process. The
    steps involved are roughly based on the process I've been following and
    should be considered a starting point, not an immutable set of laws.

[33mcommit e1a9412ec5b778fd39de5a475b928da88cd68433[m
Author: Aaron Gable <aaron@aarongable.com>
Date:   Tue Sep 1 10:28:24 2020 -0700

    Add citation for sub-CAs to ca_digital_signature_not_set (#464)

[33mcommit 01996c6fbb8372cdae2f4b2a82b121e69b789132[m
Merge: 4666bb74 9ab0643d
Author: mtg <git@mtg.de>
Date:   Wed Aug 26 08:56:33 2020 +0200

    Merge https://github.com/zmap/zlint

[33mcommit 9ab0643df8f6ad6bac722e72851a0fd3ac7f350c[m
Author: Jacob Hoffman-Andrews <github@hoffman-andrews.com>
Date:   Thu Aug 20 19:31:25 2020 -0700

    Ballot SC31 makes OCSP optional for intermediate certificates. (#463)
    
    Do not merge until Ballot SC31 successfully passes its review period.
    
    Fixes #462.

[33mcommit 3f689d276c06aad066b36967c4af3f75f6822247[m
Author: Zakir Durumeric <zakird@gmail.com>
Date:   Tue Aug 4 12:32:17 2020 -0700

    README to suggest checking x509.ParseCertificate error (#460)
    
    * example to check output
    
    * Spacing fix
    
    * Update README.md
    
    * filter error

[33mcommit ada09919b3bc00cbe01accb9fcf8f3e423892b6a[m
Author: TLD Update Robot <47792085+tld-update-bot@users.noreply.github.com>
Date:   Wed Jul 29 11:45:04 2020 -0400

    autopull: 2020-07-29T15:10:15Z (#459)
    
    Co-authored-by: tld-update-bot <cpu+tldbot@letsencrypt.org>

[33mcommit 6d02ef7694200df4173dc7e6f5b2e5dfdcbd8ebf[m
Author: BJ Cardon <cardonator@users.noreply.github.com>
Date:   Thu Jul 23 06:59:07 2020 -0600

    tests: add NA test case for e_tls_server_cert_valid_time_longer_than_398_days (#457)

[33mcommit 34310bdb6e20040c6ff901468f554b14c2f7b63f[m
Author: BJ Cardon <cardonator@users.noreply.github.com>
Date:   Tue Jul 21 19:14:49 2020 -0600

    this lint shouldn't apply to CA certs (#456)

[33mcommit ca9532d1d3e99a3c9dd2b6fba3f70ef699e1afc9[m
Author: Andrew Caird <acaird@gmail.com>
Date:   Mon Jul 20 17:56:59 2020 -0400

    Create options for human-readable output formats (#437)
    
    * Add a -summary option to print a short summary of the linting
    
    Linting the test file `testdata/utf8ControlX88.pem` results in:
    ```
    +-------+--------------+
    | LEVEL | # OCCURANCES |
    +-------+--------------+
    | info  |            0 |
    | warn  |            7 |
    | error |           15 |
    | fatal |            0 |
    +-------+--------------+
    ```
    
    * Added -longsummary option and output
    
    Running:
    ```sh
    testdata/indivValAllBad.pem | ./zlint -longsummary
    ```
    the output is:
    ```
    +-------+--------------+------------------------------------------+
    | LEVEL | # OCCURANCES |                 DETAILS                  |
    +-------+--------------+------------------------------------------+
    | info  |            0 |  -                                       |
    | warn  |            1 | w_ext_san_critical_with_subject_dn       |
    | error |            7 | e_ca_crl_sign_not_set                    |
    |       |              | e_sub_ca_crl_distribution_points_missing |
    |       |              | e_ca_country_name_missing                |
    |       |              | e_cert_policy_iv_requires_country        |
    |       |              | e_sub_cert_not_is_ca                     |
    |       |              | e_ca_key_cert_sign_not_set               |
    |       |              | e_ca_organization_name_missing           |
    | fatal |            0 |  -                                       |
    +-------+--------------+------------------------------------------+
    ```
    
    * spelling fix
    
    * Remove tablewriter dependency and reimplement the good parts
    
    * spelling fix
    
    * Fixed a missed merge :(
    
    * switched longsummary to longSummary; fixed output bug
    
     - switched `-longsummary` option to `-longSummary` to be more consistent with
       existing options
    
     - fixed an embarrassing output bug when two categories had the same number of
       errors
    
    * Cleaned up typos, variable names, formatting
    
    * parent 99579098a16c10d1d704b8b8149dd8c35329107f
    author Andrew Caird <acaird@gmail.com> 1590420366 -0400
    committer Andrew Caird <acaird@gmail.com> 1593372751 -0400
    
    Add a -summary option to print a short summary of the linting
    
    Linting the test file `testdata/utf8ControlX88.pem` results in:
    ```
    +-------+--------------+
    | LEVEL | # OCCURANCES |
    +-------+--------------+
    | info  |            0 |
    | warn  |            7 |
    | error |           15 |
    | fatal |            0 |
    +-------+--------------+
    ```
    
    and a  -longSummary option and output
    
    Running:
    ```sh
    testdata/indivValAllBad.pem | ./zlint -longsummary
    ```
    the output is:
    ```
    +-------+--------------+------------------------------------------+
    | LEVEL | # OCCURANCES |                 DETAILS                  |
    +-------+--------------+------------------------------------------+
    | info  |            0 |  -                                       |
    | warn  |            1 | w_ext_san_critical_with_subject_dn       |
    | error |            7 | e_ca_crl_sign_not_set                    |
    |       |              | e_sub_ca_crl_distribution_points_missing |
    |       |              | e_ca_country_name_missing                |
    |       |              | e_cert_policy_iv_requires_country        |
    |       |              | e_sub_cert_not_is_ca                     |
    |       |              | e_ca_key_cert_sign_not_set               |
    |       |              | e_ca_organization_name_missing           |
    | fatal |            0 |  -                                       |
    +-------+--------------+------------------------------------------+
    ```
    
    * autopull: 2020-05-27T14:34:02Z (#441)
    
    Co-authored-by: tld-update-bot <cpu+tldbot@letsencrypt.org>
    
    * gTLD autopull: 2020-05-28T14:35:00Z (#442)
    
    Co-authored-by: tld-update-bot <cpu+tldbot@letsencrypt.org>
    
    * Moved structure creation out of function into a method for reporting
    
    * Moved the formatted output routines out of main
    
    * Changed newRT to a pointer receiver
    
    * Changed output options to all them all; newlines for nice output
    
    * Changed output options to allow printing of them all; newlines for nice output
    
    Co-authored-by: Zakir Durumeric <zakird@gmail.com>
    Co-authored-by: Daniel McCarney <daniel@binaryparadox.net>
    Co-authored-by: TLD Update Robot <47792085+tld-update-bot@users.noreply.github.com>
    Co-authored-by: tld-update-bot <cpu+tldbot@letsencrypt.org>

[33mcommit 5f05d1d1ce935f63136a071a7cdb103ba0a6e235[m
Author: TLD Update Robot <47792085+tld-update-bot@users.noreply.github.com>
Date:   Sat Jul 18 11:35:26 2020 -0400

    gTLD autopull: 2020-07-18T15:05:07Z (#455)
    
    Co-authored-by: tld-update-bot <cpu+tldbot@letsencrypt.org>

[33mcommit a9b00321fcd4fddcbec3af6e7478b67dabfb7904[m
Author: TLD Update Robot <47792085+tld-update-bot@users.noreply.github.com>
Date:   Sat Jun 27 14:37:17 2020 -0400

    gTLD autopull: 2020-06-27T14:52:30Z (#452)
    
    Co-authored-by: tld-update-bot <cpu+tldbot@letsencrypt.org>

[33mcommit f530e42915b345f3cff5333ec161221e1c3a3615[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Thu Jun 25 19:17:02 2020 -0400

    docs: add Entrust Datacard to README ZLInt users. (#451)
    
    Per Bugzilla[0] Entrust reports using ZLint:
    
    > we are using both pre-issuance linting and post-issuance linting using
    > zlint.
    
    [0]: https://bugzilla.mozilla.org/show_bug.cgi?id=1648472

[33mcommit d4acbba05c50b44ae34d60d9e1b0fd3b34a8619f[m
Author: sleevi <rsleevi@chromium.org>
Date:   Fri Jun 12 17:41:17 2020 -0400

    lints: cabf_br lint to verify .onion addresses are well-formed (#450)
    
    Adds a new lint, identified as `e_san_dns_name_onion_invalid`,
    that makes sure that the `.onion` addresses present within a
    certificate are well-formed v2 or v3 addresses, according to
    the v2 or v3 Rendezvous specifications.
    
    Closes #440

[33mcommit 84a8a2047667f1d6bda3368204ff1e258f1b8b6e[m
Author: sleevi <rsleevi@chromium.org>
Date:   Wed Jun 10 12:24:48 2020 -0400

    Fix .onion tests to only apply to EV certificates (#449)
    
    Before this change, ZLint would reject .onion names in non-EV certs
    via the `lint_san_dns_name_onion_not_ev_cert` lint, and if that
    was suppressed, then complain about the missing Tor Service
    Descriptor extension. As of CA/Browser Forum Ballot SC27, it's
    allowed for v3 onion names to appear in DV/OV/IV certificates, and
    the Tor Service Descriptor extension is neither required nor
    prohibited for these.
    
    This change corrects the Tor Service Descriptor tests to properly
    account for it being mandatory for EV, while optional for DV/OV/IV.
    This does not introduce new lints to ensure that the address is
    itself a well-formed V2 (if EV) or V3 (all types) address, which
    will come in a follow-up change.
    
    Closes #440

[33mcommit ecf8678e38d1e50a2e92cb3cd35eebe12f45cc11[m
Author: sleevi <rsleevi@chromium.org>
Date:   Thu Jun 4 21:47:15 2020 -0400

    Move EV-specific tests to cabf_ev (#445)

[33mcommit c820d9566a0e295070ebe3954b5fe521862d6e34[m
Author: sleevi <rsleevi@chromium.org>
Date:   Thu Jun 4 20:58:56 2020 -0400

    Fix the EV validity check (#447)
    
    The lint lint_ev_valid_time_too_long has several issues:
    * It set the maximum validity as 825-days, rather than 27 months
      (which is 366 + 365 + 31 + 31 + 30 = 823 days) for certs issued
      before the 825-day change
    * It set the source of the requirements to the BRs, rather than
      the EVGs
    
    Co-authored-by: Zakir Durumeric <zakird@gmail.com>

[33mcommit 37a03da912510fb5b4b89a07f12d1f715c247812[m
Author: sleevi <ryan.sleevi@gmail.com>
Date:   Thu Jun 4 08:14:26 2020 -0400

    docs: correct link to integration test documentation (#446)

[33mcommit ce1631b8b1c6fce0348724d37a0d0b4ad62c73fa[m
Author: TLD Update Robot <47792085+tld-update-bot@users.noreply.github.com>
Date:   Wed Jun 3 10:55:58 2020 -0400

    autopull: 2020-06-03T14:39:17Z (#444)
    
    Co-authored-by: tld-update-bot <cpu+tldbot@letsencrypt.org>

[33mcommit de9eafbe63b6e03da98c681fd1b916da97779f1f[m
Author: Roland Bracewell Shoemaker <rolandshoemaker@gmail.com>
Date:   Mon Jun 1 13:19:04 2020 -0700

    Check tbsCertificate signature algorithm matches certificate (#436)
    
    * Check tbsCertificate signature algorithm matches certificate
    
    Per RFC 5280 section 4.1.1.2, couldn't find an existing lint.
    
    * Add ignore to integration + reuse tbsCert
    
    Co-authored-by: Daniel McCarney <daniel@binaryparadox.net>

[33mcommit 82e1f43dcbc2e6ef9aa89194c0a828487f218e9f[m
Author: TLD Update Robot <47792085+tld-update-bot@users.noreply.github.com>
Date:   Thu May 28 14:51:09 2020 -0400

    gTLD autopull: 2020-05-28T14:35:00Z (#442)
    
    Co-authored-by: tld-update-bot <cpu+tldbot@letsencrypt.org>

[33mcommit da06a3a1d029fe8b8bd6b50d7ce20ac3da0e7664[m
Author: TLD Update Robot <47792085+tld-update-bot@users.noreply.github.com>
Date:   Wed May 27 11:03:53 2020 -0400

    autopull: 2020-05-27T14:34:02Z (#441)
    
    Co-authored-by: tld-update-bot <cpu+tldbot@letsencrypt.org>

[33mcommit 99579098a16c10d1d704b8b8149dd8c35329107f[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Thu May 14 12:27:47 2020 -0400

    Deps: Update ZCrypto, fix assoc. test breakage. (#435)
    
    * deps: update zcrypto to tip of master.
    
    This pulls in ZCrypto at zmap/zcrypto@16679db
    
    * lints: remove invalid TestEtsiTypeAsQcStmt test case.
    
    The `testdata/QcStmtEtsiTaggedValueCert20.pem` test certificate has an
    invalid QCStatements extension value[0] and ZCrypto with support for
    parsing QC Statements panics reading the test cert.
    
    Since ZLint can't lint certificates that ZCrypto won't parse we must
    remove this test case.
    
    [0]: https://github.com/zmap/zlint/issues/433#issuecomment-628698981
    
    * lints: rm invalid `TestEtsiQcCompliance`, `TestEtsiQcType` certs.
    
    Similar to the prev. commit now that ZCrypto understands QCStatement
    extensions it will error parsing these test cases and so they must be
    removed. This test coverage should be handled by ZCrypto.
    
    * integration: updates for QCStatement lint expected results.
    
    With an updated ZCrypto there is now 1 certificate[0] from the integration
    test data that no longer parses. This in turn means that the
    `e_qcstatem_qctype_valid`, `n_subject_common_name_included`,
    `w_qcstatem_qcpds_lang_case` and `w_qcstatem_qctype_web` lints that
    previously had findings for this certificate need to have their expected
    result counts adjusted.
    
    [0]: https://censys.io/certificates/4712f1b2a94994b55626ecba2104bbf23d39c05e7a2751e5af8a923bac23fd8f

[33mcommit a42b7782cdb2bff02d6a18691a34611580f4b7fa[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Wed May 13 17:00:59 2020 -0400

    ci: remove vendor dir, Go 1.13.x -> 1.14.x, fix integration test data (#432)
    
    * chore: remove vendor dir.
    
    Vendoring has lost favour compared to relying on the Go 1.13+
    proxy/module checksum behaviour[0].
    
    [0]: https://proxy.golang.org/
    
    * ci: go 1.13.x -> go1.14.x
    
    Also remove setting GO111MODULE and GOFLAGS. The former is already the
    default since Go 1.12.x and the latter isn't required because we removed
    the vendor dir.
    
    * ci: update expected integration test data for Go 1.14.
    
    New lints without result cases in our integration test data are added
    with the expected set `{}`.
    
    Four existing lints have their expected error result tallies updated:
    
    1. "e_sub_cert_locality_name_must_not_appear"
      old: fatals: 0    errs: 23   warns: 0    infos: 0
      new: fatals: 0    errs: 13   warns: 0    infos: 0
    
    2. "e_sub_cert_province_must_not_appear"
      old: fatals: 0    errs: 16   warns: 0    infos: 0
      new: fatals: 0    errs: 8    warns: 0    infos: 0
    
    3. "e_sub_cert_street_address_should_not_exist"
      old: fatals: 0    errs: 8    warns: 0    infos: 0
      new: fatals: 0    errs: 0    warns: 0    infos: 0
    
    4. "e_sub_cert_postal_code_must_not_appear"
      old: fatals: 0    errs: 8    warns: 0    infos: 0
      new: fatals: 0    errs: 0    warns: 0    infos: 0
    
    These four lints previously returned an error result for certificates
    that had a Subject Organization/GivenName/Surname that were encoded as
    a BMPString. Go < 1.14.x's ASN.1 package did not support this encoding
    type and so the lints assumed the field was absent, resulting in a false
    positive. In Go 1.14.x+ the field is correctly decoded and the error
    result is no longer applicable.

[33mcommit bb6c7a74f1901fecf545c51af97e9e874a0260d9[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Wed May 13 12:46:42 2020 -0400

    docs: add ZLint announcements mailing list to README (#431)
    
    Co-authored-by: Zakir Durumeric <zakird@gmail.com>

[33mcommit ee0c915cbde37dd2cbea362a30bd9bf1dfe53819[m
Author: Zakir Durumeric <zakird@gmail.com>
Date:   Tue May 12 13:24:17 2020 -0500

    Adding mailing list link to README.

[33mcommit 1e160b10bc75c589461aa5036196f07f8f38fdcb[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Mon May 11 10:50:57 2020 -0400

    ci: update goreleaser install URL. (#429)
    
    The upstream project README[0] is using a different URL now and in my
    tests the old URL has an HTTPS subject common name mismatch preventing
    installation from succeeding and breaking the `make code-lint` phase of
    integration tests.
    
    [0]: https://github.com/golangci/golangci-lint#install

[33mcommit 3bf4bbf127fa32da92ab89abc8554d4e7d46defb[m
Author: MTG <36234449+mtgag@users.noreply.github.com>
Date:   Mon May 11 16:02:29 2020 +0200

    lints: enforce Mozilla PKI policy for ECDSA pubkey/sig alg curves/encoding. (#378)
    
    `e_mp_ecdsa_pub_key_encoding_correct`, enforces certificate ECDSA public key
    algorithm identifiers are a byte-for-byte match to the required values from
    Section 5.1.2 of the Mozilla root store policy or a `lint.Error` level finding
    is returned. The `e_mp_ecdsa_signature_encoding_correct` lint applies similar
    checks to certificate ECDSA signature algorithm identifiers. Both lints
    require that the ECDSA curve in use be one of P-256 or P-384, per Moz.
    policy.
    
    To help implement the new lints (and to simplify one existing lint), a new
    utility function `util.GetPublicKeyAidEncoded` is added. This function returns
    the encoded tag/length ASN.1 bytes of a certificate's `SubjectPublicKeyInfo`
    sequence's algorithm field (or an error if the field can not be extracted).
    
    Resolves #355, #358

[33mcommit 206df7d26e1ca081025a4314c98255c69404e7c7[m
Author: TLD Update Robot <47792085+tld-update-bot@users.noreply.github.com>
Date:   Thu Apr 2 15:03:03 2020 -0400

    gTLD autopull: 2020-04-02T17:35:25Z (#425)
    
    Co-authored-by: tld-update-bot <cpu+tldbot@letsencrypt.org>

[33mcommit d933f03c8465fd904602c64a7ae82abbc86e8833[m
Author: TLD Update Robot <47792085+tld-update-bot@users.noreply.github.com>
Date:   Sat Mar 28 13:48:42 2020 -0400

    autopull: 2020-03-28T17:34:11Z (#423)
    
    Co-authored-by: tld-update-bot <cpu+tldbot@letsencrypt.org>

[33mcommit 4ca06954c1c903a63a217ad21930947c648920a9[m
Author: John Wood <j@jdtw.us>
Date:   Mon Mar 23 12:56:49 2020 -0700

    Fix spelling of 'distinguished' in lint descriptions (#422)

[33mcommit 94d7dde424d0e8ebaee5b1abc57da7979bda55f5[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Tue Mar 17 13:28:05 2020 -0400

    util: rewrite test/prepend_testcerts_openssl.sh, update testdata (#421)
    
    * util: rewrite test/prepend_testcerts_openssl.sh
    
    The old version needed to be run from a specific directory, didn't pass
    `shellcheck`, and would unconditionally prepend the OpenSSL text output
    to all certs in the testdata dir (even if they already had it).
    
    This version should be safer and suitable to be integrated with CI in
    a later step. It also supports taking a glob as the first argument and
    only prepending certs that need it and have a filename that matches
    the glob.
    
    * testdata: add openssl -text output where missed.
    
    This catches up all of the test data to have prepended text.
    
    * testdata: add note to subCertLocalityNameDoesNotNeedToAppear.pem.
    
    This test file does not parse successfully with OpenSSL 1.1.1d on my dev
    machine. Adding a small text note about this before the PEM content
    avoids the `v2/test/prepend_testcerts_openssl.sh` script emitting
    a warning.
    
    * CI: require all testdata is prepended w/ text.
    
    This updates CI to run the `test/prepend_testcerts_openssl.sh` script
    and fail if there are any diffs to the `testdata/` directory. This would
    indicate there was a `.pem` file that didn't have text prepended to it.

[33mcommit 83d24bd1f73e7a4ba91091ed1aa894d00ab8d68b[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Tue Mar 17 12:17:42 2020 -0400

    lints: lint for upcoming Apple max cert lifetime policy. (#417)
    
    A new `e_tls_server_cert_valid_time_longer_than_398_days` lint is added for the
    Apple source category (presently named `lint.AppleCTPolicy`, see
    https://github.com/zmap/zlint/issues/418).
    
    This lint returns an error lint result if a server-auth certificate issued after
    Sept 1st, 2020 has a lifetime > 398 days. The lifetime is calculated as Apple
    specifies, e.g. "398 days is measured with a day being equal to 86,400
    seconds.".
    
    A warning result is returned if a certificate issued after Sept 1st, 2020 has
    a lifetime > 397 days and < 398 days. This matches Apple's SHOULD-equivalent
    recommendation to use a validity period <= 397 days in length.
    
    See https://support.apple.com/en-us/HT211025 for more information.
    
    Resolves https://github.com/zmap/zlint/issues/407

[33mcommit cfbfdeca3ae4c583b66a9c002275f48c478495d0[m
Author: TLD Update Robot <47792085+tld-update-bot@users.noreply.github.com>
Date:   Sat Mar 14 14:15:17 2020 -0400

    gTLD autopull: 2020-03-14T17:26:52Z (#420)
    
    Co-authored-by: tld-update-bot <cpu+tldbot@letsencrypt.org>

[33mcommit c7c6a31fe3969214b3fbf417d52137fc8743c37b[m
Author: MTG <36234449+mtgag@users.noreply.github.com>
Date:   Thu Mar 12 16:02:02 2020 +0100

    lints: enforce Mozilla PKI policy RSASSA-PSS encoding requirements (#377)
    
    This commit adds a new Mozilla source lint
    (`e_mp_rsassa-pss_parameters_encoding_in_signature_algorithm_correct`) for
    enforcing the RSASSA-PSS encoding requirements for TBSCertificate signature
    algorithm fields based on version 2.7 of the Mozilla PKI policy.
    
    It returns an `Error` result when the RSASSA-PSS parameters of
    a TBSCertificate's Signature algorithm field do not match the exact encoded
    bytes specified in the Mozilla policy.

[33mcommit b28794baf37819442e98e28655a8d4536681b095[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Wed Mar 11 13:01:07 2020 -0400

    docs: fix template to use v2 package import. (#416)
    
    The `template` file used by `v2/newLint.sh` needs to use the ZLint 2.0.0
    import path for the `lint` package or building a lint created with the
    utility will fail.

[33mcommit 19685159ea30c8546bdc0463e40dd07d01902034[m
Author: Zakir Durumeric <zakird@gmail.com>
Date:   Thu Mar 5 05:50:35 2020 -0800

    lints: disallow reserved iPAddresses in NCs (#414)
    
    Co-authored-by: Zakir Durumeric <zakird@gmail.com>

[33mcommit 48bf6ee88374e55cd233d713175690c0c241b24f[m
Author: Zakir Durumeric <zakird@gmail.com>
Date:   Wed Mar 4 11:29:45 2020 -0800

    remove lisp reserved range since no longer IANA reserved (#415)
    
    * remove lisp reserved range since no longer IANA reserved
    
    * go way of deprecating const
    
    * replacing README, bad change made it in. Will get corrected with squash

[33mcommit 3329bb69d206198706abc7c84eccb59011bd9de5[m
Author: BJ Cardon <cardonator@users.noreply.github.com>
Date:   Tue Feb 25 17:28:21 2020 -0500

    README: fix a typo and fix the example for LintCertificateEx (#409)

[33mcommit 5b2df5c915f1d73dbfcd4f2ac2a6241103fbc4a3[m
Author: MTG <36234449+mtgag@users.noreply.github.com>
Date:   Wed Feb 19 15:24:42 2020 +0100

    lints: enforce Mozilla PKI policy omission of id-RSASSA-PSS oid (#376)
    
    Adds new Mozilla sourced lint, `e_mp_rsassa-pss_in_spki`, that enforces Section 5.1.1
    of the v2.7 Mozilla PKI policy[0]:
    
    CAs MUST NOT use the id-RSASSA-PSS OID (1.2.840.113549.1.1.10) within a SubjectPublicKeyInfo to represent a RSA key.
    
    [0]: https://www.mozilla.org/en-US/about/governance/policies/security-group/certs/policy/

[33mcommit 36d042eba350cc658fc10997693ce44eb8fa616d[m
Author: Daniel McCarney <cpu@letsencrypt.org>
Date:   Fri Feb 14 14:01:39 2020 -0500

    ci: try and fix goreleaser for v2 structure (round 2) (#406)

[33mcommit a03f7226d0faf5b92aec5aaa4d6f7fd1d942a710[m
Author: Daniel McCarney <cpu@letsencrypt.org>
Date:   Fri Feb 14 12:53:54 2020 -0500

    ci: try and fix goreleaser for v2 structure (#405)
    
    * ci: try moving .goreleaser.yml to v2/
    * ci: update main path

[33mcommit fd40f579253ea1ebfb18a585ab5cd8e7dcde61aa[m
Author: David Adrian <davidcadrian@gmail.com>
Date:   Thu Feb 13 18:34:21 2020 -0500

    Fix v2 with go.mod (#398)
    
    https://github.com/zmap/zlint/pull/398
    
    Adopt the **Major Subdirectory** approach from
    https://github.com/golang/go/wiki/Modules#releasing-modules-v2-or-higher
    
    ### Effects:
    
    * Moves all code into a `v2` subdirectory. Consumers of ZLint v2 will need to change their imports in go to `github.com/zmap/zlint/v2` and update `go.mod` accordingly.
    * Old versions of ZLint are still fetchable via `go.mod` and version pinning, e.g. `require github.com/zmap/zlint@1.1.0`
    * People using `go get` without modules are going to have all their code break. However, this was going to happen regardless of what we do with our directories because we made breaking changes to our code.
    
    ### Patching old versions
    
    To patch an old version (pre-2.0), we would need to branch off of one of the old tags. To avoid this, we could attempt to maintain support for 1.1.0 by implementing another point release at the top-level directory on top of `v2`, and then exposing the old API's above the v2 directory (see https://github.com/rsc/quote/blob/master/quote.go as an example). However, I don't believe we plan on supporting old versions at all, so there's no reason to do this. On the off chance we do need to cut a point release for 1.1.0, we can use Git.
    
    ### Why not use the Major Branch approach
    
    When not maintaining support for v1 side-by-side with v2, the major branch approach is identical to the Major Subdirectory approach from a Git standpoint---point release for v1 would need to be done on branches. However, it does then have the side effect that the import path for your code no longer matches the directory structure of your code. Effectively, we could `mv` all of the `v2` directory back up to the top-level, claim the name `...v2/`, and use `v2` in all of our import paths. While this might look cleaner from a Git repo standpoint, my gut sense is that we may as well match the directories to the import path, since that's slightly easier to grok.

[33mcommit 53441bdd36c98f9d85ece445effe95d9c283f1c8[m
Author: Paschalis Korosoglou <824785+pkoro@users.noreply.github.com>
Date:   Thu Feb 13 22:38:04 2020 +0200

    misc: update newLint.sh script and contributing guide.  (#397)
    
    * Minor changes in lint generator script
    * Modifies package name per new lint
    * Adds variable LINTNAME and seperates it from actual filename

[33mcommit 24e7a0db2810c5364d9e42efbfdb6d6a6940aeca[m
Author: Daniel McCarney <cpu@letsencrypt.org>
Date:   Tue Feb 11 15:38:46 2020 -0500

    README: Update, split out a CONTRIBUTING.md (#386)

[33mcommit 79424f2a127788b83b973b5279ac7bd873705677[m
Author: Daniel McCarney <cpu@letsencrypt.org>
Date:   Tue Feb 11 15:10:50 2020 -0500

    cmd/zlint: fix panic w/ deref of nil registry. (#385)
    
    This fixes a panic that can occur when there are no filtering arguments
    provided to the `zlint` command line tool.
    
    This occurs because `setLints` returned a `nil` `Registry` when the intention
    was to use the global registry.
    
    Before fix:
    ```
    $ zlint -list-lints-source
    panic: runtime error: invalid memory address or nil pointer dereference
    [signal SIGSEGV: segmentation violation code=0x1 addr=0x38 pc=0x717bd0]
    
    goroutine 1 [running]:
    main.main()
            /home/daniel/go/src/github.com/zmap/zlint/cmd/zlint/main.go:85 +0xe0
    ```
    
    After fix:
    ```
    $ zlint -list-lints-source
        AWSLabs
        Apple
        CABF_BR
        CABF_EV
        ETSI_ESI
        Mozilla
        RFC5280
        RFC5480
        RFC5891
        ZLint
    ```

[33mcommit 7741587316b5f34b13c0f4849816dd33697f5f19[m
Author: Daniel McCarney <cpu@letsencrypt.org>
Date:   Tue Feb 11 13:43:45 2020 -0500

    zlint: refactor lint reg., allow filtering lints used. (#372)
    
    This replaces the exported `lint.LintMap` field (`map[string]*Lint`) that was used by `RegisterLint` with a more robust solution based around a `Registry` interface. This allows ZLint users to include/exclude lints by name or source.
    
    ## Top Level API
    
    The `lint.RegisterLint` function remains the same, meaning individual lints are not changed. Lints that call this function are added to the global registry. Clients can access this registry with `lint.GlobalRegistry()`. Similarly the top level `zlint.LintCertificate` remains unchanged. It lints the provided cert with all lints in the global registry.
    
    The old `zlint.LintCertificateFiltered` function that accepted a lint name regex to filter the lints applied is replaced by a new function `zlint.LintCertificateEx` that allows specifying a `lint.Registry` explicitly. The same regex filtering can be done by pre-filtering the provided registry (See notes below).
    
    The `lint.Source` type was changed from an int enum to a string enum. This makes it easier to work with as a consumer (e.g. via command line flags, and JSON output) and since the number of lints (and sources) is small the benefits to using an int enum type are minimal. The serialized form of Lints now includes the `Source` field in the output as `"source"`.
    
    ## Registry
    
    The `Registry` interface also allows finding all lint names with `Names()`, finding all lint sources with `Sources()`, finding a specific lint by name with `ByName()`, and finding all lints for a given source
    with `BySource()`.
    
    The `zlint.EncodeLintDescriptionsToJSON` function is now implemented by the `Registry` interface as `WriteJSON`. This makes it easier to encode a subset of the Registry's lints by filtering the global registry.
    
    ~Like before (with the exported `map[string]*Lint`) the registry is not safe for concurrent updates. That's fine for the current ZLint codebase but is something we may want to consider addressing in the future.~ _Edit: I decided it made sense to add locking to future proof the implementation for thread safety, see 6072e24 The implementation in this branch is now safe for concurrent access/registration_
    
    ## Registry Filtering
    
    Filtering of lints to be run is now done with the `lint.Registry.Filter` function and corresponding `lint.FilterOptions` type. This allows filtering a registry to include/exclude lints by name (or using a name regex), and to include/exclude lints by source.
    
    By filtering the global registry and then providing it explicitly to `zlint.LintCertificateEx` callers have control over exactly what lints will be applied.
    
    Filtering operations are applied with the following precedence: excludes by source > includes by source > excludes by name > includes by name.
    
    E.g. excluding a source and then trying to include a lint in that excluded source by name will not work. The source exclusion happens first.
    
    ## ZLint CMD Updates
    
    The `zlint` command (`cmd/zlint/main.go`) is updated to add four new command line flags:
    
    1. `-list-lints-sources` - Prints a list of lint sources, one per line.
    2. `-excludeSources` - Comma-separated list of lint sources to exclude.
    3. `-includeSources` - Comma-separated list of lint sources to include.
    4. `-nameFilter` - Regex used to match lint names to include (cannot be used at the same time as `-excludeSources` or `-includeSources)
    
    Two existing flags are renamed:
    
    1. `-include` becomes `-includeNames`
    2. `-exclude` becomes `-excludeNames`.
    
    Notably all three list flags (`-list-lints-json`, `-list-lints-schema` and `-list-lints-sources`) now operate **after** applying the include/exclude filters, allowing an easy way to find which lints/sources will be run with the filtered command line flags in use.
    
    ## Integration Test Updates
    
    Matching the `zlint` command the integration test (`integration/integration_test.go`) command line flags are updated to allow including/excluding lints by source.
    
    Resolves https://github.com/zmap/zlint/issues/344

[33mcommit 4666bb74318f221c77ca69616603d2e897d7cd3e[m
Author: mtg <git@mtg.de>
Date:   Tue Feb 4 17:58:04 2020 +0100

    Revert "lint about the encoding of qcstatements for PSD2"
    
    This reverts commit 6c2367080d148f4b8c01f96a4c80e3ac55d1ef26.

[33mcommit 6c2367080d148f4b8c01f96a4c80e3ac55d1ef26[m
Author: mtg <git@mtg.de>
Date:   Tue Feb 4 17:45:58 2020 +0100

    lint about the encoding of qcstatements for PSD2

[33mcommit 72fb7ad5f84659029286854606d828ead7ef38ef[m
Author: Daniel McCarney <cpu@letsencrypt.org>
Date:   Mon Feb 3 14:01:15 2020 -0500

    project: add goreleaser configuration. (#374)
    
    Adds configuration for GoReleaser to the project/CI.
    
    By default releases are added to the repository in draft status. This gives
    maintainers a chance to write the release notes/changelog and verify the build
    artifacts before publishing the release. Tags with "rc" in the name will be
    marked as pre-release candidates automatically.
    
    Updates #351

[33mcommit 8a37cc71af2ae1d62f62add5174f34626a3278d6[m
Author: TLD Update Robot <47792085+tld-update-bot@users.noreply.github.com>
Date:   Thu Jan 30 12:28:05 2020 -0500

    gTLD autopull: 2020-01-30T17:10:08Z (#375)

[33mcommit 11071233ae3b140047b71612ea934e86b0bd2d66[m
Author: Daniel McCarney <cpu@letsencrypt.org>
Date:   Wed Jan 29 11:51:51 2020 -0500

    deps: update golang.org/crypto/cryptobyte to 8b5121be2f68. (#373)
    
    This addresses CVE-2020-7919[0]:
    
    > On 32-bit architectures, a malformed input to crypto/x509 or the ASN.1
    parsing functions of golang.org/x/crypto/cryptobyte can lead to a panic.
    
    [0]: https://groups.google.com/forum/?utm_medium=email&utm_source=footer#!msg/golang-announce/Hsw4mHYc470/WJeW5wguEgAJ

[33mcommit 77026f684b414c0e84106407bf93f1cbfdba0ed8[m
Author: Jacob Hoffman-Andrews <github@hoffman-andrews.com>
Date:   Wed Jan 22 15:24:26 2020 -0800

    Add reference to RFC 6818 to clarify explicitText (#370)

[33mcommit c0407b6a75c49ca02e2dcbc9e3aad1cee89596ba[m
Author: Daniel McCarney <cpu@letsencrypt.org>
Date:   Tue Jan 21 11:24:44 2020 -0500

    lints: improve template_test.go (#367)
    
    Rather than hardcode a mapping of `LintSource` to package directory that
    needs to be maintained the `template_test.go` logic should just walk the
    filesystem under the `lints/` directory and check all `.go` files.
    
    This makes a smaller `lint` API, removes two places that need to be kept
    up to date with new `LintSource`s and results in a test that is robust
    against further subdirectory modifications (e.g. a structure deeper than
    1 package below `lints`).

[33mcommit dbb54ce28280eff52888ff9083ad3c4f26cbf214[m
Author: Daniel McCarney <cpu@letsencrypt.org>
Date:   Sun Jan 19 15:28:31 2020 -0500

    lints/mozilla: fix moz lint packages (#365)
    
    * lints/mozilla: fix package names
    
    * lint: add Moz source to list/Directory
    
    Co-authored-by: Zakir Durumeric <zakird@gmail.com>

[33mcommit cc90ed6cceeb23e77e84277c6c99d82f365bedfc[m
Author: Daniel McCarney <cpu@letsencrypt.org>
Date:   Sun Jan 19 15:10:33 2020 -0500

    test: more comments in helpers.go (#366)
    
    I had left this commit with some additional comments out of #364 by
    mistake.

[33mcommit 2cce20392ad0045f265595820b66914e2f844bd8[m
Author: Daniel McCarney <cpu@letsencrypt.org>
Date:   Fri Jan 17 17:57:32 2020 -0500

    lints: better test utils, avoid accessing lint.Lints directly (#364)
    
    * testlint: remove unused testDef dir/json data
    
    * testlint: move prepend_openssl.sh to test/
    
    * test: update paths in prepend_testcerts_openssl.sh
    
    * testlint: move all test certs to testdata/
    
    * test: fix helpers.go package/paths
    
    * lints: refactor all lints to use new test helpers.
    
    This avoids needing to access `lint.Lints` (soon to be un-exported) and
    also removes a lot of duplication (particularly of test data paths).

[33mcommit 566701eb88d3c0987bec9b8d7fa8b91eaea6202a[m
Author: Daniel McCarney <cpu@letsencrypt.org>
Date:   Thu Jan 16 12:18:29 2020 -0500

    Lints: add new lints for Mozilla Root Store Policy (adopted) (#353)
    
    * Lints: add new lints for Mozilla Root Store Policy.
    
    * Split Mozilla Root Store Policy RSA key lint
    
    Mozilla Root Store Policy contains multiple different requirements on
    RSA keys. All these were tested in a single lint. These split into two
    different lints based on the different requirement.
    
    * Deleted old Mozilla Root Store Policy RSA key lint.
    
    * Moved hasEKU() to util package.
    
    * Added fetching Mozilla Trust Store SPKIs.
    
    * Added cross-cert detection for MP EKU lint.
    
    * Added fatal error details to MP AuthKeyID lint.
    
    * Minor style change.
    
    * Added error details to MP ECDSA lint.
    
    * Renamed lint_mp_allowed_rsa_keys_exponent to e_mp_exponent_cannot_be_one.
    
    * Split RSA modulus lints into two files.
    
    * Minor fix in test function name.
    
    * Update lints/lint_mp_modulus_must_be_divisible_by_8.go
    
    Co-Authored-By: Daniel McCarney <daniel@binaryparadox.net>
    
    * Update lints/lint_mp_modulus_must_be_divisible_by_8.go
    
    Co-Authored-By: Daniel McCarney <daniel@binaryparadox.net>
    
    * Update lints/lint_mp_modulus_must_be_divisible_by_8.go
    
    Co-Authored-By: Daniel McCarney <daniel@binaryparadox.net>
    
    * Update lints/lint_mp_modulus_must_be_2048_bits_or_more.go
    
    Co-Authored-By: Daniel McCarney <daniel@binaryparadox.net>
    
    * Update lints/lint_mp_exponent_cannot_be_one.go
    
    Co-Authored-By: Daniel McCarney <daniel@binaryparadox.net>
    
    * Update lints/lint_mp_modulus_must_be_2048_bits_or_more.go
    
    Co-Authored-By: Daniel McCarney <daniel@binaryparadox.net>
    
    * Update lints/lint_mp_modulus_must_be_2048_bits_or_more.go
    
    Co-Authored-By: Daniel McCarney <daniel@binaryparadox.net>
    
    * Fixed incorrect commits by github ui.
    
    * Minor syntax change.
    
    * Removed zlint-mozilla-trusted-roots-update.
    
    * Renamed IsSPKIMozillaTrusted() to IsInMozillaRootStore().
    
    * lints: move lint_mp_* to lints/mozilla
    
    * lints: remove unneeded .gitinclude
    
    * lints/mozilla: fix build breakages from refactoring in master
    
    * lint: add src link for MozillaRootStorePolicy
    
    * zlint: run Mozilla lints
    
    * util: remove IsInMozillaRootStore and assoc. data.
    
    We'll return to this requirement in a subsequent PR when the tooling to
    generate the data can be reviewed and automated.
    
    * lints/mozilla: demote lint_mp_allowed_eku to notice
    
    * lints/mozilla; simplify lint_mp_ecdsa_allowed_algo CheckApplies.
    
    * lints/mozilla: rename ecdsa_allowed_algorithm -> ecdsa_allowed_curve_hash_pair.
    
    * lints/mozilla: add ecdsa curve/hash pair err detail.
    
    * lints/mozilla: ref trusted roots data issue num
    
    * lints/mozilla: fix mp_allowed_eku lint name
    
    * lints/mozilla: clarify allowed_eku desc
    
    * lints/mozilla: use 0 for err return from getSigningKeySize
    
    * lints/mozilla: assume CheckApplies works
    
    * lints/mozilla: remove `e_mp_ecdsa_allowed_curve_hash_pair`.
    
    * integration: add vetted expected results for Moz. lints
    
    Co-authored-by: Fotis Loukos <fotisl@users.noreply.github.com>
    Co-authored-by: Zakir Durumeric <zakird@gmail.com>

[33mcommit ea19827801ed54974eb244b531d22ff4ca585eb9[m
Author: Daniel McCarney <cpu@letsencrypt.org>
Date:   Mon Jan 13 15:21:56 2020 -0500

    README: fix crt.sh link target. (#349)

[33mcommit 4a01d2e8f105d7ff317aad339e98a5fe1e10b7d9[m
Author: Daniel McCarney <cpu@letsencrypt.org>
Date:   Mon Jan 13 14:09:57 2020 -0500

    README: Link to company sites, not bugzilla bugs. (#348)

[33mcommit 2c5688ec6e9eec503e31523ac7324922a41fc84f[m
Author: James Kasten <jdkasten@umich.edu>
Date:   Mon Jan 13 10:43:35 2020 -0800

    README: Add Google Trust Services to list of users/integrations (#347)
    
    Self reporting. There are aren't any associated bugs or posts about this, hence the lack of a link.

[33mcommit b7425cbf555a5a2b443aa1bdeea976e5e25f7065[m
Author: Daniel McCarney <cpu@letsencrypt.org>
Date:   Tue Jan 7 15:19:56 2020 -0500

    lints: add more context to `w_subject_contains_malformed_arpa_ip`. (#345)
    
    Section 7.1.4.2.1 of the BRs is a good citation for
    `e_subject_contains_reserved_arpa_ip` but isn't a great choice for
    `w_subject_contains_malformed_arpa_ip`.
    
    When the `.arpa` address doesn't have enough labels, or can't be parsed as an IP
    address it's clear that it isn't an internal IP address and so 7.1.4.2.1 isn't
    a good citation. Section 3.2.2.6 talks about wildcard domains for "registry
    controlled" zones, and `.arpa` is one of those (based on BCP49). A wildcard
    label is one way the `.arpa` domain wouldn't parse as an IP.
    
    While the larger discussion on how `.arpa` domains that aren't formatted per RFC
    3596 unfolds we can ref 3.2.2.6 and add a bit more context to the lint and
    description. It isn't perfect, but I think less confusing than ref'ing
    7.1.4.2.1, which clearly doesn't apply.
    
    See also https://github.com/zmap/zlint/issues/343

[33mcommit 9bba7b7e572cd92a5b6d74ad0520522f45277ffc[m
Author: Daniel McCarney <cpu@letsencrypt.org>
Date:   Mon Jan 6 13:56:08 2020 -0500

    lints: warn for RSA-PSS sigalg in cabf lint, not err. (#342)
    
    The `e_signature_algorithm_not_supported` lint enforces Section 6.1.5 of
    the baseline requirements by checking certificate signature algorithms
    against a fixed set. Previously this set did not include the RSA-PSS
    signature algorithms and would mistakenly flag certificates signed with
    a RSA-PSS algorithm with an error result.
    
    The BRs do not forbid using RSA-PSS signature algorithms (provided the
    associated digest algorithm is one of the three approved in 6.1.5). The
    Mozilla root program requirements do forbid RSA-PSS in v2.7+ but that
    should be checked in a separate Mozilla scoped lint.
    
    This commit adjusts the `e_signature_algorithm_not_supported` lint to
    return `lint.Warn` for RSA-PSS with SHA256, SHA384 or SHA512.
    
    See #326 for more background.

[33mcommit 359be75f66a8c0ee0ea2f7f8fcaed3df4095e32a[m
Author: TLD Update Robot <47792085+tld-update-bot@users.noreply.github.com>
Date:   Mon Jan 6 12:11:54 2020 -0500

    gTLD autopull: 2020-01-06T16:47:48Z (#341)

[33mcommit 86bcc674785a38a8bc72dcc306ee5a9572c3c0fb[m
Author: Daniel McCarney <cpu@letsencrypt.org>
Date:   Fri Jan 3 13:07:16 2020 -0500

    Misc. cleanups, unit test for finding leftover template bits. (#340)
    
    * tests: remove gofmt_test.go
    
    The golangci-lint pass run in CI includes an equivalent test. If folks
    want to test for unformatted code locally install the linter and run
    `golangci-lint run` in the root directory. This will flag findings
    above and beyond `gofmt` problems ahead of CI failing.
    
    * lints: remove commented out code.
    
    In three cases, remove a comment ahead of a return that added no
    useful context.
    
    In `lints/community/lint_rsa_exp_negative_test.go` remove a commented
    out test case for a negative RSA exponent. The test code doesn't build
    as-is and the referenced test cert (`rsaExpNegative.pem`) doesn't exist
    in-tree. A TODO is left to indicate there's missing test coverage for
    later follow-up.
    
    * lints: fix "certtificate" comment typo.
    
    * lints: fix tabs in ref text for lint_sub_cert_or_sub_ca_using_sha1.
    
    * lints: fix field name ref. in lint Descriptions.
    
    These two lints mistakenly said in their `Description` that they only
    check the `DNSNames` field of the certificate when in fact they only
    check the `IANDNSNames` field. There are two corresponding lints
    (`lints/community/lint_san_wildcard_not_first.go` and
    `lints/community/lint_san_bare_wildcard.go`) that check `DNSNames`.
    
    * lints: add slice of known LintSources, test for templating leftovers.
    
    There should never be finished lint source code that contains template
    text intended to be replaced by the programmer. A new
    `TestLeftoverTemplates` unit test is added to make sure we enforce this
    during CI to lessen the burden on code reviewers to catch this problem.
    
    * tests: use full path in TestLeftoverTemplates errs
    
    * lints: fix TestLeftoverTemplates findings
    
    Prior to these fixes all of the modified files had templating leftovers:
    ```
    === RUN   TestLeftoverTemplates
    --- FAIL: TestLeftoverTemplates (0.01s)
        template_test.go:49: Lint "cabf_br/lint_root_ca_extended_key_usage_present.go" contains template leftover "// Add actual lint here"
        template_test.go:49: Lint "cabf_br/lint_root_ca_key_usage_present.go" contains template leftover "// Add actual lint here"
        template_test.go:49: Lint "cabf_br/lint_sub_cert_cert_policy_empty.go" contains template leftover "// Add actual lint here"
        template_test.go:49: Lint "cabf_br/lint_sub_cert_certificate_policies_missing.go" contains template leftover "// Add actual lint here"
        template_test.go:49: Lint "cabf_br/lint_sub_cert_crl_distribution_points_does_not_contain_url.go" contains template leftover "// Add actual lint here"
        template_test.go:49: Lint "cabf_br/lint_sub_cert_eku_extra_values.go" contains template leftover "// Add actual lint here"
        template_test.go:49: Lint "cabf_br/lint_sub_cert_eku_missing.go" contains template leftover "// Add actual lint here"
        template_test.go:49: Lint "cabf_br/lint_sub_cert_eku_server_auth_client_auth_missing.go" contains template leftover "// Add actual lint here"
        template_test.go:49: Lint "cabf_br/lint_sub_cert_key_usage_cert_sign_bit_set.go" contains template leftover "// Add actual lint here"
        template_test.go:49: Lint "cabf_br/lint_sub_cert_key_usage_crl_sign_bit_set.go" contains template leftover "// Add actual lint here"
        template_test.go:49: Lint "rfc/lint_basic_constraints_not_critical.go" contains template leftover "// Add actual lint here"
        template_test.go:49: Lint "rfc/lint_ext_key_usage_not_critical.go" contains template leftover "// Add actual lint here"
        template_test.go:49: Lint "rfc/lint_basic_constraints_not_critical.go" contains template leftover "// Add actual lint here"
        template_test.go:49: Lint "rfc/lint_ext_key_usage_not_critical.go" contains template leftover "// Add actual lint here"
        template_test.go:49: Lint "rfc/lint_basic_constraints_not_critical.go" contains template leftover "// Add actual lint here"
        template_test.go:49: Lint "rfc/lint_ext_key_usage_not_critical.go" contains template leftover "// Add actual lint here"
    FAIL
    FAIL    command-line-arguments  0.017s
    FAIL
    ```
    
    * lints: update template test with another string, fix occurrences.
    
    ```
    === RUN   TestLeftoverTemplates
    --- FAIL: TestLeftoverTemplates (0.01s)
        template_test.go:50: Lint "cabf_br/lint_sub_ca_name_constraints_not_critical.go" contains template leftover "Change this to match source TEXT"
        template_test.go:50: Lint "community/lint_validity_time_not_positive.go" contains template leftover "Change this to match source TEXT"
        template_test.go:50: Lint "community/lint_validity_time_not_positive.go" contains template leftover "Change this to match source TEXT"
    FAIL
    FAIL    command-line-arguments  0.017s
    FAIL
    ```
    
    * lints: move lint_ian_bare_wildcard.go from RFC to community.
    
    It cites RFC 5280 but that RFC doesn't prescribe any semantics to the
    use of wildcards in DNSNames or elsewhere. I suspect this lint actually
    came from AWSLabs, similar to `lint_ian_wildcard_not_first.go` and
    `lint_san_bare_wildcard.go`, both of which are already in
    `lints/community/`.
    
    * lints: fix moved lint_ian_bare_wildcard.go source/category/package
    
    * lints: fix off-by-one in RFC max length lint Descs.
    
    The upper bounds being enforced against in the changed lints are
    inclusive. The lint tests were doing the right thing but the
    descriptions incorrectly described the boundary as if it were exclusive.
    
    For comparison the following lints already did the right thing already
    and had the UB+1 in the desc:
    ```
    lints/rfc/lint_subject_given_name_max_length.go
    lints/rfc/lint_subject_postal_code_max_length.go
    lints/rfc/lint_subject_street_address_max_length.go
    lints/rfc/lint_subject_surname_max_length.go
    ```
    
    * lint: revert accidental whitespace diff

[33mcommit e3ad0f9eba10b1fa0ee70e6581a627cdfeaa590c[m
Author: Zakir Durumeric <zakird@gmail.com>
Date:   Thu Jan 2 17:15:59 2020 -0500

    Split of lints into directories by source (#337)
    
    * initial pass at a dissection
    
    * moving in new lints
    
    * second pass on directories for lints
    
    * lints in better shape
    
    * make zlint work again
    
    * tests pass.
    
    * updating copyright while I'm making large sweeping changes
    
    * missing an important file
    
    * apparently a random util file wasn't go'fmted?????
    
    * integration tests fixes
    
    Co-authored-by: Daniel McCarney <daniel@binaryparadox.net>

[33mcommit 0ab41f2f58a96458f8311f3afbce35381d2addc1[m
Author: Zakir Durumeric <zakird@gmail.com>
Date:   Sat Dec 28 12:57:59 2019 -0600

    README: add note about small PRs (#339)

[33mcommit 257d49ddebf672fc0c581d7efdd3e62175b891e4[m
Author: TLD Update Robot <47792085+tld-update-bot@users.noreply.github.com>
Date:   Wed Dec 25 11:56:00 2019 -0500

    gTLD autopull: 2019-12-25T16:40:11Z (#338)

[33mcommit c74b45bf8ea3d3ba26201c53582bc3d9b6e0de3a[m
Author: Daniel McCarney <cpu@letsencrypt.org>
Date:   Mon Dec 9 16:40:10 2019 -0500

    CI: Add golangci-lint, enforce Go best practices (#335)
    
    * tidy: rename test functions so they are run.
    
    Unit tests functions must be named with `Test` as a prefix or they won't
    be run. This fixes an `unused` golangci-lint finding for this file.
    
    * tidy: remove unused functions.
    
    Neither of these functions are being used anywhere. Deleting them fixes
    a `golangci-lint` finding from the `unused` linter.
    
    * tidy: cleanup errcheck finding
    
    * tidy: cleanup errcheck finding
    
    * tidy: fixup all gosimple golangci-lint findings
    
    * tidy: fix ineffassign golangci-lint findings
    
    * tidy: cleanup gocritic golangci-lint finding
    
    * tidy: fix golangci-lint goimports findings
    
    * tidy: fix golangci-lint nakedret finding
    
    * tidy: ignore golangci-lint interfacer for one func
    
    * tidy: fix golangci-lint misspell finding
    
    * tidy: add some gocyclo lint ignores
    
    * CI: enforce golangci-lint.

[33mcommit 872e43139cd4269530ec2b237c5643e352a45fa3[m
Author: TLD Update Robot <47792085+tld-update-bot@users.noreply.github.com>
Date:   Fri Dec 6 11:47:41 2019 -0500

    gTLD autopull: 2019-12-06T16:32:55Z (#334)

[33mcommit 71201e7f6c374a07357066504a13577aed052cf6[m[33m ([m[1;33mtag: v1.1.0[m[33m)[m
Author: TLD Update Robot <47792085+tld-update-bot@users.noreply.github.com>
Date:   Mon Dec 2 11:56:38 2019 -0500

    gTLD autopull: 2019-12-02T16:31:54Z (#333)

[33mcommit 9f4f7099a6efd817b6a8d9c67249b6c505cea27e[m
Author: Daniel McCarney <cpu@letsencrypt.org>
Date:   Tue Nov 26 12:22:20 2019 -0500

    README: Add Camerfirma to users list. (#331)

[33mcommit 5b9959d6ee45ffa6d9368aea8c08897e9115dbea[m
Author: Daniel McCarney <cpu@letsencrypt.org>
Date:   Tue Nov 26 12:19:44 2019 -0500

    lints: add w_extra_subject_common_names lint (#330)
    
    The BRs don't expressly forbid having multiple subject common names but it
    seems counter to the intentions of section 7.1.4.2.2 and generally weird enough
    that ZLint should produce a Warn level finding (like cablint does already).
    
    To implement this lint `github.com/zmap/zcrypto` is updated to 7f2fe6f
    the tip of master at the time of writing. Notably this brings in a change that
    stores multiple subject CN values into the `Certificate.Subject` `pkix.Name`
    field.
    
    Along the way, I bumped the Go mod version to 1.13 and updated the README to
    reflect we expect Go 1.13+ to use zlint master (it's what our CI is using
    already).
    
    Integration tests flagged 16 certificates tripping the new lint. Those
    certificates are:
    
    https://censys.io/certificates/1cd00ff04092b4c2faa7becd76c44f0a7ca38fbb2269d1588d81756159ca6ec6
    https://censys.io/certificates/2242b07bb4393996a940a01eb08336849a5d199410715c0af3e8c6a5cd007932
    https://censys.io/certificates/225d57740aa0a824c164f7c5994cccbe8627310d573c793632005e170ee07699
    https://censys.io/certificates/22aa2d265a69c95792967a9c182928e01c8bfbcab1667e6e9e0259dd31041e7f
    https://censys.io/certificates/2b09e53182a4ef6e440b3c39f90f3d91d8e98a6d973233d323956981e0674deb
    https://censys.io/certificates/2cff4e44a9fb9563e55217300f4d5e49f73c20101aed5725e7b6b04d539cff9e
    https://censys.io/certificates/6b1db534a22b2a3ce9eaa54b5ce8720391db6ca60014edd2ff4b78d7bffc2cb9
    https://censys.io/certificates/7d95620a0993673e289e0ae894566dd6c0b03e936d1c0a4d3484d8898d9296ae
    https://censys.io/certificates/824c4c4f893962023dc256d48ce52ac667f7b504e9a46947c3207b57f64d7ad0
    https://censys.io/certificates/8795bcd44516c8643d2a41a6da735d556b989a4c2c96c974bb11c0811404f479
    https://censys.io/certificates/a04c2da1bccb1a97c6a882a4b8688941673adc715d8d2769eec345cb1c4e3b52
    https://censys.io/certificates/c289369fed510acf2653b8a8eb8bee949d3c18f8d8a2817c6373ad2ec7e789d9
    https://censys.io/certificates/c6bb0cb620135851db0eb8ca54d0e8b77c1565683e405a3177e460e8fd3cd9cc
    https://censys.io/certificates/e18e4a9277e96542371436a1b78f1d3e09bf1095a44d2e7c02e049945a3dd66a
    https://censys.io/certificates/f006baa460163b35cbc35d96ffb20aa1ec27f2a8edf5db99a8ad51ab7b5bb88b
    https://censys.io/certificates/f658d49a2dc3332e2e0d893cbeb3d982d404bf5bf17d7e2e925e38ba7093e174
    
    All of them were checked to ensure that they do in fact have multiple subject
    common names and are being linted correctly.
    
    I also took this as a chance to describe the process of adding a new lint and
    updating the golden data for `integration/README.md`.

[33mcommit 227314aaca984cd9137bdf34796b294566e8725f[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Sun Nov 24 16:51:13 2019 -0500

    integration: restructure expected data to use lints vs. cert FPs (#329)
    
    * integration: small TestMain tidy-up chore.
    
    * integration: rename -summary flag to be more specific
    
    * integration: rename -force to -forceDownload.
    
    * integration: add -overwriteExpected flag
    
    * integration: add -lintSummarize, results by lint name.
    
    * integration: rename result type to resultCount
    
    * integration: cleaner resultCount increment
    
    * integration: sort summary output.
    
    * integration: support filtering by fp/lint name
    
    * integration: add small certByFP.sh util script
    
    * makefile: add more integration examples.
    
    * integration: track expected by lint instead of cert fp.
    
    * README: first pass at integration test docs
    
    * integration: show # of files left to download
    
    * CI: remove data caching.
    
    Per Travis docs the cache is fetched over the network from an
    S3-like-service and isn't any faster than downloading the data directly
    from Github for our use-case. Removing the caching saves us from
    spending time at the end of builds uploading the new cache data back to
    the cloud.
    
    * integration: fix progress print format args
    
    * integration: print loading progress
    
    * integration: refine README
    
    * integration: README clarifications.

[33mcommit 6ba0b4dbc15bab5dd75635674bbe59e9b5c3ad59[m
Author: TLD Update Robot <47792085+tld-update-bot@users.noreply.github.com>
Date:   Wed Nov 20 11:59:22 2019 -0500

    gTLD autopull: 2019-11-20T16:25:28Z (#327)

[33mcommit eea5fe83935a0904234975899617923cd3a0e83e[m
Author: TLD Update Robot <47792085+tld-update-bot@users.noreply.github.com>
Date:   Fri Nov 15 11:40:49 2019 -0500

    gTLD autopull: 2019-11-15T16:23:46Z (#325)

[33mcommit 7314deb0a2a11b829a1144d21ec75d16b587e8f4[m
Author: Daniel McCarney <cpu@letsencrypt.org>
Date:   Mon Sep 30 12:22:49 2019 -0400

    tests: add large cert corpus based integration test. (#318)
    
    # Background
    
    The idea is to run `zlint` across a big corpus of certificates, ideally as
    diverse a corpus as possible, and compare the results against expected values.
    The hope is that this will catch regressions in lints that change lint results
    for the corpus in an unexpected way.
    
    To start with 60 data files (each of ~5000 certificates) are used to lint
    ~600,000 certificates during each CI run to verify lint results do not change
    unexpectedly. The tooling was written such that it is easy to run against
    larger corpuses before doing a big release, or on a less frequent schedule than
    for day to day CI.
    
    The config in this branch adds ~10 minutes runtime to CI. I verified it would
    help catch regressions by reverting the fix to the
    `e_subject_printable_string_badalpha` lint
    (`5dcecad773158b82b5e52064ee2782d1b8a79314`). The integration tests
    successfully flagged a difference that would have focused attention on the new
    lint: https://gist.github.com/cpu/3a80db08a14a9de7a56db4ff1fc821e9
    
    ## Golden state
    
    In a perfect world we would record the full results from each lint for each
    certificate and make sure they don't change unexpectedly. Recording that much
    data in a "golden" state file in the repo for a large number of certs would
    result in pretty significant bloat and slow clones. Putting it out-of-repo
    would make PRs cumbersome and adds challenges for external contributors.
    
    As a compromise the current integration test only tracks a count of fatal,
    error, warning, and notice lint results per-certificate but not details about
    which lint produced each result at that level.
    
    With the one corpus file in this branch this adds ~60mb of golden state data to
    the repo (see `integration/config.json`) which seems acceptable to me. If we
    want to trim this down we could stop tracking notice level lint results or use
    a more efficient serialized representation than JSON. One of the advantages of
    the bloated JSON form is that it is human readable and easy to update in an
    editor.
    
    To produce the golden state for the first time or for a large update you can
    edit the `integration/config.json` to delete the `"Expected"` map. When the
    test is run and there is no configured `Expected` map then the current results
    are saved as the `Expected` map to be validated against in subsequent runs.
    
    ## Corpus
    
    The cert corpus data is quite large so the integration tests are written to
    keep the data out of tree and only download it when needed. Support is included
    to automatically decompress the downloaded data using bzip2 when the URL ends
    in the `.bz2` file extension.
    
    The Travis CI config is updated to explicitly cache the downloaded corpus data
    between builds to help avoid needing to download it every build. Similarly some
    effort was made to process data in parallel to keep builds as fast as possible
    while still linting as many certs as we can. Local testing on good hardware can
    be done even faster by increasing the parallelism over what is used in CI where
    the worker machines are weak and often under-provisioned.
    
    To start with 60 corpus files (each ~5000 certs) are configured for day to day
    CI. These files (and more) live in the
    https://github.com/zmap/zlint-test-corpus repository.
    
    ## Running Integration Tests
    
    By default when running `go test` (or `make test`) the integration tests are
    not built or run. You must provide `-tag=integration` to build/run these tests.
    This should keep day to day development quick.
    
    The `makefile` is updated with a `make integration` target to run just the
    integration tests and the Travis script uses that to run integration tests. The
    `PARALLELISM` and `INT_FLAGS` makefile variables can be used to change the
    number of Go routines and additional flags (see Extra Command Line Flags)
    without modifying the Makefile.
    
    ## Extra Command Line Flags
    
    To use an alternative configuration specify the `-config` command line flag. By
    default `integration/config.json` is used.
    
    To control the number of linting Go routines used by the `TestCorpus`
    integration test change the `-parallelism` command line flag when running the
    integration tests. By default 5 Go routines are used for local integration
    tests and 3 for Travis CI.
    
    To force corpus data to be downloaded even when it exists on disk use the
    `-force` command line flag when running the integration tests. By default
    `-force` is disabled.
    
    To have a summary of certificate fingerprint and integration test results at
    the end of the integration test provide the `-summary` command line flag. Note
    that this is quite spammy and is disabled by default
    
    To change how many certificates are linted before a '.' character is printed to
    the screen use the `-outputTick` flag. By default one period is printed per
    1000 certificates to keep Travis from deciding the CI job is dead.

[33mcommit 7db289cfd3689e9ecd04b8ab31681ab69c90bf29[m[33m ([m[1;33mtag: v1.0.2[m[33m)[m
Author: bilalashraf123 <bilal.ashraf@gmail.com>
Date:   Thu Sep 26 23:20:33 2019 +0500

    Fixed two bugs in QcEuLimitValue - QC Statement (#315)
    
    * Changed to IdEtsiQcsQcLimitValue
    
    * Marked fields exportable for Unmarshal
    
    * No QcEuLimitValue in test cert
    
    * Added tests for QcStmtLimitValue

[33mcommit 00156801166b89b8cbeb1f56bd6bd720120960bf[m
Author: TLD Update Robot <47792085+tld-update-bot@users.noreply.github.com>
Date:   Sat Sep 21 17:40:57 2019 -0400

    autopull: 2019-09-21T15:56:13Z (#321)

[33mcommit 43843b085caa3465b5fe7ea4f53c62a06e67bd25[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Sun Sep 15 19:41:22 2019 -0400

    README: Add a section on users/integrations. (#320)
    
    * README: Add a section on users/integrations.
    
    * README: update crt.sh link, also note it is used for Sectico.
    
    * README: Add Izenpe to users list
    
    * README: Add GoDaddy to users list.
    
    * README: Add EJBCA integration.

[33mcommit c67053f79915a9f9edf470b522977df16ac2c07c[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Sun Sep 15 19:37:42 2019 -0400

    CI: Switch to Go 1.13. (#319)
    
    The `makefile` could have the `GO_ENV` cleaned up to remove
    `GO111MODULE="on"` but for now I've left it to aid with Go 1.12.x
    compatibility.

[33mcommit c6437affd66336f6a9bc50cc7213ec5a5e1deddd[m
Author: TLD Update Robot <47792085+tld-update-bot@users.noreply.github.com>
Date:   Thu Sep 12 12:03:52 2019 -0400

    tld autopull: 2019-09-12T15:51:26Z (#317)

[33mcommit 0d4db4102b199c2c07467c3b048ef51583535b3c[m
Author: TLD Update Robot <47792085+tld-update-bot@users.noreply.github.com>
Date:   Tue Sep 10 11:23:46 2019 -0400

    gTLD autopull: 2019-09-10T15:21:16Z (#314)

[33mcommit a0b3bc322455906a290ea52bd3064b94a403b03a[m
Author: TLD Update Robot <47792085+tld-update-bot@users.noreply.github.com>
Date:   Fri Aug 30 15:09:06 2019 -0400

    gTLD autopull: 2019-08-30T19:02:52Z (#313)

[33mcommit 5dcecad773158b82b5e52064ee2782d1b8a79314[m[33m ([m[1;33mtag: v1.0.1[m[33m)[m
Author: Zakir Durumeric <zakird@gmail.com>
Date:   Sat Aug 24 09:30:31 2019 -0400

    lints: fix e_subject_printable_string_badalpha for single quote (#311)
    
    #309 missed the single quote character in its regex of valid characters. This PR
    fixes the regex and adds a test case. This PR addresses the comment here:
    https://github.com/zmap/zlint/pull/309#discussion_r317297206

[33mcommit dc635f9345c00451bd06a5877e1ebacfb1b9b0f1[m[33m ([m[1;33mtag: v1.0.0[m[33m)[m
Author: Daniel McCarney <cpu@letsencrypt.org>
Date:   Mon Aug 19 15:45:15 2019 -0400

    README: add semver guidance (#310)

[33mcommit 3307e6abe1904cf6f0573d7c9fb35a800385f02c[m
Author: Daniel McCarney <cpu@letsencrypt.org>
Date:   Mon Aug 12 19:42:38 2019 -0400

    lints: add e_subject_printable_string_badalpha lint. (#309)
    
    When the raw Subject RDNSequence of a Certificate includes
    a PrintableString type DirectoryString attribute the value of the
    attribute must adhere to the PrintableString character set defined in
    RFC 5280 Appendix B:
    
      The character string type PrintableString supports a very basic Latin
      character set: the lowercase letters 'a' through 'z', uppercase
      letters 'A' through 'Z', the digits '0' through '9', eleven special
      characters ' = ( ) + , - . / : ? and space.
    
    If any of the PrintableString attributes in the linted Certificate's
    raw subject do not match a regexp for this character set an Error level
    lint result is returned by the `e_subject_printable_string_badalpha` lint.

[33mcommit d18ad02ac400de715e688d59f315e53de1cfe18f[m
Author: Daniel McCarney <cpu@letsencrypt.org>
Date:   Fri Aug 9 16:59:28 2019 -0400

    README: Make Go version requirement explicit. (#306)

[33mcommit 0dfef633f728201b05b16493f0b765da56469862[m
Author: TLD Update Robot <47792085+tld-update-bot@users.noreply.github.com>
Date:   Fri Aug 9 16:08:37 2019 -0400

    gTLD autopull: 2019-08-09T20:03:29Z (#308)

[33mcommit 88c3f6b6f2f5ebc573c4679e548d6f1823d89213[m
Author: TLD Update Robot <47792085+tld-update-bot@users.noreply.github.com>
Date:   Tue Aug 6 14:24:16 2019 -0400

    gTLD autopull: 2019-08-06T18:19:07Z (#304)

[33mcommit fd021b4cfbeb919cc763d1cafd1e604658a6bbe7[m
Author: TLD Update Robot <47792085+tld-update-bot@users.noreply.github.com>
Date:   Tue Aug 6 11:40:20 2019 -0400

    gTLD autopull: 2019-08-06T15:35:36Z (#303)

[33mcommit 1fdad3421775e34615dec50d234ac9cb73c9a6a1[m
Author: TLD Update Robot <47792085+tld-update-bot@users.noreply.github.com>
Date:   Sat Aug 3 11:32:02 2019 -0400

    gTLD autopull: 2019-08-03T15:29:31Z (#302)

[33mcommit 0e1f6d0520cf1137c74f40f80c44c7a1d75c38d1[m
Author: Daniel McCarney <cpu@letsencrypt.org>
Date:   Thu Aug 1 14:37:18 2019 -0400

    README: add guidance on choosing a lint result level. (#301)
    
    Clarifying the historic context on the lint levels in the README will
    help new lint contributors (and consumers of the lint results)
    understand the rationale behind which lints return which result.
    
    The text is largely stolen from @zakird's comment[0] on an unrelated
    issue.
    
    [0]: https://github.com/zmap/zlint/issues/291#issuecomment-514413055

[33mcommit 13a927f87ec7ccb59edbf529ba0a64d26621c6b0[m
Author: Daniel McCarney <cpu@letsencrypt.org>
Date:   Thu Aug 1 14:04:51 2019 -0400

    lints: consistently name lints by lint status result. (#300)
    
    * lints: subject_contains_malformed_arpa_ip -> w_subject_contains_malformed_arpa_ip
    
    * lints: onion_subject_validity_time_too_large -> e_onion_subject_validity_time_too_large
    
    * lints: ext_tor_service_descriptor_hash_invalid -> e_ext_tor_service_descriptor_hash_invalid
    
    * lints: ct_sct_policy_count_unsatisfied -> w_ct_sct_policy_count_unsatisfied
    
    * lints: san_dns_name_onion_not_ev_cert -> e_san_dns_name_onion_not_ev_cert
    
    * lints: subject_contains_reserved_arpa_ip -> e_subject_contains_reserved_arpa_ip
    
    * tests: add test to enforce lint name prefix convention.
    
    * test: fix TestLintNames allowedPrefixes
    
    * review: fix gofmt of new test file

[33mcommit b126a9b258d55b1b9621e9a16525567317d86b6e[m
Author: Daniel McCarney <cpu@letsencrypt.org>
Date:   Thu Aug 1 12:21:32 2019 -0400

    lints: ct_sct_policy_count_unsatisfied NA for precerts. (#299)
    
    The `ct_sct_policy_count_unsatisfied` lint should return NA when asked
    to lint a precertificate (e.g. a "poisoned" cert containing the
    `util.CtPoisonOID` defined in RFC 6962).

[33mcommit 9971d62266e74547157ff95b5413227e21d8fe23[m
Author: Daniel McCarney <cpu@letsencrypt.org>
Date:   Tue Jul 30 17:53:01 2019 -0400

    lints: implement json.Unmarshaler for LintStatus. (#297)
    
    The `lints.LintStatus` type implements `json.Marshaler` to marshal to
    a human readable string (e.g. `"error"` instead of `6`). Without
    providing a compatible `json.Unmarshaler` implementation downstream
    users that marshal a `lints.LintStatus` to JSON will encounter an error
    when they try to unmarshal it due to a mixmatch of types (`int` vs
    `string`).

[33mcommit 757a6bf54dd74342f378904af0265e01b26d975e[m
Author: TLD Update Robot <47792085+tld-update-bot@users.noreply.github.com>
Date:   Tue Jul 30 11:30:40 2019 -0400

    gtld autopull: 2019-07-30T15:27:51Z (#296)
    
    Marks removal date of `.bnl`.

[33mcommit d8d7761d228b61e7bd2227a7dfa9cd9b60652674[m
Author: Daniel McCarney <cpu@letsencrypt.org>
Date:   Mon Jul 29 16:14:22 2019 -0400

    util: faster GetExtFromCert using ExtensionsMap. (#295)
    
    The `github.com/zmap/zcrypto` library that provides certificate parsing
    for zlint has been updated to add a `ExtensionsMap` field on parsed
    Certificates.
    
    The `util.GetExtFromCert` function is updated to use the `ExtensionsMap`
    for O(1) extension access by OID instead of an O(n) search of the
    `Extensions` slice.
    
    This change speeds up the function and removes `util.GetExtFromCert`
    from the top10 cumulative CPU usage nodes in the zlint benchmarking.
    
    Before:
    ```
    $ cd $GOPATH/src/github.com/zmap/zlint
    $ go test --run=XXX -bench=. -cpuprofile all.profile
    <snipped>
    $ go tool pprof all.profile
    File: zlint.test
    Type: cpu
    Time: Jul 26, 2019 at 10:58am (EDT)
    Duration: 4.13mins, Total samples = 4.79mins (115.90%)
    Entering interactive mode (type "help" for commands, "o" for options)
    (pprof) top10
    Showing nodes accounting for 146.71s, 51.07% of 287.25s total
    Dropped 522 nodes (cum <= 1.44s)
    Showing top 10 nodes out of 221
          flat  flat%   sum%        cum   cum%
        45.89s 15.98% 15.98%    142.36s 49.56%  runtime.mallocgc
        31.59s 11.00% 26.97%     39.26s 13.67%  runtime.heapBitsSetType
        18.20s  6.34% 33.31%     18.20s  6.34%  runtime.nextFreeFast
        11.27s  3.92% 37.23%     20.05s  6.98%  runtime.scanobject
         9.90s  3.45% 40.68%      9.90s  3.45%  runtime.memclrNoHeapPointers
         8.31s  2.89% 43.57%      8.31s  2.89%  encoding/asn1.ObjectIdentifier.Equal
         6.92s  2.41% 45.98%     13.52s  4.71%  github.com/zmap/zlint/util.GetExtFromCert
         5.63s  1.96% 47.94%         7s  2.44%  runtime.heapBitsForAddr
         4.60s  1.60% 49.54%    130.51s 45.43%  runtime.newobject
         4.40s  1.53% 51.07%    245.67s 85.52%  github.com/zmap/zlint.BenchmarkZlint.func4
    ```
    
    After:
    ```
    $ cd $GOPATH/src/github.com/zmap/zlint
    $ go test --run=XXX -bench=. -cpuprofile all.new.profile
    <snipped>
    $ go tool pprof all.new.profile
    File: zlint.test
    Type: cpu
    Time: Jul 29, 2019 at 3:38pm (EDT)
    Duration: 4.21mins, Total samples = 4.89mins (116.10%)
    Entering interactive mode (type "help" for commands, "o" for options)
    (pprof) top10
    Showing nodes accounting for 142.94s, 48.69% of 293.55s total
    Dropped 513 nodes (cum <= 1.47s)
    Showing top 10 nodes out of 233
          flat  flat%   sum%        cum   cum%
        48.05s 16.37% 16.37%    138.73s 47.26%  runtime.mallocgc
        28.32s  9.65% 26.02%     37.09s 12.63%  runtime.heapBitsSetType
        14.54s  4.95% 30.97%     14.54s  4.95%  runtime.nextFreeFast
        12.19s  4.15% 35.12%     21.52s  7.33%  runtime.scanobject
        10.59s  3.61% 38.73%     30.52s 10.40%  runtime.concatstrings
         9.60s  3.27% 42.00%      9.60s  3.27%  runtime.memclrNoHeapPointers
         5.60s  1.91% 43.91%      6.88s  2.34%  runtime.heapBitsForAddr
         4.85s  1.65% 45.56%    252.74s 86.10%  github.com/zmap/zlint.BenchmarkZlint.func4
         4.73s  1.61% 47.17%    118.13s 40.24%  runtime.newobject
         4.47s  1.52% 48.69%      4.47s  1.52%  runtime.memmove
    ```
    
    :tada:

[33mcommit 320c5961c17adf603bfceb77dca63752740f4527[m
Author: Daniel McCarney <cpu@letsencrypt.org>
Date:   Fri Jul 26 14:25:10 2019 -0400

    deps: update zcrypto, remove govalidator. (#294)
    
    This commit updates the `github.com/zmap/zcrypto` dependency to the
    tip of master. This allows removing the
    `github.com/asakevich/govalidator` dependency by using the new
    `zcrypto/util.IsURL` function relied on in `zlint` by `util/fqdn.go`.

[33mcommit d9a29c3ddfb5b3b498dea53607b9c395267d3807[m
Author: Daniel McCarney <cpu@letsencrypt.org>
Date:   Fri Jul 26 09:29:56 2019 -0400

    lints: add Notice level lint for EE ECDSA KeyUsages. (#293)
    
    RFC 5480 Section 3 "Key Usage Bits" indicates that end-entity certificates using
    a EC public key MAY include the digitalSignature, nonRepudiation, and
    keyAgreement Key Usages.
    
    If such a certificate contains other Key Usages the new n_ecdsa_ee_invalid_ku
    lint will return a Notice level LintResult indicating the unexpected Key Usage
    bits that were included. Depending on the adoption of a clarification document
    and respective CABF BR updates it may be possible to increase the severity of
    this LintResult in the future. See #291 for more background discussion.

[33mcommit a0632adea60b9c2f9068b07d2e0a9e0dc5039744[m
Author: Daniel McCarney <cpu@letsencrypt.org>
Date:   Fri Jul 19 21:53:09 2019 -0400

    lints: remove w_serial_number_low_entropy lint. (#292)
    
    Per zlint #270:
    
    > I believe this check does more harm than good.
    >
    >
    > A fully compliant CA which generates a random serial number from
    > exactly 64 bits of entropy will produce a serial number less than
    > 8 bytes long 1 in 256 times. That means that for every million certs
    > issued, this check will cause about 4,000 false positives.
    >
    > ...
    >
    > The only sensible way to detect low entropy is to run an analysis
    > across a large corpus of certificates. If you try to detect it on
    > a cert-by-cert basis you should at least have a much smaller minimum
    > length than 8 so there's a lower false positive rate than 1/256.
    
    This commit removes the `w_serial_number_low_entropy` lint and
    associated tests/testdata.

[33mcommit dfa3ce3b1d700fbdebe28026c67479da1cb3ef9f[m
Author: Daniel McCarney <cpu@letsencrypt.org>
Date:   Thu Jul 11 14:46:53 2019 -0400

    ci: use GOFLAGS, enforce gofmt -s. (#290)
    
    * ci: use GOFLAGS, enforce gofmt -s.
    
    This commit sets up a global `GOFLAGS` to avoid needing to repeat it for
    each command.
    
    It also changes the script to exit on the first error, and to enforce
    that all non-vendored `.go` files satisfy `gofmt -s` without diffs. This
    will make sure all code is consistently formatted and help contributors
    do the right thing by default.
    
    * ci: fix format enforcement to expect silence.

[33mcommit c65cea169ca143abb31093051f96d5fed68522e5[m
Author: tadukurow <luuk.vandenbroek@globalsign.com>
Date:   Mon Jun 10 16:42:22 2019 +0100

    lints: update/expand e_subject_contains_noninformational_value
    
    Expanded the check for no metadata only in subject DN to check for subject DN
    fields containing no characters in a-Z0-9 or outside of ascii table. This
    catches more than just checking for ".", "-", " ". Also remove separate checks
    for serial and domainComponent as they are part of pkix.Names so separate
    checking was redundant.
    
    This will help zlint catch issues that CABLint catches today such as:
    
      https://crt.sh/?id=106177929&opt=zlint,cablint,x509lint
      https://crt.sh/?id=134328239&opt=cablint,zlint
      https://crt.sh/?id=26408912&opt=cablint,zlint
    
    The new check is still susceptible to UTF8 runes metadata only.

[33mcommit 46c8a3a2f9838308e46a05b45a75b699fc264473[m
Author: TLD Update Robot <47792085+tld-update-bot@users.noreply.github.com>
Date:   Thu Jun 6 16:08:37 2019 -0400

    gtld autopull: 2019-06-06T20:03:56Z (#288)
    
    Notes removal of `.honeywell` gTLD.

[33mcommit b991e17a58f1447cb63ef87a9a113c0cafba7a89[m
Author: tadukurow <luuk.vandenbroek@globalsign.com>
Date:   Thu Jun 6 13:24:19 2019 +0100

    lints: RFC4055 RSA SPKI AlgorithmIdentifier param and tbsCertificate.signature (#286)
    
    Adds two new lints:
    
    1. `e_spki_rsa_encryption_parameter_not_null`
    2. `e_tbs_signature_rsa_encryption_parameter_not_null`
    
    The first enforces that the RSA AlgorithmIdentifier in a certificate SPKI field is correctly encoded
    (particularly with respect to the mandatory NULL parameters).
    
    The second does similarly with the tbsCertificate.Signature field.

[33mcommit 64ec0afbd7174756179985c660d7e47141528626[m
Author: Daniel McCarney <cpu@letsencrypt.org>
Date:   Tue May 21 15:19:20 2019 -0400

    CI: Update to Go 1.12. (#284)
    
    Go 1.12.x is the latest stable release at the time of writing.

[33mcommit 9047d02cf65ab5251641c4ef1568761cad06685e[m
Author: MTG <36234449+mtgag@users.noreply.github.com>
Date:   Thu May 16 18:15:41 2019 +0200

    Fix for #272 (#282)
    
    * added support for qc statements according to ETSI 319 412-5
    
    * updated date of ETSI specification
    
    * go fmt
    
    * Update util/qc_stmt.go
    
    Co-Authored-By: mtgag <36234449+mtgag@users.noreply.github.com>
    
    * Update lints/lint_qcstatem_qctype_web.go
    
    Co-Authored-By: mtgag <36234449+mtgag@users.noreply.github.com>
    
    * removed cryptosource GmbH copyright
    
    * changed oi to oid and Wo to Without
    
    * deleted
    
    * fixes capitalization
    
    * typo in last commit
    
    * missing e
    
    * added fix for #272
    
    * merged changes

[33mcommit 4d94f5800b73944b340e8266580538b136057aa6[m
Author: Jaime Hablutzel <hablutzel1@gmail.com>
Date:   Thu May 2 10:16:24 2019 -0500

    Making lint apply to subscriber certificates only. (#281)

[33mcommit c46893cb03d22258e44d276d8f8ce59037df4cee[m
Author: TLD Update Robot <47792085+tld-update-bot@users.noreply.github.com>
Date:   Wed Apr 10 12:54:01 2019 -0400

    autopull: 2019-04-10T16:49:43Z (#280)

[33mcommit f13105e53ee699f027a0751ea0d25964776948e7[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Tue Apr 9 13:15:24 2019 -0400

    lints: count embedded SCTs for Apple CT policy. (#278)
    
    A new `ct_sct_policy_count_unsatisfied` lint is added that checks if
    subscriber certificates issued after October 15th 2018 have embedded
    SCTs from a sufficient number of unique CT logs to meet Apple's CT log
    policy[0].
    
    The number of required SCTs from different logs is calculated based on the
    Certificate's lifetime. If the number of required SCTs are not embedded in
    the certificate a Notice level LintResult is returned.
    
    | Certificate lifetime | # of SCTs from separate logs |
    -------------------------------------------------------
    | Less than 15 months  | 2                            |
    | 15 to 27 months      | 3                            |
    | 27 to 39 months      | 4                            |
    | More than 39 months  | 5                            |
     ------------------------------------------------------
    
    Important note 1: We can't know whether additional SCTs were presented
    alongside the certificate via OCSP stapling. The new linter assumes only
    embedded SCTs are used and ignores the portion of the Apple policy
    related to SCTs delivered via OCSP. This is one limitation that
    restricts the linter's findings to Notice level. See more background
    discussion in Issue 226[1].
    
    Important note 2: The linter doesn't maintain a list of Apple's trusted
    logs. The SCTs embedded in the certificate may not be from log's Apple
    actually trusts. Similarly the embedded SCT signatures are not validated
    in any way.
    
    [0]: https://support.apple.com/en-us/HT205280
    [1]: https://github.com/zmap/zlint/issues/226

[33mcommit 48eabaf3670895c84fe36e90ebdb092db69669de[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Sun Mar 31 17:14:50 2019 -0400

    lints: fix bug in lint_qcstatem_qclimitvalue_valid.go (#276)
    
    * Revert "temporarily pulling qcstatement_crash_zlint.crt lint until we can troubleshoot why it's causing Zlint to crash (#273)"
    
    This reverts commit 81c75536456cffe04d2af9c51ef0eaf3c6cd60d5.
    
    * lints: fix bug in lint_qcstatem_qclimitvalue_valid.go
    
    The `e_qcstatem_qclimitvalue_valid` linter was blinding performing
    a type cast that could fail in some real-world conditions. This fix
    checks if the type cast fails and returns an error lint result instead
    of panicing.
    
    A small test cases is added that lints the real-world certificate that
    triggered the bug. Before applying the fix in this branch the test
    fails as expected:
    ```
    $ go test -v --test.run TestQcStatemQcLimitValueValid ./lints/...
    === RUN   TestQcStatemQcLimitValueValid
    --- FAIL: TestQcStatemQcLimitValueValid (0.00s)
    panic: interface conversion: util.EtsiQcStmtIf is util.EtsiQcSscd, not util.EtsiQcLimitValue [recovered]
            panic: interface conversion: util.EtsiQcStmtIf is util.EtsiQcSscd, not util.EtsiQcLimitValue
    
    goroutine 5 [running]:
    testing.tRunner.func1(0xc00012a600)
            /usr/lib/go/src/testing/testing.go:830 +0x388
    panic(0x7c9040, 0xc0002dea80)
            /usr/lib/go/src/runtime/panic.go:522 +0x1b5
    github.com/zmap/zlint/lints.(*qcStatemQcLimitValueValid).Execute(0xc3f450, 0xc000272a80, 0xc0002b0101)
            /home/daniel/go/src/github.com/zmap/zlint/lints/lint_qcstatem_qclimitvalue_valid.go:60 +0x59a
    github.com/zmap/zlint/lints.(*Lint).Execute(0xc0002aac00, 0xc000272a80, 0x85b992)
            /home/daniel/go/src/github.com/zmap/zlint/lints/base.go:114 +0x94
    github.com/zmap/zlint/lints.TestQcStatemQcLimitValueValid(0xc00012a600)
            /home/daniel/go/src/github.com/zmap/zlint/lints/lint_qcstatem_qclimitvalue_valid_test.go:25 +0x1e8
    testing.tRunner(0xc00012a600, 0x878ef0)
            /usr/lib/go/src/testing/testing.go:865 +0xc0
    created by testing.(*T).Run
            /usr/lib/go/src/testing/testing.go:916 +0x357
    FAIL    github.com/zmap/zlint/lints     0.008s
    ```
    
    Afterwards, it passes :tada:
    
    ```
    $ go test -v --test.run TestQcStatemQcLimitValueValid ./lints/...
    === RUN   TestQcStatemQcLimitValueValid
    --- PASS: TestQcStatemQcLimitValueValid (0.00s)
    PASS
    ok      github.com/zmap/zlint/lints     0.005s
    ```
    
    * nit: fix shuffled import order

[33mcommit 70de5ac71c59a75671bfc3ca3bb539b6be90e7b8[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Sun Mar 31 14:23:32 2019 -0400

    lints: add lint for TorServiceDescriptorHash ext. (#275)
    
    * lints: enforce .onion certs have valid cabf-TorServiceDescriptor ext.
    
    Adds a lint (`lint_ext_tor_service_descriptor.go`) that validates
    that `.onion` subjects have the correct `cabf-TorServiceDescriptor`
    extension with a well formed `TorServiceDescriptorHash` object for each
    `.onion` subject.
    
    * lints: add lint for TorServiceDescriptorHash ext.
    
    The new `lints/lint_ext_tor_service_descriptor_hash_invalid.go` lint
    validates that certificates with `.onion` subjects include a well formed
    `TorServiceDescriptor` extension with a `TorServiceDescriptorHash` for
    each onion eTLD+1 subject as expected by CAB forum Ballot 201[0].
    
    E.g. a certificate with three onion subjects (`a.example.onion`,
    `b.example.onion`, `c.other-example.onion`) should have *two*
    `TorServiceDescriptorHash` entries in the `TorServiceDescriptor`
    extension:
    
    * One with URI `https://example.onion`
    * One with URI `https://other-example.onion`
    
    Only the eTLD+1 is relevant for Onion sites. Subdomains are resolved by
    the hidden server based on the HTTP layer `Host` header.
    
    The new linter will return a fail result if:
    
    * There is no `TorServiceDescriptor` extension present.
    * There were no `TorServiceDescriptors` parsed by zcrypto
    * There are `TorServiceDescriptorHash` entries with an invalid Onion
      URL (unparseable, missing hostname, non-HTTPS protocol, etc).
    * There are `TorServiceDescriptorHash` entries with an unknown hash
      algorithm or incorrect hash bit length.
    * There is a `TorServiceDescriptorHash` entry that doesn't correspond to
      an onion subject in the cert.
    * There is an onion subject in the cert that doesn't correspond to
      a `TorServiceDescriptorHash`.
    
    [0]: https://cabforum.org/2017/06/08/2427/
    
    * lints: typo fix for onion service desc comment

[33mcommit 81c75536456cffe04d2af9c51ef0eaf3c6cd60d5[m
Author: Zakir Durumeric <zakird@gmail.com>
Date:   Fri Mar 29 21:24:41 2019 -0400

    temporarily pulling qcstatement_crash_zlint.crt lint until we can troubleshoot why it's causing Zlint to crash (#273)

[33mcommit d326a8ac0b1def266078a292c87b581b11571f66[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Sat Mar 23 23:23:38 2019 -0400

    lints: fix comment typos for IPv6 arpa zone. (#271)

[33mcommit 50895a56a02b2e6e6f79768a2fa7a98b825e4366[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Tue Mar 12 12:35:14 2019 -0400

    lints: clarify dsa_improper_modulus description (#269)

[33mcommit 37704a2cd136bddc46feea6bd2b2ac70f6be4d45[m
Author: Zakir Durumeric <zakird@gmail.com>
Date:   Tue Mar 12 09:24:30 2019 -0400

    test for bad DSA modulus (#268)
    
    Adds a test case with an improper DSA modulus/divisor size for the
    `e_dsa_improper_modulus_or_divisor_size` lint.

[33mcommit 503aaf6c7edc0d9203de55138576be96b63b78ca[m
Author: Zakir Durumeric <zakird@gmail.com>
Date:   Mon Mar 11 05:41:34 2019 -0700

    test for lint_ext_authority_key_identifier_no_key_identifier (#267)
    
    Add test coverage for the `e_ext_authority_key_identifier_no_key_identifier` lint.

[33mcommit 9c9f067b800e54e0541b3ceccb071399a07f0edc[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Thu Mar 7 17:36:32 2019 -0500

    Add Go module support, cleanup Makefile (#266)
    
    * chore: ignore zlint-gtld-update build artifacts
    
    * chore: use lowercase import paths
    
    * project: add Go modules support.
    
    * project: clean up makefile, use go modules
    
    * chore: reformat travis.yml.
    
    * chore: remove vendor/ from gitignore
    
    * project: vendor go module deps

[33mcommit c8bc33bb3a1ca9a78d2bdf4c806bc2465a4efdce[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Wed Mar 6 17:15:22 2019 -0500

    lints: enforce .onion certs are EV, have <15mo expiry (#265)
    
    * lints: enforce .onion subjects are in EV certs.
    
    Adds a lint (`lint_san_dns_name_onion_not_ev_cert.go`) that validates
    that any subscriber certificates containing a `.onion` subject that were
    issued after May 1st, 2015 are EV certificates. Any non-EV certs issued
    after this date that contain a `.onion` subject should receive an
    `Error` lint result.
    
    * lints: enforce .onion cert maximum validity.
    
    Adds a lint (`lint_onion_subject_validity_time_too_large.go`) that
    validates certificates with one or more `.onion` subjects do not have
    a validity period larger than 15 months.
    
    * util: add CertificateSubjInTLD helper.
    
    * lints: update .onion lint Citations.
    
    * lints: cleanup duplicate loop in onion EV lint

[33mcommit 0f862af0a6db1b6bf2f7fffef5cb160583b93537[m
Author: MTG <36234449+mtgag@users.noreply.github.com>
Date:   Fri Mar 1 02:49:36 2019 +0100

    Adding support for linting QcStatements (#250)
    
    * added support for qc statements according to ETSI 319 412-5
    
    * updated date of ETSI specification
    
    * go fmt
    
    * Update util/qc_stmt.go
    
    Co-Authored-By: mtgag <36234449+mtgag@users.noreply.github.com>
    
    * Update lints/lint_qcstatem_qctype_web.go
    
    Co-Authored-By: mtgag <36234449+mtgag@users.noreply.github.com>
    
    * removed cryptosource GmbH copyright
    
    * changed oi to oid and Wo to Without
    
    * deleted
    
    * fixes capitalization
    
    * typo in last commit
    
    * missing e

[33mcommit edc14b276d77a052891716312a2b580e4bbb87c4[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Wed Feb 27 17:03:41 2019 -0500

    lints: add new lints for .arpa reverse DNS subjects. (#260)
    
    * lints: add new lint for .arpa reverse DNS subjects.
    
    A new `lint_subject_contains_reserved_arpa_ip.go` lint is added that
    checks that any subject with a domain ending in the suffix
    ".in-addr.arpa" or ".ip6.arpa":
    
    1) has the correct number of labels for the address class in question.
    2) specifies a reversed IP address that parses as a valid IP address.
    3) specifies a parsed IP address isn't in an IANA reserved IP range.
    
    * lints: split rDNS arpa lint into two separate lints.
    
    * nit: update copyright header year
    
    * nit: move comments to proper block

[33mcommit e73a2aa5897d9329b74c1ec2db9043f5b4657750[m
Author: Jonathan Rudenberg <jonathan@titanous.com>
Date:   Wed Feb 27 16:59:21 2019 -0500

    Update year in template copyright header (#262)

[33mcommit 7aa8fbb2ed3e7d5406ac55be0f462dca6b7ddb83[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Wed Feb 27 13:48:32 2019 -0500

    utils: clarify newLint.sh args. (#261)
    
    My first attempt at using `newLint.sh` gave me
    `lint_lint_subject_contains_reserved_arpa_ip` :-)

[33mcommit 007fb1dc6e36169c64fd7399ec132a956a5c6805[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Tue Feb 19 23:26:56 2019 -0500

    zlint-gtld-update: don't template generation date. (#256)
    
    Having `zlint-gtld-update` template a generation date stamp into the top
    of `util/gtld_map.go` doesn't add much value above-and-beyond the git
    commit date that last modified the file. Including the generation date
    makes automating `zlint-gtld-update` more complex because the automation
    must account for ignoring diffs that only change the generation date but
    leave the underlying `tldMap` the same.

[33mcommit f38bd223a43c3378b07b78a0bab51ac7006b8586[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Mon Feb 18 10:46:16 2019 -0500

    gtld_map: capture `.active` removal. (#255)
    
    IANA has revoked the `.active` gTLD as of 2019-02-17.
    
    ```
    $ curl https://www.icann.org/resources/registries/gtlds/v2/gtlds.json 2>/dev/null | \
      jq '.gTLDs | .[] | select(.gTLD=="active") | .removalDate'
    
    "2019-02-17"
    ```

[33mcommit b2aa7469fab8fcdcd14bdcee36d96448b7330fdf[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Fri Feb 15 13:30:03 2019 -0500

    gtld_map: capture .epost and .zippo removal. (#254)
    
    IANA has revoked the `.epost` and `.zippo` gTLDs as of 2019-02-15.

[33mcommit fbc0b698c5777242bf84986e8a52ebece97720d7[m
Author: Phil Porada <pgporada@users.noreply.github.com>
Date:   Thu Feb 14 00:53:18 2019 -0500

    Remove .blanco GTLD (#253)

[33mcommit bb32118ad3ab29c4d9a697aa1d8faa71c07e7500[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Mon Feb 11 11:26:21 2019 -0500

    ci: update to Go 1.11.x (#252)
    
    Previously CI was running tests under Go 1.9.x and this is no longer
    listed as a stable version. Using Go 1.11.x tests against the newest
    stable version available at the time of writing.
    
    In addition to changing the CI version one file (`util/ip.go`) requires a `go
    fmt -s` update using the Go 1.11.x toolchain in order for (`TestGoFmt`) to
    pass.
    
    Before the update:
    ```
    $ go version
    go1.11.5 linux/amd64
    $ make test
    GORACE=halt_on_error=1 go test -race ./...
    --- FAIL: TestGofmt (0.25s)
        gofmt_test.go:37: glob util/*.go not gofmt'ed
        FAIL
        FAIL  github.com/zmap/zlint 0.275s
    <snipped>
    ```
    
    After:
    ```
    $ go version
    go1.11.5 linux/amd64
    $ make test
    ok      github.com/zmap/zlint   0.191s
    <snipped>
    ```

[33mcommit a797fdc8b16c70478920913adc010e65f604610a[m
Author: Kiel C <kchris@letsencrypt.org>
Date:   Mon Feb 11 07:21:49 2019 -0800

    Two TLDs added via zlint-gtld-update. (#251)

[33mcommit 90b8be3e6248c4b1e464232ab077aadffb2d5cd3[m
Author: BJ Cardon <cardonator@users.noreply.github.com>
Date:   Wed Jan 23 17:02:32 2019 -0700

    Properly parse BMPStrings out of the ExplicitText userNotice field for length check (#244)
    
    * testing out some solutions
    
    * correctly parse a BMPString in the ExplicitText field
    
    * correctly check for the BMPString type before trying to parse the string
    
    * move parseBMPString to util/encodings
    
    * add a test PEM that has a BMPString ExplicitText
    
    * added openssl output to test cert
    
    * * add a const for the bmpString tag since Go does not currently provide one
    * add a comment pointing out that we are only looking at the raw bytes from the userNotice sequence
    
    * rename constant to better match Go asn1 library

[33mcommit b4a052e4ce8a0f4c412768cf5a5d802a1c028600[m
Author: BJ Cardon <cardonator@users.noreply.github.com>
Date:   Wed Jan 16 14:33:35 2019 -0700

    New lint to ensure DN fields only contain printable characters (#249)
    
    * add a new lint to make sure that all DN fields only contain printable characters
    
    * remove accidentally left in logs
    
    * normalize to hex
    
    * found several false positives in UTF8 strings because we were not looking at the runes as opposed to the raw bytes
    
    * added a test for UTF8 characters which triggered the lint to fail previously
    
    * fix test

[33mcommit 5b6682016f2105c4254cc086f26c49b9811edbf2[m
Author: tadukurow <luukvandenbroek@gmail.com>
Date:   Tue Jan 15 16:49:14 2019 +0000

    Add lints include/exclude flag to executable (#247)
    
    * Add lints include/exclude flag to executable
    
    * use struct{} over interface{} in maps

[33mcommit 7fc4ee7f2008b5dc979a17a45ccbaa27a69ff8c7[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Sun Dec 16 16:36:42 2018 -0500

    gtld_map: capture `.spiegel` removal. (#248)
    
    See http://www.iana.org/domains/root/db/spiegel.html

[33mcommit ad0c575cebb83abb73116a918a2832bc652ef57e[m
Author: Zach Peacock <1316813+thoom@users.noreply.github.com>
Date:   Fri Nov 23 11:12:12 2018 -0700

    Added notice if the DNSNames are duplicated in the SANS extextension. (#245)

[33mcommit 55c4aa8f8cddf2611ff0ba80491a56f6ffafe4e2[m
Author: Zach Peacock <1316813+thoom@users.noreply.github.com>
Date:   Tue Nov 13 15:13:09 2018 -0700

    Fixed template used to create new lints (#246)

[33mcommit 7ecc723be25df79f5438898f19dbaa3a8f2cf627[m
Author: BJ Cardon <cardonator@users.noreply.github.com>
Date:   Fri Nov 9 11:26:13 2018 -0700

    e_subject_common_name_not_from_san incorrectly fails on case sensitive match (#243)
    
    * allow cn and san to have mixed (not matching) case
    test with no pem yet
    
    * add a test PEM to catch the case
    
    * Update lints/lint_subject_common_name_not_from_san.go
    
    Co-Authored-By: cardonator <cardonator@users.noreply.github.com>

[33mcommit 709893ce67d2dc2efe09c58f6ed8e94ee1b374cf[m
Author: BJ Cardon <cardonator@users.noreply.github.com>
Date:   Sat Oct 20 14:45:41 2018 -0600

    lowercase effective domain label when checking if TLD is legitimate (#241)
    
    * lowercase effective domain label when checking if TLD is legitimate
    
    * add tests for HasValidTLD utility function

[33mcommit 61cf5c4e58bd262b241dad09b770b03348f6b561[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Tue Oct 16 19:21:43 2018 -0400

    gtld_map: capture `.statoil` removal. (#239)

[33mcommit 34b7be2e59081f4bbe6970785e021e6bf0741f2a[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Fri Sep 21 12:05:21 2018 -0400

    gtld_map: capture latest IANA removals. (#238)
    
    Removed:
    * `.goodhands`
    * `.jlc`
    * `.panerai`
    * `.vista`

[33mcommit 868b34da65defaccf9bcc04f079f9cbfeff3c678[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Tue Sep 11 17:34:09 2018 -0400

    README: Document updating generated gtld_map.go (#237)

[33mcommit e2c7d742bb02e91440ba80f02828b72faea18143[m
Author: David Adrian <davidcadrian@gmail.com>
Date:   Tue Sep 11 17:29:56 2018 -0400

    Filter existing EV lints to subscriber certs (#229)
    
    * Filter existing EV lints to subscriber certs
    
    All of the existing EV lints don't make sense for CA certificates, and
    should have been scoped to subscriber certs.
    
    * fixed tests

[33mcommit 8093f211c43679b1ade744d238a02ba1f0c07371[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Tue Sep 11 16:54:41 2018 -0400

    Give TLD data effective periods, rework generation. (#236)
    
    * Give TLD data effective periods, rework generation.
    
    A new `zlint-gtld-update` command is added which fetches ICANN/IANA TLD
    data, generating a Go map of TLD names to their effective periods. The
    `data/gtld_map.go` file is now generated using `zlint-gtld-update` using
    `go generate ./...` in the root of the project. This replaces
    `updateTLDs.sh`, `scripts/consolidate_tlds.py`, and the `data/`
    directory files.
    
    The `lint_dnsname_right_label_valid_tld.go` lint is updated to check
    that a TLD exists and was valid at the time the certificate was issued.
    Tests are added for the case where a certificate identifier referenced
    a TLD that was not yet valid, that was no longer valid, and that was
    valid at the time of issuance.
    
    * Comment getData
    
    * Handle .onion TLD correctly, add testcase.
    
    * Moving an HTTP to HTTPS

[33mcommit 848521ffb4ee6042d34faf7389e970d149a84d91[m
Author: Daniel McCarney <daniel@binaryparadox.net>
Date:   Fri Aug 31 12:22:43 2018 -0400

    util: remove updatetld.sh in favour of updateTLDs.sh (#234)
    
    There is a more robust TLD update script in the root of the project
    directory under the name 'updateTLDs.sh'.

[33mcommit 02fe9a29bbae57da0c77db7afb53734dc262b130[m
Author: Kiel C <kchris@letsencrypt.org>
Date:   Wed Aug 22 07:08:57 2018 -0700

    Update TLD map using updateTLDs.sh (#233)

[33mcommit 12b8dc0338e6261fb4ad6a623c0a4c1bc99b3dfe[m
Author: Steven <swchang10@hotmail.com>
Date:   Thu Jun 28 18:27:46 2018 -0700

    Update to tests to conform with RFC5891 (#232)
    
    * Update to tests to conform RFC5891
    
    * Add new RFC to base
    
    * Update source and reference date for new RFC
    
    * Fix failing test
    
    * Add new pem files for test

[33mcommit 9bebe5e32c2c4b27892021e8e2f0459ef3b075ab[m
Author: justinbastress <33579608+justinbastress@users.noreply.github.com>
Date:   Tue May 29 13:29:23 2018 -0400

    add test case for issue #223 (#231)

[33mcommit 50c579ea6e55a2c41adb11f8ee77ab994211d3e6[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Tue May 29 06:36:22 2018 -0500

    remove ev locality lint (#230)

[33mcommit 251516b8a38fac8140665435053dbb9921972125[m
Author: Jacob Hoffman-Andrews <github@hoffman-andrews.com>
Date:   Wed May 23 15:33:23 2018 -0700

    Run race detector in Travis. (#224)

[33mcommit 56537c7665d1cb8a6ae026f94157e1d293e1b3d2[m
Author: Zakir Durumeric <zakird@gmail.com>
Date:   Mon May 21 13:08:16 2018 -0400

    subject_multiple_rdn.go warning -> notice (#221)
    
    * subject_multiple_rdn.go warning -> notice
    
    * fixing associated test
    
    * all the tests this time
    
    * return notice
    
    * updated description

[33mcommit 1b7e944ead5d95139a90b5d2d78a3bdbc20dd529[m
Author: Zakir Durumeric <zakird@gmail.com>
Date:   Sun May 20 20:48:44 2018 -0400

    email length is 255 not 128 (#222)
    
    * email length is 255 not 128
    
    * new test cert

[33mcommit f6e1d287883a5d9a100f695b87d496a711f9d389[m
Author: Tim D. Smith <github@tim-smith.us>
Date:   Fri Apr 13 06:16:02 2018 -0700

    Fix lint description (#218)
    
    The lint description was incorrectly duplicated from `e_sub_cert_eku_server_auth_client_auth_missing`.

[33mcommit 83c5c0b1fa58465a7c4567c9f19d25d623632f7e[m
Author: justinbastress <33579608+justinbastress@users.noreply.github.com>
Date:   Mon Mar 19 13:05:19 2018 -0400

    New SANOtherName test cert with CONSTRUCTED OtherName value (#214)

[33mcommit 6ae1b281ecb761f860f06a4c7e26fa74b67487fb[m
Author: justinbastress <33579608+justinbastress@users.noreply.github.com>
Date:   Mon Mar 19 11:03:56 2018 -0400

    lint-wide basicConstraints variable leads to potential race condition (#217)

[33mcommit 77f487ab11cb395f313683a183a128c9c24a6eb3[m
Author: justinbastress <33579608+justinbastress@users.noreply.github.com>
Date:   Mon Mar 19 09:57:23 2018 -0400

    Fix other shared state bugs in time lints (#216)
    
    * Don't assume that Execute is called immediately after CheckApplies on the exact same cert
    
    * fix same problem in lint_utc_time_not_in_zulu
    
    * Remove remaining lint_generalized_time / lint_utc_time shared state bugs

[33mcommit f27cb8f8ffb6d66991d553d4e401c7a6ae128af9[m
Author: justinbastress <33579608+justinbastress@users.noreply.github.com>
Date:   Mon Mar 19 09:45:00 2018 -0400

    Fix misuse of lint-global state (#215)
    
    * Don't assume that Execute is called immediately after CheckApplies on the exact same cert
    
    * fix same problem in lint_utc_time_not_in_zulu

[33mcommit 88032a5e59f98690016d7dd312a3620cecb3e2e0[m
Author: justinbastress <33579608+justinbastress@users.noreply.github.com>
Date:   Wed Mar 14 12:54:24 2018 -0400

    Rearrange copyright blocks so as to not mangle godocs (#213)
    
    * rearrange copyright blocks to not mangle godocs
    
    * update go to 1.9

[33mcommit 13cf4d349e95ba2f7bb454cf94cbd164e10e5d17[m
Author: Jonathan Rudenberg <jonathan@titanous.com>
Date:   Thu Mar 1 21:41:54 2018 -0500

    Change EV max validity to 825 days (#208)
    
    This is very slightly greater than 27 months and came into effect on
    March 17, 2017, so changing the existing lint instead of creating
    a new one.
    
    Reference EVGL: 9.4 / Ballot 193

[33mcommit d41c3b0c6541f9034787e1a8e1bc2ac34a105737[m
Author: justinbastress <33579608+justinbastress@users.noreply.github.com>
Date:   Tue Feb 27 18:26:17 2018 -0500

    Fix for 7.1.6.1 tests -- only state/locality only required in subscribers (#207)
    
    * lint_cert_policy_iv_requires_province_or_locality only applies to subscriber certificates, per 7.1.6.1 (citing 7.1.4.2.2) -- re issue #206
    
    * lint_cert_policy_iv_requires_province_or_locality only applies to subscriber certificates, per 7.1.6.1 (citing 7.1.4.2.2) -- re issue #206

[33mcommit e0298945acad69eaee526b29cef8f1cc572b94a1[m
Author: justinbastress <33579608+justinbastress@users.noreply.github.com>
Date:   Wed Feb 21 09:42:55 2018 -0500

    Remove e_sub_ca_eku_name_constraints
    
    https://github.com/zmap/zlint/pull/203
    
    Fixes #200

[33mcommit 4b48de28dcddf65f3dec379f1ee1a8f8abcac28b[m
Author: justinbastress <33579608+justinbastress@users.noreply.github.com>
Date:   Fri Feb 9 14:17:26 2018 -0500

    Remove anyPolicy check on subordinate CAs (#202)
    
    * Per #201, remove lint_sub_ca_must_not_contain_anypolicy.go; affiliated subordinate CAs *are* allowed to have the AnyPolicy (see the second part of section 7.1.6.3), and there is no way for us to verify affiliation (see definition in section 1.6.1)
    
    * remove unused test

[33mcommit aca25bbbeca43a9cbe333d479f1684a72eef0459[m
Author: justinbastress <33579608+justinbastress@users.noreply.github.com>
Date:   Tue Jan 16 10:04:40 2018 -0500

    Address issue 198 (SAN format) (#199)
    
    * FQDN utils: authority does not need to have an @; make AUthIsFQDNOrIP take an authority, not a host; add IsFQDNOrIP that does just take a host
    
    * address issue in lint_ext_san_uri_host_not_fqdn_or_ip -- catch bad URLs, check host instead of (bogus) authority; update test certs to use proper URI as mentioned in issue (test:// instead of test//)
    
    * don't fail out on URIs with no authority
    
    * add test for no-authority URI
    
    * Update GetAuthority and GetHost in fqdn.go to match rules from rfc3986, and to use net/url.Parse() where possible.
    
    * add happy-case tests for GetAuthority and GetHost for all combinations of userinfo/port/absolute path/query/fragment
    
    * add exceptional test cases for GetAuthority/GetHost

[33mcommit 80fe8eced4b4df54e7e9b70e52bc6aa273d03a54[m
Author: justinbastress <33579608+justinbastress@users.noreply.github.com>
Date:   Thu Jan 11 08:59:06 2018 -0500

    Fix 39 month effective date (#196)
    
    * fix effective date for SubCert825Days
    
    * Push 39-month effective date forward as well ('... issued after 1 July 2016...')

[33mcommit e1cfeb895232d2f18a79c572bf816f6a274548d6[m
Author: justinbastress <33579608+justinbastress@users.noreply.github.com>
Date:   Wed Jan 10 17:17:58 2018 -0500

    fix effective date for SubCert825Days (#195)

[33mcommit 5899dfa3116b1f4c9f88e6a4dab18f72e5836812[m
Author: justinbastress <33579608+justinbastress@users.noreply.github.com>
Date:   Tue Jan 9 13:56:12 2018 -0500

    Strip filename comments (#194)
    
    * strip redundant filename comments
    
    * add copyright where missing; remove some additional redundant comments

[33mcommit 41cd1cfc0f4bf3e7a62c332a61695f8460c7ba48[m
Author: justinbastress <33579608+justinbastress@users.noreply.github.com>
Date:   Tue Jan 9 13:20:04 2018 -0500

    Issue #191: 825 day certificates (#193)
    
    * Issue #191: Update existing lint_sub_cert_valid_time_too_long to implement the language of https://cabforum.org/2017/03/17/ballot-193-825-day-certificate-lifetimes/ -- namely, subscriber certificates issued after 2018/03/01 must have a validity period no longer than 825 days (and those issued after 2016/07/01 retain the 39-month requirement). Also updated the filename in the header comment, and added a vacuous success for certificates issued prior to 2016/07/01.
    
    * gofmt lint
    
    * add new lint for 825-day validity window; update text of CABF-BR 6.3.2 in subCertValidTimeTooLong
    
    * sub_cert_valid_time_too_long -> sub_cert_valid_time_longer_than_39_months
    
    * fix name of 39 months lint test
    
    * Add tests for 825 day limit: > 825 days, > cutoff -> fail; <= 825 days, > cutoff -> pass; < cutoff -> not effective
    
    * gofmt test
    
    * shortening description to what's tested in lint
    
    * updating description to describe specific lint

[33mcommit 56d810974215ff3ac0d520b10c1708963f48a10a[m
Author: David Adrian <davidcadrian@gmail.com>
Date:   Tue Jan 2 10:56:36 2018 -0500

    Add `openssl x509 -text` to "How to write tests" (#192)

[33mcommit cfe45c5d7d7f2b3a7330889022e5b5a509a869fa[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Thu Dec 28 15:38:51 2017 -0600

    check if CN is IP address before testing for DNSName lints (#190)

[33mcommit 55d1dccbf9f6b8922e098320a4f81bb7da07c1f3[m
Author: Maciej Galkowski <maciejgalkowski@gmail.com>
Date:   Fri Dec 8 17:21:16 2017 +0000

    Add new lints for RFC 5280, fix existing ones (#184)
    
    * Add new lints for RFC 5280, fix existing ones
    
    * Add util.IsEmptyASN1Sequence func
    
    * Fix compilation error

[33mcommit 860a701c9b8c1e9648f3bec85531ed09101cebee[m
Author: Maciej Galkowski <maciejgalkowski@gmail.com>
Date:   Thu Dec 7 18:46:28 2017 +0000

    Add length check linters for the remaining subject DN fields (#185)
    
    * Add length check linters for the remaining subject DN fields
    
    * Fix typo

[33mcommit 6c4a75f84e594d40a1d88fb02dcd8cfb89330fe8[m
Author: Maciej Galkowski <maciejgalkowski@gmail.com>
Date:   Mon Dec 4 16:01:33 2017 +0000

    Enhance e_ext_san_contains_reserved_ip linter to check for more reserved IP ranges, reject 0.0.0.0 IP address (#182)

[33mcommit cfd5104da0befc1651518b1bc3b4ddbe45b55f59[m
Author: Zakir Durumeric <zakird@gmail.com>
Date:   Wed Nov 22 09:47:00 2017 -0600

    adding a test for lint_dnsname_bad_character_in_label

[33mcommit e7e179d93c7c5a991adb9486fd713271f9202086[m
Author: Rob Stradling <rob@comodo.com>
Date:   Wed Nov 22 14:09:19 2017 +0000

    Add missing [ to dnsNameRegexp. (#181)

[33mcommit 36a643fb3d1d4e92d7de5056003fda59676d18ea[m
Author: Maciej Galkowski <maciejgalkowski@gmail.com>
Date:   Mon Nov 20 14:38:18 2017 +0000

    Speed up the linters (#179)
    
    * Speed up the linters
    
    * Changes requested in the code review

[33mcommit 7057a53b1d1daf9ee70547255cf2fa516ff50c3c[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Fri Sep 15 11:37:52 2017 -0500

    Update dnsName to apply to correct EKU values (#176)
    
    * update all dnsName lints with serverAuth only
    
    * prepend openssl output

[33mcommit 7b41c239ca65ff0b6df2e226b6b19884960c1b2b[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Fri Sep 15 11:17:55 2017 -0500

    Fix san bare suffix, ca CN effective time, subCA EKU  (#174)
    
    * fix san bare suffix bug
    
    * gofmt

[33mcommit 88cddc638874e5f8b1e9899d95a68001ec35be8a[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Thu Sep 7 18:33:59 2017 -0500

    Add bare IANA Suffix lint (#172)
    
    * add bare iana suffix lint
    
    * change naming
    
    * add license
    
    * y

[33mcommit 5826a245d6a559e51c77144f7aed099189bb1ea5[m
Author: David Adrian <davidcadrian@gmail.com>
Date:   Thu Sep 7 11:41:15 2017 -0400

    Update new lint template
    
    Fill in source, don't add useless comment, add license

[33mcommit 2c8be8ad1a769ec3544da976aff3a0310f8b0628[m
Author: Jonathan Rudenberg <jonathan@titanous.com>
Date:   Thu Sep 7 11:23:55 2017 -0400

    Fix fatal errors for certificates with public suffixes (#156)
    
    * Fix fatal errors for certificates with public suffixes
    
    - Use the ICANN section of the PSL for all relevant lints, as we
    want to lint the non-ICANN labels.
    - Return NA instead of Fatal for PSL failures, the TLD lint will
    catch these if they are valid.
    
    Fixes #155
    
    Signed-off-by: Jonathan Rudenberg <jonathan@titanous.com>
    
    * update lints to fatal instead of NA
    
    * update lints
    
    * remove unused import
    
    * scope DNS lints to BRs
    
    * update to NA

[33mcommit 3081e3ed02034d502c3c4d295ccafc98a03ca427[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Wed Sep 6 16:52:36 2017 -0500

    Add question mark notice lint (#170)

[33mcommit f4040dac46c3d36f9178aa9901226abe92d7ad0e[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Wed Sep 6 14:25:56 2017 -0500

    Fix eku check for BRs (#171)
    
    * fix eku check for BRs
    
    * fix nits
    
    * update name

[33mcommit 27b67e2f61e20b09c4753633504bc0e316d6b391[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Wed Sep 6 11:08:09 2017 -0500

    Add type to lints (#169)
    
    * adding type to lints
    
    * BRs -> CABFBaselineRequirements
    
    * update names
    
    * ReadableSource -> Citation
    
    * rename to UnknownLintSource
    
    * fix README spacing
    
    * Fix nits
    
    * update base.go to remove Source from JSON
    
    * fix test
    
    * address JSON

[33mcommit 9e196af38011ae3292ac4f451ddaa6cb9e8a1385[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Tue Sep 5 10:45:36 2017 -0500

    remove old subject lints from 1.4.0 (#168)

[33mcommit 8e4deb3844d7d076e899539572e117d69b83c877[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Tue Sep 5 08:50:14 2017 -0500

    fix international names length check (#164)

[33mcommit d3378a58946b0093480dc0c9fef6aa3e6aedb70d[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Sat Sep 2 20:25:02 2017 -0500

    Set correct effective date for givenName lints (#162)

[33mcommit 66858570887619f9a5f481beb2a8fb244fce97e4[m
Author: David Adrian <davidcadrian@gmail.com>
Date:   Fri Sep 1 15:53:48 2017 -0400

    Add --list-lints-schema (#160)
    
    Output lints as ZSchema, sorted by name.

[33mcommit a2f68c6cbdd32f94b8e65306fe6a69f957fa30b0[m
Author: David Adrian <davidcadrian@gmail.com>
Date:   Fri Sep 1 14:44:43 2017 -0400

    Add licenses header to some files.
    
    Start adding license header until I got tired of it. Will do more later.

[33mcommit b7b0924fd2c7d38f9937cbb434fbfe73502135ff[m
Author: David Adrian <davidcadrian@gmail.com>
Date:   Fri Sep 1 11:21:59 2017 -0400

    Prevent panic() in example code in README

[33mcommit a23cd954a4d64f4f7f0524d94ddc90fdcb758aa5[m
Author: David Adrian <davidcadrian@gmail.com>
Date:   Fri Sep 1 11:20:45 2017 -0400

    Fix comment grammar

[33mcommit 7d1dc46016c7567abfa709343eef053bee5ae6f0[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Fri Sep 1 10:13:15 2017 -0500

    Address ambiguity in BRs (#158)
    
    * address ambiguity in BRs
    
    * add comment

[33mcommit f94d2a7c4a96821d692c492c631e35213eec56cd[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Thu Aug 31 23:08:02 2017 -0500

    Removing duplicate lints (#154)

[33mcommit 53c9753c5932714928e0141edeea1cc3de2188d3[m
Author: David Adrian <davidcadrian@gmail.com>
Date:   Thu Aug 31 23:35:50 2017 -0400

    Refactor structure to simplify types and names (#153)
    
    Add comments explaining the interfaces. Remove extraneous error
    returns. Fix test formatting.

[33mcommit d3c6ab2a20d7d252eaf6695e8d5a2321be726fd4[m
Author: Jonathan Rudenberg <jonathan@titanous.com>
Date:   Thu Aug 31 13:40:32 2017 -0400

    Fix some typos (#152)
    
    Signed-off-by: Jonathan Rudenberg <jonathan@titanous.com>

[33mcommit 806ee207c2e54cce56d07fd222e8a305d211b772[m
Author: Jonathan Rudenberg <jonathan@titanous.com>
Date:   Thu Aug 31 13:35:39 2017 -0400

    Remove useless comments (#151)
    
    Signed-off-by: Jonathan Rudenberg <jonathan@titanous.com>

[33mcommit b13f74b4631e585f321575783bde4ea3aaae175b[m
Author: Jonathan Rudenberg <jonathan@titanous.com>
Date:   Thu Aug 31 13:22:21 2017 -0400

    Change Golang -> Go (#149)
    
    Signed-off-by: Jonathan Rudenberg <jonathan@titanous.com>

[33mcommit 4373a48e3b5735a3493100358c9ba8b3be3f1304[m
Author: Alex Gaynor <alex.gaynor@gmail.com>
Date:   Thu Aug 31 13:07:35 2017 -0400

    Apply SHA1 test to DSA and ECDSA certs as well. (#150)

[33mcommit 6fb7b154abd1f7b042206894b3181636d63349c0[m
Author: Rich Salz <rsalz@akamai.com>
Date:   Thu Aug 31 12:47:28 2017 -0400

    Replace last instances of provenance with source (#148)

[33mcommit a7fda0c5a17152ea85457471cc7403e9aaac66ee[m
Author: Rich Salz <rsalz@akamai.com>
Date:   Thu Aug 31 12:40:33 2017 -0400

    Improve help message (#147)

[33mcommit 311c8541dc9c9774473ebb88659a7cadb7539d1f[m
Author: Zakir Durumeric <zakird@gmail.com>
Date:   Thu Aug 31 11:56:14 2017 -0400

    syntax highlight code example in README

[33mcommit 6ad5cc524fe90e390b421473bd30a137dbe14b1c[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Wed Aug 30 23:26:42 2017 -0500

    Add DSA unique representation lint (#142)

[33mcommit 383696bee1b56f8509369b8bd566928ea75adedb[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Wed Aug 30 23:11:54 2017 -0500

    Add DSA Correct subgroup lint  (#143)
    
    Clean up other DSA lints

[33mcommit 0ea55e3be77431cf77661498607e09e8953dc5e7[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Wed Aug 30 21:27:16 2017 -0500

    add sub cert validity check to 39 months (#146)

[33mcommit 9d433a47f773ad461320befd59247b513def5167[m
Author: Zakir Durumeric <zakird@gmail.com>
Date:   Wed Aug 30 14:45:27 2017 -0400

    changing provenance to source (#145)
    
    * changing provenance to source
    
    * providenct to source

[33mcommit 18a30b3365d737152eb9389736a861336654ead7[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Wed Aug 30 12:52:57 2017 -0500

    Update README on how to add a new lint (#144)
    
    * update README for creating a lint
    
    * update
    
    * make readme better
    
    * some backticks
    
    * Update README.md

[33mcommit 9162bd940e15bd045d2d716752d3f771d5ed94be[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Tue Aug 29 22:17:38 2017 -0500

    add Is cA check (#139)
    
    * add isCa check
    
    * add isCA check on raw extension
    
    * add std asn1 lib
    
    * remove basicConstraints declaration

[33mcommit 019b8f30d0aefd59ccad9cb900f8b6589370c490[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Tue Aug 29 19:04:02 2017 -0500

    Subscriber certs should not have is cA field set. (#140)
    
    * stuff
    
    * check actual extension for is cA field
    
    * change to default go asn1

[33mcommit be08bd22bfbea3a670a46d62603a16019aa2c669[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Tue Aug 29 18:54:55 2017 -0500

    Cleanup name constraints check (#141)
    
    * name constraints check cleanup
    
    * update pem with openssl

[33mcommit 32224f1959a8685ec1eec0c41cfcc959941518ae[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Tue Aug 29 12:45:22 2017 -0500

    Align the text of BR lints to match what we have in the public spreadsheet. (#138)
    
    * normalizing text with spreadsheet
    
    * update to match spreadsheet
    
    * gofmt

[33mcommit f5ee2167043aa0c2959754a62775f1f8d8762812[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Tue Aug 29 12:13:28 2017 -0500

    Remove san dnsname FQDN lint (#137)

[33mcommit 6188084fdb716cae13c8547ffa915d11db507812[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Tue Aug 29 09:58:05 2017 -0500

    panic when initialize throws an error (#135)
    
    * panic when initialize throws an error
    
    * add error
    
    * make error idiomatic

[33mcommit e67e60db89424af511bd25a70b4b06d07d8104df[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Tue Aug 29 09:52:42 2017 -0500

    Add bad character in label check (#134)
    
    * bad character in label lint
    
    * move regexp to init
    
    * update function signature
    
    * address zakir comments
    
    * update compile to MustCompile
    
    * add CheckApplies

[33mcommit 20cd430294359a6ce5fd5b871d47903d66a6f498[m
Author: Rich Salz <rsalz@akamai.com>
Date:   Mon Aug 28 17:38:13 2017 -0400

    Take multiple input files; intuit filetype (#132)
    
    If file ends with .der or .pem, then use that as the filetype.
    Otherwise can specify it with the -format flag.

[33mcommit dd4c0f1e4330953aa90c8c066be8b905b8ddc520[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Mon Aug 28 16:32:57 2017 -0500

    Add label length check for DNSNames. (#133)
    
    * add label too long check
    
    * add openssl output

[33mcommit fbdbc89d7d32082d84c232664caddd40c0df1388[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Mon Aug 28 16:00:22 2017 -0500

    Update TLD Script (#116)
    
    * update tld script
    
    * add check for valid tld
    
    * remove SAN extension check
    
    * update to subscriber cert
    
    * add existent check
    
    * add proper cn check

[33mcommit d50975d6892f5b28258b9d72ba8fd49674ac3a40[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Mon Aug 28 15:57:37 2017 -0500

    add empty label dnsname check (#117)
    
    * add empty label dnsname check
    
    * update to subscriber cert
    
    * add existent check
    
    * proper cn checking

[33mcommit 574726c5e32c660a4fbbcd3700488307b731bfff[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Mon Aug 28 15:55:57 2017 -0500

    adding SLD hyphen check (#119)
    
    * adding SLD hyphen check
    
    * remove SAN checkApplies
    
    * return ResultStruct
    
    * add existent check
    
    * add proper error checking
    
    * change name of func
    
    * rename again

[33mcommit 27c2a19c47e8ada659b2e7188d174d02a70ae91b[m
Author: Zakir Durumeric <zakird@gmail.com>
Date:   Mon Aug 28 16:38:23 2017 -0400

    Start a more helpful readme file (#131)

[33mcommit f80e205f38fce6e98c258a7c9a743421c4c463a8[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Mon Aug 28 15:28:37 2017 -0500

    Check if underscore in SLD (#118)
    
    * check for underscore in SLD
    
    * underscore lint
    
    * remove SAN checkApplies
    
    * add result
    
    * update to fatal
    
    * add existent check
    
    * add correct error checking

[33mcommit 8259bf877a6989aae8e2ac9d6ee662e78b11ae43[m
Author: Rich Salz <rsalz@akamai.com>
Date:   Mon Aug 28 16:27:04 2017 -0400

    Add -pretty flag (#130)

[33mcommit 8a96225c4ffd03df69a52d674ead3b098bc6689c[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Mon Aug 28 15:20:57 2017 -0500

    adding underscore in left of ETLD+1 check (#120)
    
    * adding underscore in TRD check
    
    * remove SAN checkApplies
    
    * apply to subscriber certs
    
    * return result in addition to bool
    
    * update to fatal
    
    * add existent check
    
    * add proper error checking

[33mcommit f718436a728255f134c80e2f09b98afc4fa3871b[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Mon Aug 28 14:59:05 2017 -0500

    Add test for wildcard left of ETLD (#127)
    
    * add test for left of public suffix
    
    * update helper to return Result in addition to bool
    
    * publicsuffix cannot parse is NA
    
    *  public suffix
    
    * update to fatal
    
    * add check for existent DNSNames
    
    * make helper not domain specific
    
    * remove unnecessary restriction
    
    * actually check if CN is empty

[33mcommit 8f5859ea6ac14b7abfbe5d1f0bfd9ff38ca01cea[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Mon Aug 28 14:28:50 2017 -0500

    Update helper to use the right lib. (#129)
    
    * update helper
    
    * gofmt

[33mcommit 119f503afd92adc6adba23e2167ced0531655b1c[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Mon Aug 28 14:19:36 2017 -0500

    Add helper function to determine if there are names to check.  (#128)
    
    * add helper function
    
    * gofmt

[33mcommit 29bd40e7adbeef3fca2e40238673c07a83aedee2[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Sun Aug 27 09:32:55 2017 -0500

    Add wildcard only in left label check (#121)
    
    * check wildcard only in left label
    
    * fix unsafe array index
    
    * add immediate return

[33mcommit d6de6b1c9f6c5c0f85e1b455700efe324f73ada1[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Sun Aug 27 09:29:51 2017 -0500

    Add left label wildcard check (#122)
    
    * left label wildcard check
    
    * fix bug to check the correct value
    
    * squash and clause
    
    * address titanous

[33mcommit 53bb5bfe293f6a5c768da2f728348e578cd80a68[m
Author: Alex Gaynor <alex.gaynor@gmail.com>
Date:   Sat Aug 26 23:58:40 2017 -0400

    Simplify removeQuestionMarks implementation (#126)

[33mcommit 4aa2e44bca2c12258b258472684bde865210c1a7[m
Author: David Adrian <davidcadrian@gmail.com>
Date:   Sat Aug 26 15:51:26 2017 -0400

    Add Slack notifications to TravisCI

[33mcommit e77dd5989af05ee75158d300c3d16cf86d86c001[m
Author: David Adrian <davidcadrian@gmail.com>
Date:   Sat Aug 26 15:42:11 2017 -0400

    Remove nested zlint package structure (#125)
    
    This moves main.go to a cmd/ subpackage, and moves zlint/zlint to the
    root package. It further simplifies main.go to operate only on single
    certificates. Multiple certificate linting should be accomplished using
    the ZCertificate utility.

[33mcommit d35efad6d164f53162c098fdd10fba34aaa2ca23[m
Author: Zakir Durumeric <zakird@gmail.com>
Date:   Sat Aug 26 11:57:18 2017 -0400

    a couple of other fixes from providence to provenance (#124)

[33mcommit 77bb27c31feb35e6e5365b411d48b4741980b4c4[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Fri Aug 25 10:03:29 2017 -0500

    Add malformed unicode for IDN check (#115)
    
    * malformed unicode
    
    * update description

[33mcommit c89c81c0d1494a2b1722d454fcfbc71ef920e99b[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Fri Aug 25 09:42:35 2017 -0500

    Check Internationalized DNSNames for NFKC (#114)
    
    * check if DNSNames are NFKC if internationalized
    
    * remove extra comment
    
    * gofmt
    
    * fix dadrian nits
    
    * new lint update
    
    * NA instead of Fatal when ToUnicode fails
    
    * remove comment

[33mcommit fd09e89b5d3f9cfb2fa136d1dd7cefa1c3559e79[m
Merge: c082259e 5db497f4
Author: David Adrian <davidcadrian@gmail.com>
Date:   Thu Aug 24 20:00:29 2017 -0400

    Merge branch 'titanous-simplify-report-generation'

[33mcommit 5db497f4470f6efa8ceeb4e653bf93268b087963[m
Merge: c082259e 3d4a5ba0
Author: David Adrian <davidcadrian@gmail.com>
Date:   Thu Aug 24 20:00:11 2017 -0400

    Merge branch 'simplify-report-generation' of https://github.com/titanous/zlint into titanous-simplify-report-generation

[33mcommit c082259e99a4278baa66c5c7b9657f34ed94b808[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Thu Aug 24 16:25:30 2017 -0500

    add serial number low entropy check (#112)
    
    * add low entropy check
    
    * update to warning
    
    * fix tests

[33mcommit 875c2d70ad894b23948ea7931e5622b28a94e69a[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Thu Aug 24 11:48:40 2017 -0500

    add subCA anypolicy check (#111)
    
    * add subCA anypolicy check
    
    * add subCA cert fix
    
    * remove second cert

[33mcommit 9c9a02cdf3f2b39f680c9ead3ec987193ba03da1[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Wed Aug 23 23:52:24 2017 -0500

    remove duplicate lints (#93)

[33mcommit 8ae4ade6f22a26b963690a3cc5328ccb4bb664b3[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Wed Aug 23 23:34:29 2017 -0500

    add countryName required check (#110)

[33mcommit 23045185e92ca99cde678d7ad8c706677f3c11a3[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Wed Aug 23 23:27:59 2017 -0500

    add postal code prohibited lint (#109)

[33mcommit 0c1c125b9c85ae8a37cebf189d286e6e3a94d681[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Wed Aug 23 23:04:59 2017 -0500

    Add province prohibited check (#108)
    
    * locality
    
    * province prohibited
    
    * fix conflict

[33mcommit 65a9b1c6a323df056a72913c900be0bacc9a3c23[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Wed Aug 23 22:57:27 2017 -0500

    Add province required check. (#107)
    
    * locality
    
    * update with province prohibited lint

[33mcommit 57042a7abb8f74be1316c79f40f7cde905005e1d[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Wed Aug 23 22:54:32 2017 -0500

    Add Locality name prohibited check (#106)
    
    * locality
    
    * locality name lint update
    
    * add locality prohibited lint

[33mcommit 07803469fb6df822f438e025e4289289ab80d7d6[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Wed Aug 23 22:16:38 2017 -0500

    Add locality name check (#105)
    
    * add locality name test
    
    * remove unnecessary comment
    
    * remove .

[33mcommit 57b4ca48e37957baa8220fafe8f8871f6c68c73c[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Wed Aug 23 18:34:40 2017 -0500

    add givenName and surname policy check (#103)
    
    * add givenName and surname policy check
    
    * given name test fix

[33mcommit 2f1a3c601e88277149d890e21207362dae0feb96[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Wed Aug 23 18:33:15 2017 -0500

    add AIA lint for subcert (#100)

[33mcommit 6ec6f1afe77b6310e377cdf4318f93616c7b2fe2[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Wed Aug 23 18:30:52 2017 -0500

    Check if streetAddress MUST NOT appear. (#104)
    
    * add streetAddress lint
    
    * reorder names
    
    * remove bad comment
    
    * build

[33mcommit 01b70c8eb4b384af7db781f7fdde57355ada935d[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Wed Aug 23 18:26:43 2017 -0500

    add subCA EKU missing lint (#102)

[33mcommit 6aae9006637ed745e65a7acedbad2cc1cdfbaa60[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Wed Aug 23 18:23:41 2017 -0500

    add common name missing lint (#101)

[33mcommit af3bdb17d9da9ba2c3fbcdc1e6065884f984c8f8[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Wed Aug 23 15:48:39 2017 -0500

    Add lints that are inbound to base.go to prevent merge conflicts (#99)
    
    * add lints tbd
    
    * gofmt

[33mcommit b8ccba0ed460ae2fa2720764859ce2482a0dd45c[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Wed Aug 23 09:07:31 2017 -0500

    check if aia exists before operating on it (#98)

[33mcommit d3ce2dc037ab70e8d3e07cb07cc0f79ee048727c[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Wed Aug 23 09:01:48 2017 -0500

    check if keyUsage exists before operating on it (#97)

[33mcommit 451071820f73e5db293150b6c530b4f91271885c[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Wed Aug 23 08:56:24 2017 -0500

    subCA EKU Valid Fields (#95)
    
    * update lint with tests
    
    * add tests
    
    * CAB -> BRs
    
    * change to notice
    
    * update with proper text
    
    * update name to not technically constrained
    
    * update reportStruct with correct name
    
    * update lint and test

[33mcommit 234a47481b9c821b5cde77e8fdfe6957e51225ff[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Tue Aug 22 13:18:21 2017 -0500

    Check if HTTP URL is inside AIA/CRL URLs. (#96)
    
    * actually check if an http url is inside ocsp, crl urls
    
    * hasPrefix instead of contains

[33mcommit b60d4bfd02b0347ba710bd59a46290153cf42473[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Tue Aug 22 10:03:17 2017 -0500

    SubCA AIA Marked Critical Lint (#94)

[33mcommit 09f16fb8a7d1a0ce78326623774b7ede9b8bddcf[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Mon Aug 21 21:08:21 2017 -0500

    RootCA Key Usage Critical Lint (#92)
    
    * not critical
    
    * update lints
    
    * fix broken test with not critial cert
    
    * remove fmt
    
    * remove comment

[33mcommit 020e8ba5126bd1458fd28a920cf5070f4435fa55[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Mon Aug 21 15:37:17 2017 -0500

    adding rootCA key usage present lint (#91)

[33mcommit 0b36404b8cdd3d194cc8b5e84cfcfcfcc9b18a50[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Mon Aug 21 15:05:15 2017 -0500

    Adding signature not supported lint (#90)
    
    * add algorithm not supported
    
    * update test with sha1withRSA test
    
    * update pems with stuff

[33mcommit 770b40799509a4aabfbd2f56b26bbea717521a25[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Mon Aug 21 13:42:52 2017 -0500

    These key identifier checks only apply to Subscriber Certificates (#88)
    
    * update two lints with correct applies clause
    
    * remove swp forever
    
    * update to not root CA and remove duplicate lint
    
    * fixing tests

[33mcommit 3528ed3db31839496c31baa2b4f1ffab2dc34c70[m
Author: Jonathan Rudenberg <jonathan@titanous.com>
Date:   Sun Aug 20 19:24:29 2017 -0400

    Fix 1024-bit RSA sunset dates (#80)
    
    Signed-off-by: Jonathan Rudenberg <jonathan@titanous.com>

[33mcommit 528096a36c9e7348e2b637c52205bcc670a8821d[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Sun Aug 20 18:15:58 2017 -0500

    Update subcert AIA lint prefix to "w" as it's actually a warning. (#86)
    
    * update sub cert aia lint to use proper prefix
    
    * update test
    
    * update underlying struct to use W instead of E

[33mcommit dd95a32d5bf6613426f8d98b1f5005ac116b5436[m
Author: Jonathan Rudenberg <jonathan@titanous.com>
Date:   Sun Aug 20 19:13:49 2017 -0400

    Fix report skew (#87)
    
    * Rename RDN -> rdn in lint names
    
    Signed-off-by: Jonathan Rudenberg <jonathan@titanous.com>
    
    * Remove unused EExtSanDnsSyntaxIncorrect

[33mcommit aa491a8aca7bda7c74a784213bcf71663909c265[m
Author: Jonathan Rudenberg <jonathan@titanous.com>
Date:   Sun Aug 20 19:11:51 2017 -0400

    Rename CAB -> BRs (#84)
    
    "[the] BRs" is an official short reference to our favorite document,
    The Baseline Requirements for the Issuance and Management of
    Publicly-Trusted Certificates.
    
    https://cabforum.org/pipermail/public/2017-August/011856.html
    
    Signed-off-by: Jonathan Rudenberg <jonathan@titanous.com>

[33mcommit db7435f8e6f89f06dde470249961aefc56421e96[m
Author: Jonathan Rudenberg <jonathan@titanous.com>
Date:   Sun Aug 20 19:11:08 2017 -0400

    Add LICENSE (#83)
    
    Signed-off-by: Jonathan Rudenberg <jonathan@titanous.com>

[33mcommit 49fbb4510fdd2e25b63486a933bc38f2a5646a04[m
Author: Jonathan Rudenberg <jonathan@titanous.com>
Date:   Sun Aug 20 19:09:48 2017 -0400

    gofmt -s (#81)
    
    Signed-off-by: Jonathan Rudenberg <jonathan@titanous.com>

[33mcommit e1426b4e0479d6d7e6087d48855ac4f9a35f081c[m
Author: Jonathan Rudenberg <jonathan@titanous.com>
Date:   Sun Aug 20 19:07:49 2017 -0400

    Fix typo in json tag for ENameConstraintMaximumNotAbsent (#82)
    
    Signed-off-by: Jonathan Rudenberg <jonathan@titanous.com>

[33mcommit 3d4a5ba0727486a09ccee45085e8da8d3672f5bf[m
Author: Jonathan Rudenberg <jonathan@titanous.com>
Date:   Sun Aug 20 15:41:21 2017 -0400

    Simplify lint report generation by using map
    
    Signed-off-by: Jonathan Rudenberg <jonathan@titanous.com>

[33mcommit 1f3497eb444fa211699f71b16a33a6c070aa80e2[m
Author: Jonathan Rudenberg <jonathan@titanous.com>
Date:   Sun Aug 20 15:39:42 2017 -0400

    Rename RDN -> rdn in lint names
    
    Signed-off-by: Jonathan Rudenberg <jonathan@titanous.com>

[33mcommit 7b2b2757caa3d18de37fb5eb574b19b32ccd2c14[m
Author: Jonathan Rudenberg <jonathan@titanous.com>
Date:   Sun Aug 20 14:03:07 2017 -0400

    Add `openssl x509 -text` output to test certs (#79)
    
    Signed-off-by: Jonathan Rudenberg <jonathan@titanous.com>

[33mcommit 39c9f80ce77cf420df30907b75d0782d7bda4b19[m
Author: Jonathan Rudenberg <jonathan@titanous.com>
Date:   Sun Aug 20 13:34:14 2017 -0400

    Rename Providence -> Provenance (#78)
    
    Signed-off-by: Jonathan Rudenberg <jonathan@titanous.com>

[33mcommit 57d939a687cedb4625ba110b1fef621f444dbe72[m
Author: Jonathan Rudenberg <jonathan@titanous.com>
Date:   Sun Aug 20 12:59:41 2017 -0400

    Switch AIA missing issuing CA URL from error to warn (#77)
    
    This is a SHOULD, not a MUST so it should be a warning.
    
    Signed-off-by: Jonathan Rudenberg <jonathan@titanous.com>

[33mcommit 6e5f9d572d5d6bb3e1b5c24daa881e405da62764[m
Author: David Adrian <davidcadrian@gmail.com>
Date:   Mon Jun 12 15:34:21 2017 -0400

    Clean up DNS Name Encoding Lints (#71)

[33mcommit 860b76a57d05be69e682cc8979319f61093b4f67[m
Author: David Adrian <davidcadrian@gmail.com>
Date:   Mon Jun 12 14:56:35 2017 -0400

    Make util/ Great Again (#70)
    
    - Fix CA Cert and Self-Signed and Intermediate Detection
    - Fix the tests that using the correct detection broke
    - Clean up util some to make sense

[33mcommit ccb27dab36cfca10930db51a7a679ac228db6e51[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Mon Jun 5 11:48:11 2017 -0400

    Update LintReport JSON to match Schema (#69)

[33mcommit 113adb24646c015defd9dd04ae823c3f1c0bba0a[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Fri Jun 2 12:49:32 2017 -0400

    Rewrite main.go with channels (#68)

[33mcommit 18005bc6417e59ddd659dd1ba32ca62c9b9201d0[m
Author: David Adrian <davidcadrian@gmail.com>
Date:   Wed May 31 18:05:45 2017 -0400

    Ensure Lint fields are set (#62)
    
    Test that Name, Providence, and Description are set on all Lints.

[33mcommit 193ba7d10700de382e0de6f01ed0ca0a1e580c2c[m
Author: David Adrian <davidcadrian@gmail.com>
Date:   Wed May 31 15:54:41 2017 -0400

    Add meta-lint for assigning to the correct report field (#63)
    
    This checks that the field name in ZLintReport matches Lint.name

[33mcommit 55c5af7cc87636cc0a02bc9dd44dbbc9428a7e8b[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Wed May 31 15:51:31 2017 -0400

    Use correct name for multiple RDN lints (#64)
    
    Fixes error exposed in #63

[33mcommit 2617c7e7a32f056e7bb6b86de7bdbb1b31ea8bb1[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Wed May 31 14:07:54 2017 -0400

    Add multiple RDNs lint (#59)
    
    This adds a new lint that checks for multiple RDNs in subject and issuer.

[33mcommit 089d34647a744babc05b09e84cac4d9ddf783a3a[m
Author: David Adrian <davidcadrian@gmail.com>
Date:   Wed May 31 13:42:17 2017 -0400

    Ensure lints have updateReport defined (#61)

[33mcommit 2d1ea919cc4920c5454a64a9b6553ad399268341[m
Author: zhengping <wang426@illinois.edu>
Date:   Mon Apr 17 16:33:58 2017 -0500

    Add DN attribute value space lints (#28)

[33mcommit b946258a7d1ec1ecafc0e0af70af01e6159e2d15[m
Author: Paul Murley <pmmurley830@gmail.com>
Date:   Wed May 24 12:51:45 2017 -0500

    Minor cosmetic fixes in test descriptions (#54)

[33mcommit 51d8fd000044f1cba7d2fe3a2858bd1d5d114c4d[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Tue May 30 20:12:14 2017 -0400

    Add --list-lints-json (#51)

[33mcommit 3f9b85cdc76d0bdc682566515fdb55ee86563c9e[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Fri May 26 16:26:29 2017 -0500

    update common name test to exclude CA certs (#57)
    
    * update common name test to exclude CA certs
    
    * smh gofmt

[33mcommit e94c186f143b5324d8048fe9ccc1c457f8db24bb[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Thu May 25 09:42:53 2017 -0500

    Add missing defense checks to lints (#55)

[33mcommit 961ea435888129d8977228c6780e99216f0b7306[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Wed May 24 15:45:46 2017 -0500

    Lints that only apply to subscriber certs should not apply to CA certs (#53)

[33mcommit 818f0029e9c386c49a75570b6bed4273e46df841[m
Author: Alex Holland <ajholland77@yahoo.com>
Date:   Mon May 22 17:34:33 2017 -0400

    Pedantic Fixes (#52)
    
    Adds goreportcard

[33mcommit fd5ecf2ec11d4adfd3f9fa3cd95602266b197ae7[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Sun May 21 16:59:09 2017 -0500

    ignoring error because its in the damn lint itself (#49)
    
    * ignoring error because its in the damn lint itself
    
    * update to 1.8.1
    
    * removing zcrypto
    
    * dont update twice

[33mcommit a313b0c1a68ff15362d7c6ea45a0f55583315ea3[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Sun May 21 15:29:03 2017 -0500

    removing writing to map (#48)

[33mcommit 59bf413cfb8aadc0f600780241df97c2c96898b3[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Sun May 21 15:05:47 2017 -0500

    Kumarde/update travis (#47)
    
    * removing dependency on zgrab to zcrypto
    
    * update travis

[33mcommit c36ba035744cb1fbe44c1f2b16f3c20bcb40d612[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Sun May 21 14:50:28 2017 -0500

    segfault fixing (#45)

[33mcommit a39229dd713b86021ac2cf12481e6f57a8e3a75c[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Tue May 16 11:29:44 2017 -0500

    updating base with max length lints (#44)

[33mcommit 872737683f22ece29da88c5979629b735ee5c5e6[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Fri May 12 11:41:42 2017 -0500

    Update JSON output to match ZTag schema (#43)

[33mcommit fea6847e8559d30cde77b5a58231befa7bff21e8[m
Author: jddicki2 <dickinson.joey@yahoo.com>
Date:   Thu May 11 19:59:44 2017 -0500

    max_length lints: Subject common_name, locality_name, organization_name, and state_name (#40)
    
    * Adding max_length lint for subject common name
    
    * Adding max_length lint for subject locality name
    
    * Adding max_length lint for subject state name
    
    * Adding max_length lint for subject organizational name
    
    * added organizational unit max length lint and tests
    
    * Changed lint to match style of other max_length lints

[33mcommit 6d5b9a14365da57350239328708081fa694abed0[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Wed May 10 15:33:11 2017 -0500

    Kumarde/fill out fields (#42)
    
    * updating info to notice
    
    * update base
    
    * gofmt code
    
    * update w bools in right struct
    
    * fixing nits
    
    * json removal
    
    * updating fields to the right places

[33mcommit 20fc474b7cb9e59614739c73f6085b6668a7db08[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Tue May 2 20:26:20 2017 -0500

    updating info to notice (#41)

[33mcommit 7f1506cb3dd99e62f366833c8c27e89d23afd176[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Sat Apr 29 10:20:12 2017 -0500

    fixing public suffix lint (#38)

[33mcommit 6752a0674f881e51a395a772d562ccfb3504bf5e[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Fri Apr 28 17:19:00 2017 -0500

    Kumarde/zlint refactor (#35)
    
    * update version to ZLint
    
    * gofmt
    
    * first refactor attempt
    
    * update refactor to have lint store their own state
    
    * actually setting ZLint field...
    
    * Fix fqdn update to govalidator
    
    * removing useless nil return
    
    * return nil if nil

[33mcommit 741f8e0396a9f1f73e99da6c6f90f92392605b91[m
Author: zhengping12 <wang426@illinois.edu>
Date:   Thu Apr 20 21:00:05 2017 -0500

    Added gofmt test for all source code (#33)
    
    * Added gofmt test for all source code
    
    * ran gofmt on every file

[33mcommit 279d659bf4687487ec7c4cb924711c98e2d61370[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Wed Apr 19 10:29:49 2017 -0500

    update warning to errors (#30)

[33mcommit 2d715176213e8021066586b67ba74418b491bd21[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Wed Apr 19 10:23:50 2017 -0500

    update warning to info (#29)

[33mcommit 258ee7c6bd15167e5a73d29392a5e68495ca7a07[m
Author: zhengping12 <wang426@illinois.edu>
Date:   Sun Apr 16 11:32:07 2017 -0500

    Changed .cer to .pem (#26)

[33mcommit 9a7f9d481b879c3ea258c043667fb6c08c96f400[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Sat Apr 15 09:28:08 2017 -0500

    fixing bad hostname checking for URIs (#25)
    
    * fixing very bad hostname checking for URIs
    
    * fixing certs

[33mcommit 11f225d2c2254629fbee3bd44d5a85fd197517d6[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Wed Apr 12 14:23:03 2017 -0500

    removing a lint which is encompased by lint_ext_san_dns_not_fqdn (#23)
    
    * removing an incorrect lint which is encompased by lint_ext_san_dns_not_fqdn
    
    * fixing tests

[33mcommit 6e61b3dadc70b6dd2a3591b8c75c0fb544102ad7[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Fri Mar 31 08:51:54 2017 -0500

    updating FQDN parsing in ZLint (#20)
    
    * updating FQDN parsing in ZLint
    
    * removing unnecessary comment
    
    * update fqdn
    
    * if ? is in end, this goes there so now its only the beginnign
    
    * gofmt

[33mcommit 504f3e137466d10c11049505c6f5911fb3af5d26[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Fri Mar 24 16:11:08 2017 -0400

    fixing bug in ip checking for common names (#19)

[33mcommit 2382879c3f4b75304f33077142a9f396d2b23445[m
Author: zhengping12 <wang426@illinois.edu>
Date:   Tue Mar 21 17:02:47 2017 -0500

    fixing naming, adding sanity check, adding error msg to zlint output file (#18)

[33mcommit 618a3ef21bd7932e65701955145ffb42d4a7142e[m
Author: zhengping12 <wang426@illinois.edu>
Date:   Tue Mar 21 16:24:10 2017 -0500

    fix to sanFqdn lint and 3 new tests (#16)
    
    * fix to sanFqdn lint and 3 new tests
    
    * new tests for ian/san uri host to test wildcard fqdn correctness
    
    * Fixing naming for IAN SAN FQDN and URI
    
    * fixing DNS capitalization
    
    * fixing IAN capitalization
    
    * fixing DNS capitalization
    
    * fixing IP capitalization
    
    * fixing SAN capitalization
    
    * fixing IA5String spelling
    
    * changing lint names to lowercase for those involving IAN and SAN
    
    * fixed RFC capitalization
    
    * removing duplicate files
    
    * fixed EDI capitalization
    
    * fixed capitalization
    
    * fixed SAN capitalization
    
    * fixed name constraint spelling

[33mcommit e6b50db049c4e6e35a1a5b74879d2d70d1aaabbb[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Mon Mar 13 14:40:52 2017 -0500

    Kumarde/rename lints (#15)
    
    * renaming lints that do not make sense
    
    * fixing typo

[33mcommit e859f3a11a2e52d14b74ad5e5a99bf3485b7c02c[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Sun Mar 12 17:45:44 2017 -0500

    Kumarde/modify output to match protobuf (#13)
    
    * modifying lints to match protobuf output
    
    * fixing a typo in constraints...
    
    * moving back from protobuf

[33mcommit 7862de5be4c4be17576e4ede7a658b590a11969c[m
Author: zhengping12 <wang426@illinois.edu>
Date:   Sun Mar 12 09:27:31 2017 -0500

    rsa related lints changes. (#11)
    
    * moving the error case to not apply for rsa certs with no public key
    
    * fixed to check ok first in rsa lints
    
    * fixing parenthesis

[33mcommit 5df5702a0d66d3eca73cb5843377ac7c751a47f3[m
Merge: b91e79e7 e8ca2730
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Tue Mar 7 13:45:49 2017 -0600

    Merge pull request #10 from zmap/dadrian/zcrypto
    
    Migrate x509 from ZGrab to ZCrypto

[33mcommit e8ca27301f8c6b06d006d3682e175e18b34a24fd[m
Author: David Adrian <davidcadrian@gmail.com>
Date:   Tue Mar 7 14:37:25 2017 -0500

    Migrate x509 from ZGrab to ZCrypto
    
    Implemented by running:
    
    ```
    $ find . -type f -name '*.go' -exec sed -i '' 's|github.com/zmap/zgrab/ztools/x509|github.com/zmap/zcrypto/x509|g' {} \;
    ```

[33mcommit b91e79e7ace9660f7972a193bf02520e22d277f5[m
Merge: 3636d956 e1d012eb
Author: Zakir Durumeric <zakird@gmail.com>
Date:   Wed Mar 1 13:39:55 2017 -0800

    Merge pull request #9 from zhengping12/master
    
    Ran go fmt and Lints/*

[33mcommit e1d012eb0c3ee9732fae75e6a74284607bb33529[m
Author: zhengping12 <wang426@illinois.edu>
Date:   Wed Mar 1 15:38:26 2017 -0600

    Ran go fmt and Lints/*

[33mcommit 3636d956d6bd0b0cc84bd8524aa1064b73def40c[m
Merge: 03634486 d5d0ef99
Author: Zakir Durumeric <zakird@gmail.com>
Date:   Wed Mar 1 13:37:32 2017 -0800

    Merge pull request #8 from zhengping12/master
    
    Removed debugging code

[33mcommit d5d0ef991ae0d0cd2adb2a7847c02a6e7efd54ea[m
Author: zhengping12 <wang426@illinois.edu>
Date:   Wed Mar 1 15:34:42 2017 -0600

    Removed debugging code

[33mcommit 036344869cc00556ddc4421a963acf93a6713540[m
Merge: 14d40795 14eb2712
Author: Zakir Durumeric <zakird@gmail.com>
Date:   Wed Mar 1 13:23:30 2017 -0800

    Merge pull request #7 from zhengping12/master
    
    only the rsa related changes

[33mcommit 14eb271201b97eba00beaf98c98d3956efb44de3[m
Author: zhengping12 <wang426@illinois.edu>
Date:   Wed Mar 1 15:19:35 2017 -0600

    lint_rsa_public_exponent_not_in_range.go fix

[33mcommit 755437e70b82d7c3abad2502de7ec84bc040ad1c[m
Author: zhengping12 <wang426@illinois.edu>
Date:   Wed Mar 1 15:01:53 2017 -0600

    only the rsa related changes

[33mcommit 14d407952a7255db29576fdc153e3ff51a1c86f1[m
Merge: 451b3b52 1876fd48
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Fri Feb 17 15:26:03 2017 -0600

    Merge pull request #4 from mhyder13/master
    
    added -num-threads option

[33mcommit 1876fd48e5ec001ee1df7c3a0d662c77307f6eee[m
Author: mhyder13 <mlhyder2@illinois.edu>
Date:   Fri Feb 10 16:33:24 2017 -0600

    added -num-threads option
    
    Changed threaded mode to accept a command line argument to alter the number of goroutines used for processing

[33mcommit 451b3b52ac3408f4f70bad91020f54cddd9bd2d1[m
Merge: 13b684fe 9cd525d8
Author: Zakir Durumeric <zakird@gmail.com>
Date:   Mon Dec 19 21:10:44 2016 -0600

    Merge pull request #3 from zmap/kumarde/gtld_in_struct
    
    fixing names

[33mcommit 9cd525d882093a33ba644db4a0e5f1f2a0861572[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Mon Dec 19 16:01:11 2016 -0500

    fixing names

[33mcommit 13b684fe7321b9c6bbbdc7fb5e096c9feefeee35[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Fri Dec 16 15:22:07 2016 -0600

    finally removing gtld

[33mcommit e432811bdca0a3525717df0e55c2162874118636[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Fri Dec 16 15:21:28 2016 -0600

    removing println

[33mcommit d21c7093bcf37aac3eb57497c3ca19f0f4c5b8ef[m
Merge: 71f46f29 1e5bb463
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Fri Dec 16 15:20:04 2016 -0600

    Merge pull request #2 from zmap/kumarde/update_gtld_util
    
    Kumarde/update gtld util

[33mcommit 1e5bb463aa9d1d0fe5af39143183987b5cc9e3b4[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Fri Dec 16 15:18:28 2016 -0600

    fixing map

[33mcommit 362dbebf2afc389fbe7ae9828935088d224a1f5d[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Fri Dec 16 13:24:20 2016 -0600

    adding gtld util direct HTTP request

[33mcommit 71f46f29a2a5aecffc2dbccd76d03b99c0e0a28c[m
Merge: a4b087f0 cbafd7a7
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Thu Dec 15 17:19:58 2016 -0600

    Merge pull request #1 from zmap/kumarde/json_cleanup
    
    Kumarde/json cleanup

[33mcommit cbafd7a7a99d06633856b6a349f01905b8172e3d[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Thu Dec 15 11:00:41 2016 -0600

    moved enumToString in right place

[33mcommit 12b0ba98556ef0b77f11508258e95dcefa9a0c90[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Thu Dec 15 10:56:06 2016 -0600

    json cleanup

[33mcommit a4b087f0e4b71d59fa46684c11ed306212746444[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Wed Dec 14 17:34:41 2016 -0600

    enumToString

[33mcommit 897779809fcd720e6b45cf77b41e2c533d85919c[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Tue Dec 13 10:34:26 2016 -0600

    adding gitignore to ignore idea files

[33mcommit 7b2398b406e6b116afd8423028ab70f7ad135d01[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Tue Dec 6 15:07:52 2016 -0600

    adding travis CI to Readme

[33mcommit cffe92b08b3dfeb21228e2cd0145e411a609b04c[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Tue Dec 6 15:04:35 2016 -0600

    again, casing

[33mcommit b52d470e9d19beb585c135cde7ce1cb5d588bdd2[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Tue Dec 6 15:01:40 2016 -0600

    casing issue fix

[33mcommit 2e3c84532728fc11264602b6fd4a2522b85eb1fd[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Tue Dec 6 14:57:24 2016 -0600

    first travis.yml attempt

[33mcommit 637e564518df8595cb671609c1318c70263435cb[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Tue Dec 6 11:48:03 2016 -0600

    purging old testing framework

[33mcommit 27d47cb83dfba54cfad6790b5669c0ed94a08870[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Mon Dec 5 22:43:34 2016 -0600

    update teamnsrg->zmap

[33mcommit fa41d29c28e934a68a62ce4a5d82e65e85ec06f6[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Mon Dec 5 22:35:00 2016 -0600

    fixing README

[33mcommit 303acf0753ff52b2fb8f564c9783d1076a522de9[m
Author: Deepak Kumar <kumarde@umich.edu>
Date:   Mon Dec 5 22:34:16 2016 -0600

    adding zmap org to README.md

[33mcommit 8a0f765f8dbabc601a7c4bea0dba0e5189089684[m
Author: Deepak Kumar <dkumar11@illinois.edu>
Date:   Mon Dec 5 22:30:32 2016 -0600

    initial commit from teamnsrg
