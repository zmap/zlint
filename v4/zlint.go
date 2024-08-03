//
// ZLint Copyright 2024 Regents of the University of Michigan
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy
// of the License at http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.
//

// Package zlint contains linters used to check parsed info from X.509
// certificates for compliance. ZLint has a focus on the Web PKI. Version 4.0
// introduces support for Version 2.0+ of the CA/Browser Forum Baseline
// Requirements ("Profiles").

package zlint

// Version is the major version of ZLint.
const Version = 4

// LintCertificateForProfile runs all known lints for c, treating c as a
// certificate of the provided Profile.
// func LintCertificateForProfile(c *x509.Certificate, profile lint.Profile) (*lint.ResultSet, error) {
// 	return nil, errors.New("unimplemented")
// }
