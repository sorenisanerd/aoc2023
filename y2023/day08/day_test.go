package main

import (
	"testing"

	"github.com/sorenisanerd/aoc2023/utils"
	"gotest.tools/assert"
)

func TestPart1Sample(t *testing.T) {
	p1 := Part(sample1, false)
	assert.Equal(t, 2, p1)

	p1 = Part(sample2, false)
	assert.Equal(t, 6, p1)
}

func TestPart2Sample(t *testing.T) {
	p2 := Part(sample3, true)
	assert.Equal(t, 6, p2)
}

func TestFull(t *testing.T) {
	p1, p2 := Parts(utils.GetData("2023/day8.txt"))
	assert.Equal(t, 14257, p1)
	assert.Equal(t, 16187743689077, p2)
}
