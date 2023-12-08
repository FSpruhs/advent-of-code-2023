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

func handtypeForPartTwo(cards [5]card) handType {
	cardMap := createCardMap(cards)
	if cardMap[joker] == 0 || cardMap[joker] == 5 {
		return findHandType(cardMap)
	}
	maxCards := 0
	tempCardKey := joker
	for key, value := range cardMap {
		if key != joker && value > maxCards {
			tempCardKey = key
			maxCards = value
		}
	}
	cardMap[tempCardKey] = cardMap[tempCardKey] + cardMap[joker]
	delete(cardMap, joker)
	return findHandType(cardMap)
}

func createCardMap(cards [5]card) map[card]int {
	cardMap := map[card]int{}
	for _, c := range cards {
		cardMap[c] = cardMap[c] + 1
	}
	return cardMap
}

func findHandType(cardMap map[card]int) handType {
	log.Println(cardMap)
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
		log.Printf("Can not find card type for %d", cardMap)
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
		handtype: findHandType(createCardMap(initCards)),
		cards:    initCards,
	}
}

func (h *hand) setHandtype(ht handType) {
	h.handtype = ht
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
		solutionPartOne += h.bid * (i + 1)
	}

	solutionPartTwo := 0
	for index := range hands {
		hands[index].setHandtype(handtypeForPartTwo(hands[index].cards))
	}
	sort.SliceStable(hands, func(i, j int) bool {
		if hands[i].handtype != hands[j].handtype {
			return hands[i].handtype < hands[j].handtype
		}
		for index := 0; index < 5; index++ {
			if hands[i].cards[index] != hands[j].cards[index] {
				if hands[i].cards[index] == joker {
					return true
				}
				if hands[j].cards[index] == joker {
					return false
				}
				return hands[i].cards[index] < hands[j].cards[index]
			}
		}
		return false
	})
	for i, h := range hands {
		solutionPartTwo += h.bid * (i + 1)
	}
	log.Println(hands)
	fmt.Printf("Solution of day 7 part 1: %d\n", solutionPartOne)
	fmt.Printf("Solution of day 7 part 2: %d\n", solutionPartTwo)
}

func readLine(line string) (string, int) {
	splittedLine := strings.Split(line, " ")
	return splittedLine[0], util.ConvertToInt(splittedLine[1])

}
