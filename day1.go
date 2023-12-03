package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var digitRE, digitsAndWordsRE *regexp.Regexp

func init() {
	digitRE = regexp.MustCompile("[0-9]")
	digitsAndWordsRE = regexp.MustCompile("([0-9]|one|two|three|four|five|six|seven|eight|nine|zero)")
}

func getFirstDigit(s string) int {
	return getDigit(digitRE, s, false, false)
}

func getFirstDigit2(s string) int {
	return getDigit(digitsAndWordsRE, s, false, true)
}

func getLastDigit(s string) int {
	return getDigit(digitRE, s, true, false)
}

func getLastDigit2(s string) int {
	return getDigit(digitsAndWordsRE, s, true, true)
}

func getDigit(r *regexp.Regexp, s string, last bool, allowWords bool) int {
	matches := r.FindAllStringSubmatch(s, -1)
	if len(matches) > 0 {
		var v string
		if last {
			v = matches[len(matches)-1][0]
		} else {
			v = matches[0][0]
		}
		rv, err := strconv.Atoi(v)
		if err != nil {
			if allowWords {
				if rv, ok := wordToNum[v]; ok {
					return rv
				}
			}
			panic(err)
		}
		return rv
	}
	return -1
}

func lval(l string, allowWords bool) int {
	if allowWords {
		return getFirstDigit2(l)*10 + getLastDigit2(l)
	}
	return getFirstDigit(l)*10 + getLastDigit(l)
}

var wordToNum = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func day1part1(input string) int {
	return day1part(input, false)
}

func day1part2(input string) int {
	return day1part(input, true)
}

func day1part(input string, allowWords bool) int {
	total := 0
	for _, l := range strings.Split(input, "\n") {
		if v := lval(l, allowWords); v > 0 {
			total += v
			fmt.Println(v, l)
		} else {
			fmt.Println("SADFASDFASDF", l)
		}
	}
	return total
}

func mustRead(fname string) string {
	fp, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	buf, err := io.ReadAll(fp)
	if err != nil {
		panic(err)
	}
	return string(buf)
}

func main() {
	//	fmt.Println(day1part1(mustRead("day1.txt")))
	fmt.Println(day1part2(mustRead("day1.txt")))
}
