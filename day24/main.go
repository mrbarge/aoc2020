package main

import (
	"aoc2020/helper"
	"fmt"
	"os"
)

func partOne(data []string) (int, error) {

	return 0, nil
}

func main() {
	fh, _ := os.Open("input.txt")
	data, _ := helper.ReadLines(fh, true)
	p1, err := partOne(data)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part one: %d\n", p1)

}
