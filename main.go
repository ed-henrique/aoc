package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Day One
	numbers := map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}

	inputDayOne, err := os.ReadFile("./input.txt")

	if err != nil {
		os.Exit(1)
	}

	words := strings.Split(string(inputDayOne), "\n")

	firstResult := DayOnePartOne(words)
	secondResult := DayOnePartTwo(words, numbers)

	fmt.Printf("Day One\n\nFirst Result: %6d\nSecond Result: %5d\n\n", firstResult, secondResult)

	// Day Two

	secondDayInput, err := os.ReadFile("./input2.txt")

	if err != nil {
		os.Exit(1)
	}

	lines := strings.Split(string(secondDayInput), "\n")
	newLines := []string{}

	for _, line := range lines {
		newLine := strings.Split(string(line), ": ")

		if len(newLine) > 1 {
			newLines = append(newLines, newLine[1])
		}
	}

	fmt.Println(newLines)
	packages := strings.Split(string(secondDayInput), "; ")

	fmt.Println(packages)
}
