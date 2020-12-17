package main

import (
	"aoc2020/helper"
	"fmt"
	"os"
)

type Coord4D struct {
	X int
	Y int
	Z int
	W int
	V int
}

func (c Coord4D) AsString() string {
	return fmt.Sprintf("%d/%d/%d/%d", c.X,c.Y,c.Z,c.W)
}

func makeNeighbours(data []string) map[string]Coord4D {

	ret := make(map[string]Coord4D)
	z := 0
	y := 0
	w := 0
	for _, line := range data {
		for x, c := range line {
			coord := Coord4D{X:x, Y:y, Z:z, W: w}
			if c == '#' {
				coord.V = 1
			}
			ret[coord.AsString()] = coord
		}
		y++
	}

	for w := -1; w <= 1; w++ {
		for z := -1; z <= 1; z++ {
			for y = -1; y <= len(data); y++ {
				for x := -1; x <= len(data[0]); x++ {
					coord := Coord4D{X: x, Y: y, Z: z, W: w, V: 0}
					if _, ok := ret[coord.AsString()]; !ok {
						ret[coord.AsString()] = coord
					}
				}
			}
		}
	}
	return ret
}

func partTwo(data []string) (int, error) {

	coords := makeNeighbours(data)

	lastActiveCount := 0
	for tick := 0; tick < 6; tick++ {
		activeCount := 0
		nextCoords := make(map[string]Coord4D)
		for _, coord := range coords {
			neighbours := 0

			for w := coord.W-1; w <= coord.W+1; w++ {
				for z := coord.Z - 1; z <= coord.Z+1; z++ {
					for y := coord.Y - 1; y <= coord.Y+1; y++ {
						for x := coord.X - 1; x <= coord.X+1; x++ {
							ncs := Coord4D{X: x, Y: y, Z: z, W: w}
							if _, ok := coords[ncs.AsString()]; !ok {
								coords[ncs.AsString()] = ncs
							}
							if coords[ncs.AsString()].V == 1 {
								neighbours++
							}
						}
					}
				}
			}
			//fmt.Println(neighbours)
			if coord.V == 1 {
				if neighbours >= 3 && neighbours <= 4 {
					newnc := Coord4D{X: coord.X, Y: coord.Y, Z: coord.Z, W: coord.W, V: 1}
					nextCoords[newnc.AsString()] = newnc
				} else {
					newnc := Coord4D{X: coord.X, Y: coord.Y, Z: coord.Z, W: coord.W, V: 0}
					nextCoords[newnc.AsString()] = newnc
				}
			} else {
				if neighbours == 3 {
					newnc := Coord4D{X: coord.X, Y: coord.Y, Z: coord.Z, W: coord.W, V: 1}
					nextCoords[newnc.AsString()] = newnc
				} else {
					newnc := Coord4D{X: coord.X, Y: coord.Y, Z: coord.Z, W: coord.W, V: 0}
					nextCoords[newnc.AsString()] = newnc
				}
			}

			if nextCoords[coord.AsString()].V == 1 {
				activeCount++
				for w := coord.W-1; w <= coord.W+1; w++ {
					for z := coord.Z - 1; z <= coord.Z+1; z++ {
						for y := coord.Y - 1; y <= coord.Y+1; y++ {
							for x := coord.X - 1; x <= coord.X+1; x++ {
								nc := Coord4D{X: x, Y: y, Z: z, W: w}
								if _, ok := nextCoords[nc.AsString()]; !ok {
									nextCoords[nc.AsString()] = nc
								}
							}
						}
					}
				}
			}
		}
		coords = nextCoords
		fmt.Println(activeCount)
		lastActiveCount = activeCount
	}
	return lastActiveCount, nil
}

func main() {
	fh, _ := os.Open("input.txt")
	data, _ := helper.ReadLines(fh, true)
	ans, err := partTwo(data)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part one: %d\n", ans)
}