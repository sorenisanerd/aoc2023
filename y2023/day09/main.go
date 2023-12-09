package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/sorenisanerd/aoc2023/utils"
)

func main() {
	fmt.Println(Parts(utils.GetData("2023/day9.txt")))
}

func extrapolate(nums []int, part2 bool) int {
	levels := [][]int{nums}
	for {
		if slices.Max(nums) == slices.Min(nums) && nums[0] == 0 {
			break
		}
		level := []int{}
		prev := nums[0]
		for _, n := range nums[1:] {
			level = append(level, n-prev)
			prev = n
		}
		levels = append(levels, level)
		nums = level
	}
	nums = append(nums, 0)
	for i := len(levels) - 1; i >= 0; i-- {
		level := levels[i]
		if !part2 {
			level = slices.Insert(level, len(level), level[len(level)-1]+nums[len(nums)-1])
		} else {
			level = slices.Insert(level, 0, level[0]-nums[0])
		}
		nums = level
	}
	if !part2 {
		return nums[len(nums)-1]
	} else {
		return nums[0]
	}
}

func Parts(s string) (int, int) {
	p1 := 0
	p2 := 0
	for _, line := range strings.Split(s, "\n") {
		if line == "" {
			continue
		}
		nums := utils.ExtractInts(line)
		p1 += extrapolate(nums, false)
		p2 += extrapolate(nums, true)
	}
	return p1, p2
}
