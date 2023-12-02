package days

import (
	"testing"

	"github.com/tuxxi/aoc2023/util"
)

func TestD1_1(t *testing.T) {
	var input = []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}
	p1, _ := day1(input)
	util.AssertEq(t, p1, 142)
}

func TestD1_2(t *testing.T) {
	var input = []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}
	_, p2 := day1(input)
	util.AssertEq(t, p2, 281)
}
