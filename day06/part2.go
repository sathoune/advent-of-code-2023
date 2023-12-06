package day06

import (
	"advent-of-code-2023/utils"
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func parseDataAsSingleInt(text string) (parsed int) {
	parts := strings.Split(text, ": ")
	numberString := strings.Replace(
		parts[1],
		" ",
		"",
		-1,
	)
	if numberString != "" {
		parsed, _ = strconv.Atoi(numberString)
	}
	return
}

func Part2() {
	_, thisFilepath, _, _ := runtime.Caller(0)
	dataFilepath := filepath.Join(filepath.Dir(thisFilepath), "input.txt")
	data := utils.ReadFile(dataFilepath)

	topScores := parseTopScores(
		[]int{parseDataAsSingleInt(data[0])},
		[]int{parseDataAsSingleInt(data[1])},
	)

	ans := 1
	for _, topScore := range topScores {
		solutions := findStrategies(topScore)
		ans *= solutions
	}
	fmt.Println(ans)
}
