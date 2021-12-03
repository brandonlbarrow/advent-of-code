package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	fInputFile = flag.String("input-file", "", "file provided as an input to the puzzle")
)

func main() {
	flag.Parse()

	inputs, err := os.ReadFile(*fInputFile)
	if err != nil {
		log.Fatal(fmt.Errorf("error opening or reading input file: %w", err))
	}

	fmt.Printf("Solution to part 1 is: %d\n", solvePart1(inputs))
	fmt.Printf("Solution to part 2 is: %d\n", solvePart2(inputs))
}

type submarinePosition struct {
	horizontal int
	depth      int
	aim        int
}

func solvePart1(inputs []byte) int {
	subPosition := &submarinePosition{}
	movements := parseInputs(inputs)
	fmt.Printf("Processing %d movements...\n", len(movements))
	for _, movement := range movements {
		var change string
		var delta int
		// fmt.Sscan reads space delimited strings into variable pointers
		fmt.Sscan(movement, &change, &delta)
		subPosition = subPosition.move(change, delta)
	}
	return subPosition.location()
}

func solvePart2(inputs []byte) int {
	subPosition := &submarinePosition{}
	movements := parseInputs(inputs)
	fmt.Printf("Processing %d movements with aim factored in\n", len(movements))
	for _, movement := range movements {
		var change string
		var delta int
		fmt.Sscan(movement, &change, &delta)
		subPosition = subPosition.moveWithAim(change, delta)
	}
	return subPosition.location()
}

func (p *submarinePosition) move(change string, delta int) *submarinePosition {
	var newHorizontal, newDepth int
	switch change {
	case "forward":
		newDepth = p.depth
		newHorizontal = p.horizontal + delta
		return &submarinePosition{
			horizontal: newHorizontal,
			depth:      newDepth,
		}
	case "down":
		newDepth = p.depth + delta
		newHorizontal = p.horizontal
		return &submarinePosition{
			horizontal: newHorizontal,
			depth:      newDepth,
		}
	case "up":
		newDepth = p.depth - delta
		newHorizontal = p.horizontal
		return &submarinePosition{
			horizontal: newHorizontal,
			depth:      newDepth,
		}
	default:
		{
			return &submarinePosition{
				horizontal: p.horizontal, depth: p.depth,
			}
		}
	}
}

func (p *submarinePosition) moveWithAim(change string, delta int) *submarinePosition {
	var newHorizontal, newDepth, newAim int
	switch change {
	case "forward":
		newAim = p.aim
		newDepth = p.depth + (p.aim * delta)
		newHorizontal = p.horizontal + delta
		return &submarinePosition{
			aim:        newAim,
			horizontal: newHorizontal,
			depth:      newDepth,
		}
	case "down":
		newAim = p.aim + delta
		newDepth = p.depth
		newHorizontal = p.horizontal
		return &submarinePosition{
			aim:        newAim,
			horizontal: newHorizontal,
			depth:      newDepth,
		}
	case "up":
		newAim = p.aim - delta
		newDepth = p.depth
		newHorizontal = p.horizontal
		return &submarinePosition{
			horizontal: newHorizontal,
			depth:      newDepth,
			aim:        newAim,
		}
	default:
		{
			return &submarinePosition{
				horizontal: p.horizontal,
				depth:      p.depth,
				aim:        p.aim,
			}
		}
	}
}

func (p *submarinePosition) location() int {
	return p.horizontal * p.depth
}

func parseInputs(raw []byte) []string {
	rawString := string(raw)
	rawStringSlice := strings.Split(rawString, "\n")
	return rawStringSlice
}
