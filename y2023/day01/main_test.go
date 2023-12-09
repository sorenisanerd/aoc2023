package main

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
	assert.Equal(t, 142, part(day1sample1, false))
}

func Test_Part2(t *testing.T) {
	assert.Equal(t, 281, part(day1sample2, true))
}

func Test_FullPart1(t *testing.T) {
	assert.Equal(t, 54697, part(utils.GetDayData(2023, 1), false))
}

func Test_FullPart2(t *testing.T) {
	assert.Equal(t, 54885, part(utils.GetDayData(2023, 1), true))
}
