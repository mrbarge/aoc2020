package main

import "fmt"

func partTwo(data []int) int {

	spokenOrder := make(map[int][]int)
	limit := 30000000

	// seed starting numbers
	for turn := 1; turn <= len(data); turn++ {
		spokenOrder[data[turn-1]] = append(spokenOrder[data[turn-1]], turn)
	}

	lastSpokenNumber := data[len(data)-1]
	for turn := len(data)+1; turn <= limit; turn++ {
		if _, ok := spokenOrder[lastSpokenNumber]; !ok {
			spokenOrder[lastSpokenNumber] = []int{turn-1}
		} else {
			spokenOrder[lastSpokenNumber] = append(spokenOrder[lastSpokenNumber], turn-1)
		}

		if len(spokenOrder[lastSpokenNumber]) > 1 {
			// already been spoken, current player announces turn difference
			order := spokenOrder[lastSpokenNumber]
			lastSpokenNumber = order[len(order)-1] - order[len(order)-2]
		} else {
			lastSpokenNumber = 0
		}
	}
	return lastSpokenNumber
}

func main() {
	//data := []int{0,3,6}
	data := []int{8, 13, 1, 0, 18, 9}
	ans := partTwo(data)
	fmt.Printf("Part one: %v\n", ans)

}
