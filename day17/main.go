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
	ans, err := partOne(data)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part one: %d\n", ans)

	c := helper.Coord3D{1,2,3}
	d := c.GetNeighbours()
	for dc, _ := range d {
		fmt.Println(dc)
	}
}