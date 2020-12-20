package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Answer map[rune]struct{}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	groupAnswers := [][]string{}
	answers := []string{}
	for s.Scan() {
		l := s.Text()
		if l == "" {
			groupAnswers = append(groupAnswers, answers)
			answers = []string{}
		} else {
			answers = append(answers, l)
		}
	}
	groupAnswers = append(groupAnswers, answers)

	c := 0
	for _, ga := range groupAnswers {
		c = c + countAny(ga)
	}

	fmt.Println("Part 1")
	fmt.Println(c)

	c = 0
	for _, ga := range groupAnswers {
		c = c + countEvery(ga)
	}

	fmt.Println("Part 2")
	fmt.Println(c)

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}

func countEvery(answers []string) int {
	var all Answer
	for _, ans := range answers {
		answer := make(Answer)
		for _, c := range ans {
			answer.Add(c)
		}
		if all == nil {
			all = answer
		} else {
			all = all.Intersection(answer)
		}
	}
	return all.Len()
}

func countAny(answers []string) int {
	all := make(Answer)
	for _, ans := range answers {
		answer := make(Answer)
		for _, c := range ans {
			answer.Add(c)
		}
		all = all.Union(answer)
	}
	return all.Len()
}

func (a Answer) Values() []rune {
	ks := []rune{}
	for k, _ := range a {
		ks = append(ks, k)
	}
	return ks
}

func (a Answer) Len() int {
	return len(a.Values())
}

func (a Answer) Include(c rune) bool {
	_, b := a[c]
	return b
}

func (a Answer) Add(c rune) {
	a[c] = struct{}{}
}

func (a Answer) Union(b Answer) (res Answer) {
	res = make(Answer)
	for _, v := range a.Values() {
		res.Add(v)
	}
	for _, v := range b.Values() {
		res.Add(v)
	}
	return res
}

func (a Answer) Intersection(b Answer) (res Answer) {
	res = make(Answer)
	for _, v := range b.Values() {
		if a.Include(v) {
			res.Add(v)
		}
	}
	return res
}
