package days

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/tuxxi/aoc2023/util"
)

func init() {
	util.RegisterDay(6, day6)
}

func day6(input []string) (any, any) {
	var part1, part2 int = 1, 0

	// input is only 2 lines
	if len(input) != 2 {
		panic("unexpected input")
	}

	var times, distances []int
	for _, time := range strings.Fields(input[0])[1:] {
		t, _ := strconv.Atoi(time)
		times = append(times, t)
	}
	for _, dist := range strings.Fields(input[1])[1:] {
		t, _ := strconv.Atoi(dist)
		distances = append(distances, t)
	}

	// part1:
	for i := 0; i < len(times); i++ {
		// skip 0 and max time
		var exceeded int
		for held := 1; held < times[i]; held++ {
			remaining := times[i] - held
			dist := remaining * held
			if dist > distances[i] {
				// fmt.Printf("race %d: holding for %d resulted in %d (max: %d)\n", i+1, held, dist, distances[i])
				exceeded++
			}
		}
		part1 *= exceeded
	}

	// part2:
	var time2, maxDist2 int
	for i := len(times) - 1; i >= 0; i-- {
		var tmul, dmul int = 1, 1 // first digit multiplier is just 1
		if i != len(times)-1 {
			tmul = int(math.Pow10(len(fmt.Sprintf("%d", time2))))
			dmul = int(math.Pow10(len(fmt.Sprintf("%d", maxDist2))))
		}

		time2 += (tmul * times[i])
		maxDist2 += (dmul * distances[i])
	}

	for held := 1; held < time2; held++ {
		remaining := time2 - held
		dist := remaining * held
		if dist > maxDist2 {
			// fmt.Printf("race p2: holding for %d resulted in %d (max: %d)\n", held, dist, maxDist2)
			part2++
		}
	}

	return part1, part2
}
