package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	ln, c1, c2, c3, c4, c5 := 0, 0, 0, 0, 0, 0
	for s.Scan() {
		ln = ln + 1
		ls := s.Text()
		if isTree1(ln, ls, 1) {
			c1 = c1 + 1
		}
		if isTree1(ln, ls, 3) {
			c2 = c2 + 1
		}
		if isTree1(ln, ls, 5) {
			c3 = c3 + 1
		}
		if isTree1(ln, ls, 7) {
			c4 = c4 + 1
		}
		if isTree2(ln, ls) {
			c5 = c5 + 1
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(c2)
	fmt.Println(c1 * c2 * c3 * c4 * c5)
}
func isTree1(px int, s string, r int) bool {
	py := r * (px - 1)
	sy := py % len(s)
	if sy >= len(s) {
		log.Fatalf("line %d: invalid, %s", px, s)
		return false
	}

	return string(s[sy]) == "#"
}

func isTree2(px int, s string) bool {
	if px%2 == 0 {
		return false
	}
	py := px / 2
	sy := py % len(s)
	if sy >= len(s) {
		log.Fatalf("line %d: invalid, %s", px, s)
		return false
	}

	return string(s[sy]) == "#"
}
