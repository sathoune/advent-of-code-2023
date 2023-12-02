package day02

import (
	"advent-of-code-2023/utils"
	"fmt"
	"path/filepath"
	"runtime"
)

func findMinimalBallsCount(game Game) (minimal Stage) {
	minimal = Stage{0, 0, 0}
	for _, stage := range game {
		if stage.Red > minimal.Red {
			minimal.Red = stage.Red
		}
		if stage.Green > minimal.Green {
			minimal.Green = stage.Green
		}
		if stage.Blue > minimal.Blue {
			minimal.Blue = stage.Blue
		}
	}
	return
}

func Part2() {
	_, thisFilepath, _, _ := runtime.Caller(0)
	dataFilepath := filepath.Join(filepath.Dir(thisFilepath), "input.txt")
	txt := utils.ReadFile(dataFilepath)

	games := parseGames(txt)

	power := 0
	for _, game := range games {
		minimal := findMinimalBallsCount(game)
		power += minimal.Red * minimal.Green * minimal.Blue
	}

	fmt.Println(power)
}
