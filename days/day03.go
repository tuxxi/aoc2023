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
			if unicode.IsDigit(rune(c)) {
				// consume the whole range of digits
				j := i
				for ; j < len(line) && unicode.IsDigit(rune(line[j])); j++ {
				}
				digits, _ := strconv.Atoi(line[i:j])
				j--

				validPart := false
				// 1 line above and below
				for y := lineNum - 1; y <= lineNum+1; y++ {
					if y < 0 || y >= len(input) {
						continue
					}
					// 1 char before and after
					for x := i - 1; x <= j+1; x++ {
						if x < 0 || x >= len(line) {
							continue
						}
						if isPart(input[y][x]) {
							validPart = true
							goto done
						}
						// check diagonals at i-1 and j+1
						if y != lineNum {
							if (i != 0 && isPart(input[y][i-1])) || (j != len(line)-1 && isPart(input[y][j+1])) {
								validPart = true
								goto done
							}
						}

					}
				}

			done:
				if validPart {
					part1 += digits
					// fmt.Printf("✔️ valid part number at (%d, %d:%d): %d\n", lineNum, i, j, digits)
				} else {
					// fmt.Printf("❌ invalid part number at (%d, %d:%d): %d\n", lineNum, i, j, digits)
				}

				// skip to next on this line
				i = j
			}
		}
	}
	return part1, part2
}

func isPart(b byte) bool {
	return !(b == '.' || unicode.IsDigit(rune(b)))
}
