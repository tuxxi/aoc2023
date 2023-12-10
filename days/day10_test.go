package days

import (
	"testing"

	"github.com/tuxxi/aoc2023/util"
)

func TestD10(t *testing.T) {
	input1 := []string{
		".....",
		".S-7.",
		".|.|.",
		".L-J.",
		".....",
	}
	p1, _ := day10(input1)
	util.AssertEq(t, p1, 4)

	input2 := []string{
		"..F7.",
		".FJ|.",
		"SJ.L7",
		"|F--J",
		"LJ...",
	}
	p1_2, _ := day10(input2)
	util.AssertEq(t, p1_2, 8)
}
