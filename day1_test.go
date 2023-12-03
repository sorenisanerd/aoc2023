package main

import (
	"testing"

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

func Test_day1part1(t *testing.T) {
	assert.Equal(t, 142, day1part1(day1sample1))
}

func Test_day2part2(t *testing.T) {
	assert.Equal(t, 281, day1part2(day1sample2))
}

func Test_getLastDigit(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"123456", 6},
		{"abdsfasdljfhlasdfhljksadhflasjkdhf1", 1},
		{"abdsfasdljfhlasdfh3ljksadhflasjkdhf1", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, getLastDigit(tt.name))
		})
	}
}

func Test_getFirstDigit(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"123456", 1},
		{"abdsfasdljfhlasdfhljksadhflasjkdhf1", 1},
		{"abdsfasdljfhlasdfh3ljksadhflasjkdhf1", 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, getFirstDigit(tt.name))
		})
	}
}
