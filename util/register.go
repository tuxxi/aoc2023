package util

import "fmt"

type Day interface {
	Part1([]string) any
	Part2([]string) any
}

var (
	days = make(map[int]Day)
)

func RegisterDay[T Day](d int) {
	var day T
	days[d] = day
}

func GetDay(d int) Day {
	day, ok := days[d]
	if !ok {
		panic(fmt.Sprintf("could not find day: %d", d))
	}
	return day
}
