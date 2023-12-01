package day01

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

func thisDirectory() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("no caller info")
	}
	return filepath.Dir(filename)
}

func readFile(fileName string) []string {
	data, err := os.ReadFile(
		fileName,
	)
	if err != nil {
		fmt.Printf("Error reading a file: %s\n", err)
		panic(err)
	}
	return strings.Split(string(data), "\n")
}

func findFirstAndLastDigit(text string) (first int, last int, notFound error) {
	pattern, _ := regexp.Compile(`[0-9]`)
	result := pattern.FindAll([]byte(text), -1)
	if len(result) == 0 {
		return 0, 0, errors.New("not found")
	}
	first, _ = strconv.Atoi(string(result[0]))
	last, _ = strconv.Atoi(string(result[len(result)-1]))
	return
}

func Part1() {
	filePath := filepath.Join(thisDirectory(), "input.txt")
	formattedData := readFile(filePath)
	sum := 0
	for _, line := range formattedData {
		firstInteger, lastInteger, _ := findFirstAndLastDigit(line)
		sum += firstInteger * 10
		sum += lastInteger
	}
	fmt.Println(sum)
}
