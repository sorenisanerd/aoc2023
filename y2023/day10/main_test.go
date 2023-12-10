package main

import (
	"testing"

	"github.com/sorenisanerd/aoc2023/utils"
	"gotest.tools/assert"
)

func TestParts(t *testing.T) {
	p1, p2 := Parts(sample)
	assert.Equal(t, 23, p1)
	assert.Equal(t, 4, p2)

	p1, p2 = Parts(utils.GetDayData(2023, 10))
	assert.Equal(t, 7093, p1)
	assert.Equal(t, 407, p2)
}
