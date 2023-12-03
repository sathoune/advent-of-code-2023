package day03

import (
	"advent-of-code-2023/utils"
	"fmt"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

type Matrix = [][]string

func sliceToMatrix(text []string) (matrix Matrix) {
	for _, row := range text {
		matrix = append(matrix, strings.Split(row, ""))
	}
	return
}

type Number struct {
	row   int
	start int
	end   int
	value int
}

func findDigitsInRows(matrix []string) (digits []Number) {
	pattern := regexp.MustCompile(`\d+`)

	for rowIndex, row := range matrix {
		foundCoordinates := pattern.FindAllStringIndex(row, -1)
		for _, coordinates := range foundCoordinates {
			start, end := coordinates[0], coordinates[1]
			parsedNumber, _ := strconv.Atoi(row[start:end])

			digits = append(digits, Number{
				rowIndex,
				start,
				end,
				parsedNumber,
			})
		}
	}
	return
}

func Part1() {
	_, thisFilepath, _, _ := runtime.Caller(0)
	dataFilepath := filepath.Join(filepath.Dir(thisFilepath), "demo.txt")
	txt := utils.ReadFile(dataFilepath)
	digitsCoordinates := findDigitsInRows(txt)
	fmt.Println(digitsCoordinates)
}
