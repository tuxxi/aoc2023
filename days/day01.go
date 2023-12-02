package days

import (
	"github.com/tuxxi/aoc2023/util"
)

type day1 struct{}

func init() {
	util.RegisterDay[day1](1)
}

func asciiByteToInt(b byte) int {
	return int(b - 0x30)
}

func asciiByteIsDigit(b byte) bool {
	return 0x30 <= b && b <= 0x39
}

func (_ day1) Part1(input []string) any {
	var total int

	for _, line := range input {
		var lineTotal int
		// 1st number
		for i := 0; i < len(line); i++ {
			if asciiByteIsDigit(line[i]) {
				lineTotal += 10 * asciiByteToInt(line[i])
				break
			}
		}

		// 2nd number, may be the same as the first
		for i := len(line) - 1; i >= 0; i-- {
			if asciiByteIsDigit(line[i]) {
				lineTotal += asciiByteToInt(line[i])
				break
			}
		}
		total += lineTotal
	}
	return total
}

func (_ day1) Part2(input []string) any {
	var total int
	nums := map[string]int{
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

	for _, line := range input {
		var lineTotal int
	outerfwd:
		for i := 0; i < len(line); i++ {
			// check if digit is numeric
			if asciiByteIsDigit(line[i]) {
				lineTotal += 10 * asciiByteToInt(line[i])
				break
			}
			// otherwise, it's spelled out, use sliding window
			for j := i; j < len(line); j++ {
				num, ok := nums[line[i:j]]
				if ok {
					lineTotal += 10 * num
					break outerfwd
				}
			}
		}

	outerback:
		for i := len(line) - 1; i >= 0; i-- {
			if asciiByteIsDigit(line[i]) {
				lineTotal += asciiByteToInt(line[i])
				break
			}
			// backwards sliding window
			for j := i; j <= len(line); j++ {
				num, ok := nums[line[i:j]]
				if ok {
					lineTotal += num
					break outerback
				}
			}
		}

		total += lineTotal
	}
	return total
}
