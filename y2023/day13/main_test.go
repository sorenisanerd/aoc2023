package main

import (
	"testing"

	"github.com/sorenisanerd/aoc2023/utils"
	"gotest.tools/assert"
)

func TestParts(t *testing.T) {
	p1, p2 := Parts(sample)
	assert.Equal(t, 0, p1)
	assert.Equal(t, 0, p2)

	p1, p2 = Parts(utils.GetDayData(2023, 13))
	assert.Equal(t, 0, p1)
	assert.Equal(t, 0, p2)
}
