package utils

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile(fileName string) []string {
	data, err := os.ReadFile(
		fileName,
	)
	if err != nil {
		fmt.Printf("Error reading a file: %s\n", err)
		panic(err)
	}
	return strings.Split(string(data), "\n")
}
