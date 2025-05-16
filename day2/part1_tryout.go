package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var nums [5]int
var matrix [6][5]int
var output = 0

func main() {
	fmt.Println("Hello Advent of Code Day 2 - Part 1")

	// Step 1: Open the file
	file, err := os.Open("./day2/shorterInput.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	//var numberOfSafeReports int

	i := 0

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Use strings.Fields to split the line by any whitespace
		parts := strings.Fields(line)

		// Ensure there are exactly two parts
		if len(parts) != 5 {
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

		num3, err3 := strconv.Atoi(parts[2])
		if err2 != nil {
			fmt.Println("Error converting second number:", err3)
			continue
		}

		num4, err4 := strconv.Atoi(parts[3])
		if err2 != nil {
			fmt.Println("Error converting second number:", err4)
			continue
		}

		num5, err5 := strconv.Atoi(parts[4])
		if err2 != nil {
			fmt.Println("Error converting second number:", err5)
			continue
		}

		nums[0] = num1
		nums[1] = num2
		nums[2] = num3
		nums[3] = num4
		nums[4] = num5

		for j := 0; j < len(parts); j++ {
			matrix[i][j] = nums[j]
		}

		// Do the logic to determine if a line (report) is safe

		// A report is safe if the level has numbers that
		// -> are all increasing or decreasing
		// -> two adjacent levels differ by at least one or at least 3

		isIncreasing := true
		isSafe := true

		for j := 1; j < len(parts); j++ {
			difference := matrix[i][j] - matrix[i][j-1]

			if difference == 0 {
				isSafe = false
				break
			}

			if difference < -3 || difference > 3 {
				isSafe = false
				break
			}

			// check to see if it's increasing
			if j == 0 && difference < 0 {
				isIncreasing = false
			} else if (isIncreasing && difference < 0) || (!isIncreasing && difference > 0) {
				isSafe = false
				break
			}
		}

		if isSafe {
			output += 1
		}

		i += 1
	}

	// Print the matrix
	fmt.Println("Matrix (2D slice):")
	for _, row := range matrix {
		fmt.Println(row)
	}

	fmt.Println("Report is: ", output)
}
