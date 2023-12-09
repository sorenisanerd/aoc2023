package main

import (
	"testing"

	"github.com/sorenisanerd/aoc2023/utils"
	"gotest.tools/assert"
)

func TestParts(t *testing.T) {
	p1, p2 := Parts(sample)
	assert.Equal(t, 6440, p1)
	assert.Equal(t, 5905, p2)

	p1, p2 = Parts(utils.GetDayData(2023, 7))
	assert.Equal(t, 246912307, p1)
	assert.Equal(t, 246894760, p2)
}
