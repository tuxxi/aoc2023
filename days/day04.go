package days

import (
	"math"
	"slices"
	"strings"

	"github.com/tuxxi/aoc2023/util"
)

func init() {
	util.RegisterDay(4, day4)
}

func day4(input []string) (any, any) {
	var part1, part2 int

	cardCounts := map[int]int{}
	// each card has one original copy
	for i := range input {
		cardCounts[i+1] = 1
	}

	for i, line := range input {
		l := strings.Split(line, ":")[1]
		l2 := strings.Split(l, "|")
		winning := util.NumsToInts(strings.Fields(l2[0]))
		have := util.NumsToInts(strings.Fields(l2[1]))

		matching := 0
		for _, h := range have {
			if slices.Contains(winning, h) {
				matching++
			}
		}
		if matching != 0 {
			part1 += int(math.Exp2(float64(matching - 1)))
		}

		cardNum := i + 1
		// next N cards are also winning
		for n := cardNum + 1; n <= cardNum+matching; n++ {
			copies := cardCounts[cardNum]
			cardCounts[n] += copies
		}
	}
	for _, count := range cardCounts {
		part2 += count
	}
	return part1, part2
}
