package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getNums(s string) int {
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}
	return -1
}

type ScratchCard struct {
	winners, numbers, wins []int
	cardId                 int
}

func (sc *ScratchCard) computeWins() {
	var wins []int
	for _, winner := range sc.winners {
		for _, number := range sc.numbers {
			if number == winner {
				wins = append(wins, number)

			}
		}
		sc.wins = wins
	}
}
func (sc ScratchCard) cardScore() int {
	sc.computeWins()
	score := 0
	for i, _ := range sc.wins {
		if i == 0 {
			score++
		} else {
			score = score * 2
		}
	}
	return score
}

func (sc ScratchCard) cardCopies(baseList []ScratchCard) []ScratchCard {
	var copies []ScratchCard
	getFirst := func(index int) ScratchCard {
		for _, card := range baseList {
			if card.cardId == index {
				return card
			}
		}
		return ScratchCard{}
	}
	for i, _ := range sc.wins {
		copy := getFirst(sc.cardId + i + 1)
		copies = append(copies, copy)
	}

	return copies
}

type ScratchCardBuilder struct {
	ScratchCard ScratchCard
}

func NewScratchCardBuilder() *ScratchCardBuilder {
	return &ScratchCardBuilder{
		ScratchCard: ScratchCard{},
	}
}
func (sc *ScratchCardBuilder) Build() ScratchCard {
	return sc.ScratchCard
}
func splitnums(strnums string) []int {
	var nums []int
	splits := strings.Split(strnums, " ")

	for _, num := range splits {
		numVal := getNums(num)
		if numVal > 0 {
			nums = append(nums, numVal)

		}
	}
	return nums
}

func parseScratchCard(index int, input string) ScratchCard {
	card := NewScratchCardBuilder()
	card.ScratchCard.cardId = index

	idChaff := strings.Split(input, ":")
	numbers := strings.Split(idChaff[1], "|")

	card.ScratchCard.winners = splitnums(numbers[0])
	card.ScratchCard.numbers = splitnums(numbers[1])

	return card.Build()
}

func countPoints(cards []ScratchCard) int {
	cardScores := 0
	for _, card := range cards {
		cardScores += card.cardScore()
	}
	return cardScores
}

func CountCopies(cards []ScratchCard) {
	var allcards []ScratchCard = cards
	var usedCards []int
	var continueLoop bool = true
	for continueLoop {
		card := allcards[0]
		if len(allcards) != 1 {
			allcards = allcards[1:]
		} else {
			continueLoop = false
		}
		card.computeWins()
		allcards = append(allcards, card.cardCopies(allcards)...)
		usedCards = append(usedCards, card.cardId)
	}
	fmt.Println(len(usedCards))

}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var cards []ScratchCard
	idx := 0
	for scanner.Scan() {
		//	fmt.Println(scanner.Text())
		cards = append(cards, parseScratchCard(idx, scanner.Text()))
		idx++
	}
	//fmt.Println(countPoints(cards))
	CountCopies(cards)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
