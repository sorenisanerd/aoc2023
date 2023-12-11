package main

import (
	"fmt"

	"slices"

	"github.com/hashicorp/go-set/v2"
	"github.com/sorenisanerd/aoc2023/utils"
)

var sample = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

func main() {
	fmt.Println(Parts(utils.GetDayData(2023, 11)))
}

func expand(galaxies *set.Set[utils.XY], clearRows, clearCols *set.Set[int], mult int) *set.Set[utils.XY] {
	rows := clearRows.Slice()
	slices.Sort(rows)
	slices.Reverse(rows)
	cols := clearCols.Slice()
	slices.Sort(cols)
	slices.Reverse(cols)

	newGalaxies := set.From([]utils.XY{})
	galaxies.ForEach(func(xy utils.XY) bool {
		dx, dy := 0, 0
		for _, row := range rows {
			if row < xy.Y {
				dy += mult - 1
			}
		}
		for _, col := range cols {
			if col < xy.X {
				dx += mult - 1
			}
		}
		newGalaxies.Insert(xy.Add(utils.XY{X: dx, Y: dy}))

		return true
	})
	return newGalaxies
}

func Parts(s string) (int, int) {
	return Part(s, false), Part(s, true)
}

func Part(s string, part2 bool) int {
	mult := 2
	if part2 {
		mult = 1_000_000
	}
	res := 0
	g := utils.ParseGrid[byte](s)
	clearRows := set.From([]int{})
	clearCols := set.From([]int{})
	for y := 0; y < len(g); y++ {
		clearRows.Insert(y)
	}

	for x := 0; x < len(g[0]); x++ {
		clearCols.Insert(x)
	}

	for y, l := range g {
		for x, c := range l {
			if c != '.' {
				clearCols.Remove(x)
				clearRows.Remove(y)
			}
		}
	}

	galaxiess := set.From([]utils.XY{})
	for y, l := range g {
		for x, c := range l {
			if c != '.' {
				galaxiess.Insert(utils.XY{X: x, Y: y})
			}
		}
	}

	galaxiess = expand(galaxiess, clearRows, clearCols, mult)

	galaxies := galaxiess.Slice()
	for i := 0; i < len(galaxies)-1; i++ {
		g1 := galaxies[i]
		for j := i + 1; j < len(galaxies); j++ {
			g2 := galaxies[j]
			dist := g1.MD(g2)
			res += dist
		}
	}

	return res
}
