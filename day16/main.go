package main

import (
	"aoc2020/helper"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Boundary struct {
	Lower int
	Upper int
}

type Field struct {
	Name string
	Boundaries []Boundary
}

func (f Field) IsValid(id int) bool {
	for _, boundary := range f.Boundaries {
		if id >= boundary.Lower && id <= boundary.Upper {
			return true
		}
	}
	return false
}

func readFields(data []string) []Field {
	fieldRE := regexp.MustCompile(`^(.+): (\d+)-(\d+) or (\d+)-(\d+)$`)

	fields := make([]Field, 0)
	for _, line := range data {
		fieldMatch := fieldRE.FindStringSubmatch(line)
		if len(fieldMatch) > 0 {
			l1, _ := strconv.Atoi(fieldMatch[2])
			u1, _ := strconv.Atoi(fieldMatch[3])
			l2, _ := strconv.Atoi(fieldMatch[4])
			u2, _ := strconv.Atoi(fieldMatch[5])
			f := Field{
				Name:       fieldMatch[1],
				Boundaries: []Boundary{
					{
						Lower: l1,
						Upper: u1,
					},
					{
						Lower: l2,
						Upper: u2,
					},
				},
			}
			fields = append(fields, f)
		}
	}
	return fields
}

func getMyTicket(data []string) ([]int, error) {
	myTicket := false
	for _, line := range data {
		if strings.HasPrefix(line, "your ticket") {
			myTicket = true
			continue
		}
		if myTicket {
			ret, err := helper.StrCsvToIntArray(line)
			if err != nil {
				return nil, err
			}
			return ret, nil
		}
	}
	return nil, fmt.Errorf("can't find my ticket")
}

func getNearbyTickets(data []string) ([][]int, error) {
	nearbyTicket := false
	tickets := make([][]int, 0)
	for _, line := range data {
		if strings.HasPrefix(line, "nearby ticket") {
			nearbyTicket = true
			continue
		}
		if nearbyTicket {
			ret, err := helper.StrCsvToIntArray(line)
			if err != nil {
				return nil, err
			}
			tickets = append(tickets, ret)
		}
	}
	return tickets, nil
}


func partOne(data []string) (int, error) {
	fields := readFields(data)

	tickets, err := getNearbyTickets(data)
	if err != nil {
		return 0, err
	}

	errorRate := 0
	for _, ticket := range tickets {
		for _, fieldVal := range ticket {
			isValid := false
			for i := 0; i < len(fields) && !isValid; i++ {
				field := fields[i]
				if field.IsValid(fieldVal) {
					// once this value is valid for a field, we can quit
					isValid = true
				}
			}
			if !isValid {
				errorRate += fieldVal
			}
		}
	}
	return errorRate, nil
}

func partTwo(data []string) (int, error) {
	fields := readFields(data)

	myticket, _ := getMyTicket(data)
	tickets, err := getNearbyTickets(data)
	if err != nil {
		return 0, err
	}

	categoryPotentials := make([][]string, len(fields))
	for i, _ := range fields {
		categoryPotentials[i] = make([]string, 0)
	}

	validTickets := make([][]int, 0)
	for _, ticket := range tickets {
		isTicketValid := true
		for _, fieldVal := range ticket {
			isValid := false
			for i := 0; i < len(fields); i++ {
				field := fields[i]
				isValid = isValid || field.IsValid(fieldVal)
			}
			if !isValid {
				isTicketValid = false
				break
			}
		}
		if isTicketValid {
			validTickets = append(validTickets, ticket)
		}
	}

	// for each valid ticket
	validTickets = append(validTickets, myticket)
	for _, field := range fields {

		// if the field validates against every ticket number in this pos, it's a candidate
		for ticketIndex := 0; ticketIndex < len(validTickets[0]); ticketIndex++ {
			candidate := true
			for _, ticket := range validTickets {
				if !field.IsValid(ticket[ticketIndex]) {
					if field.Name == "arrival station" {
						fmt.Printf("Found breaking candidate: %v\n", ticket[ticketIndex])
					}

					candidate = false
					break
				}
			}
			if candidate {
				found := false
				for _, fn := range categoryPotentials[ticketIndex] {
					if fn == field.Name {
						found = true
					}
				}
				if !found {
					categoryPotentials[ticketIndex] = append(categoryPotentials[ticketIndex], field.Name)
				}
			}
		}
	}
	// by process of elimination, keep on ruling out duplicates until only 1 candidate per category
	fieldMap := make(map[string]Field)
	for _, field := range fields {
		fieldMap[field.Name] = field
	}

	done := false
	for !done {

		for i, v := range categoryPotentials {
			if len(v) == 1 {
				removeVal := v[0]
				// remove this vlaue from all other potential positions
				for j := 0; j < len(categoryPotentials); j++ {
					if i == j {
						continue
					}
					tmpSlice := make([]string, 0)
					for _, fv := range categoryPotentials[j] {
						if fv != removeVal {
							tmpSlice = append(tmpSlice, fv)
						}
					}
					categoryPotentials[j] = tmpSlice
				}
			}
		}
		done = isDone(categoryPotentials)
	}

	answer := 1
	for i, v := range categoryPotentials {
		if strings.HasPrefix(v[0], "departure") {
			answer *= myticket[i]
		}
	}
	return answer, nil
}

func isDone(l [][]string) bool {
	for _, v := range l {
		if len(v) > 1 {
			return false
		}
	}
	return true
}
func main() {
	fh, _ := os.Open("input.txt")
	data, _ := helper.ReadLines(fh, true)
	ans, err := partOne(data)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part one: %v\n", ans)
	ans, err = partTwo(data)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part two: %v\n", ans)

}
