package lints

import "testing"

func TestAllLintsHaveUpdateReport(t *testing.T) {
	for name, lint := range Lints {
		if lint.updateReport == nil {
			t.Errorf("lint %s has nil updateReport", name)
		}
	}
}
