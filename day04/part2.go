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

func main() {
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

func (p Passport) Valid() bool {
	return p.ValidByr() &&
		p.ValidIyr() &&
		p.ValidEyr() &&
		p.ValidHgt() &&
		p.ValidHcl() &&
		p.ValidEcl() &&
		p.ValidPid()
}

func (p Passport) ValidByr() bool {
	if p.byr == "" {
		return false
	}
	byr, err := strconv.Atoi(p.byr)
	if err != nil {
		return false
	}
	if byr < 1920 {
		return false
	}
	if byr > 2002 {
		return false
	}
	return true
}

func (p Passport) ValidIyr() bool {
	if p.iyr == "" {
		return false
	}
	iyr, err := strconv.Atoi(p.iyr)
	if err != nil {
		return false
	}
	if iyr < 2010 {
		return false
	}
	if iyr > 2020 {
		return false
	}
	return true
}

func (p Passport) ValidEyr() bool {
	if p.eyr == "" {
		return false
	}
	eyr, err := strconv.Atoi(p.eyr)
	if err != nil {
		return false
	}
	if eyr < 2020 {
		return false
	}
	if eyr > 2030 {
		return false
	}
	return true
}

func (p Passport) ValidHgt() bool {
	hgt := p.hgt
	if hgt == "" {
		return false
	}
	if len(hgt) < 4 {
		return false
	}
	u := hgt[len(hgt)-2:]
	v, err := strconv.Atoi(hgt[0 : len(hgt)-2])
	if err != nil {
		return false
	}
	switch u {
	case "in":
		if v < 59 || v > 76 {
			return false
		}
	case "cm":
		if v < 150 || v > 193 {
			return false
		}
	default:
		return false
	}

	return true
}

func (p Passport) ValidHcl() bool {
	m, err := regexp.Match(`#[0-9a-f]{6}`, []byte(p.hcl))
	if err != nil {
		return false
	}
	if !m {
		return false
	}
	return true
}

func (p Passport) ValidEcl() bool {
	for _, v := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
		if p.ecl == v {
			return true
		}
	}
	return false
}

func (p Passport) ValidPid() bool {
	if len(p.pid) != 9 {
		return false
	}
	m, err := regexp.Match(`[0-9]{9}`, []byte(p.pid))
	if err != nil {
		return false
	}
	return m
}
