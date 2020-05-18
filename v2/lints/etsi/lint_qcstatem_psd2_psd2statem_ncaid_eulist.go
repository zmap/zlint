package etsi

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

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v2/lint"
	"github.com/zmap/zlint/v2/util"
)

type countryAndNCAIdPair struct {
	country string
	ncaid   string
}

type qcStatemPsd2Psd2StatemNcaidEulist struct{}

func (l *qcStatemPsd2Psd2StatemNcaidEulist) Initialize() error {
	return nil
}

func (l *qcStatemPsd2Psd2StatemNcaidEulist) CheckApplies(c *x509.Certificate) bool {
	_, isPresent := util.IsQcStatemPresent(c, &util.IdEtsiPsd2Statem)
	return isPresent
}

func (l *qcStatemPsd2Psd2StatemNcaidEulist) Execute(c *x509.Certificate) *lint.LintResult {

	countryAndNCAIdMap := []countryAndNCAIdPair{
		{"AT", "FMA"},
		{"BE", "NBB"},
		{"BG", "BNB"},
		{"HR", "CNB"},
		{"CY", "CBC"},
		{"CZ", "CNB"},
		{"DK", "DFSA"},
		{"EE", "FI"},
		{"FI", "FINFSA"},
		{"FR", "ACPR"},
		{"DE", "BAFIN"},
		{"GF", "BOG"},
		{"HU", "CBH"},
		{"IS", "FME"},
		{"IE", "CBI"},
		{"IT", "BI"},
		{"LI", "FMA"},
		{"LV", "FCMC"},
		{"LT", "BL"},
		{"LU", "CSSF"},
		{"NO", "FSA"},
		{"MT", "MFSA"},
		{"NL", "DNB"},
		{"PL", "PFSA"},
		{"PT", "BP"},
		{"RO", "NBR"},
		{"SK", "NBS"},
		{"SI", "BS"},
		{"ES", "BE"},
		{"SE", "FINA"},
		{"GB", "FCA"},
	}
	ext := util.GetExtFromCert(c, util.QcStateOid)
	s := util.ParseQcStatem(ext.Value, util.IdEtsiPsd2Statem)
	if s.GetErrorInfo() != "" {
		return &lint.LintResult{Status: lint.Warn, Details: "parsing error for PSD2 QcStatement, cannot properly apply this lint: " + s.GetErrorInfo()}
	}
	psd2Statem, ok := s.(util.EtsiPsd2)
	if !ok {
		return &lint.LintResult{Status: lint.Fatal, Details: "parsed QcStatem is not of type EtsiPsd2"}
	}
	if psd2Statem.GetNcaCountry() == "" || psd2Statem.GetNcaId() == "" {
		return &lint.LintResult{Status: lint.Warn, Details: "NCAId field (country-NcaId) in PSD2 QcStatement has invalid format, cannot properly apply this lint"}
	}
	for _, countryAndNCAId := range countryAndNCAIdMap {
		if psd2Statem.GetNcaCountry() == countryAndNCAId.country {
			if psd2Statem.GetNcaId() != countryAndNCAId.ncaid {

				return &lint.LintResult{Status: lint.Warn, Details: "NCAId in PSD2 QcStatement for given country is not conforming to (informative) EU List"}
			}
		}
	}
	return &lint.LintResult{Status: lint.Pass}
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "w_qcstatem_psd2_psd2statem_ncaid_eulist",
		Description:   "If the country given in the PSD2 QcStatement NCAId field is found in the ETSI list in Annex D (see citation), then this lint checks that the corresponding NcaId is contained in the NcaId field of the PSD2 QcStatement",
		Citation:      "ETSI TS 119 495, Annex D",
		Source:        lint.EtsiEsi,
		EffectiveDate: util.EtsiEn319_412_5_V2_2_1_Date,
		Lint:          &qcStatemPsd2Psd2StatemNcaidEulist{},
	})
}
