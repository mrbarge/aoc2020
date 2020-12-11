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

func partTwo(data []string) int {

	plane := convertData(data)

	diff := 1
	for diff > 0 {
		plane, diff = roundTwo(plane)
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

func isOccupied(s SeatState) (occupied bool, keepLooking bool) {
	if s == OCCUPIED {
		return true, false
	}
	if s == FLOOR {
		return false, true
	}
	return false, false
}

func countOccupiedNeighboursLOS(plane [][]SeatState, c helper.Coord) int {

	neighbours := 0
	// count left
	x := 0
	y := 0
	for x = c.X - 1; x >= 0 && plane[c.Y][x] == FLOOR; x-- {}
	if x >= 0 && plane[c.Y][x] == OCCUPIED {
		neighbours++
	}
	for x = c.X + 1; x < len(plane[0]) && plane[c.Y][x] == FLOOR; x++ {}
	if x < len(plane[0]) && plane[c.Y][x] == OCCUPIED {
		neighbours++
	}
	for y = c.Y - 1; y >= 0 && plane[y][c.X] == FLOOR; y-- {}
	if y >= 0 && plane[y][c.X] == OCCUPIED {
		neighbours++
	}
	for y = c.Y + 1; y < len(plane) && plane[y][c.X] == FLOOR; y++ {}
	if y < len(plane) && plane[y][c.X] == OCCUPIED {
		neighbours++
	}
	for x, y = c.X-1, c.Y-1; x >= 0 && y >= 0 && plane[y][x] == FLOOR; {x--;y--}
	if x >= 0 && y >= 0 && plane[y][x] == OCCUPIED {
		neighbours++
	}
	for x, y = c.X-1, c.Y+1; x >= 0 && y < len(plane) && plane[y][x] == FLOOR; {x--;y++	}
	if x >= 0 && y < len(plane) && plane[y][x] == OCCUPIED {
		neighbours++
	}
	for x, y = c.X+1, c.Y+1; x < len(plane[0]) && y < len(plane) && plane[y][x] == FLOOR; {x++;y++}
	if x < len(plane[0]) && y < len(plane) && plane[y][x] == OCCUPIED {
		neighbours++
	}
	for x, y = c.X+1, c.Y-1; x < len(plane[0]) && y >= 0 && plane[y][x] == FLOOR; {x++;y--}
	if x < len(plane[0]) && y >= 0 && plane[y][x] == OCCUPIED {
		neighbours++
	}
	return neighbours
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

func roundTwo(plane [][]SeatState) ([][]SeatState, int) {
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
				occupiedNeighbours := countOccupiedNeighboursLOS(plane, coord)
				if occupiedNeighbours == 0 {
					nextPlane[y][x] = OCCUPIED
					numChanges++
				} else {
					nextPlane[y][x] = EMPTY
				}
			case OCCUPIED:
				occupiedNeighbours := countOccupiedNeighboursLOS(plane, coord)
				if occupiedNeighbours >= 5 {
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
	ans2 := partTwo(data)
	fmt.Printf("Part two: %v\n", ans2)

}
