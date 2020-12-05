package main

import (
	"aoc2020/helper"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type passport struct {
	birthYear      string
	eyeColour      string
	issueYear      string
	expirationYear string
	height         string
	hairColour     string
	id             string
	countryId      string
}

func (p *passport) read(s string) error {
	components := strings.Split(s, ":")
	if len(components) != 2 {
		return fmt.Errorf("improperly formatted element: %s", s)
	}
	switch components[0] {
	case "byr":
		p.birthYear = components[1]
	case "iyr":
		p.issueYear = components[1]
	case "eyr":
		p.expirationYear = components[1]
	case "hgt":
		p.height = components[1]
	case "hcl":
		p.hairColour = components[1]
	case "ecl":
		p.eyeColour = components[1]
	case "pid":
		p.id = components[1]
	case "cid":
		p.countryId = components[1]
	default:
		return fmt.Errorf("invalid data: %s", s)
	}
	return nil
}

func (p *passport) valid() bool {
	if len(p.eyeColour) == 0 ||
		len(p.hairColour) == 0 ||
		len(p.height) == 0 ||
		len(p.birthYear) == 0 ||
		len(p.issueYear) == 0 ||
		len(p.expirationYear) == 0 ||
		len(p.id) == 0 {
		return false
	}
	return true
}

func (p *passport) strictValid() (bool, error) {
	by, err := strconv.Atoi(p.birthYear)
	if err != nil {
		return false, fmt.Errorf("invalid birthyear: %v", p.birthYear)
	}
	if by < 1920 || by > 2002 {
		return false, fmt.Errorf("invalid birthyear: %v", p.birthYear)
	}

	iy ,err := strconv.Atoi(p.issueYear)
	if err != nil {
		return false, fmt.Errorf("invalid issueyear: %v", p.issueYear)
	}
	if iy < 2010 || iy > 2020 {
		return false, fmt.Errorf("invalid issueyear: %v", p.issueYear)
	}

	ey, err := strconv.Atoi(p.expirationYear)
	if err != nil {
		return false, fmt.Errorf("invalid expirationyear: %v", p.expirationYear)
	}
	if ey < 2020 || ey > 2030 {
		return false, fmt.Errorf("invalid expirationyear: %v", p.expirationYear)
	}

	hre := regexp.MustCompile(`(\d+)(cm|in)`)
	hmatch := hre.FindStringSubmatch(p.height)
	if len(hmatch) == 0 {
		return false, fmt.Errorf("invalid height: %v", p.height)
	} else {
		hv, err := strconv.Atoi(hmatch[1])
		if err != nil {
			return false, fmt.Errorf("bad height")
		}
		if hmatch[2] == "in" && (hv < 59 || hv > 76) {
			return false, fmt.Errorf("invalid height range")
		}
		if hmatch[2] == "cm" && (hv < 150 || hv > 193) {
			return false, fmt.Errorf("invalid height range")
		}
	}

	hairmatch, _ := regexp.MatchString(`#[0-9a-f]{6}`, p.hairColour)
	if !hairmatch {
		return false, fmt.Errorf("invalid hair: %v", p.hairColour)
	}

	eyematch, _ := regexp.MatchString(`(amb|blu|brn|gry|grn|hzl|oth)`, p.eyeColour)
	if !eyematch {
		return false, fmt.Errorf("invalid eyecolour: %v", p.eyeColour)
	}

	pidmatch, _ := regexp.MatchString(`^\d{9}$`, p.id)
	if !pidmatch {
		return false, fmt.Errorf("invalid passportid: %v", p.id)
	}

	return true, nil
}

func problem(data []string) (part1 int, part2 int) {

	passports := make([]passport, 0)

	var tmpPass passport
	for _, entry := range data {
		entry = strings.TrimSpace(entry)
		if len(entry) == 0 {
			passports = append(passports, tmpPass)
			tmpPass = passport{}
			continue
		}

		fields := strings.Split(entry, " ")
		for _, field := range fields {
			err := tmpPass.read(field)
			if err != nil {
				fmt.Printf("Error reading field: %v\n", err)
			}
		}
	}

	for _, p := range passports {
		if p.valid() {
			part1++
		}

		v, _ := p.strictValid()
		if v {
			part2++
		}
	}

	return part1, part2
}

func main() {
	fh, _ := os.Open("input.txt")
	data, _ := helper.ReadLines(fh, false)
	p1, p2 := problem(data)
	fmt.Printf("Part one: %v\n", p1)
	fmt.Printf("Part two: %v\n", p2)

}
