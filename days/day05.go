package days

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/tuxxi/aoc2023/util"
	"golang.org/x/exp/slices"
)

func init() {
	util.RegisterDay(5, day5)
}

type mapEntry struct {
	dstStart, srcStart, len_ int
}

type mapping struct {
	entries []mapEntry
}

var (
	seeds        []int = nil
	metaMapping        = make(map[string]mapping)
	mappingOrder       = []string{
		"seed", "soil", "fertilizer", "water", "light", "temperature", "humidity",
	}
)

func day5(input []string) (any, any) {
	var part1, part2 int

	var (
		currMapSrc  string
		currMapText []string
	)
	// pad end of input with blank line to make parsing work
	input = append(input, "")

	for _, line := range input {
		if line == "" {
			// end of map, parse what we have
			fmt.Printf("finalizing map for %s with %d lines\n", currMapSrc, len(currMapText))
			metaMapping[currMapSrc] = parseMap(currMapText)
			currMapText = nil
			continue
		}
		if unicode.IsDigit(rune(line[0])) {
			currMapText = append(currMapText, line)
			continue
		}

		if strings.HasPrefix(line, "seeds:") {
			seeds = util.NumsToInts(strings.Fields(strings.Split(line, ":")[1]))
			continue
		}
		if strings.Contains(line, "-to-") {
			currMapSrc, _, _ = strings.Cut(line, "-to-")
			continue
		}
	}

	var locations []int
	for _, seed := range seeds {
		v := seed
		for _, mapping := range mappingOrder {
			fmt.Printf("%s %d,", mapping, v)
			v = metaMapping[mapping].readMapVal(v)
		}
		fmt.Printf("location %d\n", v)
		locations = append(locations, v)
	}
	part1 = slices.Min(locations)

	return part1, part2
}

func parseMap(mapText []string) mapping {
	output := mapping{}

	for _, line := range mapText {
		nums := util.NumsToInts(strings.Fields(line))
		dstStart := nums[0]
		srcStart := nums[1]
		len_ := nums[2]
		output.entries = append(output.entries, mapEntry{dstStart, srcStart, len_})
	}
	return output
}

// read from the sparse map.
func (m mapping) readMapVal(k int) int {
	for _, e := range m.entries {
		if e.srcStart <= k && k <= e.srcStart+e.len_ {
			return e.dstStart + (k - e.srcStart)
		}
	}
	return k
}
