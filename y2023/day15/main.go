package main

import (
	"fmt"

	"github.com/sorenisanerd/aoc2023/utils"
)

var sample = ``

func main() {
	fmt.Println(Parts(utils.GetDayData(2023, 15)))
}

func Parts(s string) (int, int) {
	return Part(s, false), Part(s, true)
}

func Part(s string, part2 bool) int {
	g := utils.ParseGrid[byte](s)
	fmt.Println(g)
	return 0
}
