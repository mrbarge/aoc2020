package helper

import "strconv"

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

func ContainsInt(i int, a []int) bool {
	for _, v := range a {
		if v == i {
			return true
		}
	}
	return false
}