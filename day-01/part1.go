package main

import (
	"fmt"
	"os"
	"strings"
)

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

func main() {
	inputName := "input.txt"
	formattedData := readFile(inputName)
	fmt.Println(formattedData)
	fmt.Println(len(formattedData))
}
