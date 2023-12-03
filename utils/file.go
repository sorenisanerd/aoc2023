package utils

import (
	"io"
	"os"

	"github.com/sorenisanerd/aoc2023/data"
)

func GetData(fname string) string {
	buf, err := data.Content.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	return string(buf)
}

func MustRead(fname string) string {
	fp, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	buf, err := io.ReadAll(fp)
	if err != nil {
		panic(err)
	}
	return string(buf)
}
