package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	a := []int{}
	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatal(err)
		}
		for _, j := range a {
			if i+j == 2020 {
				fmt.Printf("%d", i*j)
			}
		}

		a = append(a, i)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
