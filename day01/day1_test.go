package day01

import (
	"testing"

	"github.com/sorenisanerd/aoc2023/utils"
	"gotest.tools/assert"
)

var day1sample1 = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

var day1sample2 = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

func Test_Part1(t *testing.T) {
	assert.Equal(t, 142, Part1(day1sample1))
}

func Test_Part2(t *testing.T) {
	assert.Equal(t, 281, Part2(day1sample2))
}

func Test_FullPart1(t *testing.T) {
	assert.Equal(t, 54697, Part1(utils.GetData("day1.txt")))
}

func Test_FullPart2(t *testing.T) {
	assert.Equal(t, 54885, Part2(utils.GetData("day1.txt")))
}
