package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	c := 0
	for s.Scan() {
		if check(s.Bytes()) {
			c = c + 1
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)
}

func check(s []byte) bool {
	r := regexp.MustCompile("([0-9]+)-([0-9]+) ([a-z]): ([a-z]+)")
	g := r.FindSubmatch(s)
	if len(g) != 5 {
		log.Fatal("cannot compile", g)
		return false
	}
	minStr, maxStr, c, pass := string(g[1]), string(g[2]), string(g[3]), string(g[4])
	min, err := strconv.Atoi(minStr)
	if err != nil {
		log.Fatal(err)
		return false
	}
	max, err := strconv.Atoi(maxStr)
	if err != nil {
		log.Fatal(err)
		return false
	}

	l := strings.Count(pass, c)

	if min <= l && l <= max {
		return true
	}
	return false
}
