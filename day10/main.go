package main

import (
	"aoc2020/helper"
	"fmt"
	"os"
	"sort"
)

func partOne(data []int) int {

	sort.Ints(data)

	oneJolts := 1
	threeJolts := 1
	for i := 0; i < len(data)-1; i++ {
		diff := data[i+1] - data[i]
		if diff == 1 {
			oneJolts++
		} else if diff == 3 {
			threeJolts++
		}
	}

	//fmt.Printf("%d vs %d\n", oneJolts, threeJolts)
	return oneJolts*threeJolts
}

func partTwo(data []int) int {

	sort.Ints(data)

	// add start and end adaptors
	procData := make([]int, len(data)+2)
	for i, r := range data {
		procData[i+1] = r
	}
	procData[len(data)+1] = data[len(data)-1]+3

	adaptors := make([]int, procData[len(procData)-1]+1)
	for _, r := range procData {
		adaptors[r] = 1
	}

	seen := make(map[int]int, 0)

	s := numValidRoutes(0, &seen, adaptors)
	return s
}

func numValidRoutes(j int, seen *map[int]int, adaptors []int) int {
	// has this adaptor been seen yet
	if numPaths, ok := (*seen)[j]; ok {
		// It has, so we don't need to waste time recomputing hop paths
		return numPaths
	}

	// check all potential hops
	foundHop := false
	allPaths := 0
	for i := 1; i < 4; i++ {
		// is this a valid hop
		hopTo := j+i
		if hopTo < len(adaptors) && adaptors[hopTo] == 1 {
			// sum valid paths stemming from taking this hop
			allPaths += numValidRoutes(hopTo, seen, adaptors)
			foundHop = true
		}
	}

	if !foundHop {
		// at the end, so pass that valid path back up the chain
		allPaths++
	}

	// store total number of possible paths from this point
	(*seen)[j] = allPaths

	return allPaths
}

func main() {
	fh, _ := os.Open("test.txt")
	data, _ := helper.ReadLinesAsInt(fh)
	ans := partOne(data)
	fmt.Printf("Part one: %v\n", ans)
	ans2 := partTwo(data)
	fmt.Printf("Part two: %v\n", ans2)
}

