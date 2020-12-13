package main

import (
	"aoc2020/helper"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func partOne(data []string) int {

	earliest, _ := strconv.Atoi(data[0])
	busStr := strings.Split(data[1], ",")

	busses := make([]int, 0)
	for _, bus := range busStr {
		if bus == "x" {
			continue
		}
		busId, _ := strconv.Atoi(bus)
		busses = append(busses,  busId)
	}
	sort.Ints(busses)

	times := make(map[int]int, 0)
	for _, busId := range busses {
		runs := int(math.Floor(float64(earliest) / float64(busId)))
		times[busId] = (runs+1) * busId
	}

	smallestTime := math.MaxInt64
	smallestBusId := 0
	for busId, time := range times {
		if time < smallestTime {
			smallestTime = time
			smallestBusId = busId
		}
	}

	return (smallestTime - earliest) * smallestBusId
}

func partTwo(data string) int {
	busStr := strings.Split(data, ",")

	busses := make([]int, 0)
	busIdToOrder := make(map[int]int, 0)
	orderToBusId := make(map[int]int, 0)
	for i, bus := range busStr {
		if bus != "x" {
			busId, _ := strconv.Atoi(bus)
			busIdToOrder[busId] = i
			orderToBusId[i] = busId
			busses = append(busses, busId)
		}
	}

	syncInterval := orderToBusId[0]
	timeTick := 0
	// in order of successive minutes of bus travel
	for busMinute := 1; busMinute < len(busses); {
		// what bus is travelling at this minute
		busId := busses[busMinute]
		busOrder := busIdToOrder[busId]
		// time when this bus should go
		targetTime := timeTick + busOrder
		// can it actually travel at this time?
		if targetTime % busId == 0 {
			syncInterval = helper.LCM(syncInterval,busId)
			busMinute++
			continue
		}
		timeTick += syncInterval
	}
	return timeTick
}


func main() {
	fh, _ := os.Open("test.txt")
	data, _ := helper.ReadLines(fh, true)
	ans := partOne(data)
	fmt.Printf("Part one: %v\n", ans)
	ans2 := partTwo(data[1])
	fmt.Printf("Part two: %v\n", ans2)
}
