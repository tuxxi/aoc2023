package days

import (
	"testing"

	"github.com/tuxxi/aoc2023/util"
)

func TestD7(t *testing.T) {
	input := []string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}
	p1, p2 := day7(input)
	util.AssertEq(t, p1, 6440)
	util.AssertEq(t, p2, 5905)
}

func TestParseCard(t *testing.T) {
	util.AssertEq(t, parseHand2("JJJKQ"), handPowerFour)
	util.AssertEq(t, parseHand2("QJJKJ"), handPowerFour)
	util.AssertEq(t, parseHand2("JJQQK"), handPowerFour)
	util.AssertEq(t, parseHand2("JAJAJ"), handPowerFive)
	util.AssertEq(t, parseHand2("JJJJJ"), handPowerFive)
}
