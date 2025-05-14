package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

func checkStr(line, str string, idx int) bool {
	if idx+len(str)-1 < len(line) {
		for i := range len(str) {
			if line[idx+i] != str[i] {
				return false
			}
		}
		return true
	}
	return false
}

func parseLine(line string, part string) int {
	var sum int
	do := true
	for idx := range len(line) {
		if checkStr(line, "do()", idx) {
			do = true
		} else if checkStr(line, "don't()", idx) {
			do = false
		} else if checkStr(line, "mul(", idx) {
			left := 0
			curr := idx+4
			for ; curr < len(line); curr++ {
				if 47 <= line[curr] && line[curr] <= 57 {
					value, _ := strconv.Atoi(string(line[curr])) 
					left = 10*left + value
				} else {
					break
				}
			}

			if curr < len(line) && line[curr] == ',' {
				curr++
			} else {
				continue
			}

			right := 0
			for ; curr < len(line); curr++ {
				if 47 <= line[curr] && line[curr] <= 57 {
					value, _ := strconv.Atoi(string(line[curr])) 
					right = 10*right + value
				} else {
					break
				}
			}

			if !(curr < len(line) && line[curr] == ')') {
				continue
			}

			if(do || part == "1") {
				sum = sum + (left*right)
			}
		}
	}
	return sum
}

func main() {
	part := os.Args[1]

	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text());
	}

	fmt.Println(parseLine(strings.Join(lines,""), part));
}
