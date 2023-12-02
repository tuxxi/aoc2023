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
	util.AssertEq(t, p2, 467835)
}

func TestD3_GearRatio(t *testing.T) {
	// adjacent gear is not double counted
	input := []string{
		"....*.....",
		"..35.633..",
		"..........",
		"..114.....",
		"...*......",
	}
	_, p2 := day3(input)
	util.AssertEq(t, p2, 22155)
}
