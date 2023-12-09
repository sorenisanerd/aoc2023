package main

import (
	"fmt"
	"strings"

	"github.com/sorenisanerd/aoc2023/utils"
)

func main() {
	fmt.Println(Parts(utils.GetDayData(2023, 1)))
}

func Parts(s string) (int, int) {
	return part(s, false), part(s, true)
}

func part(input string, allowWords bool) int {
	total := 0
	for _, l := range strings.Split(input, "\n") {
		digits := getDigits(l, allowWords)
		if len(digits) > 0 {
			v := digits[0]*10 + digits[len(digits)-1]
			total += v
		}
	}
	return total
}

func getDigits(l string, allowWords bool) []int {
	rv := []int{}
	for len(l) > 0 {
		if x := l[0]; utils.IsDigit(x) {
			rv = append(rv, utils.MustAtoi(string(x)))
		} else if allowWords {
			for i, w := range strings.Split("one two three four five six seven eight nine", " ") {
				if strings.HasPrefix(l, w) {
					rv = append(rv, i+1)
				}
			}
		}
		l = l[1:]
	}
	return rv
}
