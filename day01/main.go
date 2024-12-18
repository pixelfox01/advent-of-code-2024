package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part1(col1, col2 []int) {
	var diffs []int

	sort.Ints(col1)
	sort.Ints(col2)

	for i := 0; i < len(col1); i++ {
		var diff int
		if col1[i] > col2[i] {
			diff = col1[i] - col2[i]
		} else {
			diff = col2[i] - col1[i]
		}
		diffs = append(diffs, diff)
	}

	sum := 0
	for _, d := range diffs {
		sum += d
	}

	fmt.Println("Part 1 answer:", sum)
}

func part2(col1, col2 []int) {
	counts := make(map[int]int)
	var similarities []int

	for _, num := range col2 {
		counts[num] = counts[num] + 1
	}

	for _, num := range col1 {
		similarities = append(similarities, num*counts[num])
	}

	sum := 0
	for _, s := range similarities {
		sum += s
	}

	fmt.Println("Part 2 answer:", sum)
}

func main() {
	f, err := os.Open("part1.txt")
	check(err)

	defer f.Close()

	var col1, col2 []int

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if len(fields) >= 2 {
			num1, err := strconv.Atoi(fields[0])
			check(err)
			num2, err := strconv.Atoi(fields[1])
			check(err)

			col1 = append(col1, num1)
			col2 = append(col2, num2)
		}
	}

	part1(col1, col2)
	part2(col1, col2)
}
