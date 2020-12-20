package main

import (
	"testing"
)

func Test_detectSeat(t *testing.T) {
	cases := []struct {
		definition string
		expectRow  int
		expectCol  int
	}{
		{"FBFBBFFRLR", 44, 5},
	}

	for i, c := range cases {
		row, col := detectSeat(c.definition)
		if row != c.expectRow || col != c.expectCol {
			t.Errorf("case #%d is failed\n\texpted\t(%d,%d)\n\tgot\t(%d,%d)\n",
				i, c.expectRow, c.expectCol, row, col)
		}
	}
}

func Test_detect(t *testing.T) {
	cases := []struct {
		isLower   bool
		min       int
		max       int
		expectMin int
		expectMax int
	}{
		{true, 0, 127, 0, 63},
		{false, 0, 63, 32, 63},
		{true, 32, 63, 32, 47},
		{false, 32, 47, 40, 47},
		{false, 40, 47, 44, 47},
		{true, 44, 47, 44, 45},
		{true, 44, 45, 44, 44},

		{false, 0, 7, 4, 7},
		{true, 4, 7, 4, 5},
		{false, 4, 5, 5, 5},
	}

	for i, c := range cases {
		min, max := detect(c.isLower, c.min, c.max)
		if min != c.expectMin || max != c.expectMax {
			t.Errorf("case #%d is failed\n\texpted\t(%d,%d)\n\tgot\t(%d,%d)\n",
				i, c.expectMin, c.expectMax, min, max)
		}
	}
}
