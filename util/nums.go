package util

import "strconv"

func NumsToInts(nums []string) []int {
	var ret []int
	for _, num := range nums {
		i, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		ret = append(ret, i)
	}
	return ret
}
