package main

import (
	"github.com/brandonlbarrow/advent-of-code/2022/common"
	"log"
)

const (
	PlayerRock       = "A"
	PlayerPaper      = "B"
	PlayerScissors   = "C"
	OpponentRock     = "X"
	OpponentPaper    = "Y"
	OpponentScissors = "Z"
)

func main() {

	input, err := common.GetInput("input")
	if err != nil {
		log.Fatal(err)
	}

}
