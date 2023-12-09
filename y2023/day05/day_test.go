package main

import (
	"testing"

	"github.com/sorenisanerd/aoc2023/utils"
	"gotest.tools/assert"
)

func TestParts(t *testing.T) {
	p1, p2 := Parts(sample)
	assert.Equal(t, 35, p1)
	assert.Equal(t, 46, p2)

	p1, p2 = Parts(utils.GetDayData(2023, 5))
	assert.Equal(t, 346433842, p1)
	assert.Equal(t, 60294664, p2)
}
