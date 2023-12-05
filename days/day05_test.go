package days

import (
	"testing"

	"github.com/tuxxi/aoc2023/util"
)

func TestParseMap(t *testing.T) {
	mapping := []string{
		"50 98 2",
		"52 50 48",
	}
	m := parseMap(mapping)
	util.AssertEq(t, m.readMapVal(79), 81)
	util.AssertEq(t, m.readMapVal(14), 14)
	util.AssertEq(t, m.readMapVal(55), 57)
	util.AssertEq(t, m.readMapVal(13), 13)
}

func TestD5(t *testing.T) {
	input := []string{
		"seeds: 79 14 55 13",
		"",
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
		"",
		"soil-to-fertilizer map:",
		"0 15 37",
		"37 52 2",
		"39 0 15",
		"",
		"fertilizer-to-water map:",
		"49 53 8",
		"0 11 42",
		"42 0 7",
		"57 7 4",
		"",
		"water-to-light map:",
		"88 18 7",
		"18 25 70",
		"",
		"light-to-temperature map:",
		"45 77 23",
		"81 45 19",
		"68 64 13",
		"",
		"temperature-to-humidity map:",
		"0 69 1",
		"1 0 69",
		"",
		"humidity-to-location map:",
		"60 56 37",
		"56 93 4",
	}
	p1, p2 := day5(input)
	util.AssertEq(t, p1, 35)
	util.AssertEq(t, p2, 46)
}
