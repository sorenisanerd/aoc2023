package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/bobg/go-generics/v3/maps"
	"github.com/sorenisanerd/aoc2023/utils"
)

var sample = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

func main() {
	fmt.Println(Parts(sample))
	fmt.Println(Parts(utils.GetData("2023/day7.txt")))
}

const (
	fiveOfAKind  = iota // 0
	fourOfAKind         // 1
	fullHouse           // 2
	threeOfAKind        // 3
	twoPair             // 4
	onePair             // 5
	highCard            // 6
)

func handType(hand string, part2 bool) int {
	cards := strings.Split(hand, "")
	slices.Sort(cards)

	counter := make(map[string]int)
	for _, c := range cards {
		counter[c] += 1
	}

	if part2 {
		mostCommon := ""
		prevMost := 0
		for c, count := range counter {
			if c == "J" {
				continue
			}
			if count > prevMost {
				mostCommon = c
				prevMost = count
			} else if count == prevMost {
				if toNum(c[0], true) > toNum(mostCommon[0], true) {
					mostCommon = c
				}
			}
		}

		// They're all J's!!! Let's lie and say they're all aces
		if prevMost == 0 {
			mostCommon = "A"
		}
		for idx, c := range cards {
			if c == "J" {
				cards[idx] = mostCommon
			}
		}
		return handType(strings.Join(cards, ""), false)
	}

	counts := maps.Values(counter)
	slices.Sort(counts)
	slices.Reverse(counts)

	if len(counter) == 1 {
		return fiveOfAKind
	}

	if counts[0] == 4 {
		return fourOfAKind
	}

	if counts[0] == 3 {
		if len(counts) == 2 {
			return fullHouse
		} else {
			return threeOfAKind
		}
	}

	if counts[0] == 2 {
		if counts[1] == 2 {
			return twoPair
		} else {
			return onePair
		}
	}

	return highCard
}

func (h hand) handType(part2 bool) int {
	return handType(h.cards, part2)
}

func (h hand) Compare(other hand) int {
	return compHand(h.cards, other.cards, false)
}

func toNum(c byte, part2 bool) int {
	switch c {
	case 'T':
		return 10
	case 'J':
		if part2 {
			return 1
		}
		return 11
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	default:
		return int(c - '0')
	}
}

func compHand(a, b string, part2 bool) int {
	for i := 0; i < len(a); i++ {
		if toNum(a[i], part2) > toNum(b[i], part2) {
			return 1
		} else if toNum(a[i], part2) < toNum(b[i], part2) {
			return -1
		}
	}
	return 0
}

type hand struct {
	cards string
	bid   int
}

func Parts(s string) (int, int) {
	return Part(s, false), Part(s, true)
}

func Part(s string, part2 bool) int {
	lines := strings.Split(s, "\n")
	hands := []hand{}
	for _, l := range lines {
		if l == "" {
			continue
		}

		elems := strings.Split(l, " ")
		cards := elems[0]
		bidString := elems[len(elems)-1]
		bid, _ := strconv.Atoi(bidString)

		hands = append(hands, hand{cards, bid})
	}

	slices.SortFunc(hands, func(a, b hand) int {
		if a.handType(part2) > b.handType(part2) {
			return -1
		} else if b.handType(part2) > a.handType(part2) {
			return 1
		}
		return compHand(a.cards, b.cards, part2)
	})

	p1 := 0
	for rank, h := range hands {
		p1 += h.bid * (rank + 1)
	}

	return p1
}
