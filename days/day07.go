package days

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/tuxxi/aoc2023/util"
)

func init() {
	util.RegisterDay(7, day7)
}

type handPower int

const (
	handPowerHigh      handPower = 0
	handPowerPair      handPower = 1
	handPowerTwoPair   handPower = 2
	handPowerThree     handPower = 3
	handPowerFullHouse handPower = 4
	handPowerFour      handPower = 5
	handPowerFive      handPower = 6
)

var cardRank = map[byte]int{'A': 14, 'K': 13, 'Q': 12, 'J': 11, 'T': 10, '9': 9, '8': 8, '7': 7, '6': 6, '5': 5, '4': 4, '3': 3, '2': 2}
var cardRankP2 = map[byte]int{'A': 14, 'K': 13, 'Q': 12, 'T': 10, '9': 9, '8': 8, '7': 7, '6': 6, '5': 5, '4': 4, '3': 3, '2': 2, 'J': 1}

type hand struct {
	cards string
	power handPower
	bid   int
}

func day7(input []string) (any, any) {
	var hands1, hands2 []hand
	for _, line := range input {
		ss := strings.Fields(line)
		cards := ss[0]
		bid, _ := strconv.Atoi(ss[1])

		hands1 = append(hands1, hand{cards, parseHand(cards), bid})
		hands2 = append(hands2, hand{cards, parseHand2(cards), bid})
	}

	p1 := run7(hands1, cardRank)
	p2 := run7(hands2, cardRankP2)
	return p1, p2
}

func run7(hands []hand, ranking map[byte]int) int {
	maxRank := len(hands)
	// sort hands in descending order (func has inverted cmp result)
	slices.SortFunc(hands, func(a hand, b hand) int {
		if a.power > b.power {
			return -1
		} else if a.power < b.power {
			return 1
		}
		for i := 0; i < 5; i++ {
			if ranking[a.cards[i]] > ranking[b.cards[i]] {
				// fmt.Printf("hands have same power %v, but card %d (%c > %c)\n", a.power, i, a.cards[i], b.cards[i])
				return -1
			} else if ranking[a.cards[i]] < ranking[b.cards[i]] {
				// fmt.Printf("hands have same power %v, but card %d (%c < %c)\n", a.power, i, a.cards[i], b.cards[i])
				return 1
			}
		}
		fmt.Println("Hands cards identical!")
		return 0
	})

	var res int
	// multiply hands bid by rank
	for i, hand := range hands {
		rank := maxRank - i
		res += hand.bid * rank
		// fmt.Printf("Hand %v: rank %d\n", hand, rank)
	}
	return res
}

func parseHand(cards string) handPower {
	seen := make(map[rune]int)
	for _, card := range cards {
		seen[card]++
	}
	power := handPowerHigh
	for _, count := range seen {
		switch count {
		case 2:
			switch power {
			case handPowerPair:
				power = handPowerTwoPair
			case handPowerThree:
				power = handPowerFullHouse
			default:
				power = handPowerPair
			}
		case 3:
			switch power {
			case handPowerPair:
				power = handPowerFullHouse
			default:
				power = handPowerThree
			}
		case 4:
			power = handPowerFour
		case 5:
			power = handPowerFive
		}
	}
	return power
}

func parseHand2(cards string) handPower {
	seen := make(map[rune]int)
	for _, card := range cards {
		seen[card]++
	}

	_, hasJack := seen['J']
	// base case, or edge case: J only
	if !hasJack || (hasJack && len(seen) == 1) {
		return parseHand(cards)
	}

	// greedy recursion algorithm
	maxPower := handPowerHigh
	for i := range util.StrIndexes(cards, 'J') {
		// try replacing each J with one of the other cards
		for c := range seen {
			if c == 'J' {
				continue
			}
			replaced := cards[:i] + string(c) + cards[i+1:]
			maxPower = max(parseHand2(replaced), maxPower)
		}
	}
	return maxPower
}
