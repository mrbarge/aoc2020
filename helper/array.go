package helper

import (
	"strconv"
	"strings"
)

func StrArrayToInt(arr []string) ([]int, error) {
	ret := make([]int, len(arr))
	for i, v := range arr {
		n, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		ret[i] = n
	}
	return ret, nil
}

func StrCsvToIntArray(s string) ([]int, error) {
	ret := make([]int, 0)
	nums := strings.Split(s, ",")
	for _, num := range nums {
		n, err := strconv.Atoi(num)
		if err != nil {
			return nil, err
		}
		ret = append(ret, n)
	}
	return ret, nil
}

func ContainsInt(i int, a []int) bool {
	for _, v := range a {
		if v == i {
			return true
		}
	}
	return false
}

func KeysStr(m map[string]int) (r []string) {
	r = make([]string, len(m))
	i := 0
	for k := range m {
		r[i] = k
	}
	return r
}
