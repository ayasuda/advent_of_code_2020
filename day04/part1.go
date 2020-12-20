package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	c := 0

	var p *Passport
	var ps []Passport

	for s.Scan() {
		l := s.Text()
		if l == "" {
			if p != nil {
				ps = append(ps, *p)
			}
			p = nil
			continue
		}
		if p == nil {
			p = new(Passport)
		}
		parse(p, l)
	}
	if p != nil {
		ps = append(ps, *p)
	}

	for _, d := range ps {
		if d.Valid() {
			c = c + 1
		}
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)
}

func (p Passport) Valid() bool {
	if p.byr == "" {
		return false
	}
	if p.iyr == "" {
		return false
	}
	if p.eyr == "" {
		return false
	}
	if p.hgt == "" {
		return false
	}
	if p.hcl == "" {
		return false
	}
	if p.ecl == "" {
		return false
	}
	if p.pid == "" {
		return false
	}

	return true
}

func parse(p *Passport, line string) {
	lva := strings.Split(line, " ")
	for _, lv := range lva {
		a := strings.Split(lv, ":")
		l, v := a[0], a[1]
		switch l {
		case "byr":
			p.byr = v
		case "iyr":
			p.iyr = v
		case "eyr":
			p.eyr = v
		case "hgt":
			p.hgt = v
		case "hcl":
			p.hcl = v
		case "ecl":
			p.ecl = v
		case "pid":
			p.pid = v
		case "cid":
			p.cid = v
		}
	}
}
