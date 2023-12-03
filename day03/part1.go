package day03

import (
	"advent-of-code-2023/utils"
	"fmt"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
)

type Number struct {
	row   int
	start int
	end   int
	value int
}

type Coordinate struct {
	x int
	y int
}

func findNumbersInRows(matrix []string) (numbers []Number) {
	pattern := regexp.MustCompile(`\d+`)

	for rowIndex, row := range matrix {
		foundCoordinates := pattern.FindAllStringIndex(row, -1)
		for _, coordinates := range foundCoordinates {
			start, end := coordinates[0], coordinates[1]
			parsedNumber, _ := strconv.Atoi(row[start:end])

			numbers = append(numbers, Number{
				rowIndex,
				start,
				end,
				parsedNumber,
			})
		}
	}
	return
}

func findCoordinatesAround(
	partCandidate Number,
	maxX int,
	maxY int,
) (coordinates []Coordinate) {
	yBeforeExists := partCandidate.row-1 >= 0
	yAfterExists := partCandidate.row+1 < maxY
	xBeforeExists := partCandidate.start-1 >= 0
	xAfterExists := partCandidate.end < maxX

	// Before
	if xBeforeExists {
		coordinates = append(coordinates, Coordinate{
			partCandidate.start - 1,
			partCandidate.row,
		})
		// Diagonal <^
		if yBeforeExists {
			coordinates = append(coordinates, Coordinate{
				partCandidate.start - 1,
				partCandidate.row - 1,
			})
		}
		// Diagonal <\/
		if yAfterExists {
			coordinates = append(coordinates, Coordinate{
				partCandidate.start - 1,
				partCandidate.row + 1,
			})
		}

	}
	// After
	if xAfterExists {
		coordinates = append(coordinates, Coordinate{
			partCandidate.end,
			partCandidate.row,
		})
		// Diagonal ^>
		if yBeforeExists {
			coordinates = append(coordinates, Coordinate{
				partCandidate.end,
				partCandidate.row - 1,
			})
		}
		// Diagonal \/>
		if yAfterExists {
			coordinates = append(coordinates, Coordinate{
				partCandidate.end,
				partCandidate.row + 1,
			})
		}
	}

	// Above
	if yBeforeExists {
		for x := partCandidate.start; x < partCandidate.end; x++ {
			coordinates = append(coordinates, Coordinate{
				x,
				partCandidate.row - 1,
			})
		}
	}
	// Below
	if yAfterExists {
		for x := partCandidate.start; x < partCandidate.end; x++ {
			coordinates = append(coordinates, Coordinate{
				x,
				partCandidate.row + 1,
			})
		}
	}
	return
}

func checkIfAnEnginePart(partCandidate Number, matrix []string) bool {
	coordinates := findCoordinatesAround(
		partCandidate,
		len(matrix[0]),
		len(matrix),
	)
	for _, coordinate := range coordinates {
		if string(matrix[coordinate.y][coordinate.x]) != "." {
			return true
		}
	}
	return false
}

func Part1() {
	_, thisFilepath, _, _ := runtime.Caller(0)
	dataFilepath := filepath.Join(filepath.Dir(thisFilepath), "input.txt")
	txt := utils.ReadFile(dataFilepath)

	numbersCoordinates := findNumbersInRows(txt)
	sum := 0
	for _, partCandidate := range numbersCoordinates {
		if checkIfAnEnginePart(partCandidate, txt) {
			sum += partCandidate.value
		}
	}
	fmt.Println(sum)
}
