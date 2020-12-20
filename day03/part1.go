package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	ln, c := 0, 0
	for s.Scan() {
		ln = ln + 1
		if isTree(ln, s.Text()) {
			c = c + 1
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)
}

func isTree(px int, s string) bool {
	py := 3 * (px - 1)
	sy := py % len(s)
	t := s[:sy] + "O" + s[sy+1:]
	if sy >= len(s) {
		log.Fatalf("line %d: invalid, %s", px, s)
		return false
	}
	fmt.Println(t)

	return string(s[sy]) == "#"
}
