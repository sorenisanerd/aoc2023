package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/sorenisanerd/aoc2023/utils"
)

var sample = `Time:      7  15   30
Distance:  9  40  200`

func main() {
	Parts(sample)
	Parts(utils.GetDayData(2023, 6))
}

func ExtractInts(s string) []int {
	rv := []int{}
	for _, word := range regexp.MustCompile(`\s+`).Split(s, -1) {
		if num, err := strconv.Atoi(word); err == nil {
			rv = append(rv, num)
		}
	}
	return rv
}

func dist(total, button int) int {
	return button * (total - button)
}

func Parts(s string) (int, int) {
	lines := strings.Split(s, "\n")
	t := strings.Split(lines[0], ":")
	r := strings.Split(lines[1], ":")
	times := ExtractInts(t[1])
	records := ExtractInts(r[1])
	time := ExtractInts(strings.Replace(t[1], " ", "", -1))[0]
	record := ExtractInts(strings.Replace(r[1], " ", "", -1))[0]

	p1 := 1
	p2 := 0

	minn, maxx := 0, 0
	for idx := range times {
		c := 0
		for i := 0; i < times[idx]; i++ {
			d := dist(times[idx], i)
			if d > records[idx] {
				c += 1
			}
		}
		p1 *= c
	}

	for i := 0; i < time; i++ {
		d := dist(time, i)
		minn = min(d, minn)
		maxx = max(d, maxx)
		if d > record {
			p2 += 1
		}
	}
	fmt.Println(p1, p2, minn, maxx)
	return p1, p2
}
