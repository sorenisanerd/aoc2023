package main

import (
	"testing"

	"github.com/sorenisanerd/aoc2023/utils"
	"gotest.tools/assert"
)

func TestParts(t *testing.T) {
	p1, p2 := Parts(utils.GetDayData(2015, 3))
	assert.Equal(t, 2592, p1)
	assert.Equal(t, 2360, p2)
}
