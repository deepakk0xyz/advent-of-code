package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(levels []int) bool {
	previous := -1
	increasing := true
	decreasing := true

	for _, value := range levels {
		if previous != -1 {
			if value == previous || value > previous+3 || value < previous-3 {
				return false
			} else if value > previous {
				decreasing = false
			} else {
				increasing = false
			}
		}
		previous = value
	}

	return increasing || decreasing
}

func part2(levels []int) bool {
	if part1(levels) {
		return true
	}

	for skip := range len(levels) {
		var levelsWithSkip []int
		for index, value := range levels {
			if index != skip {
				levelsWithSkip = append(levelsWithSkip, value)
			}
		}
		if part1(levelsWithSkip) {
			return true
		}
	}

	return false
}

func main() {
	part := os.Args[1]
	safe := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var levels []int
		for _, level := range strings.Fields(scanner.Text()) {
			value, err := strconv.Atoi(level)
			if err != nil {
				fmt.Println(err)
			}
			levels = append(levels, value)
		}

		if part == "1" && part1(levels) {
			safe += 1
		} else if part2(levels) {
			safe += 1
		}

	}
	fmt.Println(safe)
}
