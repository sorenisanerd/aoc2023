package day01

import (
	"strings"
)

func Part1(input string) int {
	return day1part(input, false)
}

func Part2(input string) int {
	return day1part(input, true)
}

func day1part(input string, allowWords bool) int {
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
		x := l[0]
		if x >= '0' && x <= '9' {
			rv = append(rv, int(x-'0'))
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
