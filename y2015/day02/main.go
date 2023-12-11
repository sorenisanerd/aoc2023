package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/sorenisanerd/aoc2023/utils"
)

func main() {
	fmt.Println(Parts(utils.GetDayData(2015, 2)))
}

func Parts(s string) (int, int) {
	p1 := 0
	p2 := 0
	for _, l := range strings.Split(s, "\n") {
		if l == "" {
			continue
		}
		dims := utils.ExtractInts(l)
		a1 := dims[0] * dims[1]
		a2 := dims[1] * dims[2]
		a3 := dims[2] * dims[0]

		minSide := slices.Min([]int{a1, a2, a3})
		p1 += minSide + 2*a1 + 2*a2 + 2*a3

		slices.Sort(dims)
		p2 += dims[0]*dims[1]*dims[2] + 2*dims[0] + 2*dims[1]
	}
	return p1, p2
}
