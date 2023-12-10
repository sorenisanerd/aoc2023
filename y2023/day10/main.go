package main

import (
	"fmt"

	"github.com/hashicorp/go-set/v2"
	"github.com/sorenisanerd/aoc2023/utils"
)

var sample = `...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`

func main() {
	fmt.Println(Parts(sample))
	fmt.Println(Parts(utils.GetDayData(2023, 10)))
}

type XY = utils.XY

var directions = map[byte][2]XY{
	'|': {utils.N, utils.S},
	'-': {utils.E, utils.W},
	'L': {utils.N, utils.E},
	'J': {utils.N, utils.W},
	'7': {utils.W, utils.S},
	'F': {utils.S, utils.E},
}

// Determine which kind of pipe S really is so
// we can substitute it
func findSub(g utils.Grid[byte], start utils.XY) byte {
	canN := false
	canE := false
	canS := false
	canW := false

	p := g.Get(start.Add(utils.N))
	if (p == 'F') || (p == '7') || (p == '|') {
		canN = true
	}

	p = g.Get(start.Add(utils.E))
	if (p == '7') || (p == 'J') || (p == '-') {
		canE = true
	}

	p = g.Get(start.Add(utils.S))
	if (p == 'L') || (p == 'J') || (p == '|') {
		canS = true
	}

	p = g.Get(start.Add(utils.W))
	if (p == 'L') || (p == 'F') || (p == '-') {
		canS = true
	}

	if canE && canW {
		return '-'
	}

	if canN && canW {
		return 'J'
	}

	if canN && canS {
		return '|'
	}

	if canE && canS {
		return 'F'
	}

	if canN && canE {
		return 'L'
	}

	if canW && canS {
		return '7'
	}
	panic("wtf?")
}

func Parts(s string) (int, int) {
	p1 := 0
	p2 := 0
	g := utils.ParseGrid[byte](s)

	start := g.FindFirst('S')
	sub := findSub(g, start)
	g.Set(start, sub)

	allPipes := set.From([]XY{start})
	curPos := start
	ds := directions[sub]
	step := ds[0]
	length := 1
	for {
		nextPos := curPos.Add(step)
		allPipes.Insert(nextPos)
		if nextPos == start {
			break
		}
		ds = directions[g.Get(nextPos)]
		s := set.From([]XY{ds[0], ds[1]})
		s.Remove(XY{X: -step.X, Y: -step.Y})
		if s.Size() != 1 {
			panic("asdfasdf")
		}
		step = s.Slice()[0]
		length += 1
		curPos = nextPos
	}
	p1 = length
	for y, l := range g {
		// Starting from the edge of the grid, we're always outside the pipe
		inside := false
		onpipe := false
		fromUp := false
		for x := range l {
			xy := XY{X: x, Y: y}
			if allPipes.Contains(xy) {
				shape := g.Get(xy)
				// Moving horizontally, crossing a | means toggling `inside`
				if shape == '|' {
					inside = !inside
				}
				// Examples:
				// In `.L----J.` both '.' are outside.
				// In `.L----7.` the first '.' is outside while the second is inside.
				if (shape == 'L') || (shape == 'F') || (shape == '7') || (shape == 'J') {
					if !onpipe {
						// we just stepped onto the pipe. Record whether the pipe bent up or down
						fromUp = (shape == 'L') || (shape == 'J')
						onpipe = true
					} else {
						// We're now stepping off the pipe
						onpipe = !onpipe
						// Whether we toggle `inside` depends on whether we started with a
						// turn that bent up or down. See examples above.
						if fromUp != ((shape == 'L') || (shape == 'J')) {
							inside = !inside
						}
					}
				}
			} else {
				if inside {
					p2 += 1
				}
			}
		}
	}
	return p1 / 2, p2
}
