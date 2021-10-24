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
