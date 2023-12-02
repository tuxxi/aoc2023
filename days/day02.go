package days

import (
	"fmt"
	"strings"

	"github.com/tuxxi/aoc2023/util"
)

type day2 struct{}

var _ util.Day = (*day1)(nil)

func init() {
	util.RegisterDay[day2](2)
}

func (_ day2) Part1(input []string) any {
	maxs := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	total := 0
	for _, line := range input {
		var (
			gameNum int
		)
		_, err := fmt.Sscanf(line, "Game %d:", &gameNum)
		if err != nil {
			panic(fmt.Sprintf("could not scan line: %v", err))
		}
		valid := true
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

				if max_, ok := maxs[color]; ok {
					if num > max_ {
						valid = false
						goto done // don't need to process more, turn is invalid
					}
				}
			}
		}
	done:
		if valid {
			total += gameNum
		}
	}
	return total
}

func (_ day2) Part2(input []string) any {
	power := 0
	for _, line := range input {
		mins := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		var (
			gameNum int
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
			}
		}
		power += (mins["red"] * mins["blue"] * mins["green"])
	}
	return power
}
