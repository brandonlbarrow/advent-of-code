package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
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
	inputSlice := parseInputs(inputs)

	part1Solution := runPart1(inputSlice)
	fmt.Printf("Part 1 Solution is %d\n", part1Solution)

	part2Solution := runPart2(inputSlice)
	fmt.Printf("Part 2 Solution is %d\n", part2Solution)
}

func runPart1(inputs []int) int {
	return calculateDepthIncreaseCount(inputs)
}

func runPart2(inputs []int) int {
	return calculateDepthSumIncreaseCount(inputs)
}

func parseInputs(raw []byte) []int {
	output := []int{}
	rawString := string(raw)
	rawStringSlice := strings.Split(rawString, "\n")
	for _, s := range rawStringSlice {
		i, _ := strconv.Atoi(s)
		output = append(output, i)
	}
	return output
}

func calculateDepthIncreaseCount(depths []int) int {
	var count, previous, current int
	for i := range depths {
		current = depths[i]
		if i > 0 {
			previous = depths[i-1]
		} else {
			continue
		}
		if current > previous {
			fmt.Printf("Depth increased from %d to %d.\n", previous, current)
			fmt.Println(count)
			count++
		}
	}
	return count
}

func calculateDepthSumIncreaseCount(depths []int) int {
	var count, currentSum, previousSum int
	for i := range depths {
		if depths[i] == depths[len(depths)-2] {
			fmt.Println("cannot create new window with remaining elements. exiting loop.")
			break
		} else {
			previousSum = currentSum
			fmt.Printf("GROUP %d: %d: %d\n\n", depths[i], depths[i+1], depths[i+2])
			currentSum = depths[i] + depths[i+1] + depths[i+2]
			if currentSum > previousSum && previousSum != 0 {
				fmt.Printf("Sum of window increased from %d to %d.\n", previousSum, currentSum)
				count++
			}

		}
	}
	return count
}
