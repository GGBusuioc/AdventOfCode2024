package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day2_1(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	output := 0

	for _, line := range lines {
		parts := strings.Fields(line)
		levels := make([]int, len(parts))

		for i, part := range parts {
			num, _ := strconv.Atoi(part)
			levels[i] = num
		}

		isIncreasing := true
		isSafe := true

		for i := 1; i < len(levels); i++ {
			diff := levels[i] - levels[i-1]

			if diff == 0 {
				isSafe = false
				break
			}

			if diff < -3 || diff > 3 {
				isSafe = false
				break
			}

			if i == 1 && diff < 0 {
				isIncreasing = false
			} else if (isIncreasing && diff < 0) || (!isIncreasing && diff > 0) {
				isSafe = false
				break
			}
		}

		if isSafe {
			output += 1
		}
	}

	fmt.Println("Output Day 2 Part 1:", output)
}
