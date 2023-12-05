package main

import (
	"advent-of-code-2023/cli"
	"advent-of-code-2023/day01"
	"advent-of-code-2023/day02"
	"advent-of-code-2023/day03"
	"advent-of-code-2023/day04"
	"advent-of-code-2023/day05"
	"fmt"
)

var availableSolutions = map[int]map[int]func(){
	1: {
		1: day01.Part1,
		2: day01.Part2,
	},
	2: {
		1: day02.Part1,
		2: day02.Part2,
	},
	3: {
		1: day03.Part1,
		2: day03.Part2,
	},
	4: {
		1: day04.Part1,
		2: day04.Part2,
	},
	5: {
		1: day05.Part1,
		2: day05.Part2,
		3: day05.Part2reimagined,
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
