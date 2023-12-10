package days

import (
	"testing"

	"github.com/tuxxi/aoc2023/util"
)

func TestD9(t *testing.T) {
	input := []string{
		"0 3 6 9 12 15",
		"1 3 6 10 15 21",
		"10 13 16 21 30 45",
	}
	p1, p2 := day9(input)
	util.AssertEq(t, p1, 114)
	util.AssertEq(t, p2, 2)
}

func TestD9Custom(t *testing.T) {
	p1, _ := day9([]string{"9 12 17 23 28 40 101 333 1036 2906 7508 18265 42466 95247 207297 439378 908879 1836876 3628951 7009831 13241340"})
	util.AssertEq(t, p1, 24465918)

	p1_2, _ := day9([]string{"-3 -2 11 59 179 432 931 1903 3821 7673 15477 31204 62335 122353 234557 437682 793917 1400032 2402455 4017281 6556347"})
	util.AssertEq(t, p1_2, 10460670)

	p1_3, _ := day9([]string{"-3 -5 1 23 66 145 306 652 1373 2778 5315 9530 15842 23878 30896 28497 -3643 -108939 -372747 -952815 -2126694"})
	util.AssertEq(t, p1_3, -4363273)
}
