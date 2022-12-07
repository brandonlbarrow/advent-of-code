package main

import (
	"fmt"
	"github.com/brandonlbarrow/advent-of-code/2022/common"
	"log"
	"strings"
)

const (
	WIN  = 6
	DRAW = 3
	LOSS = 0
)

var (
	getPoints = map[Symbol]int{
		rock:     1,
		paper:    2,
		scissors: 3,
	}
	getSymbol = map[string]Symbol{
		"A": rock,
		"B": paper,
		"C": scissors,
		"X": rock,
		"Y": paper,
		"Z": scissors,
	}
)

type Symbol string

var (
	rock     Symbol = "rock"
	paper    Symbol = "paper"
	scissors Symbol = "scissors"
)

func main() {

	input, err := common.GetInput("input")
	if err != nil {
		log.Fatal(err)
	}
	parsedInput := strings.Split(strings.Join(strings.Split(string(input), "\n\n"), " "), "\n")
	playerScore := 0
	for _, x := range parsedInput {
		y := strings.Split(x, " ")
		if len(y) != 2 {
			log.Fatal(err)
		}
		opponentInput := y[0]
		playerInput := y[1]
		fmt.Printf("playing round, opponent plays %s, player plays %s\n", getSymbol[opponentInput], getSymbol[playerInput])
		currentScore := playRound(opponentInput, playerInput)
		fmt.Println("score was:", currentScore)
		fmt.Println("current total score:", playerScore)
		playerScore += currentScore
	}
	fmt.Println(playerScore)

}

func playRound(opponentInput, playerInput string) int {
	switch getSymbol[opponentInput] {
	case rock:
		switch getSymbol[playerInput] {
		case rock:
			return DRAW + getPoints[rock]
		case paper:
			return WIN + getPoints[paper]
		case scissors:
			return LOSS + getPoints[scissors]
		}
	case paper:
		switch getSymbol[playerInput] {
		case rock:
			return LOSS + getPoints[rock]
		case paper:
			return DRAW + getPoints[paper]
		case scissors:
			return WIN + getPoints[scissors]
		}
	case scissors:
		switch getSymbol[playerInput] {
		case rock:
			return WIN + getPoints[rock]
		case paper:
			return LOSS + getPoints[paper]
		case scissors:
			return DRAW + getPoints[scissors]
		}
	}
	return 0
}
