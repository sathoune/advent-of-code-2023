package day04

import (
	"advent-of-code-2023/utils"
	"fmt"
	"path/filepath"
	"runtime"
	"sort"
)

type ScratchpadSet struct {
	Scratchpad
	Count int
}

func NewScratchpadSet(s Scratchpad) *ScratchpadSet {
	return &ScratchpadSet{
		s,
		1,
	}
}

func Part2() {
	_, thisFilepath, _, _ := runtime.Caller(0)
	dataFilepath := filepath.Join(filepath.Dir(thisFilepath), "input.txt")
	data := utils.ReadFile(dataFilepath)
	scratchpads := toDayMap(data)

	scratchpadSets := make(map[int]*ScratchpadSet)
	for cardNumber, scratchpad := range scratchpads {
		scratchpadSets[cardNumber] = NewScratchpadSet(scratchpad)
	}

	keys := make([]int, 0)
	for k := range scratchpadSets {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, cardKey := range keys {
		matches := scratchpadSets[cardKey].Scratchpad.countMatches()
		for x := 1; x < matches+1; x++ {
			if entry, ok := scratchpadSets[cardKey+x]; ok {
				entry.Count += 1 * scratchpadSets[cardKey].Count
			}
		}
	}

	sum := 0
	for _, scratchpadSet := range scratchpadSets {
		sum += scratchpadSet.Count
	}
	fmt.Println(sum)
}
