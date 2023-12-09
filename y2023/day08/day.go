package main

import (
	"fmt"
	"strings"

	"github.com/sorenisanerd/aoc2023/utils"
	"golang.org/x/exp/maps"
)

var sample1 = `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`

var sample2 = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

var sample3 = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

func main() {
	fmt.Println(Parts(utils.GetData("2023/day8.txt")))
}

func Parts(s string) (int, int) {
	return Part(s, false), Part(s, true)
}

func Part(s string, part2 bool) int {
	lines := strings.Split(s, "\n")
	directions := lines[0]
	next := map[string][2]string{}
	positions := []string{}

	for _, line := range lines[2:] {
		if line == "" {
			continue
		}
		this := line[0:3]
		if line[2] == 'A' {
			positions = append(positions, this)
		}
		left := line[7:10]
		right := line[12:15]
		next[this] = [2]string{left, right}
	}

	if !part2 {
		positions = []string{"AAA"}
	}

	periods := map[int]int{}

	fmt.Println(positions)

	rv := 0
	for ; ; rv++ {
		dirIdx := rv % len(directions)
		dir := string(directions[dirIdx])
		for idx, pos := range positions {
			if dir == "R" {
				positions[idx] = next[pos][1]
			} else if dir == "L" {
				positions[idx] = next[pos][0]
			} else {
				panic("asdfasdf")
			}
			if !part2 {
				if positions[idx] == "ZZZ" {
					return rv + 1
				}
			} else {
				if positions[idx][2] == 'Z' {
					periods[idx] = rv + 1
					if len(periods) == len(positions) {
						ps := maps.Values(periods)
						return utils.LCM(ps[0], ps[1:]...)
					}
				}
			}
		}
	}
}
