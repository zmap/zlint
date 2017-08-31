package lints

import (
	"time"

	"github.com/zmap/zcrypto/x509"
)

var (
	// Lints is a map of all known lints by name. Add a Lint to the map by calling
	// RegisterLint.
	Lints = make(map[string]*Lint)
)

// LintInterface is implemented by each Lint.
type LintInterface interface {
	// Initialize runs once per-lint. It is called during RegisterLint().
	Initialize() error

	// CheckApplies runs once per certificate. It returns true if the Lint should
	// run on the given certificate. If CheckApplies returns false, the Lint
	// result is automatically set to NA without calling CheckEffective() or
	// Run().
	CheckApplies(c *x509.Certificate) bool

	// Execute() is the body of the lint. It is called for every certificate for
	// which CheckApplies() returns true.
	Execute(c *x509.Certificate) *LintResult
}

// A Lint struct represents a single lint, e.g.
// "e_basic_constraints_not_critical". It contains an implementation of LintInterface.
type Lint struct {

	// Name is a lowercase underscore-separated string describing what a given
	// Lint checks. If Name beings with "w", the lint MUST NOT return Error, only
	// Warn. If Name beings with "e", the Lint MUST NOT return Warn, only Error.
	Name string `json:"name,omitempty"`

	// A human-readable description of what the Lint checks. Usually copied
	// directly from the CA/B Baseline Requirements or RFC 5280.
	Description string `json:"description,omitempty"`

	// The source of the check, e.g. "BRs: 6.1.6" or "RFC 5280: 4.1.2.6".
	Source string `json:"source,omitempty"`

	// Lints automatically returns NE for all certificates where CheckApplies() is
	// true but with NotBefore < EffectiveDate. This check is bypassed if
	// EffectiveDate is zero.
	EffectiveDate time.Time `json:"-"`

	// The implementation of the lint logic.
	Lint LintInterface `json:"-"`
}

// CheckEffective returns true if c was issued on or after the EffectiveDate. If
// EffectiveDate is zero, CheckEffective always returns true.
func (l *Lint) CheckEffective(c *x509.Certificate) bool {
	if l.EffectiveDate.IsZero() || !l.EffectiveDate.After(c.NotBefore) {
		return true
	}
	return false
}

// Execute runs the lint against a certificate. See LintInterface for details
// about the methods called. The ordering is as follows:
//
// CheckApplies()
// CheckEffective()
// Execute()
func (l *Lint) Execute(cert *x509.Certificate) *LintResult {
	if !l.Lint.CheckApplies(cert) {
		return &LintResult{Status: NA}
	} else if !l.CheckEffective(cert) {
		return &LintResult{Status: NE}
	}
	res := l.Lint.Execute(cert)
	return res
}

// RegisterLint must be called once for each lint to be excuted. Duplicate lint
// names are squashed. Normally, RegisterLint is called during init().
func RegisterLint(l *Lint) {
	if err := l.Lint.Initialize(); err != nil {
		panic("could not initialize lint: " + l.Name + ": " + err.Error())
	}
	Lints[l.Name] = l
}
