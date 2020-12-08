package main

import (
	"aoc2020/helper"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type BagType string
type BagRequirement struct {
	bagType BagType
	quantity int
}

func canContain(bagType BagType, target BagType, allreqs map[BagType][]BagRequirement) bool {
	bagReqs := allreqs[bagType]
	// can it contain the bag type directly?
	for _, t := range bagReqs {
		if t.bagType == target {
			return true
		}
	}
	// can it contain the bag type indirectly?
	for _, t := range bagReqs {
		if canContain(t.bagType, target, allreqs) {
			return true
		}
	}
	return false
}

func countBags(bagType BagType, quantity int, allreqs map[BagType][]BagRequirement) (totalBags int) {
	bagReqs := allreqs[bagType]
	if len(bagReqs) == 0 {
		return quantity
	}
	for _, req := range bagReqs {
		totalBags += countBags(req.bagType, req.quantity, allreqs)
	}
	return quantity + (quantity*totalBags)
}

func problem(data []string, targetBag BagType) (part1 int, part2 int, err error) {

	// read and process input data
	bagReqs := make(map[BagType][]BagRequirement, 0)
	for _, d := range data {
		bagType, bagRequirements, err := process(d)
		if err != nil {
			return 0, 0, err
		}

		bagReqs[bagType] = bagRequirements
	}

	// part 1
	for t := range bagReqs {
		if canContain(t, targetBag, bagReqs) {
			part1++
		}
	}

	// part 2
	targetReqs := bagReqs[targetBag]
	for _, req := range targetReqs {
		part2 += countBags(req.bagType, req.quantity, bagReqs)
	}

	return part1, part2, nil
}

func process(s string) (BagType, []BagRequirement, error) {

	ruleRE := regexp.MustCompile(`^(.+?) bags contain (.+)$`)
	bagRE := regexp.MustCompile(`^(\d+) (.+?) bag.*$`)

	hmatch := ruleRE.FindStringSubmatch(s)
	if len(hmatch) == 0 {
		return "", nil, fmt.Errorf("invalid data format: %v", s)
	}

	bagType := hmatch[1]
	bagContents := hmatch[2]

	bagRequirements := make([]BagRequirement, 0)
	if bagContents == "no other bags." {
		return BagType(bagType), bagRequirements, nil
	}

	splitReqs := strings.Split(bagContents, ", ")
	for _, req := range splitReqs {
		bmatch := bagRE.FindStringSubmatch(req)
		if len(bmatch) == 0 {
			return "", nil, fmt.Errorf("invalid bag requirement: %v (%s)", req, s)
		}
		quantity, err := strconv.Atoi(bmatch[1])
		if err != nil {
			return "", nil, fmt.Errorf("error parsing requirement quantity: %v (%s)", bmatch[1], s)
		}

		br := BagRequirement{
			quantity: quantity,
			bagType: BagType(bmatch[2]),
		}

		bagRequirements = append(bagRequirements, br)
	}

	return BagType(bagType), bagRequirements, nil
}

func main() {
	fh, _ := os.Open("input.txt")
	data, _ := helper.ReadLines(fh, true)
	p1, p2, err := problem(data, BagType("shiny gold"))
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("Part one: %d\n", p1)
	fmt.Printf("Part two: %d\n", p2)
}
