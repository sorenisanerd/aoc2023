package utils

import (
	"regexp"
	"strconv"

	"slices"
)

func ReplaceAtIndex(in string, i int, r rune) string {
	if i < 0 {
		i = len(in) + i
	}
	out := []rune(in)
	out[i] = r
	return string(out)
}

func Atoi(s string) (int, string) {
	rv := 0
	for l := 1; l <= len(s); l++ {
		if v, err := strconv.Atoi(s[:l]); err != nil {
			return rv, s[l-1:]
		} else {
			rv = v
		}
	}
	return rv, ""
}

func MustAtoi(s string) int {
	if n, err := strconv.Atoi(s); err != nil {
		panic(err)
	} else {
		return n
	}
}

var nonDigits = regexp.MustCompile(`[^\d-]+`)

func ExtractInts(s string) []int {
	rv := []int{}
	for _, s_ := range nonDigits.Split(s, -1) {
		if s_ == "" {
			continue
		}
		rv = append(rv, MustAtoi(s_))
	}
	return rv
}

var whitespace = regexp.MustCompile(`\s+`)

func Words(s string) []string {
	return slices.DeleteFunc(whitespace.Split(s, -1), func(s string) bool {
		return s == ""
	})
}
