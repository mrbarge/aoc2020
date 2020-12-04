package main

import (
	"aoc2020/helper"
	"fmt"
	"os"
)

type slope struct {
	right int
	down int
}

func partTwo(s []string) (sum int, err error) {
	slopes := []slope {
		{1,1},
		{3,1},
		{5,1},
		{7,1},
		{1,2},
	}

	sum = 1
	for _, slope := range slopes {
		n, err := partOne(s, slope)
		if err != nil {
			return 0, err
		}
		sum *= n
	}

	return sum, nil
}

func partOne(data []string, s slope) (count int, err error) {

	if len(data) == 0 {
		return 0, fmt.Errorf("empty data")
	}

	pos := helper.Coord{
		X: 0,
		Y: 0,
	}

	maxX := len(data[0])
	maxY := len(data)

	for pos.Y < maxY {

		row := data[pos.Y]
		if row[pos.X] == '#' {
			count++
		}
		pos.X = (pos.X + s.right) % maxX
		pos.Y += s.down
	}

	return count, nil
}

func main() {
	fh, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("can't read input file: %v", err)
		os.Exit(1)
	}
	data, _ := helper.ReadLines(fh, false)
	ans, err := partOne(data, slope{3,1})
	fmt.Printf("Part one: %v\n", ans)

	ans, err = partTwo(data)
	fmt.Printf("Part two: %v\n", ans)

}
