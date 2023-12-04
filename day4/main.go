package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	Id             int
	Numbers        []int
	WinningNumbers []int
	Matches        []int
	Points         int
	Played         int
}

func main() {
	file, err := os.Open("part1.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cards := make([]*Card, 0)
	for scanner.Scan() {
		value := scanner.Text()
		card := parseCard(&value)
		cards = append(cards, card)
	}

	sum := 0
	p2Sum := 0
	for i := 0; i < len(cards); i++ {
		card := *cards[i]
		sum += card.Points
		matches := len(card.Matches)
		value := card.Played
		p2Sum += value

		for j := 1; j <= matches; j++ {
			idx := card.Id + j - 1
			if idx >= len(cards) {
				continue
			}

			cards[idx].Played += value
		}
	}

	fmt.Println("Part 1: ", sum)
	fmt.Println("Part 2: ", p2Sum)
}

func parseCard(line *string) *Card {
	split := strings.Split(*line, ": ")
	nameIdSplit := strings.Split(split[0], " ")
	id, _ := strconv.Atoi(nameIdSplit[len(nameIdSplit)-1])

	numbers := strings.Split(split[1], " | ")
	myNumbers := make([]int, 0)
	winningNumbers := make([]int, 0)

	for _, number := range strings.Split(numbers[0], " ") {
		num, _ := strconv.Atoi(number)
		myNumbers = append(myNumbers, num)
	}

	for _, number := range strings.Split(numbers[1], " ") {
		num, _ := strconv.Atoi(number)
		if num == 0 {
			continue
		}
		winningNumbers = append(winningNumbers, num)
	}

	matches := make([]int, 0)
	for _, myNumber := range myNumbers {
		for _, winningNumber := range winningNumbers {
			if myNumber == winningNumber {
				matches = append(matches, myNumber)
			}
		}
	}

	points := math.Pow(1.0*2, float64(len(matches))-1)

	return &Card{
		Id:             id,
		Numbers:        myNumbers,
		WinningNumbers: winningNumbers,
		Matches:        matches,
		Points:         int(points),
		Played:         1,
	}
}
