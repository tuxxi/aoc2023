package days

import (
	"fmt"
	"slices"

	"github.com/tuxxi/aoc2023/util"
)

func init() {
	util.RegisterDay(10, day10)
}

type tile struct {
	c         rune
	reachable []*tile
	x, y      int
	dist      int
}

func (t *tile) String() string {
	return fmt.Sprintf("%c@(%d, %d)", t.c, t.x, t.y)
}

type maze struct {
	tiles [][]tile
	start *tile
}

func day10(input []string) (any, any) {
	var part1, part2 int

	// parse input into a graph structure
	m := parseMaze(input)

	// part1: find the maximum distance from the starting point of the graph.
	// use depth-first floodfill algorithm
	var seen []*tile = []*tile{m.start}
	var tileStack []*tile = []*tile{m.start}

	for len(tileStack) != 0 {
		curr := tileStack[0]
		tileStack = tileStack[1:] // pop
		// fmt.Printf("processing: %s\n", curr)
		for _, next := range curr.reachable {
			if !slices.Contains(seen, next) {
				// fmt.Printf("new tile reachable from '%s': '%s'\n", curr, next)
				tileStack = append(tileStack, next) // push

				// calculate distances
				next.dist = curr.dist + 1
				if next.dist > part1 {
					part1 = next.dist
				}
				seen = append(seen, next)
			}
		}
		// fmt.Printf("tile stack: %v\n", tileStack)
	}
	return part1, part2
}

func parseMaze(input []string) *maze {
	m := &maze{}

	// prefil with valid memory
	for y, line := range input {
		m.tiles = append(m.tiles, []tile{})
		for x, c := range line {
			m.tiles[y] = append(m.tiles[y], tile{c, nil, x, y, 0})
		}
	}

	// parse
	for y := range m.tiles {
		for x := range m.tiles[y] {
			switch m.tiles[y][x].c {
			case '|': // N-S vertical connector
				if y > 0 {
					m.tiles[y][x].reachable = append(m.tiles[y][x].reachable, &m.tiles[y-1][x])
				}
				if y < len(m.tiles)-1 {
					m.tiles[y][x].reachable = append(m.tiles[y][x].reachable, &m.tiles[y+1][x])
				}
			case '-': // E-W vertical connector
				if x > 0 {
					m.tiles[y][x].reachable = append(m.tiles[y][x].reachable, &m.tiles[y][x-1])
				}
				if x < len(m.tiles[y])-1 {
					m.tiles[y][x].reachable = append(m.tiles[y][x].reachable, &m.tiles[y][x+1])
				}
			case 'L': // N-E 90deg connector
				if y > 0 {
					m.tiles[y][x].reachable = append(m.tiles[y][x].reachable, &m.tiles[y-1][x])
				}
				if x < len(m.tiles[y])-1 {
					m.tiles[y][x].reachable = append(m.tiles[y][x].reachable, &m.tiles[y][x+1])
				}
			case 'J': // N-W 90deg connector
				if y > 0 {
					m.tiles[y][x].reachable = append(m.tiles[y][x].reachable, &m.tiles[y-1][x])
				}
				if x > 0 {
					m.tiles[y][x].reachable = append(m.tiles[y][x].reachable, &m.tiles[y][x-1])
				}
			case '7': //S-W 90deg connector
				if x > 0 {
					m.tiles[y][x].reachable = append(m.tiles[y][x].reachable, &m.tiles[y][x-1])
				}
				if y < len(m.tiles)-1 {
					m.tiles[y][x].reachable = append(m.tiles[y][x].reachable, &m.tiles[y+1][x])
				}
			case 'F': //S-E 90deg connector
				if x < len(m.tiles[y])-1 {
					m.tiles[y][x].reachable = append(m.tiles[y][x].reachable, &m.tiles[y][x+1])
				}

				if y < len(m.tiles)-1 {
					m.tiles[y][x].reachable = append(m.tiles[y][x].reachable, &m.tiles[y+1][x])
				}
			case 'S':
				m.start = &m.tiles[y][x]
			case '.':
				// ground
			}
		}
	}

	// check reachable from start
	if m.start.y > 0 {
		if slices.Contains(m.tiles[m.start.y-1][m.start.x].reachable, m.start) {
			m.start.reachable = append(m.start.reachable, &m.tiles[m.start.y-1][m.start.x])
		}
	}
	if m.start.y < len(m.tiles)-1 {
		if slices.Contains(m.tiles[m.start.y+1][m.start.x].reachable, m.start) {
			m.start.reachable = append(m.start.reachable, &m.tiles[m.start.y+1][m.start.x])
		}
	}
	if m.start.x > 0 {
		if slices.Contains(m.tiles[m.start.y][m.start.x-1].reachable, m.start) {
			m.start.reachable = append(m.start.reachable, &m.tiles[m.start.y][m.start.x-1])
		}
	}
	if m.start.x < len(m.tiles[m.start.y])-1 {
		if slices.Contains(m.tiles[m.start.y][m.start.x+1].reachable, m.start) {
			m.start.reachable = append(m.start.reachable, &m.tiles[m.start.y][m.start.x+1])
		}
	}
	return m
}
