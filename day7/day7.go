package day7

import (
	"fmt"
	"github.com/FSpruhs/advent-of-code-2023/util"
	"log"
	"os"
	"sort"
	"strings"
)

type card int

const (
	two card = iota
	three
	four
	five
	six
	seven
	eight
	nine
	ten
	joker
	queen
	king
	ass
)

type handType int

const (
	highCard handType = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

type hand struct {
	cards    [5]card
	handtype handType
	bid      int
}

func findHandType(cards [5]card) handType {
	cardMap := map[card]int{}
	for _, c := range cards {
		cardMap[c] = cardMap[c] + 1
	}
	switch len(cardMap) {
	case 1:
		return fiveOfAKind
	case 2:
		return handleTwoDifferentCardTypes(cardMap)
	case 3:
		return handleThreeDifferentCardTypes(cardMap)
	case 4:
		return onePair
	case 5:
		return highCard
	default:
		log.Printf("Can not find card type for %s", cards)
		os.Exit(1)
	}
	return fiveOfAKind
}

func handleThreeDifferentCardTypes(cardMap map[card]int) handType {
	for key := range cardMap {
		if cardMap[key] == 3 {
			return threeOfAKind
		}
	}
	return twoPair
}

func handleTwoDifferentCardTypes(cardMap map[card]int) handType {
	for key := range cardMap {
		if cardMap[key] == 1 || cardMap[key] == 4 {
			return fourOfAKind
		}
	}
	return fullHouse
}

func newHand(cardsString string, initBid int) hand {
	initCards := toCards(cardsString)
	return hand{
		bid:      initBid,
		handtype: findHandType(initCards),
		cards:    initCards,
	}
}

func toCards(cardString string) [5]card {
	cards := [5]card{}
	for i := 0; i < len(cardString); i++ {
		switch fmt.Sprintf("%c", cardString[i]) {
		case "2":
			cards[i] = two
		case "3":
			cards[i] = three
		case "4":
			cards[i] = four
		case "5":
			cards[i] = five
		case "6":
			cards[i] = six
		case "7":
			cards[i] = seven
		case "8":
			cards[i] = eight
		case "9":
			cards[i] = nine
		case "T":
			cards[i] = ten
		case "J":
			cards[i] = joker
		case "Q":
			cards[i] = queen
		case "K":
			cards[i] = king
		case "A":
			cards[i] = ass
		default:
			log.Printf("Can not read card %c", cardString[i])
			os.Exit(1)
		}
	}
	return cards
}

func Solve(filePath string) {
	input := util.ReadFile(filePath)
	hands := []hand{}
	for _, line := range *input {
		initHand := newHand(readLine(line))
		hands = append(hands, initHand)
	}
	sort.SliceStable(hands, func(i, j int) bool {
		if hands[i].handtype != hands[j].handtype {
			return hands[i].handtype < hands[j].handtype
		}
		for index := 0; index < 5; index++ {
			if hands[i].cards[index] != hands[j].cards[index] {
				return hands[i].cards[index] < hands[j].cards[index]
			}
		}
		return false
	})
	solutionPartOne := 0
	for i, h := range hands {
		log.Println(h)
		solutionPartOne += h.bid * (i + 1)
	}

	fmt.Printf("Solution of day 7 part 1: %d\n", solutionPartOne)
}

func readLine(line string) (string, int) {
	splittedLine := strings.Split(line, " ")
	return splittedLine[0], util.ConvertToInt(splittedLine[1])
}
