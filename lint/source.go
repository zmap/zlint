package lint

/*
 * ZLint Copyright 2020 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

// An Enum to programmatically represent the source of a lint
type LintSource int

// NOTE(@cpu): If you are adding a new LintSource make sure you have considered
// updating the Directory() function.
const (
	UnknownLintSource LintSource = iota
	CABFBaselineRequirements
	RFC5280
	RFC5480
	RFC5891
	ZLint
	AWSLabs
	EtsiEsi // ETSI - Electronic Signatures and Infrastructures (ESI)
	CABFEVGuidelines
	AppleCTPolicy          // https://support.apple.com/en-us/HT205280
	MozillaRootStorePolicy // https://github.com/mozilla/pkipolicy
)

// LintSources contains a list of the valid lint sources we expect to be used
// by ZLint lints.
var LintSources = []LintSource{
	CABFBaselineRequirements,
	CABFEVGuidelines,
	RFC5280,
	RFC5480,
	RFC5891,
	AppleCTPolicy,
	EtsiEsi,
	ZLint,
	AWSLabs,
}

// Directory returns the directory name in `lints/` for the LintSource.
func (l LintSource) Directory() string {
	switch l {
	case CABFBaselineRequirements:
		return "cabf_br"
	case CABFEVGuidelines:
		return "cabf_ev"
	case RFC5280, RFC5480, RFC5891:
		return "rfc"
	case AppleCTPolicy:
		return "apple"
	case EtsiEsi:
		return "etsi"
	default:
		return "community"
	}
}
