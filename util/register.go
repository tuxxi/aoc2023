package util

import "fmt"

type Day func([]string) (any, any)

var (
	days = make(map[int]Day)
)

func RegisterDay(d int, impl Day) {
	days[d] = impl
}

func GetDay(d int) Day {
	impl, ok := days[d]
	if !ok {
		panic(fmt.Sprintf("could not find day: %d", d))
	}
	return impl
}
