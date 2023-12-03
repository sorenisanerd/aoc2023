package day02

import (
	"strconv"
	"strings"
)

type cubes struct {
	red   int
	green int
	blue  int
}

func (inventory cubes) isPossible(cubes cubes) bool {
	return inventory.red >= cubes.red && inventory.green >= cubes.green && inventory.blue >= cubes.blue
}

func parseGameLine(l string) (id int, rounds []cubes) {
	split1 := strings.SplitN(l[5:], ": ", 2)
	id, err := strconv.Atoi(split1[0])
	if err != nil {
		panic(err)
	}

	for _, round := range strings.Split(split1[1], "; ") {
		r := cubes{}
		for _, col := range strings.Split(round, ", ") {
			split2 := strings.Split(col, " ")
			count, err := strconv.Atoi(split2[0])
			if err != nil {
				panic(err)
			}

			if strings.HasSuffix(split2[1], "red") {
				r.red = count
			} else if strings.HasSuffix(split2[1], "green") {
				r.green = count
			} else if strings.HasSuffix(split2[1], "blue") {
				r.blue = count
			} else {
				panic("wtf")
			}
		}
		rounds = append(rounds, r)
	}
	return
}

func Part1(s string) int {
	inventory := cubes{
		red:   12,
		green: 13,
		blue:  14,
	}

	rv := 0

LineLoop:
	for _, l := range strings.Split(s, "\n") {
		if l == "" {
			continue
		}
		id, rounds := parseGameLine(l)
		for _, round := range rounds {
			if !inventory.isPossible(round) {
				continue LineLoop
			}
		}
		rv += id
	}
	return rv
}

func Part2(s string) int {
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	rv := 0
	for _, l := range strings.Split(s, "\n") {
		needed := cubes{}

		if l == "" {
			continue
		}
		_, rounds := parseGameLine(l)
		for _, round := range rounds {
			needed.red = max(needed.red, round.red)
			needed.green = max(needed.green, round.green)
			needed.blue = max(needed.blue, round.blue)
		}
		power := needed.red * needed.green * needed.blue
		rv += power
	}
	return rv
}
