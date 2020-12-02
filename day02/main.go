package main

import (
	"aoc2020/helper"
	"fmt"
	"os"
)

func partOne(arr []string) (int, error) {
	var (
		valid  int
		min    int
		max    int
		letter string
		passwd string
	)

	for _, rule := range arr {
		_, err := fmt.Sscanf(rule, "%d-%d %s %s", &min, &max, &letter, &passwd)
		if err != nil {
			return -1, err
		}
		// get rid of the : from the matched letter
		letter = string(letter[0])
		freq := countFrequency(passwd)
		if val, ok := freq[letter]; ok && val >= min && val <= max {
			valid++
		}
	}
	return valid, nil
}

func partTwo(arr []string) (int, error) {
	var (
		valid  int
		pos1   int
		pos2   int
		letter string
		passwd string
	)

	for _, rule := range arr {
		_, err := fmt.Sscanf(rule, "%d-%d %s %s", &pos1, &pos2, &letter, &passwd)
		if err != nil {
			return -1, err
		}
		// get rid of the : from the matched letter
		r := letter[0]
		m1 := (passwd[pos1-1] == r)
		m2 := (passwd[pos2-1] == r)
		if m1 != m2 {
			valid++
		}
	}
	return valid, nil
}

func countFrequency(s string) map[string]int {
	ret := make(map[string]int, 0)
	for _, c := range s {
		ret[string(c)] += 1
	}
	return ret
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh)
	if err != nil {
		fmt.Printf("unable to read file %v", err)
		os.Exit(1)
	}
	ans, err := partOne(lines)
	if err != nil {
		fmt.Printf("Error in part 1: %v", err)
		os.Exit(1)
	}
	fmt.Println(ans)

	ans, err = partTwo(lines)
	if err != nil {
		fmt.Printf("Error in part 2: %v", err)
		os.Exit(1)
	}
	fmt.Println(ans)

}
