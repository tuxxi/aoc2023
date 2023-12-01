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

	input := util.GetInputStdin()
	fmt.Printf("Running day: %d. Input is %d lines long\n", *day, len(input))

	fmt.Println("Part1: ", util.GetDay(*day).Part1(input))
	fmt.Println("Part2: ", util.GetDay(*day).Part2(input))
}
