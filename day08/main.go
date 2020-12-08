package main

import (
	"aoc2020/console"
	"aoc2020/helper"
	"fmt"
	"os"
)

func partOne(data []string) (int64, error) {
	c, err := console.New(data)
	if err != nil {
		return 0, err
	}
	c.Run()
	return c.Accumulator(), nil
}

func partTwo(data []string) (int64, error) {
	c, err := console.New(data)
	if err != nil {
		return 0, err
	}
	c.FixBrokenCode()
	return c.Accumulator(), nil
}

func main() {
	fh, _ := os.Open("input.txt")
	data, _ := helper.ReadLines(fh, true)
	ans, err := partOne(data)
	if err != nil {
		fmt.Printf("Error in part one: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Part one: %d\n", ans)
	ans, err = partTwo(data)
	if err != nil {
		fmt.Printf("Error in part two: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Part two: %d\n", ans)

}
