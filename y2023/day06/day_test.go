package main

import (
	"testing"

	"github.com/sorenisanerd/aoc2023/utils"
	"gotest.tools/assert"
)

func TestParts(t *testing.T) {
	p1, p2 := Parts(sample)
	assert.Equal(t, 288, p1)
	assert.Equal(t, 71503, p2)

	p1, p2 = Parts(utils.GetDayData(2023, 6))
	assert.Equal(t, 781200, p1)
	assert.Equal(t, 49240091, p2)
}
