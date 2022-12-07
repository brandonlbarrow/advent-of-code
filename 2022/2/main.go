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
	getWinDecision = map[string]int{
		"X": LOSS,
		"Y": DRAW,
		"Z": WIN,
	}
	winDecisionString = map[int]string{
		WIN:  "WIN",
		LOSS: "LOSS",
		DRAW: "DRAW",
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
	fmt.Println("player score round 1:", playerScore)
	playerScore = 0

	for _, x := range parsedInput {
		y := strings.Split(x, " ")
		if len(y) != 2 {
			log.Fatal(err)
		}
		opponentInput := y[0]
		playerInput := y[1]
		fmt.Printf("playing round 2, opponent plays %s, must have outcome %s\n", getSymbol[opponentInput], winDecisionString[getWinDecision[playerInput]])
		currentScore := playRoundForOutcome(opponentInput, playerInput)
		fmt.Println("score was:", currentScore)
		fmt.Println("current total score:", playerScore)
		playerScore += currentScore
	}
	fmt.Println("player score round 2:", playerScore)

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

func playRoundForOutcome(opponentInput, desiredOutcome string) int {
	switch getSymbol[opponentInput] {
	case rock:
		switch getWinDecision[desiredOutcome] {
		case LOSS:
			return LOSS + getPoints[scissors]
		case WIN:
			return WIN + getPoints[paper]
		case DRAW:
			return DRAW + getPoints[rock]
		}
	case paper:
		switch getWinDecision[desiredOutcome] {
		case LOSS:
			return LOSS + getPoints[rock]
		case DRAW:
			return DRAW + getPoints[paper]
		case WIN:
			return WIN + getPoints[scissors]
		}
	case scissors:
		switch getWinDecision[desiredOutcome] {
		case LOSS:
			return LOSS + getPoints[paper]
		case DRAW:
			return DRAW + getPoints[scissors]
		case WIN:
			return WIN + getPoints[rock]
		}
	}
	return 0
}
