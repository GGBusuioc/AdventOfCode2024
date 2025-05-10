package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello Advent of Code Day 1 - Part 1")

	// Step 1: Open the file
	file, err := os.Open("./day1/input1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var holder1 [1000]int
	var holder2 [1000]int

	var solutionArray [1000]int

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

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// convert holder1 & holder2 to slice before performing the sorting
	sliceHolder1 := holder1[:]
	sliceHolder2 := holder2[:]

	// Sort in ascending order first array
	sort.Slice(sliceHolder1, func(i, j int) bool {
		return sliceHolder1[i] < sliceHolder1[j] // '>' for descending order
	})

	// Sort in ascending order second array
	sort.Slice(sliceHolder2, func(i, j int) bool {
		return sliceHolder2[i] < sliceHolder2[j] // '>' for descending order
	})

	// FINAL PART: Calculate the outcome by subtracting sliceHolder1[i] from sliceHolder2[i] or vice versa and store it in the totalDistance array
	for i := 0; i < 1000; i++ {
		if sliceHolder1[i] > sliceHolder2[i] {
			solutionArray[i] = sliceHolder1[i] - sliceHolder2[i]
		} else {
			solutionArray[i] = sliceHolder2[i] - sliceHolder1[i]
		}

		totalDistance += solutionArray[i]
	}

	// fmt.Println("Sorted array in ascending order:", solutionArray)
	fmt.Println("Total distance:", totalDistance)
}
