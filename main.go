package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/sorenisanerd/aoc2023/day01"
	"github.com/sorenisanerd/aoc2023/day02"
	"github.com/sorenisanerd/aoc2023/utils"
)

func getInput(day int) io.Reader {
	url := fmt.Sprintf("https://adventofcode.com/2023/day/%d/input", day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: aocCookie,
	})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	return resp.Body
}

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "getInput" {
			day, err := strconv.Atoi(os.Args[2])
			if err != nil {
				panic(err)
			}

			inputReader := getInput(day)
			fp, err := os.Create(fmt.Sprintf("data/day%d.txt", day))
			if err != nil {
				panic(err)
			}
			io.Copy(fp, inputReader)
			return
		}
	}
	fmt.Println(day01.Part1(utils.GetData("day1.txt")))
	fmt.Println(day01.Part2(utils.GetData("day1.txt")))
	fmt.Println(day02.Part1(utils.GetData("day2.txt")))
	fmt.Println(day02.Part2(utils.GetData("day2.txt")))
}
