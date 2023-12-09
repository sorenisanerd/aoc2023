package main

import (
	"testing"

	"github.com/sorenisanerd/aoc2023/utils"
	"gotest.tools/assert"
)

var sample = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

func TestParts(t *testing.T) {
	p1, p2 := Parts(sample)
	assert.Equal(t, 114, p1)
	assert.Equal(t, 2, p2)

	p1, p2 = Parts(utils.GetDayData(2023, 9))
	assert.Equal(t, 1980437560, p1)
	assert.Equal(t, 977, p2)
}
