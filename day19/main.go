package main

import (
	"aoc2020/helper"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Rule struct {
	rulesets [][]int
	v uint8
	id int
}

func readRules(data []string) map[int]Rule {
	ruleRE := regexp.MustCompile(`^(\d+): "(.+)"$`)
	rulesetRE := regexp.MustCompile(`^(\d+): (.+)$`)

	ruleMap := make(map[int]Rule)
	for _, line := range data {
		ruleMatch := ruleRE.FindStringSubmatch(line)
		if len(ruleMatch) > 0 {
			ruleid, _ := strconv.Atoi(ruleMatch[1])
			ruleMap[ruleid] = Rule{
				rulesets: [][]int{},
				v:        ruleMatch[2][0],
				id:       ruleid,
			}
			continue
		}
		rulesetMatch := rulesetRE.FindStringSubmatch(line)
		if len(rulesetMatch) > 0 {
			ruleid, _ := strconv.Atoi(rulesetMatch[1])
			rd := strings.Split(rulesetMatch[2],"|")
			rndsets := make([][]int, 0)
			for _, ruleset := range rd {
				rnd, _ := helper.StrCsvToIntArray(strings.Trim(ruleset, " "), " ")
				rndsets = append(rndsets, rnd)
			}
			ruleMap[ruleid] = Rule{
				rulesets: rndsets,
				v:        0,
				id:       ruleid,
			}
		}
	}
	return ruleMap
}

func canParse(s string, rules map[int]Rule) bool {

	base := rules[0]

	combos := build(base, rules)
	for _, l := range combos {
		if l == s {
			return true
		}
	}
	return false
}

func canParseTwo(s string, rules map[int]Rule) bool {

	base := rules[0]

	combos := buildTwo(base, 0, len(s), rules)
	for _, l := range combos {
		if l == s {
			return true
		}
	}
	return false
}

// if the number of rulesets is 0
//    return a string array of the value

// if the number of rulesets is 1
//    for each rule in the ruleset
//        get combinations for that rule
//	  permutestrings on combinations
//    return all combinations

// if the number of rulesets is 2
//    for each ruleset in the rule
//       build(ruleset) to get all combinations
//    permutestrings on all combinations
//    return all combinations


// if the rule has no rulesets
//
func build(r Rule, rules map[int]Rule) []string {
	if len(r.rulesets) == 0 {
		return []string{string(r.v)}
	}

	rcombos := make([]string, 0)
	if len(r.rulesets[0]) == 1 {
		rid := r.rulesets[0][0]
		rrule := rules[rid]
		rcombos = append(rcombos, build(rrule, rules)...)
	} else {
		rcr1 := rules[r.rulesets[0][0]]
		rcr2 := rules[r.rulesets[0][1]]
		rc1 := build(rcr1, rules)
		rc2 := build(rcr2, rules)
		rc := helper.PermuteStrings(rc1, rc2)
		rcombos = append(rcombos, rc...)
	}

	if len(r.rulesets) > 1 {
		if len(r.rulesets[1]) == 1 {
			rid := r.rulesets[1][0]
			rrule := rules[rid]
			rcombos = append(rcombos, build(rrule, rules)...)
		} else {
			rcr1 := rules[r.rulesets[1][0]]
			rcr2 := rules[r.rulesets[1][1]]
			rc1 := build(rcr1, rules)
			rc2 := build(rcr2, rules)
			rc := helper.PermuteStrings(rc1, rc2)
			rcombos = append(rcombos, rc...)
		}
	}

	return rcombos
}

func buildTwo(r Rule, depth int, ln int, rules map[int]Rule) []string {
	if len(r.rulesets) == 0 {
		return []string{string(r.v)}
	}

	//fmt.Printf("Called with %v deth and %v len\n", depth, ln)
	//if depth > ln {
	//	return []string{}
	//}

	rcombos := make([]string, 0)
	if len(r.rulesets[0]) == 1 {
		rid := r.rulesets[0][0]
		rrule := rules[rid]
		rcombos = append(rcombos, buildTwo(rrule, depth+1, ln, rules)...)
	} else {
		rcr1 := rules[r.rulesets[0][0]]
		rcr2 := rules[r.rulesets[0][1]]
		rc1 := buildTwo(rcr1, depth+1, ln, rules)
		rc2 := buildTwo(rcr2, depth+2, ln, rules)
		if len(rc1) > 0 && len(rc2) > 0 && (len(rc1) + len(rc2)) < ln {
			rc := helper.PermuteStrings(rc1, rc2)
			rcombos = append(rcombos, rc...)
		}
	}

	if len(r.rulesets) > 1 {
		if len(r.rulesets[1]) == 1 {
			rid := r.rulesets[1][0]
			rrule := rules[rid]
			rcombos = append(rcombos, buildTwo(rrule, depth+1, ln, rules)...)
		} else {
			rcr1 := rules[r.rulesets[1][0]]
			rcr2 := rules[r.rulesets[1][1]]
			rc1 := buildTwo(rcr1, depth+1, ln, rules)
			rc2 := buildTwo(rcr2, depth+2, ln, rules)
			if len(rc1) > 0 && len(rc2) > 0  && (len(rc1) + len(rc2)) < ln {
				rc := helper.PermuteStrings(rc1, rc2)
				rcombos = append(rcombos, rc...)
			}
		}
	}

	return rcombos
}

func partOne(ruledata []string, data []string) int {
	rules := readRules(ruledata)

	count := 0
	for _, line := range data {
		if canParse(line, rules) {
			count++
		}
	}
	return count
}

func partTwo(ruledata []string, data []string) int {
	rules := readRules(ruledata)

	count := 0
	for _, line := range data {
		if canParseTwo(line, rules) {
			count++
		}
	}
	return count
}

func main() {
	rfh, _ := os.Open("rules2.txt")
	ruledata, _ := helper.ReadLines(rfh, true)
	fh, _ := os.Open("input.txt")
	data, _ := helper.ReadLines(fh, true)
	ans := partTwo(ruledata, data)
	fmt.Printf("Part two: %v\n", ans)
}
