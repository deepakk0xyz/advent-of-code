package main

import (
	"fmt"
	"os"
	"bufio"
)

func part1(lines []string) int {
	rows := len(lines)
	columns := len(lines[0])

	directions := [][]int{{0,1},{0,-1},{1,0},{-1,0},{1,1},{1,-1},{-1,1},{-1,-1}};
	xmas := "XMAS"
	count := 0

	for row := range rows {
		for column := range columns {
			for _, direction := range directions {
				matched := true
				for idx := range 4 {
					curr := []int{ row + idx * direction[0], column + idx * direction[1] };
					if !(0 <= curr[0] && curr[0] < rows && 0 <= curr[1] && curr[1] < columns && lines[curr[0]][curr[1]] == xmas[idx]) {
						matched = false
					}
				}

				if(matched) {
					count++
				}
			}
		}
	}

	return count
}

func part2(lines []string) int {
	rows := len(lines)
	columns := len(lines[0])

	count := 0

	for row := range rows {
		for column := range columns {
			if row == 0 || row == rows-1 || column == 0  || column == columns-1 || lines[row][column] != 'A' {
				continue
			}
			left := lines[row-1][column-1] == 'M' && lines[row+1][column-1] == 'M' && lines[row-1][column+1] == 'S' && lines[row+1][column+1] == 'S'
			right := lines[row-1][column-1] == 'S' && lines[row+1][column-1] == 'S' && lines[row-1][column+1] == 'M' && lines[row+1][column+1] == 'M'
			top := lines[row-1][column-1] == 'M' && lines[row-1][column+1] == 'M' && lines[row+1][column-1] == 'S' && lines[row+1][column+1] == 'S'
			bottom := lines[row-1][column-1] == 'S' && lines[row-1][column+1] == 'S' && lines[row+1][column-1] == 'M' && lines[row+1][column+1] == 'M'

			if top || bottom || left || right {
				count++
			}
		}
	}

	return count
}


func main() {
	part := os.Args[1]

	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text());
	}

	if part == "1" {
		fmt.Println(part1(lines));
	} else {
		fmt.Println(part2(lines));
	}
}
