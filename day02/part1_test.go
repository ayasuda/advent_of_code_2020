package main

import (
	"testing"
)

func TestCheck(t *testing.T) {
	res := check([]byte("1-3 a: abcde"))
	if res != true {
		t.Fatal("mis check1 1")
	}
	res = check([]byte("1-3 b: cdefg"))
	if res != false {
		t.Fatal("mis check1 2")
	}
	res = check([]byte("2-9 c: ccccccccc"))
	if res != true {
		t.Fatal("mis check1 3")
	}
}

func TestCheck2(t *testing.T) {
	res := check2([]byte("1-3 a: abcde"))
	if res != true {
		t.Fatal("mis check2 1")
	}
	res = check2([]byte("1-3 b: cdefg"))
	if res != false {
		t.Fatal("mis check2 2")
	}
	res = check2([]byte("2-9 c: ccccccccc"))
	if res != false {
		t.Fatal("mis check2 3")
	}
}
