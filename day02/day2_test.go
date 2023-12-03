package day02

import (
	"testing"

	"github.com/sorenisanerd/aoc2023/utils"
	"gotest.tools/assert"
)

var day2sample = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

func Test_Part1(t *testing.T) {
	assert.Equal(t, 8, Part1(day2sample))
}

func Test_Part2(t *testing.T) {
	assert.Equal(t, 2286, Part2(day2sample))
}

func Test_Part1Full(t *testing.T) {
	assert.Equal(t, 2283, Part1(utils.GetData("day2.txt")))
}

func Test_Part2Full(t *testing.T) {
	assert.Equal(t, 78669, Part2(utils.GetData("day2.txt")))
}
