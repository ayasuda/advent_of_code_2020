package main

import (
	"testing"
)

func Test_countAny(t *testing.T) {
	cases := []struct {
		answers     []string
		expectCount int
	}{
		{[]string{"abc"}, 3},
		{[]string{"a", "b", "c"}, 3},
		{[]string{"ab", "ac"}, 3},
		{[]string{"ab", "ca"}, 3},
		{[]string{"a", "a", "a", "a"}, 1},
		{[]string{"b"}, 1},
	}

	for i, c := range cases {
		cnt := countAny(c.answers)
		if cnt != c.expectCount {
			t.Errorf("case #%d is failed\n\texpect\t%d\n\tgot\t%d", i, c.expectCount, cnt)
		}
	}
}

func Test_countEvery(t *testing.T) {
	cases := []struct {
		answers     []string
		expectCount int
	}{
		{[]string{"abc"}, 3},
		{[]string{"a", "b", "c"}, 0},
		{[]string{"ab", "ac"}, 1},
		{[]string{"ab", "ca"}, 1},
		{[]string{"a", "a", "a", "a"}, 1},
		{[]string{"b"}, 1},
	}

	for i, c := range cases {
		cnt := countEvery(c.answers)
		if cnt != c.expectCount {
			t.Errorf("case #%d is failed\n\texpect\t%d\n\tgot\t%d", i, c.expectCount, cnt)
		}
	}
}
