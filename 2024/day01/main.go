package main

import (
	"fmt"
	"os"
	"sort"
)

func part1(left, right []int) int {
	ans := 0

	sort.Ints(left)
	sort.Ints(right)

	for idx, _ := range left {
		if left[idx] > right[idx] {
			ans += left[idx] - right[idx]
		} else {
			ans += right[idx] - left[idx]
		}
	}

	return ans
}

func part2(left, right []int) int {
	ans := 0
	counter := make(map[int]int)

	for _, value := range right {
		counter[value]++
	}

	for _, value := range left {
		ans += value * counter[value]
	}

	return ans
}

func main() {
	part := os.Args[1]
	var left []int
	var right []int

	for {
		var l, r int
		_, err := fmt.Scanf("%d %d", &l, &r)

		if err != nil {
			break
		}

		left = append(left, l)
		right = append(right, r)
	}

	if part == "1" {
		fmt.Println(part1(left, right))
	} else {
		fmt.Println(part2(left, right))
	}
}
