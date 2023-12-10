package utils

import "strings"

type Char interface {
	byte | rune
}

type Grid[T Char] []string

func (g Grid[T]) Get(xy XY) T {
	return T(g[xy.Y][xy.X])
}

func (g Grid[T]) Set(xy XY, r T) T {
	g[xy.Y] = ReplaceAtIndex(g[xy.Y], xy.X, rune(r))
	return r
}

func (g Grid[T]) FindFirst(c T) XY {
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[y]); x++ {
			xy := XY{x, y}
			if g.Get(xy) == c {
				return xy
			}
		}
	}
	return NotFound
}

var NotFound = XY{-2134124124211, -124312412341}

func GetLines(s string) []string {
	return strings.Split(s, "\n")
}

func ParseGrid[T Char](s string) Grid[T] {
	lines := GetLines(s)
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	return Grid[T](lines)
}

func (g Grid[T]) String() string {
	return strings.Join(g, "\n")
}

func (g Grid[T]) WithinBounds(xy XY) bool {
	if xy.Y < 0 || xy.Y >= len(g) {
		return false
	}
	if xy.X < 0 || xy.X >= len(g[xy.Y]) {
		return false
	}
	return true
}

type XY struct {
	X int
	Y int
}

func NewXY(x int, y int) XY {
	return XY{X: x, Y: y}
}

var Up = XY{0, -1}
var Down = XY{0, 1}
var Left = XY{-1, 0}
var Right = XY{1, 0}

var North = Up
var South = Down
var East = Right
var West = Left

var N = North
var S = South
var E = East
var W = West

var NE = XY{1, -1}
var SE = XY{1, 1}
var NW = XY{-1, -1}
var SW = XY{-1, 1}

var FourDirections = []XY{Up, Down, Left, Right}
var EightDirections = append(FourDirections, NW, NE, SE, SW)

func (xy XY) Add(other XY) XY {
	xy.X += other.X
	xy.Y += other.Y
	return xy
}
