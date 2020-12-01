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
			for _, k := range a {
				if i+j+k == 2020 {
					fmt.Printf("%d\n", i*j*k)
					return
				}
			}
		}

		a = append(a, i)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
