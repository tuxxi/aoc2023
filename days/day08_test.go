package days

import (
	"testing"

	"github.com/tuxxi/aoc2023/util"
)

func TestDay8(t *testing.T) {
	in1 := []string{
		"RL",
		"",
		"AAA = (BBB, CCC)",
		"BBB = (DDD, EEE)",
		"CCC = (ZZZ, GGG)",
		"DDD = (DDD, DDD)",
		"EEE = (EEE, EEE)",
		"GGG = (GGG, GGG)",
		"ZZZ = (ZZZ, ZZZ)",
	}
	p1, _ := day8(in1)
	util.AssertEq(t, p1, 2)

	in2 := []string{
		"LLR",
		"",
		"AAA = (BBB, BBB)",
		"BBB = (AAA, ZZZ)",
		"ZZZ = (ZZZ, ZZZ)",
	}
	p1, _ = day8(in2)
	util.AssertEq(t, p1, 6)

}

func TestDay8Part2(t *testing.T) {
	in3 := []string{
		"LR",
		"",
		"11A = (11B, XXX)",
		"11B = (XXX, 11Z)",
		"11Z = (11B, XXX)",
		"22A = (22B, XXX)",
		"22B = (22C, 22C)",
		"22C = (22Z, 22Z)",
		"22Z = (22B, 22B)",
		"XXX = (XXX, XXX)",
	}
	instructions, m := parseInput(in3)
	// run p2 separately, p1 depends on starting on AAA which this input doesn't have
	util.AssertEq(t, d8part2(instructions, m), 6)
}
