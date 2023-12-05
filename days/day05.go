package days

import (
	"fmt"
	"math"
	"strings"
	"sync"
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

type seedsRange struct {
	start, len_ int
}

var (
	seeds   []int = nil
	seedsP2 []seedsRange

	metaMapping  = make(map[string]mapping)
	mappingOrder = []string{
		"seed", "soil", "fertilizer", "water", "light", "temperature", "humidity",
	}
)

func day5(input []string) (any, any) {
	var (
		currMapSrc  string
		currMapText []string
	)
	// pad end of input with blank line to make parsing work
	input = append(input, "")

	for _, line := range input {
		if line == "" {
			// end of map, parse what we have
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
			for i := 0; i < len(seeds); i += 2 {
				seedsP2 = append(seedsP2, seedsRange{seeds[i], seeds[i+1]})
			}

			continue
		}
		if strings.Contains(line, "-to-") {
			currMapSrc, _, _ = strings.Cut(line, "-to-")
			continue
		}
	}

	part1 := math.MaxInt
	for _, seed := range seeds {
		v := seed
		for _, mapping := range mappingOrder {
			// fmt.Printf("%s %d,", mapping, v)
			v = metaMapping[mapping].readMapVal(v)
		}
		// fmt.Printf("location %d\n", v)
		if v < part1 {
			part1 = v
		}
	}

	// bruteforce p2 in parallel to speedup
	var (
		wg        sync.WaitGroup
		p2minsMtx sync.Mutex
		p2mins    []int
	)
	for _, rang := range seedsP2 {
		wg.Add(1)
		go func(rang seedsRange) {
			defer wg.Done()

			eachMin := math.MaxInt
			for seed := rang.start; seed < rang.start+rang.len_; seed++ {
				v := seed
				for _, mapping := range mappingOrder {
					v = metaMapping[mapping].readMapVal(v)
				}
				if v < eachMin {
					eachMin = v
				}
			}
			p2minsMtx.Lock()
			p2mins = append(p2mins, eachMin)
			p2minsMtx.Unlock()
			fmt.Println("completed range:", rang, "minimum:", eachMin)
		}(rang)
	}
	wg.Wait()

	return part1, slices.Min(p2mins)
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
		if e.srcStart <= k && k < e.srcStart+e.len_ {
			return e.dstStart + (k - e.srcStart)
		}
	}
	return k
}
