package main

import "fmt"

func findLoopCount(target int) (loopcount int) {
	subjectNo := 7
	key := 1
	for key != target {
		key = (key * subjectNo) % 20201227
		loopcount += 1
	}
	return loopcount
}

func problem(cardPubKey int, doorPubKey int) int {
	lc := findLoopCount(doorPubKey)
	doorKey := 1
	for i := 0; i < lc; i++ {
		doorKey = (doorKey * cardPubKey) % 20201227
	}
	return doorKey
}

func main() {
	ans := problem(6929599, 2448427)
	fmt.Printf("Part one: %v\n", ans)
}
