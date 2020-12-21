package main

import (
	"aoc2020/helper"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func readData(data []string) (sharedIngredients map[string][]string, ingredientCount map[string]int, err error) {
	recipeRE := regexp.MustCompile(`^(.+) \(contains (.+)\)$`)

	sharedIngredients = make(map[string][]string)
	ingredientCount = make(map[string]int)
	for _, line := range data {
		recipeMatch := recipeRE.FindStringSubmatch(line)
		if len(recipeMatch) > 0 {
			ingredients := strings.Split(recipeMatch[1], " ")
			allergies := strings.Split(strings.ReplaceAll(recipeMatch[2], " ", ""), ",")
			for _, allergy := range allergies {
				if _, ok := sharedIngredients[allergy]; !ok {
					sharedIngredients[allergy] = ingredients
				}
				sharedIngredients[allergy] = helper.Intersection(sharedIngredients[allergy], ingredients)
			}
			for _, ingredient := range ingredients {
				ingredientCount[ingredient] += 1
			}
		} else {
			return nil, nil, fmt.Errorf("unable to parse line: %s", line)
		}
	}

	return sharedIngredients, ingredientCount, nil
}

func mapAllergyIngredient(sharedAllergyIngredients map[string][]string) map[string]string {

	notDone := true
	ret := make(map[string]string)
	for notDone {
		for allergy, ingredients := range sharedAllergyIngredients {
			// find a 1-to-1 mapping that we can eliminate from others
			if len(ingredients) == 1 {
				ret[allergy] = ingredients[0]
				targetIngredient := ingredients[0]
				for targetAllergy, targetIngredients := range sharedAllergyIngredients {
					if allergy == targetAllergy {
						continue
					}
					if helper.ContainsString(targetIngredient, targetIngredients) {
						deleteIdx := 0
						for i, v := range targetIngredients {
							if v == targetIngredient {
								deleteIdx = i
								break
							}
						}
						sharedAllergyIngredients[targetAllergy] = append(targetIngredients[:deleteIdx], targetIngredients[deleteIdx+1:]...)
					}
				}
			}
		}

		// only finished when each allergy can be mapped
		if len(ret) == len(sharedAllergyIngredients) {
			notDone = false
		}
	}
	return ret
}

func problem(data []string) (p1 int, p2 string, err error) {

	sharedAllergyIngredients, ingredientCount, err := readData(data)
	if err != nil {
		return 0, "", err
	}

	seen := make(map[string]bool)
	for _, sharedIngredients := range sharedAllergyIngredients {
		for _, ingredient := range sharedIngredients {
			seen[ingredient] = true
		}
	}

	for ingredient, count := range ingredientCount {
		if _, ok := seen[ingredient]; !ok {
			p1 += count
		}
	}

	allergyMap := mapAllergyIngredient(sharedAllergyIngredients)
	sortedKeys := make([]string, 0)
	for allergy, _ := range allergyMap {
		sortedKeys = append(sortedKeys, allergy)
	}
	sort.Strings(sortedKeys)

	for _, allergy := range sortedKeys {
		p2 += allergyMap[allergy] + ","
	}
	p2 = p2[0:len(p2)-2]

	return p1, p2, nil
}

func main() {
	fh, _ := os.Open("input.txt")
	data, _ := helper.ReadLines(fh, true)
	p1, p2, err := problem(data)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part one: %d\n", p1)
	fmt.Printf("Part two: %s\n", p2)
}
