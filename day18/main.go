package main

import (
	"aoc2020/helper"
	"fmt"
	"go/ast"
	"go/parser"
	"os"
	"strconv"
	"strings"
)

func processPartOne(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	return parseAstPartOne(s)
}

func processPartTwo(s string) int {
	// perform the grand switcheroo of plus and multiplication
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "+", "T")
	s = strings.ReplaceAll(s, "*", "+")
	s = strings.ReplaceAll(s, "T", "*")

	return parseAstPartTwo(s)
}

func parseAstPartOne(s string) int {
	tr, _ := parser.ParseExpr(s)
	v := unrollPartOne(tr.(ast.Node))
	vn := calc(v)
	return vn
}

func parseAstPartTwo(s string) int {
	fmt.Printf("Processing line %v\n", s)
	tr, _ := parser.ParseExpr(s)
	v := unrollPartTwo(tr.(ast.Node))
	fmt.Println(v)
	return v
}

func unrollPartOne(n ast.Node) []string {
	ret := make([]string, 0)
	switch nt := n.(type) {
	case *ast.BinaryExpr:
		left := nt.X
		right := nt.Y
		op := nt.Op
		switch lt := left.(type) {
		case *ast.ParenExpr:
			v := unrollPartOne(left.(ast.Node))
			vn := calc(v)
			ret = append(ret, strconv.Itoa(vn))
		case *ast.BinaryExpr:
			v := unrollPartOne(left)
			ret = append(ret, v...)
		case *ast.BasicLit:
			ret = append(ret, lt.Value)
		}
		ret = append(ret, op.String())
		switch rt := right.(type) {
		case *ast.ParenExpr:
			v := unrollPartOne(right.(ast.Node))
			vn := calc(v)
			ret = append(ret, strconv.Itoa(vn))
		case *ast.BinaryExpr:
			v := unrollPartOne(right)
			ret = append(ret, v...)
		case *ast.BasicLit:
			ret = append(ret, rt.Value)
		}
	case *ast.ParenExpr:
		return unrollPartOne(nt.X)
	case *ast.BasicLit:
		return []string{nt.Value}
	}
	return ret
}

func unrollPartTwo(n ast.Node) int {
	ret := make([]string, 0)
	switch nt := n.(type) {
	case *ast.BinaryExpr:
		left := nt.X
		right := nt.Y
		leftValue := 0
		rightValue := 0
		op := nt.Op
		switch lt := left.(type) {
		case *ast.ParenExpr:
			leftValue = unrollPartTwo(left.(ast.Node))
		case *ast.BinaryExpr:
			leftValue = unrollPartTwo(left)
		case *ast.BasicLit:
			leftValue, _ = strconv.Atoi(lt.Value)
		}
		ret = append(ret, op.String())
		switch rt := right.(type) {
		case *ast.ParenExpr:
			rightValue = unrollPartTwo(right.(ast.Node))
		case *ast.BinaryExpr:
			rightValue = unrollPartTwo(right.(ast.Node))
		case *ast.BasicLit:
			rightValue, _ = strconv.Atoi(rt.Value)
		}

		if op.String() == "+" {
			// switcheroo..
			return leftValue * rightValue
		} else if op.String() == "*" {
			return leftValue + rightValue
		}
	case *ast.ParenExpr:
		return unrollPartTwo(nt.X)
	case *ast.BasicLit:
		vn, _ := strconv.Atoi(nt.Value)
		return vn
	}
	return 0
}

func calc(d []string) int {
	if len(d) == 0 {
		return 0
	}
	start, _ := strconv.Atoi(d[0])
	for i := 1; i < len(d); {
		next, _ := strconv.Atoi(d[i+1])
		switch d[i] {
		case "+":
			start += next
		case "*":
			start *= next
		}
		i += 2
	}
	return start
}


func partOne(data []string) int {
	sum := 0
	for _, line := range data {
		sum += processPartOne(line)
	}
	return sum
}

func partTwo(data []string) int {
	sum := 0
	for _, line := range data {
		sum += processPartTwo(line)
	}
	return sum
}

func main() {
	fh, _ := os.Open("input.txt")
	data, _ := helper.ReadLines(fh, true)
	ans := partOne(data)
	fmt.Printf("Part one: %d\n", ans)
	ans = partTwo(data)
	fmt.Printf("Part two: %d\n", ans)

}