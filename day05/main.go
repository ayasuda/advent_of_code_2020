package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

var MAX_ROW int = 127
var MAX_COL int = 7

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	max := 0
	var seatIdList []int
	for s.Scan() {
		l := s.Text()
		if len(l) != 10 {
			log.Fatalf("length is not 10, got %v", l)
		}
		row, col := detectSeat(l)
		seatId := row*8 + col
		seatIdList = append(seatIdList, seatId)
		if seatId > max {
			max = seatId
		}
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("part1")
	fmt.Println(max)

	sort.Ints(seatIdList)
	end := len(seatIdList)
	for i, n := range seatIdList {
		if i+1 == end {
			break
		}
		if seatIdList[i+1]-n > 1 {
			fmt.Println("part2")
			fmt.Println(n + 1)
		}
	}
}

func detectSeat(defenition string) (row int, col int) {
	rd, cd := defenition[0:7], defenition[7:]

	min, max := 0, MAX_ROW
	for _, d := range rd {
		switch d {
		case 'F':
			min, max = detect(true, min, max)
		case 'B':
			min, max = detect(false, min, max)
		default:
			log.Fatalf("detection is not F or B : got %c", d)
		}
	}
	if min != max {
		log.Fatalf("detection failed, min and max is different at %s\n: got (min, max) = (%d, %d)",
			rd, min, max)
	}
	row = min

	min, max = 0, MAX_COL
	for _, d := range cd {
		switch d {
		case 'L':
			min, max = detect(true, min, max)
		case 'R':
			min, max = detect(false, min, max)
		default:
			log.Fatalf("detection is not L or R : got %c", d)
		}
	}
	if min != max {
		log.Fatalf("detection failed, min and max is different at %s\n: got (min, max) = (%d, %d)",
			cd, min, max)
	}
	col = min
	return row, col
}

func detect(isLower bool, min, max int) (int, int) {
	if isLower {
		return min, max - (max-min)/2 - 1
	} else {
		return 1 + min + (max-min)/2, max
	}
}
