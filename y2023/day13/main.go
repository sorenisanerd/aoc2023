package main

import (
	"fmt"
	"strings"

	"github.com/sorenisanerd/aoc2023/utils"
)

type XY = utils.XY

var sample = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

func main() {
	fmt.Println(Parts(utils.GetDayData(2023, 13)))
}

func Parts(s string) (int, int) {
	return Part(s, false), Part(s, true)
}

func MirrorScore(g utils.Grid[byte], part2 bool) int {
	rv := 0
	wantSmudges := 0
	if part2 {
		wantSmudges = 1
	}
	for x := 0; x < len(g[0])-1; x++ {
		smudges := 0
		for dx := 0; dx <= x; dx++ {
			left := x - dx
			right := x + dx + 1
			if 0 <= left && right < len(g[0]) && left < right {
				for y := range g {
					if g.Get(XY{left, y}) != g.Get(XY{right, y}) {
						smudges += 1
					}
				}
			}
		}
		if smudges == wantSmudges {
			rv += x + 1
		}
	}
	for y := 0; y < len(g)-1; y++ {
		smudges := 0
		for dy := 0; dy <= y; dy++ {
			up := y - dy
			down := y + dy + 1
			if 0 <= up && down < len(g) && up < down {
				for x := range g[y] {
					if g.Get(XY{x, up}) != g.Get(XY{x, down}) {
						smudges += 1
					}
				}
			}
		}
		if smudges == wantSmudges {
			rv += 100 * (y + 1)
		}
	}
	return rv
}

func Part(s string, part2 bool) int {
	ans := 0
	for _, grid := range strings.Split(s, "\n\n") {
		g := utils.ParseGrid[byte](grid)
		ans += MirrorScore(g, part2)
	}
	return ans
}
