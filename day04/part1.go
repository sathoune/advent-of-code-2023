package day04

import (
	"advent-of-code-2023/utils"
	"fmt"
	"math"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

type Scratchpad struct {
	winning   map[int]int
	scratched []int
}

func (s Scratchpad) countMatches() (matches int) {
	for _, number := range s.scratched {
		_, ok := s.winning[number]
		if ok {
			matches += 1
		}
	}
	return
}

func parseGameNumbers(numbers string) (parsed []int) {
	split := strings.Split(numbers, " ")
	for _, stringNumber := range split {
		if stringNumber != "" {
			converted, _ := strconv.Atoi(stringNumber)
			parsed = append(parsed, converted)
		}
	}
	return
}

func toDayMap(data []string) (scratchPads map[int]Scratchpad) {
	scratchPads = make(map[int]Scratchpad)

	for index, line := range data {
		res := strings.Split(line, ": ")
		gameNumbers := strings.Split(res[1], " | ")
		winning := parseGameNumbers(gameNumbers[0])
		scratched := parseGameNumbers(gameNumbers[1])
		winningMap := make(map[int]int)

		for _, number := range winning {
			winningMap[number] = number
		}

		scratchPads[index] = Scratchpad{
			winningMap,
			scratched,
		}
	}
	return
}

func Part1() {
	_, thisFilepath, _, _ := runtime.Caller(0)
	dataFilepath := filepath.Join(filepath.Dir(thisFilepath), "input.txt")
	data := utils.ReadFile(dataFilepath)
	scratchpads := toDayMap(data)

	points := 0.0
	for _, scratchpad := range scratchpads {
		matches := scratchpad.countMatches()
		if matches > 0 {
			points += math.Pow(float64(2), float64(matches-1))
		}
	}

	fmt.Println(points)
}
