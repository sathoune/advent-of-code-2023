package day02

import (
	"advent-of-code-2023/utils"
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

type Stage struct {
	Red   int
	Green int
	Blue  int
}

type Game = []Stage

var limits = Stage{
	12,
	13,
	14,
}

func parseId(text string) (id int) {
	id, _ = strconv.Atoi(strings.Replace(text, "Game ", "", 1))
	return
}

func splitPick(pick string) (count int, colour string) {
	split := strings.Split(pick[1:], " ")
	count, _ = strconv.Atoi(split[0])
	colour = split[1]
	return
}

func parseGame(text string) (result Stage) {
	foundPicks := strings.Split(text, ",")

	for _, pick := range foundPicks {
		count, colour := splitPick(pick)
		if colour == "blue" {
			result.Blue = count
		} else if colour == "red" {
			result.Red = count
		} else if colour == "green" {
			result.Green = count
		}
	}
	return
}

func splitIntoChunks(text string) (id int, results []Stage) {
	parts := strings.Split(text, ":")
	if len(parts) > 2 {
		panic("To many colons!")
	}
	id = parseId(parts[0])
	for _, gameText := range strings.Split(parts[1], ";") {
		results = append(results, parseGame(gameText))
	}

	return
}

func validateStage(stage Stage) (valid bool) {
	if stage.Red > limits.Red {
		return false
	}
	if stage.Green > limits.Green {
		return false
	}
	if stage.Blue > limits.Blue {
		return false
	}
	return true
}

func validateGame(game Game) (valid bool) {
	for _, stage := range game {
		stageValid := validateStage(stage)
		if !stageValid {
			return false
		}
	}
	return true
}

func Part1() {
	_, thisFilepath, _, _ := runtime.Caller(0)
	dataFilepath := filepath.Join(filepath.Dir(thisFilepath), "input.txt")
	txt := utils.ReadFile(dataFilepath)

	grouped := make(map[int]Game)
	for _, line := range txt {
		id, game := splitIntoChunks(line)
		grouped[id] = game
	}
	sum := 0
	for id, game := range grouped {
		valid := validateGame(game)
		if valid {
			sum += id
		}
	}
	fmt.Println(sum)
}
