package main

import (
	md5_ "crypto/md5"
	"fmt"

	"github.com/sorenisanerd/aoc2023/utils"
)

var sample = `OOOO.#.O..
OO..#....#
OO..O##..O
O..#.OO...
........#.
..#....#.#
..O..#.O.O
..O.......
#....###..
#....#....`

func main() {
	//	fmt.Println(Parts(sample))
	fmt.Println(Parts(utils.GetDayData(2023, 14)))
}

func Parts(s string) (int, int) {
	return Part(s, false), Part(s, true)
}

func rollGrid(g utils.Grid[byte], dir utils.XY) {
	var startPos, iter utils.XY

	if dir == utils.N {
		startPos = utils.XY{0, 0}
		iter = utils.S
	} else if dir == utils.E {
		startPos = utils.XY{len(g[0]) - 1, 0}
		iter = utils.W
	} else if dir == utils.S {
		startPos = utils.XY{0, len(g) - 1}
		iter = utils.N
	} else if dir == utils.W {
		startPos = utils.XY{0, 0}
		iter = utils.E
	}

	pos := startPos
	for {
		c := g.Get(pos)
		switch c {
		case '#':
			/* nothing */
		case 'O':
			newPos := pos
		Move:
			for {
				tryPos := newPos.Add(dir)
				if !g.WithinBounds(tryPos) {
					break
				}
				c := g.Get(tryPos)
				switch c {
				case '.':
					newPos = tryPos
				case '#':
					break Move
				case 'O':
					break Move
				default:
					panic("unknown char")
				}
			}
			//				fmt.Println("Moved rock from", xy, "to", newPos)
			g.Set(pos, '.') // remove rock
			g.Set(newPos, 'O')
		case '.':
			/* nothing */
		default:
			panic("unknown char")
		}
		pos = pos.Add(iter)
		if !g.WithinBounds(pos) {
			if dir == utils.N {
				pos = utils.XY{pos.X + 1, 0}
			} else if dir == utils.E {
				pos = utils.XY{len(g[0]) - 1, pos.Y + 1}
			} else if dir == utils.S {
				pos = utils.XY{pos.X + 1, len(g) - 1}
			} else if dir == utils.W {
				pos = utils.XY{0, pos.Y + 1}
			}
			if !g.WithinBounds(pos) {
				break
			}
		}
	}
}

func md5(s string) string {
	hasher := md5_.New()
	hasher.Write([]byte(s))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

func Part(s string, part2 bool) int {
	rv := 0
	g := utils.ParseGrid[byte](s)
	if !part2 {
		rollGrid(g, utils.N)
	} else {
		seen := map[string]int{}
		for i := 0; i < 2000; i++ {
			sum := md5(g.String())
			if prev, ok := seen[sum]; ok {
				period := i - prev
				if 1000000000%period == i%period {
					break
				}
			} else {
				seen[sum] = i
			}
			rollGrid(g, utils.N)
			rollGrid(g, utils.W)
			rollGrid(g, utils.S)
			rollGrid(g, utils.E)
		}
	}
	g.FindAll('O').ForEach(func(rock utils.XY) bool {
		rv += (len(g) - rock.Y)
		return true
	})
	return rv
}
