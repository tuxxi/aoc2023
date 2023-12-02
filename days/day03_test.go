package days

import (
	"testing"

	"github.com/tuxxi/aoc2023/util"
)

func TestD3(t *testing.T) {
	input := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}
	p1, p2 := day3(input)
	util.AssertEq(t, p1, 4361)
	util.AssertEq(t, p2, 0)
}

func TestD3_Adjacent(t *testing.T) {
	input := []string{
		"467.*114..",
		"..........",
		"...*.1....",
	}
	p1, _ := day3(input)
	util.AssertEq(t, p1, 114)

}
