package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		panic(fmt.Errorf("no command line arguments passed"))
	}

	input := os.Args[1]

	// Remove brackets
	input = strings.ReplaceAll(input, "[", "")
	input = strings.ReplaceAll(input, "]", "")
	input = strings.TrimSpace(input)

	var arr []int

	if input != "" {
		strNums := strings.Split(input, ",")
		for _, s := range strNums {
			num, err := strconv.Atoi(strings.TrimSpace(s))
			if err == nil {
				arr = append(arr, num)
			}
		}
	}

	output := handle(arr)
	fmt.Println(output)
}

/*
 * Implement your solution inside this method.
 * The output return value should be a string.
 */
func handle(input []int) string {
	// WRITE YOUR CODE HERE
	return "Method Not Implemented."
}