package day02

import (
	"advent-of-code-2023/utils"
	"fmt"
	"path/filepath"
	"runtime"
)

func Part2() {
	_, thisFilepath, _, _ := runtime.Caller(0)
	dataFilepath := filepath.Join(filepath.Dir(thisFilepath), "demo.txt")
	txt := utils.ReadFile(dataFilepath)
	games := parseGames(txt)
	fmt.Println(games)
}
