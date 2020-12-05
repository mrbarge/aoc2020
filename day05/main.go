package main

import (
	"aoc2020/helper"
	"fmt"
	"math"
	"os"
	"sort"
)

func getSeatId(seat string) (int, error) {

	var (
		row, minRow, maxRow int
		col, minCol, maxCol int
	)

	if len(seat) < 10 {
		return 0, fmt.Errorf("invalid seat: %s", seat)
	}

	maxRow = 127
	maxCol = 7
	for r := 0; r < 7; r++ {
		switch seat[r] {
		case 'F':
			midPoint := math.Floor(((float64(minRow) + float64(maxRow) + 2.0) / 2.0) - 1.0)
			maxRow = int(midPoint)
		case 'B':
			midPoint := math.Ceil(((float64(minRow) + float64(maxRow) + 2.0) / 2.0) - 1.0)
			minRow = int(midPoint)
		}
	}
	row = minRow

	for r := 7; r < 10; r++ {
		switch seat[r] {
		case 'L':
			midPoint := math.Floor(((float64(minCol) + float64(maxCol) + 2.0) / 2.0) - 1.0)
			maxCol = int(midPoint)
		case 'R':
			midPoint := math.Ceil(((float64(minCol) + float64(maxCol) + 2.0) / 2.0) - 1.0)
			minCol = int(midPoint)
		}
	}
	col = minCol

	return (row * 8) + col, nil
}

func problem(data []string) (part1 int, part2 int) {

	seatIds := make([]int, 0)

	for _, seat := range data {
		seatId, err := getSeatId(seat)
		if err != nil {
			fmt.Printf("Error for seat %v: %v\n", seat, err)
			continue
		}
		seatIds = append(seatIds, seatId)
	}

	sort.Ints(seatIds)
	fmt.Println(seatIds)
	part1 = seatIds[len(seatIds)-1]

	i := 0
	for v := seatIds[0]; v < seatIds[len(seatIds)-1]; v++ {
		if seatIds[i] != v {
			part2 = v
			break
		}
		i++
	}
	return part1, part2
}

func main() {
	fh, _ := os.Open("input.txt")
	data, _ := helper.ReadLines(fh, true)
	p1, p2 := problem(data)
	fmt.Printf("Part one: %v\n", p1)
	fmt.Printf("Part two: %v\n", p2)
}
