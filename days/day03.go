package days

import (
	"strconv"
	"unicode"

	"github.com/tuxxi/aoc2023/util"
)

func init() {
	util.RegisterDay(3, day3)
}

func day3(input []string) (any, any) {
	var part1, part2 int

	for lineNum, line := range input {
		for i := 0; i < len(line); i++ {
			c := line[i]
			if isPart(c) {
				adjacenctNums := make([]int, 0)

				// check in a box around the part
				for y := lineNum - 1; y <= lineNum+1; y++ {
					if y < 0 || y >= len(input) {
						continue
					}
					x := i - 1
					for ; x <= i+1; x++ {
						if x < 0 || x >= len(line) {
							continue
						}
						char := rune(input[y][x])
						if unicode.IsDigit(char) {
							// consume digits to find the range [start, end] on line y
							var start, end int = x, x
							for ; start >= 0 && unicode.IsDigit(rune(input[y][start])); start-- {
							}
							for ; end < len(line) && unicode.IsDigit(rune(input[y][end])); end++ {
							}
							partNum, _ := strconv.Atoi(input[y][start+1 : end])
							adjacenctNums = append(adjacenctNums, partNum)
							// skip forward to end of the range to avoid double-counting a part
							x = end
						}
					}
				}
				for _, num := range adjacenctNums {
					part1 += num
				}
				if isGear(c) && len(adjacenctNums) == 2 {
					part2 += (adjacenctNums[0] * adjacenctNums[1])
				}
			}
		}
	}
	return part1, part2
}

func isPart(b byte) bool {
	return !(b == '.' || unicode.IsDigit(rune(b)))
}

func isGear(b byte) bool {
	return b == '*'
}
