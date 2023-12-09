package main

import (
	"fmt"

	"github.com/hashicorp/go-set/v2"
	"github.com/sorenisanerd/aoc2023/utils"
)

func main() {
	fmt.Println(Parts(utils.GetData("2023/day3.txt")))
}

type XY = utils.XY

func Part1(s string) int {
	rv, _ := Parts(s)
	return rv
}

func Part2(s string) int {
	_, rv := Parts(s)
	return rv
}

func Parts(s string) (int, int) {
	g := utils.ParseGrid[byte](s)

	p1 := 0
	gears := map[XY][]int{}
	for y, l := range g {
		n := 0
		isPart := false
		gearSet := set.New[XY](0)
		for x := 0; x <= len(l); x++ {
			xy := utils.NewXY(x, y)
			if x < len(l) && utils.IsDigit(l[x]) {
				n = n*10 + int(l[x]-'0')
				for _, dxdy := range utils.EightDirections {
					// Using := here instead of = means we're
					// creating a new xy variable scoped to this
					// block
					if xy := xy.Add(dxdy); g.WithinBounds(xy) {
						c := g.Get(xy)
						if !utils.IsDigit(c) && c != '.' {
							isPart = true
							if c == '*' {
								gearSet.Insert(xy)
							}
						}
					}
				}
			} else if n > 0 {
				if isPart {
					p1 += n
				}
				gearSet.ForEach(func(gear XY) bool {
					gears[gear] = append(gears[gear], n)
					return true
				})
				n = 0
				isPart = false
				gearSet = set.New[XY](0)
			}
		}
	}
	p2 := 0
	for gear := range gears {
		if len(gears[gear]) == 2 {
			p2 += gears[gear][0] * gears[gear][1]
		}
	}

	return p1, p2
}
