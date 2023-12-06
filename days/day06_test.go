package days

import (
	"testing"

	"github.com/tuxxi/aoc2023/util"
)

func TestD6(t *testing.T) {
	input := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}
	p1, p2 := day6(input)
	util.AssertEq(t, p1, 288)
	util.AssertEq(t, p2, 71503)
}
