package main

import (
	"aoc2020/helper"
	"fmt"
	"math"
	"os"
)

func canCalc(data []int, start int, end int, target int) bool {
	for i := start; i <= end && i < len(data); i++ {
		for j := start; j <= end && j < len(data); j++ {
			if i == j {
				continue
			}
			sum := data[i] + data[j]
			if sum == target {
				return true
			}
		}
	}
	return false
}

func partOne(data []int, span int) int {

	startPos := 0
	endPos := span
	for i := span+1; i < len(data); i++ {
		if !canCalc(data, startPos, endPos, data[i]) {
			return data[i]
		}
		startPos++
		endPos++
	}
	return -1
}

func partTwo(data []int, target int) int {

	runningSum := 0
	startRange := 0
	for i := 0; i < len(data)-1; i++ {
		runningSum = data[i]
		startRange = i
		for j := i+1; j < len(data); j++ {
			runningSum += data[j]
			if runningSum == target {
				smallest := math.MaxInt64
				largest := math.MinInt64
				for x := startRange; x <= j; x++ {
					if data[x] < smallest {
						smallest = data[x]
					}
					if data[x] > largest {
						largest = data[x]
					}
				}
				return smallest + largest
			}

			if runningSum > target {
				runningSum = 0
				break
			}
		}
	}
	return -1
}

func main() {
	fh, _ := os.Open("input.txt")
	data, _ := helper.ReadLinesAsInt(fh)
	ans := partOne(data, 25)
	fmt.Printf("Part one: %v\n", ans)
	ans2 := partTwo(data, ans)
	fmt.Printf("Part two: %v\n", ans2)
}
