package util

import (
	"testing"
	"time"
)

func TestBeforeOrOn(t *testing.T) {
	data := []struct {
		left  time.Time
		right time.Time
		want  bool
	}{
		{
			time.Date(2021, time.February, 1, 1, 1, 58, 0, time.UTC),
			time.Date(2021, time.February, 1, 1, 1, 59, 0, time.UTC),
			true,
		},
		{
			time.Date(2021, time.February, 1, 1, 1, 59, 0, time.UTC),
			time.Date(2021, time.February, 1, 1, 1, 59, 0, time.UTC),
			true,
		}, {
			time.Date(2021, time.February, 1, 1, 2, 0, 0, time.UTC),
			time.Date(2021, time.February, 1, 1, 1, 59, 0, time.UTC),
			false,
		}}
	for _, test := range data {
		got := BeforeOrOn(test.left, test.right)
		if got != test.want {
			t.Errorf("Got %v from %v", got, test)
		}
	}
}

func TestOnOrAfter(t *testing.T) {
	data := []struct {
		left  time.Time
		right time.Time
		want  bool
	}{
		{
			time.Date(2021, time.February, 1, 1, 1, 58, 0, time.UTC),
			time.Date(2021, time.February, 1, 1, 1, 59, 0, time.UTC),
			false,
		},
		{
			time.Date(2021, time.February, 1, 1, 1, 59, 0, time.UTC),
			time.Date(2021, time.February, 1, 1, 1, 59, 0, time.UTC),
			true,
		}, {
			time.Date(2021, time.February, 1, 1, 2, 0, 0, time.UTC),
			time.Date(2021, time.February, 1, 1, 1, 59, 0, time.UTC),
			true,
		}}
	for _, test := range data {
		got := OnOrAfter(test.left, test.right)
		if got != test.want {
			t.Errorf("Got %v from %v", got, test)
		}
	}
}

func TestInclusiveSecondsBetween(t *testing.T) {
	data := []struct {
		left  time.Time
		right time.Time
		want  float64
	}{
		{
			time.Date(2021, time.February, 1, 1, 1, 58, 0, time.UTC),
			time.Date(2021, time.February, 1, 1, 1, 58, 0, time.UTC),
			1,
		},
		{
			time.Date(2021, time.February, 1, 1, 1, 58, 0, time.UTC),
			time.Date(2021, time.February, 1, 1, 1, 59, 0, time.UTC),
			2,
		},
		{
			time.Date(2021, time.February, 1, 1, 1, 58, 0, time.UTC),
			time.Date(2021, time.February, 2, 1, 1, 57, 0, time.UTC),
			86400,
		},
		{
			time.Date(2021, time.February, 1, 1, 1, 58, 0, time.UTC),
			time.Date(2021, time.February, 2, 1, 1, 58, 0, time.UTC),
			86401,
		},
	}
	for _, test := range data {
		got := InclusiveSecondsBetween(test.left, test.right)
		if got != test.want {
			t.Errorf("Got %v from %v", got, test)
		}
	}
}

func TestInclusiveDaysBetween(t *testing.T) {
	data := []struct {
		left  time.Time
		right time.Time
		want  float64
	}{
		{
			time.Date(2021, time.February, 1, 1, 1, 58, 0, time.UTC),
			time.Date(2021, time.February, 1, 1, 1, 58, 0, time.UTC),
			1,
		},
		{
			time.Date(2021, time.February, 1, 1, 1, 58, 0, time.UTC),
			time.Date(2021, time.February, 1, 1, 1, 59, 0, time.UTC),
			1,
		},
		{
			time.Date(2021, time.February, 1, 1, 1, 58, 0, time.UTC),
			time.Date(2021, time.February, 2, 1, 1, 57, 0, time.UTC),
			1,
		},
		{
			time.Date(2021, time.February, 1, 1, 1, 58, 0, time.UTC),
			time.Date(2021, time.February, 2, 1, 1, 58, 0, time.UTC),
			2,
		},
	}
	for _, test := range data {
		got := InclusiveDaysBetween(test.left, test.right)
		if got != test.want {
			t.Errorf("Got %v from %v", got, test)
		}
	}
}

func TestInclusiveSpanExceedsDays(t *testing.T) {
	data := []struct {
		left           time.Time
		right          time.Time
		maxDaysAllowed float64
		want           bool
	}{
		{
			time.Date(2021, time.February, 1, 1, 1, 58, 0, time.UTC),
			time.Date(2021, time.February, 1, 1, 1, 58, 0, time.UTC),
			1,
			false,
		},
		{
			time.Date(2021, time.February, 1, 1, 1, 58, 0, time.UTC),
			time.Date(2021, time.February, 1, 1, 1, 59, 0, time.UTC),
			1,
			false,
		},
		{
			time.Date(2021, time.February, 1, 1, 1, 58, 0, time.UTC),
			time.Date(2021, time.February, 2, 1, 1, 57, 0, time.UTC),
			1,
			false,
		},
		{
			time.Date(2021, time.February, 1, 1, 1, 58, 0, time.UTC),
			time.Date(2021, time.February, 2, 1, 1, 58, 0, time.UTC),
			1,
			true,
		},
		{
			time.Date(2021, time.February, 1, 1, 1, 58, 0, time.UTC),
			time.Date(2021, time.February, 2, 1, 1, 58, 0, time.UTC),
			2,
			false,
		},
		{
			time.Date(2021, time.February, 1, 1, 1, 58, 0, time.UTC),
			time.Date(2021, time.February, 2, 1, 1, 58, 0, time.UTC),
			199,
			false,
		},
		//Last Update: Oct 26 12:20:26 2026 GMT
		//Next Update: Nov  5 12:20:26 2026 GMT
		// from https://github.com/zmap/zlint/pull/1119#discussion_r4437503214
		{
			time.Date(2026, time.October, 26, 12, 20, 26, 0, time.UTC),
			time.Date(2026, time.November, 5, 12, 20, 25, 0, time.UTC),
			10,
			false,
		},
		{
			time.Date(2026, time.October, 26, 12, 20, 26, 0, time.UTC),
			time.Date(2026, time.November, 5, 12, 20, 26, 0, time.UTC),
			10,
			true,
		},
		{
			time.Date(2026, time.October, 26, 12, 20, 26, 0, time.UTC),
			time.Date(2026, time.November, 5, 12, 20, 26, 0, time.UTC),
			11,
			false,
		},
	}
	for _, test := range data {
		got := InclusiveSpanExceedsDays(test.left, test.right, test.maxDaysAllowed)
		if got != test.want {
			t.Errorf("Got %v from %v", got, test)
		}
	}
}
