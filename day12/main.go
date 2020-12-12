package main

import (
	"aoc2020/helper"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Direction string
type Ship struct {
	c   helper.Coord
	dir Direction
}

func move(ship Ship, instruction string) (Ship, error) {
	dir := instruction[0]
	steps, err := strconv.Atoi(instruction[1:])
	if err != nil {
		return ship, fmt.Errorf("invalid steps: %v", instruction)
	}

	switch dir {
	case 'N':
		ship.c.Y += steps
	case 'S':
		ship.c.Y -= steps
	case 'E':
		ship.c.X += steps
	case 'W':
		ship.c.X -= steps
	case 'L':
		ship.dir = turn(ship.dir, steps/90, true)
	case 'R':
		ship.dir = turn(ship.dir, steps/90, false)
	case 'F':
		nextInst := fmt.Sprintf("%s%d", ship.dir, steps)
		ship, err = move(ship, nextInst)
	}
	return ship, err
}

func moveWithWaypoint(ship Ship, waypoint Ship, instruction string) (Ship, Ship, error) {
	dir := instruction[0]
	steps, err := strconv.Atoi(instruction[1:])
	if err != nil {
		return ship, waypoint, fmt.Errorf("invalid steps: %v", instruction)
	}

	switch dir {
	case 'N':
		waypoint.c.Y += steps
	case 'S':
		waypoint.c.Y -= steps
	case 'E':
		waypoint.c.X += steps
	case 'W':
		waypoint.c.X -= steps
	case 'L':
		waypoint = rotate(waypoint, steps/90, true)
	case 'R':
		waypoint = rotate(waypoint, steps/90, false)
	case 'F':
		ship.c.Y += waypoint.c.Y * steps
		ship.c.X += waypoint.c.X * steps
	}
	return ship, waypoint, err
}

func turn(facing Direction, times int, left bool) Direction {
	if times > 1 {
		facing = turn(facing, times-1, left)
	}
	if left {
		var left = map[Direction]Direction{"N": "W", "S": "E", "W": "S", "E": "N"}
		return left[facing]
	}
	var right = map[Direction]Direction{"N": "E", "S": "W", "W": "N", "E": "S"}
	return right[facing]
}

func rotate(waypoint Ship, times int, left bool) Ship {
	if times > 1 {
		waypoint = rotate(waypoint, times-1, left)
	}
	if left {
		t := waypoint.c.X
		waypoint.c.X = -1 * waypoint.c.Y
		waypoint.c.Y = t
	} else {
		t := waypoint.c.X
		waypoint.c.X = waypoint.c.Y
		waypoint.c.Y = -1 * t
	}
	return waypoint
}

func partOne(data []string) int {
	var err error
	s := Ship{
		c:   helper.Coord{0, 0},
		dir: "E",
	}
	for _, line := range data {
		s, err = move(s, line)
		if err != nil {
			fmt.Printf("Error moving: %v\n", err)
		}
	}

	return int(math.Abs(float64(s.c.X)) + math.Abs(float64(s.c.Y)))
}

func partTwo(data []string) int {
	var err error
	s := Ship{
		c:   helper.Coord{0, 0},
		dir: "E",
	}
	waypoint := Ship{
		c:   helper.Coord{10, 1},
		dir: "E",
	}
	for _, line := range data {
		s, waypoint, err = moveWithWaypoint(s, waypoint, line)
		if err != nil {
			fmt.Printf("Error moving: %v\n", err)
		}
	}

	return int(math.Abs(float64(s.c.X)) + math.Abs(float64(s.c.Y)))
}

func main() {
	fh, _ := os.Open("input.txt")
	data, _ := helper.ReadLines(fh, true)
	ans := partOne(data)
	fmt.Printf("Part one: %v\n", ans)
	ans2 := partTwo(data)
	fmt.Printf("Part two: %v\n", ans2)
}
