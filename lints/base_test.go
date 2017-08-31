package lints

import "testing"

func TestAllLintsHaveNameDescriptionSource(t *testing.T) {
	for name, lint := range Lints {
		if lint.Name == "" {
			t.Errorf("lint %s has empty name", name)
		}
		if lint.Description == "" {
			t.Errorf("lint %s has empty description", name)
		}
		if lint.Source == "" {
			t.Errorf("lint %s has empty provenance", name)
		}
	}
}
