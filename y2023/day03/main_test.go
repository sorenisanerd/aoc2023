package main

import (
	"testing"

	"github.com/sorenisanerd/aoc2023/utils"
	"gotest.tools/assert"
)

var day3sample = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func Test_Part1(t *testing.T) {
	assert.Equal(t, 4361, Part1(day3sample))
}

func Test_Part1Full(t *testing.T) {
	assert.Equal(t, 527144, Part1(utils.GetData("2023/day3.txt")))
}

func Test_Part2(t *testing.T) {
	assert.Equal(t, 467835, Part2(day3sample))
}

func Test_Part2Full(t *testing.T) {
	assert.Equal(t, 81463996, Part2(utils.GetData("2023/day3.txt")))
}
