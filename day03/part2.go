package day03

import (
	"advent-of-code-2023/utils"
	"fmt"
	"path/filepath"
	"regexp"
	"runtime"
)

func findAllGears(text []string) (gears []Number) {
	pattern := regexp.MustCompile(`[*]`)
	for rowIndex, row := range text {
		foundCoordinates := pattern.FindAllStringIndex(row, -1)
		for _, coordinate := range foundCoordinates {
			start, end := coordinate[0], coordinate[1]
			gears = append(gears, Number{
				rowIndex,
				start,
				end,
				0,
			})
		}
	}
	return
}

func findNumbersInCoordinates(
	coordinates []Coordinate,
	matrix []string,
) (numbersAround []Coordinate) {
	pattern := regexp.MustCompile("[0-9]")

	for _, coordinate := range coordinates {
		if pattern.FindString(string(matrix[coordinate.y][coordinate.x])) != "" {
			numbersAround = append(numbersAround, coordinate)
		}
	}
	return
}

func findDigitsAroundGear(
	gearCandidate Number, matrix []string,
) (numbersAround []Coordinate) {
	coordinatesAround := findCoordinatesAround(
		gearCandidate,
		len(matrix[0]),
		len(matrix),
	)

	return findNumbersInCoordinates(coordinatesAround, matrix)
}

func findDigitInNumbers(coordinate Coordinate, numbers []Number) Number {
	for _, number := range numbers {
		yMatches := coordinate.y == number.row
		startMatches := coordinate.x < number.start
		endMatches := coordinate.x >= number.end
		if yMatches && startMatches && endMatches {
			return number
		}
	}
	return Number{}
}

func numberExist(
	number Number,
	numbers []Number,
) bool {
	for _, n := range numbers {
		if number == n {
			return true
		}
	}
	return false
}

func findNumbersForDigits(
	digits []Coordinate,
	numbers []Number,
) (matchedNumbers []Number) {
	for _, digit := range digits {
		matchedNumber := findDigitInNumbers(digit, numbers)
		if !numberExist(matchedNumber, matchedNumbers) {
			matchedNumbers = append(matchedNumbers, matchedNumber)
		}
	}
	return
}

func Part2() {
	_, thisFilepath, _, _ := runtime.Caller(0)
	dataFilepath := filepath.Join(filepath.Dir(thisFilepath), "input.txt")
	txt := utils.ReadFile(dataFilepath)

	gears := findAllGears(txt)
	numbers := findNumbersInRows(txt)
	numbersAroundGears := make([][]Number, len(gears))
	for gearIndex, gear := range gears {
		matchedNumbers := findNumbersForDigits(
			findDigitsAroundGear(gear, txt),
			numbers,
		)
		numbersAroundGears[gearIndex] = matchedNumbers
	}

	gearsPower := 0
	for _, numbersAroundGear := range numbersAroundGears {
		if len(numbersAroundGear) > 1 {
			gearPower := 1
			for _, number := range numbersAroundGear {
				gearPower *= number.value
			}
			gearsPower += gearPower
		}
	}

	fmt.Println(gearsPower)
}
