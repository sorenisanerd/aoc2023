package utils

import (
	"testing"

	"gotest.tools/assert"
)

func TestParseGrid(t *testing.T) {
	g := ParseGrid[byte](``)
	assert.Equal(t, 0, len(g))

	g = ParseGrid[byte](`abc`)
	assert.Equal(t, 1, len(g))
	assert.Equal(t, byte('a'), g.Get(NewXY(0, 0)))
	assert.Equal(t, byte('c'), g.Get(NewXY(2, 0)))

	assert.Equal(t, "abc", g.String())

	g.Set(NewXY(1, 0), 'm')
	assert.Equal(t, "amc", g.String())
}

func TestXYAdd(t *testing.T) {
	xy := XY{10, 10}
	assert.Equal(t, XY{10, 10}, xy)

	saved_xy := xy

	assert.Equal(t, NewXY(11, 11), xy.Add(XY{1, 1}))
	assert.Equal(t, NewXY(10, 10), saved_xy)
}
