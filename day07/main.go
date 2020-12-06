package main

import (
	"aoc2020/helper"
	"fmt"
	"os"
)

func main() {
	fh, _ := os.Open("test.txt")
	data, _ := helper.ReadLines(fh, true)
	fmt.Println(data)
}
