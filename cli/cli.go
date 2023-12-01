package cli

import (
	"fmt"
	"os"
	"strconv"
)

func ParseArguments() (day int, part int) {
	var err error = nil
	if len(os.Args) < 3 {
		fmt.Println("Please run with: go run main.go <day> <part>")
		os.Exit(1)
	}

	day, err = strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid day number.")
		os.Exit(1)
	}

	part, err = strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid part number.")
		os.Exit(1)
	}
	return
}

func SolutionToExecute(
	availableSolutions map[int]map[int]func(),
	day int,
	part int,
) (
	solution func(),
) {
	ok := false

	solution, ok = availableSolutions[day][part]
	if !ok {
		fmt.Printf("Solution for day %v, part %v doesn't exist yet.\n", day, part)
		os.Exit(1)
	}
	return
}
