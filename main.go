package main

import (
	"fmt"

	"flag"

	_ "github.com/tuxxi/aoc2023/days" // ensure each day is registered
	"github.com/tuxxi/aoc2023/util"
)

var (
	day = flag.Int("day", 1, "which day to run")
)

func main() {
	flag.Parse()

	input := util.GetPuzzleInput(*day)
	fmt.Printf("Running day: %d. Input is %d lines long\n", *day, len(input))
	p1, p2 := util.GetDay(*day)(input)

	fmt.Println("Part1: ", p1)
	fmt.Println("Part2: ", p2)
}
