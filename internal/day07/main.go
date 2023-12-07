// Package day07 is our package for the 7th AoC day
package day07

import (
	"sort"
	"strconv"
	"strings"
)

// Hand hold all data we need for the hand (ever row)
type Hand struct {
	Index      int
	Cards      Cards
	CardsValue []int
	Sum        int
	Bid        int
}

// Cards is just a string array
// build a struct, so I can write functions on it
type Cards []string

// Count gives us the number of cards matching the input character
func (c Cards) Count(i string) int {
	sum := 0
	for _, v := range c {
		if v == i {
			sum++
		}
	}
	return sum
}

// Unique returns the unique set of cards
func (c Cards) Unique() Cards {
	var uniqueCards Cards
	for _, v := range c {
		if uniqueCards.Count(v) == 0 {
			uniqueCards = append(uniqueCards, v)
		}
	}
	return uniqueCards
}

// Sum calculates the value of the set of cards
func (c Cards) Sum() int {
	uniqueCards := c.Unique()
	cardSum := 0
	for _, number := range uniqueCards {
		switch c.Count(number) {
		case 2:
			cardSum++
		case 3:
			cardSum += 3
		case 4:
			cardSum += 6
		case 5:
			cardSum += 7
		}
	}
	return cardSum
}

func getCardValue(card string) int {
	switch card {
	case "2":
		return 2
	case "3":
		return 3
	case "4":
		return 4
	case "5":
		return 5
	case "6":
		return 6
	case "7":
		return 7
	case "8":
		return 8
	case "9":
		return 9
	case "T":
		return 10
	case "J":
		return 11
	case "Q":
		return 12
	case "K":
		return 13
	case "A":
		return 14
	}
	return -1
}

// RunPart1 is for the first star of the day
func RunPart1(input []string) int {
	var hands []Hand
	for lineIdx, line := range input {
		var cards Cards
		var newHand Hand
		newHand.Index = lineIdx
		hand := strings.Fields(line)[0]
		bid := strings.Fields(line)[1]
		newHand.Bid, _ = strconv.Atoi(bid)
		cards = strings.Split(hand, "")
		newHand.Cards = cards
		for _, c := range cards {
			newHand.CardsValue = append(newHand.CardsValue, getCardValue(c))
		}
		newHand.Sum = cards.Sum()
		hands = append(hands, newHand)
	}
	sort.SliceStable(hands, func(a, b int) bool {
		if hands[a].Sum < hands[b].Sum {
			return true
		} else if hands[a].Sum == hands[b].Sum {
			for i := 0; i < len(hands[a].CardsValue); i++ {
				if hands[a].CardsValue[i] == hands[b].CardsValue[i] {
					continue
				} else if hands[a].CardsValue[i] < hands[b].CardsValue[i] {
					return true
				}
				return false

			}
		}
		return false
	})
	sum := 0
	for handIdx, hand := range hands {
		sum += (handIdx + 1) * hand.Bid
	}
	return sum
}

// RunPart2 is for the second star of the day
func RunPart2(input []string) int {
	var hands []Hand
	for lineIdx, line := range input {
		var cards Cards
		var newHand Hand
		newHand.Index = lineIdx
		hand := strings.Fields(line)[0]
		bid := strings.Fields(line)[1]
		newHand.Bid, _ = strconv.Atoi(bid)
		cards = strings.Split(hand, "")
		newHand.Cards = append([]string{}, cards...)
		for _, c := range cards {
			value := getCardValue(c)
			if value == 11 {
				value = 1
			}
			newHand.CardsValue = append(newHand.CardsValue, value)
		}
		maxCardSum := 0
		replaceableCards := []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"}
		for _, r := range replaceableCards {
			for k := range cards {
				if cards[k] == "J" {
					cards[k] = r
				}
			}
			cardSum := cards.Sum()
			if cardSum > maxCardSum {
				maxCardSum = cardSum
			}
			cards = append([]string{}, newHand.Cards...)
		}
		newHand.Sum = maxCardSum
		hands = append(hands, newHand)
	}
	sort.SliceStable(hands, func(a, b int) bool {
		if hands[a].Sum < hands[b].Sum {
			return true
		} else if hands[a].Sum == hands[b].Sum {
			for i := 0; i < len(hands[a].CardsValue); i++ {
				if hands[a].CardsValue[i] == hands[b].CardsValue[i] {
					continue
				} else if hands[a].CardsValue[i] < hands[b].CardsValue[i] {
					return true
				}
				return false

			}
		}
		return false
	})
	sum := 0
	for handIdx, hand := range hands {
		sum += (handIdx + 1) * hand.Bid
	}
	return sum
}
