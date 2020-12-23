package main

import (
	"aoc2020/helper"
	"fmt"
)

func partOne(data []int) (ans []int) {

	cups := make(map[int]int)
	for i, _ := range data {
		if i+1 < len(data) {
			cups[data[i]] = data[i+1]
		} else {
			cups[data[i]] = data[0]
		}
	}
	current := data[0]
	for i := 0; i < 100; i++ {
		cups, current = round(cups, current, 9)
	}

	v := cups[1]
	ans = make([]int, 0)
	ans = append(ans, v)
	for v != 1 {
		v = cups[v]
		ans = append(ans, v)
	}
	return ans
}

func partTwo(data []int) int {

	for i := 10; i <= 1000000; i++ {
		data = append(data, i)
	}
	cups := make(map[int]int)
	for i, _ := range data {
		if i+1 < len(data) {
			cups[data[i]] = data[i+1]
		} else {
			cups[data[i]] = data[0]
		}
	}
	current := data[0]
	for i := 0; i < 10000000; i++ {
		cups, current = round(cups, current, 1000000)
	}

	t1 := cups[1]
	t2 := cups[t1]
	ans := t1 * t2
	return ans
}

func round(cups map[int]int, current int, max int) (map[int]int, int) {

	chosenCups := make([]int, 0)
	cPos := current
	for i := 0; i < 3; i++ {
		cPos = cups[cPos]
		chosenCups = append(chosenCups, cPos)
	}
	cups[current] = cups[cPos]

	destCup := current-1
	found := false
	for !found {
		if destCup <= 0 {
			destCup = max
		}
		if !helper.ContainsInt(destCup, chosenCups) {
			// not in the seen pile, must be valid
			found = true
			break
		}
		destCup--
	}
	cups[destCup], cups[chosenCups[len(chosenCups)-1]] = chosenCups[0], cups[destCup]

	return cups, cups[current]
}

func main() {
	input := []int {3,6,4,2,8,9,7,1,5}
	//input := []int {3,8,9,1,2,5,4,6,7}
	p1 := partOne(input)
	fmt.Printf("Part one: %v\n", p1)
	p2 := partTwo(input)
	fmt.Printf("Part two: %v\n", p2)

}
