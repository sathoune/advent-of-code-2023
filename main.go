package main

import (
	"advent-of-code-2023/cli"
	"advent-of-code-2023/day01"
	"fmt"
)

var availableSolutions = map[int]map[int]func(){
	1: {
		1: day01.Part1,
		2: day01.Part2,
	},
}

func main() {
	day, part := cli.ParseArguments()
	solution := cli.SolutionToExecute(availableSolutions, day, part)

	fmt.Printf("Executing day %v, part %v\n", day, part)
	fmt.Println("Solution:")
	fmt.Println()

	solution()
}