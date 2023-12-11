package main

import (
	"fmt"
	"strings"

	"github.com/sorenisanerd/aoc2023/utils"
)

func main() {
	fmt.Println(Parts(utils.GetDayData(2015, 1)))
}

func Parts(s string) (int, int) {
	p1 := 0
	p2 := -1
	for _, l := range strings.Split(s, "\n") {
		if l == "" {
			continue
		}
		for idx, c := range l {
			if c == ')' {
				p1 -= 1
			} else if c == '(' {
				p1 += 1
			}
			if p2 < 0 && p1 < 0 {
				p2 = idx + 1
			}
		}
	}
	return p1, p2
}
