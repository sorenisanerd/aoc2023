package utils

import "fmt"

// GCD finds the Greatest Common Divisor (GCD) of two integers
func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

// LCM finds the Least Common Multiple of 2 or more integers
func LCM(a int, rest ...int) int {
	rv := a

	for _, n := range rest {
		rv = (n * rv) / GCD(n, rv)
	}

	return rv
}

func IsDigit[T byte | rune | string](r T) bool {
	switch v := any(r).(type) {
	case rune:
		return v >= '0' && v <= '9'
	case byte:
		return v >= '0' && v <= '9'
	case string:
		if len(v) != 1 {
			panic(fmt.Sprintf("IsDigit: string must be length 1, got %d", len(v)))
		}
		return IsDigit(v[0])
	}
	panic("How did we end up here?")
}
