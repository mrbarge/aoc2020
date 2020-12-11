package main

import (
	"aoc2020/helper"
	"fmt"
	"os"
)

type SeatState int

const (
	EMPTY = iota
	OCCUPIED
	FLOOR
)

func convertData(data []string) [][]SeatState {

	ss := make([][]SeatState, 0)

	for _, row := range data {
		srow := make([]SeatState, len(row))
		for i, c := range row {
			switch c {
			case 'L':
				srow[i] = EMPTY
			case '#':
				srow[i] = OCCUPIED
			case '.':
				srow[i] = FLOOR
			}
		}
		ss = append(ss, srow)
	}
	return ss
}

func partOne(data []string) int {

	plane := convertData(data)

	diff := 1
	for diff > 0 {
		plane, diff = round(plane)
	}
	occ := countOccupied(plane)

	return occ
}

func countOccupied(plane [][]SeatState) (occupied int) {
	for _, row := range plane {
		for _, seat := range row {
			if seat == OCCUPIED {
				occupied++
			}
		}
	}
	return occupied
}

func countOccupiedNeighbours(plane [][]SeatState, c helper.Coord) int {
	neighbours := c.GetNeighboursPos()
	occupied := 0

	for _, n := range neighbours {
		if n.X >= len(plane[0]) || n.Y >= len(plane) {
			continue
		}
		neighbourSeat := plane[n.Y][n.X]
		if neighbourSeat == OCCUPIED {
			occupied++
		}
	}

	//fmt.Printf("Neighbours for %v are %v (%v occupied)\n", c, neighbours, occupied)

	return occupied
}

func round(plane [][]SeatState) ([][]SeatState, int) {
	nextPlane := make([][]SeatState, len(plane))
	for y, r := range plane {
		nr := make([]SeatState, len(r))
		for i, ns := range r {
			nr[i] = ns
		}
		nextPlane[y] = nr
	}

	numChanges := 0
	for y, row := range plane {
		for x, seat := range row {
			coord := helper.Coord{x, y}

			switch seat {
			case EMPTY:
				occupiedNeighbours := countOccupiedNeighbours(plane, coord)
				if occupiedNeighbours == 0 {
					nextPlane[y][x] = OCCUPIED
					numChanges++
				} else {
					nextPlane[y][x] = EMPTY
				}
			case OCCUPIED:
				occupiedNeighbours := countOccupiedNeighbours(plane, coord)
				if occupiedNeighbours >= 4 {
					nextPlane[y][x] = EMPTY
					numChanges++
				} else {
					nextPlane[y][x] = OCCUPIED
				}
			case FLOOR:
				nextPlane[y][x] = FLOOR
			}
		}
	}
	return nextPlane, numChanges
}

func printSeatMap(plane [][]SeatState) {
	for _, row := range plane {
		for _, seat := range row {
			switch seat {
			case EMPTY:
				fmt.Print("L")
			case FLOOR:
				fmt.Print(".")
			case OCCUPIED:
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	fh, _ := os.Open("input.txt")
	data, _ := helper.ReadLines(fh, true)
	ans := partOne(data)
	fmt.Printf("Part one: %v\n", ans)
	//ans2 := partTwo(data, ans)
	//fmt.Printf("Part two: %v\n", ans2)

}
