package main

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-set/v2"
	"github.com/sorenisanerd/aoc2023/utils"
)

func main() {
	fmt.Println(Parts(utils.GetDayData(2015, 3)))
}

func Parts(s string) (int, int) {
	return Part1(s), Part2(s)
}

var instructions = map[rune]utils.XY{
	'>': utils.E,
	'<': utils.W,
	'^': utils.N,
	'v': utils.S,
}

func Part1(s string) int {
	curPos := utils.XY{X: 0, Y: 0}
	houses := set.From([]utils.XY{curPos})

	for _, l := range strings.Split(s, "\n") {
		if l == "" {
			continue
		}
		for _, c := range l {
			dir := instructions[c]
			newPos := curPos.Add(dir)
			houses.Insert(newPos)
			curPos = newPos
		}
	}
	return houses.Size()
}

func Part2(s string) int {
	santa := utils.XY{X: 0, Y: 0}
	robot := utils.XY{X: 0, Y: 0}
	houses := set.From([]utils.XY{santa})

	for _, l := range strings.Split(s, "\n") {
		if l == "" {
			continue
		}
		for idx, c := range l {
			dir := instructions[c]
			if idx%2 == 0 {
				santa = santa.Add(dir)
				houses.Insert(santa)
			} else {
				robot = robot.Add(dir)
				houses.Insert(robot)
			}
		}
	}
	return houses.Size()
}
