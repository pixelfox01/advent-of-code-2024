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

	var correctUpdates []string

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
					break
				}
				bp--
			}
			for ap < len(nums) {
				if slices.Contains(before[v], nums[ap]) {
					safe = false
					break
				}
				ap++
			}
			if !safe {
				break
			}
		}
		if safe {
			correctUpdates = append(correctUpdates, line)
		}
	}

	var midValues []int

	for _, u := range correctUpdates {
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

func part2() int {
	before := make(map[string][]string)
	after := make(map[string][]string)

	var incorrectUpdates []string

	f, e := os.Open("part2.txt")
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
					break
				}
				bp--
			}
			for ap < len(nums) {
				if slices.Contains(before[v], nums[ap]) {
					safe = false
					break
				}
				ap++
			}
			if !safe {
				break
			}
		}
		if !safe {
			incorrectUpdates = append(incorrectUpdates, line)
		}
	}

	var midValues []int

	for _, u := range incorrectUpdates {
		var indices []int
		nums := strings.Split(u, ",")
		for i, n := range nums {
			curIdx := 0
			for j := range nums {
				if j == i {
					continue
				}
				if slices.Contains(before[n], nums[j]) {
					curIdx++
				}
			}
			indices = append(indices, curIdx)
		}
		midIdx := slices.Index(indices, len(nums)/2)
		midNum, e := strconv.Atoi(nums[midIdx])
		check(e)
		midValues = append(midValues, midNum)
	}

	sum := 0
	for _, m := range midValues {
		sum += m
	}

	return sum
}

func main() {
	part1 := part1()
	part2 := part2()
	fmt.Println("Part 2 Answer:", part1)
	fmt.Println("Part 2 Answer:", part2)
}
