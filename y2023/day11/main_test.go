package main

import (
	"testing"

	"github.com/sorenisanerd/aoc2023/utils"
	"gotest.tools/assert"
)

func TestParts(t *testing.T) {
	p1, p2 := Parts(sample)
	assert.Equal(t, 374, p1)
	assert.Equal(t, 82000210, p2)

	p1, p2 = Parts(utils.GetDayData(2023, 11))
	assert.Equal(t, 9312968, p1)
	assert.Equal(t, 597714117556, p2)
}
