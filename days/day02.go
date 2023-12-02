package days

import (
	"fmt"
	"strings"

	"github.com/tuxxi/aoc2023/util"
)

type day2 struct{}

func init() {
	util.RegisterDay[day2](2)
}

func (_ day2) Part1(input []string) any {
	part1, _ := impl(input)
	return part1
}

func (_ day2) Part2(input []string) any {
	_, part2 := impl(input)
	return part2
}

func impl(input []string) (part1, part2 int) {
	maxs := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	for _, line := range input {
		mins := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		var (
			gameNum int
			valid   bool = true
		)
		_, err := fmt.Sscanf(line, "Game %d:", &gameNum)
		if err != nil {
			panic(fmt.Sprintf("could not scan line: %v", err))
		}
		turns := strings.TrimSpace(strings.Split(line, ":")[1])
		for _, turn := range strings.Split(turns, ";") {
			for _, cube := range strings.Split(turn, ",") {
				var (
					num   int
					color string
				)
				_, err := fmt.Sscanf(cube, "%d %s", &num, &color)
				if err != nil {
					panic(fmt.Sprintf("could not scan cube: %v", err))
				}
				if mins[color] < num {
					mins[color] = num
				}

				if max_, ok := maxs[color]; ok {
					if num > max_ {
						valid = false
					}
				}
			}
		}
		part2 += (mins["red"] * mins["blue"] * mins["green"])
		if valid {
			part1 += gameNum
		}
	}
	return part1, part2
}
