package days

import (
	"fmt"
	"path/filepath"
	"regexp"

	"github.com/tuxxi/aoc2023/util"
)

func init() {
	util.RegisterDay(8, day8)
}

type entry struct {
	L string
	R string
}

func day8(input []string) (any, any) {
	instructions, m := parseInput(input)

	part1 := d8part1(instructions, m, "AAA", "ZZZ")
	part2 := d8part2(instructions, m)

	return part1, part2
}

func d8part1(instructions string, m map[string]entry, start, target string) int {
	node := start
	for step := 0; ; step++ {
		if matched, _ := filepath.Match(target, node); matched {
			return step
		}
		next := m[node]
		instr := instructions[step%len(instructions)]
		if instr == 'L' {
			node = next.L
		} else if instr == 'R' {
			node = next.R
		}
	}
}

func d8part2(instructions string, m map[string]entry) int {
	// p2 - start on nodes ending in 'A', end on any node ending in 'Z'
	// the total time required is the product of all cycles
	var allsteps []int
	for node := range m {
		if node[2] == 'A' {
			allsteps = append(allsteps, d8part1(instructions, m, node, "??Z"))
		}
	}
	// fmt.Println("allsteps:", allsteps)

	part2 := 1
	for _, step := range allsteps {
		lcm := LCM(part2, step)
		fmt.Printf("LCM(%d,%d) = %d\n", part2, step, lcm)
		if lcm > part2 {
			part2 = lcm
		}
	}

	return part2
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	return a * b / GCD(a, b)
}

func parseInput(input []string) (string, map[string]entry) {
	instructions := input[0]

	r, err := regexp.Compile(`[A-Z1-9]{3}`)
	if err != nil {
		panic(err)
	}

	m := make(map[string]entry)

	for _, line := range input[2:] {
		elems := r.FindAllStringSubmatch(line, -1)
		m[elems[0][0]] = entry{elems[1][0], elems[2][0]}
	}
	return instructions, m
}
