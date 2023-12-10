package days

import (
	"slices"
	"strings"

	"github.com/tuxxi/aoc2023/util"
)

func init() {
	util.RegisterDay(9, day9)
}

func day9(input []string) (any, any) {
	var part1, part2 int = 0, 0

	for _, line := range input {
		ints := util.NumsToInts(strings.Fields(line))

		// stack of the difference between each seq, recursively.
		var sequences [][]int
		sequences = append(sequences, ints)

		seq := ints
		for {
			var diffs []int
			for i := range seq[:len(seq)-1] {
				d := seq[i+1] - seq[i]
				diffs = append(diffs, d)
			}
			// all diffs were 0, bottom out
			if slices.Max(diffs) == 0 && slices.Min(diffs) == 0 {
				break
			}
			// recurse
			sequences = append(sequences, diffs)
			seq = diffs
		}

		slices.Reverse(sequences)
		var predicted1, predicted2 int
		// "bubble up" sequences starting from the last
		for idx := range sequences[:len(sequences)-1] {
			curr := sequences[idx]
			next := sequences[idx+1]

			predicted1 = next[len(next)-1] + curr[len(curr)-1]
			predicted2 = next[0] - curr[0]
			// fmt.Println(idx, curr)
			// fmt.Println(idx+1, next, predicted1)
			// fmt.Println(idx+1, next, predicted2)

			sequences[idx+1] = append(sequences[idx+1], predicted1)           // next += predicted
			sequences[idx+1] = append([]int{predicted2}, sequences[idx+1]...) // next = prepend(predicted2)
		}
		// terminal
		part1 += predicted1
		part2 += predicted2
	}
	return part1, part2
}
