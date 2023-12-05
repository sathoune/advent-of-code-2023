package day05

import (
	"advent-of-code-2023/utils"
	"fmt"
	"path/filepath"
	"runtime"
)

type SeedRange struct {
	Start int
	Stop  int
}

func parseSeedRanges(text string) (seedRanges []SeedRange) {
	numbers := parseSeeds(text)
	seedRange := SeedRange{}
	for index, number := range numbers {
		if index%2 == 1 {
			seedRange.Stop = number + seedRange.Start
			seedRanges = append(seedRanges, seedRange)
		} else {
			seedRange.Start = number
		}
	}
	return
}

func Part2() {
	_, thisFilepath, _, _ := runtime.Caller(0)
	dataFilepath := filepath.Join(filepath.Dir(thisFilepath), "input.txt")
	data := utils.ReadFile(dataFilepath)

	seedsData := data[0]
	seedRanges := parseSeedRanges(seedsData)
	rest := data[1:]
	maps := gatherMaps(rest)
	locations := make([]int, 0)
	for _, seedRange := range seedRanges {
		fmt.Println("Handling range: ", seedRange)
		locationsForRange := make([]int, 0)
		for v := seedRange.Start; v < seedRange.Stop; v++ {
			locationsForRange = append(locationsForRange, findLocation(v, maps))
		}
		locations = append(locations, findMinimum(locationsForRange))
	}
	fmt.Println(findMinimum(locations))
}
