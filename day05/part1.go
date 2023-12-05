package day05

import (
	"advent-of-code-2023/utils"
	"fmt"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

type Mapper struct {
	from int
	to   int
	step int
}

type ConversionMap struct {
	Mappers []Mapper
}

func (c ConversionMap) mapToDestination(source int) (destination int) {
	for _, mapper := range c.Mappers {
		matches := source >= mapper.from && source < mapper.from+mapper.step
		if matches {
			return mapper.to - mapper.from + source
		}
	}
	return source
}

func parseStringToNumbers(text string) (numbers []int) {
	for _, number := range strings.Split(text, " ") {
		if value, err := strconv.Atoi(number); err == nil {
			numbers = append(numbers, value)
		}
	}
	return
}

func parseSeeds(text string) (seeds []int) {
	parts := strings.Split(text, ": ")
	return parseStringToNumbers(parts[1])
}

func gatherMaps(text []string) (converters []ConversionMap) {
	pattern := regexp.MustCompile(`^[a-z]`)
	currentConverter := ConversionMap{}
	for _, line := range text {
		if pattern.MatchString(line) {
			converters = append(converters, currentConverter)
			currentConverter = ConversionMap{}
		} else if line == "" {

		} else {
			parameters := parseStringToNumbers(line)
			currentConverter.Mappers = append(
				currentConverter.Mappers,
				Mapper{parameters[1],
					parameters[0],
					parameters[2],
				},
			)
		}
	}
	converters = append(converters, currentConverter)
	return converters[1:]
}

func findLocation(seed int, conversions []ConversionMap) (destination int) {
	destination = seed
	for _, conversion := range conversions {
		destination = conversion.mapToDestination(destination)
	}
	return
}

func findMinimum(numbers []int) int {
	minimal := numbers[0]
	for _, number := range numbers[1:] {
		if number < minimal {
			minimal = number
		}
	}
	return minimal
}

func Part1() {
	_, thisFilepath, _, _ := runtime.Caller(0)
	dataFilepath := filepath.Join(filepath.Dir(thisFilepath), "input.txt")
	data := utils.ReadFile(dataFilepath)

	seedsData := data[0]
	seeds := parseSeeds(seedsData)
	rest := data[1:]
	maps := gatherMaps(rest)

	locations := make([]int, 0)
	for _, seed := range seeds {
		locations = append(locations, findLocation(seed, maps))
	}

	fmt.Println(findMinimum(locations))

}
