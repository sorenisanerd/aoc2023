package main

import (
	"testing"

	"github.com/sorenisanerd/aoc2023/utils"
	"gotest.tools/assert"
)

func TestParts(t *testing.T) {
	p1, p2 := Parts(utils.GetDayData(2015, 2))
	assert.Equal(t, 1606483, p1)
	assert.Equal(t, 3842356, p2)
}
