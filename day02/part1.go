package day02

import (
	"advent-of-code-2023/utils"
	"fmt"
	"path/filepath"
	"runtime"
)

var maxRed = 12
var maxGreen = 13
var maxBlue = 14

func Part1() {
	_, thisFilepath, _, _ := runtime.Caller(0)
	dataFilepath := filepath.Join(filepath.Dir(thisFilepath), "demo.txt")
	txt := utils.ReadFile(dataFilepath)
	fmt.Println(txt)
}
