package util

import (
	"strconv"
	"strings"
)

// Convert a list of stringified numbers into a list of integers
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

// Get an "iterator" of all the indexes of character 'r' in string 's'
func StrIndexes(s string, r rune) <-chan int {
	c := make(chan int)
	go func() {
		var itr int = -1
		for {
			i := strings.IndexRune(s[itr+1:], r)
			if i == -1 {
				close(c)
				return
			}
			itr += i + 1
			c <- itr
		}
	}()
	return c
}
