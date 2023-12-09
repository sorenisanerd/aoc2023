package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/hashicorp/go-set/v2"
	"github.com/sorenisanerd/aoc2023/utils"
)

func main() {
	fmt.Println(Parts(utils.InputData(2023, 4)))
}

func Parts(s string) (int, int) {
	p1 := 0
	p2 := 0
	counts := map[int]int{}
	cards := strings.Split(s, "\n")
	for _, card := range cards {
		if card == "" {
			continue
		}
		split1 := strings.Split(card, ": ")
		split2 := strings.Split(split1[0], " ")
		cardno, _ := strconv.Atoi(split2[len(split2)-1])
		counts[cardno] += 1

		drawn := set.New[int](0)
		mine := set.New[int](0)
		split3 := strings.Split(split1[1], "|")
		for _, no := range strings.Split(split3[0], " ") {
			if no == "" {
				continue
			}
			drawn.Insert(utils.MustAtoi(no))
		}

		for _, no := range strings.Split(split3[1], " ") {
			if no == "" {
				continue
			}
			no_int, _ := strconv.Atoi(no)
			mine.Insert(no_int)
		}

		correct := mine.Intersect(drawn).Size()
		if correct > 0 {
			score := int(math.Pow(2, float64(correct-1)))
			p1 += score
			for n := 1; n <= correct; n++ {
				if cardno+n > len(cards) {
					break
				}
				counts[cardno+n] += counts[cardno]
			}
		}
	}
	for c := range counts {
		p2 += counts[c]
	}
	return p1, p2
}
