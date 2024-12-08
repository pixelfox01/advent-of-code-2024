package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part1() int {
	before := make(map[string][]string)
	after := make(map[string][]string)

	var safeUpdates []string

	f, e := os.Open("part1.txt")
	check(e)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		a, b := strings.Split(line, "|")[0], strings.Split(line, "|")[1]

		after[a] = append(after[a], b)
		before[b] = append(before[b], a)
	}

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, ",")
		safe := true
		for i, v := range nums {
			bp, ap := i-1, i+1
			for bp >= 0 {
				if slices.Contains(after[v], nums[bp]) {
					safe = false
				}
				bp--
			}
			for ap < len(nums) {
				if slices.Contains(before[v], nums[ap]) {
					safe = false
				}
				ap++
			}
		}
		if safe {
			safeUpdates = append(safeUpdates, line)
		}
	}

	var midValues []int

	for _, u := range safeUpdates {
		nums := strings.Split(u, ",")
		mid, e := strconv.Atoi(nums[len(nums)/2])
		check(e)
		midValues = append(midValues, mid)
	}

	sum := 0
	for _, m := range midValues {
		sum += m
	}

	return sum
}

func main() {
	part1 := part1()
	fmt.Println("Part 1 Ans:", part1)
}
