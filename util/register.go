package util

import "fmt"

type dayImpl func([]string) (any, any)

type Day struct {
	dayImpl
}

func (d Day) Part1(lines []string) any {
	p1, _ := d.dayImpl(lines)
	return p1
}

func (d Day) Part2(lines []string) any {
	_, p2 := d.dayImpl(lines)
	return p2
}

var (
	days = make(map[int]dayImpl)
)

func RegisterDay(d int, impl dayImpl) {
	days[d] = impl
}

func GetDay(d int) Day {
	impl, ok := days[d]
	if !ok {
		panic(fmt.Sprintf("could not find day: %d", d))
	}
	return Day{impl}
}
