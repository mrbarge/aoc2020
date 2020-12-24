package main

import (
	"aoc2020/helper"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type HexDir int
const (
	E = iota
	NE
	NW
	W
	SE
	SW
)

func parseLine(l string) (dir []HexDir) {

	dir = make([]HexDir, 0)

	i := 0
	for i < len(l) {
		switch l[i] {
		case 'e':
			dir = append(dir, E)
		case 'w':
			dir = append(dir, W)
		case 'n':
			i++
			if i == len(l) {
				// this seems bad but let's handle it anyway
				break
			}
			switch l[i] {
			case 'e':
				dir = append(dir, NE)
			case 'w':
				dir = append(dir, NW)
			}
		case 's':
			i++
			if i == len(l) {
				// this seems bad but let's handle it anyway
				break
			}
			switch l[i] {
			case 'e':
				dir = append(dir, SE)
			case 'w':
				dir = append(dir, SW)
			}
		}
		i++
	}

	return dir
}

func getCoord(dir []HexDir) helper.Coord {

	c := helper.Coord{
		X: 0, Y: 0,
	}

	for _, d := range dir {
		switch d {
		case E:
			c.X += 2
		case W:
			c.X -= 2
		case NE:
			c.X += 1
			c.Y += 1
		case NW:
			c.X -= 1
			c.Y += 1
		case SE:
			c.X += 1
			c.Y -= 1
		case SW:
			c.X -= 1
			c.Y -= 1
		}
	}
	return c
}

func coordToKey(c helper.Coord) string {
	s := fmt.Sprintf("%v:%v", c.X, c.Y)
	return s
}

func problem(data []string) (p1 int, p2 int) {

	coordState := make(map[string]bool)
	for _, line := range data {
		dirs := parseLine(line)
		coord := getCoord(dirs)
		coordKey := coordToKey(coord)

		if _, ok := coordState[coordKey]; !ok {
			coordState[coordKey] = true
			continue
		}
		coordState[coordKey] = !coordState[coordKey]
	}

	for _, v := range coordState {
		if v {
			p1++
		}
	}

	for i:= 0; i < 100; i++ {
		coordState = enrich(coordState)
		coordState = swap(coordState)
	}

	for _, v := range coordState {
		if v {
			p2++
		}
	}

	return p1, p2
}

func keyToCoord(s string) helper.Coord {
	fieldRE := regexp.MustCompile(`^(.+):(.+)$`)
	fieldMatch := fieldRE.FindStringSubmatch(s)
	if len(fieldMatch) > 0 {
		x, _ := strconv.Atoi(fieldMatch[1])
		y, _ := strconv.Atoi(fieldMatch[2])
		c := helper.Coord{
			X: x,
			Y: y,
		}
		return c
	}
	return helper.Coord{0,0}
}

func enrich(coords map[string]bool) map[string]bool {
	retMap := make(map[string]bool)
	for k, v := range coords {
		retMap[k] = v
		c := keyToCoord(k)
		neighbours := make([]helper.Coord, 0)
		neighbours = append(neighbours, helper.Coord{c.X + 2, c.Y})
		neighbours = append(neighbours, helper.Coord{c.X - 2, c.Y})
		neighbours = append(neighbours, helper.Coord{c.X + 1, c.Y + 1})
		neighbours = append(neighbours, helper.Coord{c.X + 1, c.Y - 1})
		neighbours = append(neighbours, helper.Coord{c.X - 1, c.Y + 1})
		neighbours = append(neighbours, helper.Coord{c.X - 1, c.Y - 1})

		for _, n := range neighbours {
			key := coordToKey(n)
			if _, ok := coords[key]; !ok {
				retMap[key] = false
			}
		}
	}
	return retMap
}

func swap(coords map[string]bool) map[string]bool {

	retMap := make(map[string]bool)
	for k, v := range coords {
		c := keyToCoord(k)
		neighbours := make([]helper.Coord, 0)
		neighbours = append(neighbours, helper.Coord{c.X+2, c.Y})
		neighbours = append(neighbours, helper.Coord{c.X-2, c.Y})
		neighbours = append(neighbours, helper.Coord{c.X+1, c.Y+1})
		neighbours = append(neighbours, helper.Coord{c.X+1, c.Y-1})
		neighbours = append(neighbours, helper.Coord{c.X-1, c.Y+1})
		neighbours = append(neighbours, helper.Coord{c.X-1, c.Y-1})

		blackTiles := 0
		for _, n := range neighbours {
			key := coordToKey(n)
			if _, ok := coords[key]; ok {
				if coords[key] {
					blackTiles++
				}
			}
		}

		if v {
			// black
			if blackTiles == 0 || blackTiles > 2 {
				retMap[k] = false
			} else {
				retMap[k] = true
			}
		} else {
			// white
			if blackTiles == 2{
				retMap[k] = true
			} else {
				retMap[k] = false
			}
		}
	}

	return retMap
}

func main() {
	fh, _ := os.Open("input.txt")
	data, _ := helper.ReadLines(fh, true)
	p1, p2 := problem(data)
	fmt.Printf("Part one: %d\n", p1)
	fmt.Printf("Part two: %d\n", p2)

}
