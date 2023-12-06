package day06

import (
	"advent-of-code-2023/utils"
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

type TopScore struct {
	time     int
	distance int
}

type RaceAttempt struct {
	acceleration int
	distance     int
	winning      bool
}

func parseData(text string) (parsed []int) {
	parts := strings.Split(text, ": ")
	numberStrings := strings.Split(parts[1], " ")
	for _, numberString := range numberStrings {
		if numberString != "" {
			number, _ := strconv.Atoi(numberString)
			parsed = append(parsed, number)
		}
	}
	return
}

func findStrategies(topScore TopScore) (winnings int) {
	for accelerate := 1; accelerate < topScore.time; accelerate++ {
		distance := accelerate * (topScore.time - accelerate)
		if distance > topScore.distance {
			winnings += 1
		}
	}
	return
}

func parseTopScores(times []int, distances []int) (races []TopScore) {
	for idx, time := range times {
		races = append(races, TopScore{time, distances[idx]})
	}
	return
}

func Part1() {
	_, thisFilepath, _, _ := runtime.Caller(0)
	dataFilepath := filepath.Join(filepath.Dir(thisFilepath), "input.txt")
	data := utils.ReadFile(dataFilepath)

	topScores := parseTopScores(
		parseData(data[0]),
		parseData(data[1]),
	)

	ans := 1
	for _, topScore := range topScores {
		solutions := findStrategies(topScore)
		ans *= solutions
	}
	fmt.Println(ans)
}
