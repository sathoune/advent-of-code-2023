package day01

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var numbersLookup = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func convertFoundNumber(number string) (converted int) {
	converted = numbersLookup[number]
	if converted == 0 {
		converted, _ = strconv.Atoi(number)
	}
	return
}

func findFirstOccurrence(text string, pattern *regexp.Regexp) string {
	return pattern.FindString(text)
}

func reverseString(s string) (reversed string) {
	for _, v := range s {
		reversed = string(v) + reversed
	}
	return
}

func createSearchPatterns() (forward *regexp.Regexp, reverse *regexp.Regexp) {
	var numbers = make([]string, len(numbersLookup)+1)
	var reverseNumbers = make([]string, len(numbersLookup)+1)

	index := 0
	for key := range numbersLookup {
		numbers[index] = key
		reverseNumbers[index] = reverseString(key)
		index++
	}

	numbers[len(numbersLookup)] = "[0-9]"
	forward = regexp.MustCompile(strings.Join(numbers, "|"))

	reverseNumbers[len(numbersLookup)] = "[0-9]"
	reverse = regexp.MustCompile(strings.Join(reverseNumbers, "|"))

	return
}

func Part2() {
	filePath := filepath.Join(thisDirectory(), "input.txt")
	formattedData := readFile(filePath)

	forwardPattern, reversePattern := createSearchPatterns()
	sum := 0
	for _, line := range formattedData {
		firstInteger := findFirstOccurrence(
			line,
			forwardPattern,
		)
		lastInteger := reverseString(
			findFirstOccurrence(
				reverseString(line),
				reversePattern,
			),
		)
		sum += convertFoundNumber(firstInteger) * 10
		sum += convertFoundNumber(lastInteger)

	}
	fmt.Println(sum)
}
