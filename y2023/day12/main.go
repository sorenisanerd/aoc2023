package main

import (
	"fmt"
	"strings"

	"github.com/sorenisanerd/aoc2023/utils"
)

func main() {
	fmt.Println(Parts(utils.GetDayData(2023, 12)))
}

func Parts(s string) (int, int) {
	return Part(s, false), Part(s, true)
}

var memo map[string]int

func init() {
	memo = map[string]int{}
}

func solve(s string, nums []int, counter int) int {
	key := fmt.Sprintf("%v%v%v", s, nums, counter)
	if val, ok := memo[key]; ok {
		return val
	}
	memo[key] = solve_(s, nums, counter)
	return memo[key]
}

func solve_(s string, nums []int, counter int) int {
	if s == "" {
		if counter > 0 {
			if len(nums) == 1 && nums[0] == counter {
				return 1
			}
		}
		if len(nums) == 0 {
			return 1
		} else {
			return 0
		}
	}
	if len(nums) == 0 {
		if strings.Contains(s, "#") {
			return 0
		}
	}
	switch s[0] {
	case '.':
		if counter > 0 {
			if counter == nums[0] {
				return solve(s[1:], nums[1:], 0)
			} else {
				return 0
			}
		}
		return solve(s[1:], nums, 0)
	case '#':
		if counter > nums[0] {
			return 0
		}
		return solve(s[1:], nums, counter+1)
	case '?':
		return solve("."+s[1:], nums, counter) + solve("#"+s[1:], nums, counter)
	}
	panic("whaat?")
}

func Part(s string, part2 bool) int {
	ans := 0
	for _, l := range strings.Split(s, "\n") {
		if l == "" {
			continue
		}
		words := strings.Split(l, " ")
		groups := utils.ExtractInts(words[1])
		springs := words[0]
		if part2 {
			springs = strings.Join([]string{springs, springs, springs, springs, springs}, "?")
			groups = utils.ExtractInts(strings.Join([]string{words[1], words[1], words[1], words[1], words[1]}, ","))
		}
		ans += solve(springs, groups, 0)
	}
	return ans
}
