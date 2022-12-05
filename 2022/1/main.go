package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	m, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		return
	}
	contents, err := io.ReadAll(m)
	if err != nil {
		fmt.Println(err)
		return
	}

	var currentHighest, previousHighest, result int
	parsedContents := strings.Split(string(contents), "\n\n")
	for _, g := range parsedContents {
		parsedGroup := strings.Split(strings.TrimSpace(g), "\n")
		currentHighest = 0
		for _, gg := range parsedGroup {
			i, _ := strconv.Atoi(gg)
			currentHighest = i + currentHighest
		}
		if currentHighest >= previousHighest {
			result = currentHighest
			previousHighest = currentHighest
		} else {
			continue
		}
	}
	fmt.Println(result)
}
