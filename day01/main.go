package main

import (
	"aoc2020/helper"
	"fmt"
	"os"
	"sort"
)

func partOne(nums []int) int {
	sort.Ints(nums)

	// Work from the front
	for i, _ := range nums {
		// Compare from the back
		for j := len(nums)-1; j > i; j-- {
			sum := nums[i] + nums[j]
			if sum == 2020 {
				return nums[i] * nums[j]
			}
			if sum < 2020 {
				// we'll never find a match, so abort
				break
			}
		}
	}
	return -1
}

func partTwo(nums []int) int {
	sort.Ints(nums)

	// Work from the front
	for i, _ := range nums {
		// Work from the front again
		for j, _ := range nums {
			if i == j {
				// don't look at the same number twice
				continue
			}
			msum := nums[i] + nums[j]
			if msum >= 2020 {
				// We'll never find a match
				break
			}
			for k := len(nums) - 1; k >= 0; k-- {
				if k == i || k == j {
					// don't look at the same number twice
					continue
				}
				sum := msum + nums[k]
				if sum == 2020 {
					return nums[i] * nums[j] * nums[k]
				}
			}
		}
	}
	return -1
}

func main() {
	fh, _ := os.Open("input.txt")
	nums, err := helper.ReadLinesAsInt(fh)
	if err != nil {
		fmt.Println("Unable to read input: %v", err)
	}
	ans := partOne(nums)
	fmt.Printf("Part one: %v\n",ans)
	ans = partTwo(nums)
	fmt.Printf("Part two: %v\n",ans)

}
