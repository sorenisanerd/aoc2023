package main

import (
	"testing"

	"github.com/sorenisanerd/aoc2023/utils"
	"gotest.tools/assert"
)

var sample = `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`

func TestParts(t *testing.T) {
	p1, p2 := Parts(sample)
	assert.Equal(t, 21, p1)
	assert.Equal(t, 525152, p2)

	p1, p2 = Parts(utils.GetDayData(2023, 12))
	assert.Equal(t, 7195, p1)
	assert.Equal(t, 33992866292225, p2)
}
