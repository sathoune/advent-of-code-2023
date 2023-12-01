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

func findFirstNumber(text string, numbers *regexp.Regexp) int {
	matched := numbers.FindString(text)
	if len(matched) == 0 {
		return 0
	}
	return convertFoundNumber(matched)
}
func findLastNumber(text string, reverseNumbers *regexp.Regexp) int {
	r := reverseString(text)
	matched := reverseNumbers.FindString(r)
	if len(matched) == 0 {
		return 0
	}
	return convertFoundNumber(reverseString(matched))
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
	reverseNumbers[len(numbersLookup)] = "[0-9]"
	forward = regexp.MustCompile(strings.Join(numbers, "|"))
	reverse = regexp.MustCompile(strings.Join(reverseNumbers, "|"))

	return

}

func Part2() {
	forwardPattern, reversePattern := createSearchPatterns()
	filePath := filepath.Join(thisDirectory(), "input.txt")
	formattedData := readFile(filePath)

	sum := 0
	for _, line := range formattedData {

		firstInteger := findFirstNumber(line, forwardPattern)
		lastInteger := findLastNumber(line, reversePattern)
		sum += firstInteger * 10
		sum += lastInteger

	}
	fmt.Println(sum)
}
