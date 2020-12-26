package main

import (
	"aoc2020/helper"
	"fmt"
)

type Cup struct {
	val int
	next *Cup
}

func round(current *Cup, cups map[int]*Cup, max int) (*Cup) {

	chosenCup := current.next
	chosenCupVals := make([]int, 0)
	chosenCupVals = append(chosenCupVals, chosenCup.val)
	chosenCupVals = append(chosenCupVals, chosenCup.next.val)
	chosenCupVals = append(chosenCupVals, chosenCup.next.next.val)
	current.next = chosenCup.next.next.next

	// find dest cup
	targetLabel := current.val - 1
	found := false
	for !found {
		if targetLabel < 1 {
			targetLabel = max
		}
		if !helper.ContainsInt(targetLabel, chosenCupVals) {
			found = true
			break
		}
		targetLabel--
	}

	destCup := cups[targetLabel]
	tmpPtr := destCup.next
	destCup.next = chosenCup
	chosenCup.next.next.next = tmpPtr

	return current.next
}

func partOne(data []int) (ans []int) {

	// make the ring'o'cups
	cupMap := make(map[int]*Cup)
	currentCup := Cup{
		val: data[0],
	}
	cupMap[currentCup.val] = &currentCup
	lastCup := &currentCup
	for i := 1; i < len(data); i++ {
		nextCup := Cup{
			val: data[i],
		}
		cupMap[data[i]] = &nextCup
		lastCup.next = &nextCup
		lastCup = &nextCup
	}
	lastCup.next = &currentCup

	loopCup := &currentCup
	for i := 0; i < 100; i++ {
		loopCup = round(loopCup, cupMap, 9)
	}

	ans = make([]int, 0)
	loopCup = cupMap[1].next
	for len(ans) != 8 {
		ans = append(ans, loopCup.val)
		loopCup = loopCup.next
	}
	return ans
}

func partTwo(data []int) (ans int) {

	for i := 10; i <= 1000000; i++ {
		data = append(data, i)
	}

	// make the ring'o'cups
	cupMap := make(map[int]*Cup)
	currentCup := Cup{
		val: data[0],
	}
	cupMap[currentCup.val] = &currentCup
	lastCup := &currentCup
	for i := 1; i < len(data); i++ {
		nextCup := Cup{
			val: data[i],
		}
		cupMap[data[i]] = &nextCup
		lastCup.next = &nextCup
		lastCup = &nextCup
	}
	lastCup.next = &currentCup

	loopCup := &currentCup
	for i := 0; i < 10000000; i++ {
		loopCup = round(loopCup, cupMap, 1000000)
	}

	ans = cupMap[1].next.val * cupMap[1].next.next.val
	return ans
}

func main() {
	input := []int {3,6,4,2,8,9,7,1,5}
	//input := []int {3,8,9,1,2,5,4,6,7}
	p1 := partOne(input)
	fmt.Printf("Part one: %v\n", p1)
	p2 := partTwo(input)
	fmt.Printf("Part two: %v\n", p2)

}
