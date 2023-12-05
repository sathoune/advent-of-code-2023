package day05

import (
	"advent-of-code-2023/utils"
	"fmt"
	"path/filepath"
	"runtime"
	"slices"
	"time"
)

func compareMappers(a Mapper, b Mapper) int {
	return a.to - b.to
}

func (c ConversionMap) sortMappers() {
	slices.SortFunc(c.Mappers, compareMappers)
}

func (c ConversionMap) mapBackwards(destination int) (source int) {
	for _, mapper := range c.Mappers {
		matches := destination >= mapper.to && destination < mapper.to+mapper.step
		if matches {
			return destination - mapper.to + mapper.from
		}
	}
	return destination
}

func seedExists(potentialSeed int, seedRanges []SeedRange) bool {
	for _, seedRange := range seedRanges {
		matches := potentialSeed >= seedRange.Start && potentialSeed < seedRange.Stop
		if matches {
			return true
		}
	}
	return false
}

func findSeed(converters []ConversionMap, seedRanges []SeedRange) int {
	locations := converters[0]
	rest := converters[1:]
	for _, locationMap := range locations.Mappers {
		for v := locationMap.to; v < locationMap.to+locationMap.step; v++ {
			source := locations.mapBackwards(v)
			for _, conversion := range rest {
				source = conversion.mapBackwards(source)
			}
			if seedExists(source, seedRanges) {
				return v
			}
		}
	}
	return -1
}

func Part2reimagined() {
	start := time.Now()
	_, thisFilepath, _, _ := runtime.Caller(0)
	dataFilepath := filepath.Join(filepath.Dir(thisFilepath), "input.txt")
	data := utils.ReadFile(dataFilepath)

	seedsData := data[0]
	seedRanges := parseSeedRanges(seedsData)
	rest := data[1:]
	maps := gatherMaps(rest)
	slices.Reverse(maps)
	maps[0].sortMappers()

	fmt.Println(findSeed(maps, seedRanges))
	fmt.Println(time.Since(start))
}
