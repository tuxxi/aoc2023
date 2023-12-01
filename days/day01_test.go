package days

import (
	"testing"

	"github.com/tuxxi/aoc2023/util"
)

func TestPart1(t *testing.T) {
	var input = []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}
	res := util.GetDay(1).Part1(input)
	if res != 142 {
		t.Fatal("expected 142, got", res)
	}
}

func TestPart2(t *testing.T) {
	var input = []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}
	res := util.GetDay(1).Part2(input)
	if res != 281 {
		t.Fatal("expected 281, got", res)
	}
}
