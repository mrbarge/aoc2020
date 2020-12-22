package main

import (
	"aoc2020/helper"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readData(data []string) (player1 []int, player2 []int, err error) {

	player1 = make([]int, 0)
	player2 = make([]int, 0)

	doPlayer1 := true
	for _, line := range data {
		if strings.HasPrefix(line, "Player 1") {
			continue
		} else if strings.HasPrefix(line, "Player 2") {
			doPlayer1 = false
			continue
		} else if line == "" {
			continue
		}

		n, err := strconv.Atoi(line)
		if err != nil {
			return player1, player2, err
		}

		if doPlayer1 {
			player1 = append(player1, n)
		} else {
			player2 = append(player2, n)
		}
	}

	return player1, player2, nil
}

func round(players *map[int][]int) {
	p1card := (*players)[1][0]
	p2card := (*players)[2][0]

	(*players)[1] = (*players)[1][1:]
	(*players)[2] = (*players)[2][1:]
	if p1card > p2card {
		(*players)[1] = append((*players)[1], p1card)
		(*players)[1] = append((*players)[1], p2card)
	} else if p2card > p1card {
		(*players)[2] = append((*players)[2], p2card)
		(*players)[2] = append((*players)[2], p1card)
	} else {
		// a draw? just put them both back in the deck I guess?
		(*players)[1] = append((*players)[1], p1card)
		(*players)[2] = append((*players)[2], p2card)
	}
}

func isWinner(p1 []int, p2 []int) (p1win bool, p2win bool) {
	if len(p1) == 0 {
		return false, true
	}
	if len(p2) == 0 {
		return true, false
	}
	return false, false
}

func partOne(data []string) (ans int, err error) {

	p1, p2, err := readData(data)
	if err != nil {
		return 0, err
	}

	players := make(map[int][]int)
	players[1] = p1
	players[2] = p2
	p1win, p2win := false, false
	for !(p1win || p2win) {
		round(&players)
		p1win, p2win = isWinner(players[1], players[2])
		if p1win {
			for i, n := range players[1] {
				ans += n * (len(players[1])-i)
			}
		}
		if p2win {
			for i, n := range players[2] {
				ans += n * (len(players[2])-i)
			}
		}
	}

	return ans, nil
}

func makeSeenKey(p1 []int, p2 []int) (key string) {
	key = "p1-"
	for _, n := range p1 {
		key += strconv.Itoa(n) + "-"
	}
	key += "p2-"
	for _, n := range p2 {
		key += strconv.Itoa(n) + "-"
	}
	return key
}

func roundTwo(players *map[int][]int) (p1win bool, p2win bool) {

	seen := make(map[string]bool)

	for !(p1win || p2win) {
		key := makeSeenKey((*players)[1], (*players)[2])
		if _, ok := seen[key]; ok {
			p1win = true
			p2win = false
			continue
		}
		seen[key] = true

		p1card := (*players)[1][0]
		p2card := (*players)[2][0]
		(*players)[1] = (*players)[1][1:]
		(*players)[2] = (*players)[2][1:]

		if len((*players)[1]) >= p1card && len((*players)[2]) >= p2card {
			// recursive combat time
			// make map copy
			nextRound := make(map[int][]int)
			nextRound[1] = make([]int, 0)
			nextRound[2] = make([]int, 0)
			for i := 0; i < p1card; i++ {
				nextRound[1] = append(nextRound[1], (*players)[1][i])
			}
			for i := 0; i < p2card; i++ {
				nextRound[2] = append(nextRound[2], (*players)[2][i])
			}
			p1win, p2win = roundTwo(&nextRound)
		} else {
			if p1card > p2card {
				p1win = true
				p2win = false
			} else {
				p2win = true
				p1win = false
			}
		}

		if p1win {
			(*players)[1] = append((*players)[1], p1card)
			(*players)[1] = append((*players)[1], p2card)
		} else {
			(*players)[2] = append((*players)[2], p2card)
			(*players)[2] = append((*players)[2], p1card)
		}

		p1win, p2win = isWinner((*players)[1], (*players)[2])
	}

	return p1win, p2win
}

func partTwo(data []string) (ans int, err error) {

	p1, p2, err := readData(data)
	if err != nil {
		return 0, err
	}

	players := make(map[int][]int)

	players[1] = p1
	players[2] = p2
	p1win, p2win := false, false
	for !(p1win || p2win) {
		roundTwo(&players)
		p1win, p2win = isWinner(players[1], players[2])
		if p1win {
			for i, n := range players[1] {
				ans += n * (len(players[1])-i)
			}
		}
		if p2win {
			for i, n := range players[2] {
				ans += n * (len(players[2])-i)
			}
		}
	}

	return ans, nil
}

func main() {
	fh, _ := os.Open("input.txt")
	data, _ := helper.ReadLines(fh, false)
	p1, err := partOne(data)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part one: %v\n", p1)
	p2, err := partTwo(data)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part two: %v\n", p2)
}
