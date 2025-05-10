package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello Advent of Code Day 1 - Part 2")

	// Step 1: Open the file
	file, err := os.Open("./day1/input1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var holder1 [1000]int
	var holder2 [1000]int

	var similarityScoreArray [1000]int

	var totalDistance int

	i := 0

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Use strings.Fields to split the line by any whitespace
		parts := strings.Fields(line)

		// Ensure there are exactly two parts
		if len(parts) != 2 {
			fmt.Println("Invalid input format:", line)
			continue
		}

		// Convert the first part to an integer
		num1, err1 := strconv.Atoi(parts[0])
		if err1 != nil {
			fmt.Println("Error converting first number:", err1)
			continue
		}

		// Convert the second part to an integer
		num2, err2 := strconv.Atoi(parts[1])
		if err2 != nil {
			fmt.Println("Error converting second number:", err2)
			continue
		}

		holder1[i] = num1
		holder2[i] = num2

		i += 1
	}

	// find similarityScore
	for i := 0; i < 1000; i++ {
		numberOfOccurances := 0
		for j := 0; j < 1000; j++ {
			if holder1[i] == holder2[j] {
				numberOfOccurances += 1
			}
		}
		similarityScoreArray[i] = numberOfOccurances * holder1[i]
		totalDistance += similarityScoreArray[i]
	}
	fmt.Println("Total distance: ", totalDistance)
}
