package main

import (
	"aoc2020/helper"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func padZeros(v string) string {
	r := v
	for i := 0; i < 36-len(v); i++ {
		r = "0" + r
	}
	return r
}

func partOne(data []string) int64 {

	maskRE := regexp.MustCompile(`^mask = (.+)$`)
	memRE := regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)

	addresses := make(map[int]string)
	var currentMask string
	for _, line := range data {
		maskmatch := maskRE.FindStringSubmatch(line)
		if len(maskmatch) > 0 {
			currentMask = maskmatch[1]
		}

		memmatch := memRE.FindStringSubmatch(line)
		if len(memmatch) > 0 {
			address, _ := strconv.Atoi(memmatch[1])
			memval, _ := strconv.Atoi(memmatch[2])

			if _, ok := addresses[address]; !ok {
				addresses[address] = "000000000000000000000000000000000000"
			}

			memvalBinary := strconv.FormatInt(int64(memval), 2)
			paddedMemVal := padZeros(memvalBinary)
			addresses[address] = applyMask(currentMask, paddedMemVal, false)
		}
	}

	var sum int64
	for _, v := range addresses {
		tval, _ := strconv.ParseInt(v, 2, 64)
		sum += tval
	}
	return sum
}

func partTwo(data []string) int64 {

	maskRE := regexp.MustCompile(`^mask = (.+)$`)
	memRE := regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)

	addresses := make(map[int]string)
	var currentMask string
	for _, line := range data {
		maskmatch := maskRE.FindStringSubmatch(line)
		if len(maskmatch) > 0 {
			currentMask = maskmatch[1]
		}

		memmatch := memRE.FindStringSubmatch(line)
		if len(memmatch) > 0 {
			address, _ := strconv.Atoi(memmatch[1])
			memval, _ := strconv.Atoi(memmatch[2])

			addressBinary := strconv.FormatInt(int64(address), 2)
			paddedAddress := padZeros(addressBinary)
			maskedValue := applyMask(currentMask, paddedAddress, true)
			candidates := getCandidates(maskedValue)
			for _, candidate := range candidates {
				tval, _ := strconv.ParseInt(candidate, 2, 64)
				if _, ok := addresses[int(tval)]; !ok {
					addresses[int(tval)] = "000000000000000000000000000000000000"
				}
				memvalBinary := strconv.FormatInt(int64(memval), 2)
				paddedMemval := padZeros(memvalBinary)
				addresses[int(tval)] = paddedMemval
			}
		}
	}

	var sum int64
	for _, v := range addresses {
		tval, _ := strconv.ParseInt(v, 2, 64)
		sum += tval
	}
	return sum
}

func getCandidates(mask string) []string {
	if len(mask) == 1 {
		if mask[0] == 'X' {
			return []string{"1","0"}
		} else {
			return []string{string(mask[0])}
		}
	}
	candidates := getCandidates(mask[1:])
	ret := make([]string, 0)
	for _, c := range candidates {
		if mask[0] == 'X' {
			ret = append(ret, "1" + c)
			ret = append(ret, "0" + c)
		} else {
			ret = append(ret, string(mask[0]) + c)
		}
	}
	return ret
}

func applyMask(mask string, value string, floating bool) string {
	var retstr string
	for i := len(mask)-1; i >= 0; i-- {
		if i >= len(value) {
			retstr = string(mask[i]) + retstr
		} else {
			switch mask[i] {
			case 'X':
				if floating {
					retstr = string(mask[i]) + retstr
				} else {
					retstr = string(value[i]) + retstr
				}
			case '1':
				retstr = string(mask[i]) + retstr
			case '0':
				if floating {
					retstr = string(value[i]) + retstr
				} else {
					retstr = string(mask[i]) + retstr
				}
			}
		}
	}
	return retstr
}

func main() {
	fh, _ := os.Open("input.txt")
	data, _ := helper.ReadLines(fh, true)
	ans := partOne(data)
	fmt.Printf("Part one: %v\n", ans)
	ans = partTwo(data)
	fmt.Printf("Part two: %v\n", ans)

}
