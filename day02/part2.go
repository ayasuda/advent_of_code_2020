package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	c := 0
	for s.Scan() {
		if check2(s.Bytes()) {
			c = c + 1
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)
}

func check2(s []byte) bool {
	r := regexp.MustCompile("([0-9]+)-([0-9]+) ([a-z]): ([a-z]+)")
	g := r.FindSubmatch(s)
	if len(g) != 5 {
		log.Fatal("cannot compile", g)
		return false
	}
	p1Str, p2Str, c, pass := string(g[1]), string(g[2]), string(g[3]), string(g[4])
	p1, err := strconv.Atoi(p1Str)
	if err != nil {
		log.Fatal(err)
		return false
	}
	p2, err := strconv.Atoi(p2Str)
	if err != nil {
		log.Fatal(err)
		return false
	}
	if len(pass) < p1 {
		log.Fatal("password length too small")
		return false
	}
	if len(pass) < p2 {
		log.Fatal("password length too small")
		return false
	}
	if string(pass[p1-1]) == c && string(pass[p2-1]) != c {
		return true
	}
	if string(pass[p1-1]) != c && string(pass[p2-1]) == c {
		return true
	}
	return false
}
