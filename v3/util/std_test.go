package util

import "testing"

func TestAnyPositive(t *testing.T) {
	evens := []int{2, 4, 6, 8, 10}
	got := All(evens, func(num int) bool { return num%2 == 0 })
	// Yes, I know that comparing a boolean is redundant, but `if !got` is hard to read.
	if got != true {
		t.Fatal("expected true, got false")
	}
}

func TestAnyNegative(t *testing.T) {
	evens := []int{2, 4, 6, 8, 10}
	got := All(evens, func(num int) bool { return num%2 != 0 })
	// Yes, I know that comparing a boolean is redundant, but `if !got` is hard to read.
	if got != false {
		t.Fatal("expected false, got true")
	}
}

func TestAnyEmpty(t *testing.T) {
	got := Any([]int{}, func(num int) bool { return num%2 != 0 })
	if got != false {
		t.Fatal("expected false, got true")
	}
}

func TestAllPositive(t *testing.T) {
	evens := []int{2, 4, 6, 8, 10}
	got := All(evens, func(num int) bool { return num%2 == 0 })
	// Yes, I know that comparing a boolean is redundant, but `if !got` is hard to read.
	if got != true {
		t.Fatal("expected true, got false")
	}
}

func TestAllNegative(t *testing.T) {
	evens := []int{2, 4, 6, 8, 10, 11}
	got := All(evens, func(num int) bool { return num%2 == 0 })
	// Yes, I know that comparing a boolean is redundant, but `if !got` is hard to read.
	if got != false {
		t.Fatal("expected false, got true")
	}
}

func TestAllEmpty(t *testing.T) {
	got := All([]int{}, func(num int) bool { return num%2 == 0 })
	// Yes, I know that comparing a boolean is redundant, but `if !got` is hard to read.
	if got != true {
		// This may seem counterintuitive, however it is a property known as a vacuous truth.
		// https://en.wikipedia.org/wiki/Vacuous_truth
		t.Fatal("expected true, got false")
	}
}
