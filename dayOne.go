package main

import (
	"strings"
)

func DayOnePartOne(words []string) (total int) {
	for _, word := range words {
		var firstDigit, secondDigit byte
		foundFD, foundSD := false, false

		for i := 0; i < len(word); i++ {
			if !foundFD && word[i] >= 48 && word[i] <= 57 {
				firstDigit = word[i]
				foundFD = true
			}

			if !foundSD && word[len(word)-i-1] >= 48 && word[len(word)-i-1] <= 57 {
				secondDigit = word[len(word)-i-1]
				foundSD = true
			}

			if foundFD && foundSD {
				break
			}
		}

		if foundFD && foundSD {
			number := int(firstDigit)*10 + int(secondDigit) - 528
			total += number
		}
	}

	return
}

func DayOnePartTwo(words []string, numbers map[string]string) int {
	newWords := []string{}

	for _, word := range words {
		newWord := word
		firstIndex := len(word)
		lastIndex := -1
		var firstDigit, secondDigit string
		firstDigitNumber := false

		for name := range numbers {
			for i := 0; i < len(word)-len(name)+1; i++ {
				if i < firstIndex && word[i:len(name)+i] == name {
					firstIndex = i
					firstDigit = name
				}

				if i > lastIndex && word[i:len(name)+i] == name {
					lastIndex = i
					secondDigit = name
				}
			}
		}

		for i := 0; i < firstIndex; i++ {
			if i < firstIndex && word[i] >= 48 && word[i] <= 57 {
				firstIndex = i
				firstDigitNumber = true

			}
		}

		if !firstDigitNumber && firstIndex != len(word) {
			newWord = strings.Replace(newWord, firstDigit, numbers[firstDigit], 1)
		}

		if lastIndex != -1 {
			newWord = strings.ReplaceAll(newWord, secondDigit, numbers[secondDigit])
		}

		newWords = append(newWords, newWord)
	}

	return DayOnePartOne(newWords)
}
