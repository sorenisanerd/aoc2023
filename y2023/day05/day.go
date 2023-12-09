package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/sorenisanerd/aoc2023/utils"
)

type myMap []myMapEntry

type myMapEntry struct {
	to   int
	from int
	len  int
}

func (m myMap) Map(id int, bound int) (int, int) {
	for _, e := range m {
		if e.from > id {
			bound = min(e.from, bound)
		}
		if id >= e.from && id < (e.from+e.len) {
			return e.to + (id - e.from), min(bound, (e.len - (id - e.from)))
		}
	}
	return id, bound
}

var sample = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
`

func main() {
	Parts(utils.GetDayData(2023, 5))
}

func ExtractInts(s string) []int {
	rv := []int{}
	for _, word := range regexp.MustCompile(`\s+`).Split(s, -1) {
		if num, err := strconv.Atoi(word); err == nil {
			rv = append(rv, num)
		}
	}
	return rv
}

func Parts(s string) (int, int) {
	p1 := 99999999999999
	p2 := 99999999999999
	var seeds []int
	var seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemperature, temperatureToHumidity, humidityToLocation myMap
	var curMap *myMap
	for _, l := range strings.Split(s, "\n") {
		switch true {
		case l == "":
			continue
		case strings.HasPrefix(l, "seeds:"):
			seeds = ExtractInts(l)
		case strings.HasPrefix(l, "seed-to-soil map"):
			curMap = &seedToSoil
		case strings.HasPrefix(l, "soil-to-fertilizer map"):
			curMap = &soilToFertilizer
		case strings.HasPrefix(l, "fertilizer-to-water map"):
			curMap = &fertilizerToWater
		case strings.HasPrefix(l, "water-to-light map"):
			curMap = &waterToLight
		case strings.HasPrefix(l, "light-to-temperature map"):
			curMap = &lightToTemperature
		case strings.HasPrefix(l, "temperature-to-humidity map"):
			curMap = &temperatureToHumidity
		case strings.HasPrefix(l, "humidity-to-location map"):
			curMap = &humidityToLocation
		default:
			nums := ExtractInts(l)
			if len(nums) != 3 {
				panic("argh")
			}
			*curMap = append(*curMap, myMapEntry{nums[0], nums[1], nums[2]})
		}
	}
	for i := 0; i < len(seeds); i += 2 {
		for j := seeds[i]; j < (seeds[i] + seeds[i+1]); {
			val, inc := humidityToLocation.Map(
				temperatureToHumidity.Map(
					lightToTemperature.Map(
						waterToLight.Map(
							fertilizerToWater.Map(
								soilToFertilizer.Map(
									seedToSoil.Map(j, 9999999999999)))))))
			p2 = min(p2, val)
			j += inc
		}

	}
	for _, seed := range seeds {
		val, _ := humidityToLocation.Map(
			temperatureToHumidity.Map(
				lightToTemperature.Map(
					waterToLight.Map(
						fertilizerToWater.Map(
							soilToFertilizer.Map(
								seedToSoil.Map(seed, 999999999999)))))))
		p1 = min(p1, val)
	}
	fmt.Println(p1, p2)
	return p1, p2
}
