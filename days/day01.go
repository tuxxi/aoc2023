package days

import (
	"github.com/tuxxi/aoc2023/util"
)

func init() {
	util.RegisterDay(1, day1)
}

func asciiByteToInt(b byte) int {
	return int(b - 0x30)
}

func asciiByteIsDigit(b byte) bool {
	return 0x30 <= b && b <= 0x39
}

func day1(input []string) (any, any) {
	var (
		part1, part2 int
		nums         = map[string]int{
			"one":   1,
			"two":   2,
			"three": 3,
			"four":  4,
			"five":  5,
			"six":   6,
			"seven": 7,
			"eight": 8,
			"nine":  9,
		}
	)

	for _, line := range input {
		var line1, line2 []int
		for i := 0; i < len(line); i++ {
			// check if digit is numeric
			if asciiByteIsDigit(line[i]) {
				line1 = append(line1, asciiByteToInt(line[i]))
				line2 = append(line2, asciiByteToInt(line[i]))
			}
			// otherwise, it's spelled out, use sliding window to find any number word starting at `i`
			for j := i; j <= len(line); j++ {
				num, ok := nums[line[i:j]]
				if ok {
					line2 = append(line2, num)
					break
				}
			}
		}

		if len(line1) != 0 {
			part1 += line1[0]*10 + line1[len(line1)-1]
		}
		if len(line2) != 0 {
			part2 += line2[0]*10 + line2[len(line2)-1]
		}
	}
	return part1, part2
}
