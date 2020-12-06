package main

import (
	"aoc2020/helper"
	"fmt"
	"os"
)

func partOne(data []string) (part1 int) {

	groupQuestions := make(map[string]int, 0)
	for _, personAnswer := range data {
		if len(personAnswer) == 0 {
			// Reached new group, process the old one
			uniqueQuestions := helper.KeysStr(groupQuestions)
			part1 += len(uniqueQuestions)
			groupQuestions = make(map[string]int, 0)
			continue
		}
		for _, question := range personAnswer {
			groupQuestions[string(question)] += 1
		}
	}
	return part1
}

func partTwo(data []string) (part2 int) {

	groupQuestions := make(map[string]int, 0)
	numPeople := 0

	for _, personAnswer := range data {
		if len(personAnswer) == 0 {
			// Reached new group, process the old one
			sameAnswers := 0
			for _, numAnswers := range groupQuestions {
				if numAnswers == numPeople {
					sameAnswers++
				}
			}
			part2 += sameAnswers
			groupQuestions = make(map[string]int, 0)
			numPeople = 0
			continue
		}
		for _, question := range personAnswer {
			groupQuestions[string(question)] += 1
		}
		numPeople += 1
	}
	return part2
}

func main() {
	fh, _ := os.Open("input.txt")
	data, _ := helper.ReadLines(fh, false)
	ans := partOne(data)
	fmt.Printf("Part one: %v\n", ans)
	ans = partTwo(data)
	fmt.Printf("Part two: %v\n", ans)

}
