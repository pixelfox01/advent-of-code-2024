package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isSafe(levels []int) bool {
	var prevDiff int
	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]
		if math.Abs(float64(diff)) < 1 || math.Abs(float64(diff)) > 3 {
			return false
		}
		if i == 1 {
			prevDiff = diff
		} else if math.Signbit(float64(prevDiff)) != math.Signbit(float64(diff)) {
			return false
		}
		prevDiff = diff
	}
	return true
}

func part2() {
	f, err := os.Open("part2.txt")
	check(err)
	var report string
	var levels []int
	var safeReports int

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		report = scanner.Text()
		fields := strings.Fields(report)
		for _, f := range fields {
			num, err := strconv.Atoi(f)
			check(err)
			levels = append(levels, num)
		}

		if safe := isSafe(levels); safe {
			safeReports++
		} else {
			for i := range levels {
				temp := append([]int(nil), levels[:i]...)
				temp = append(temp, levels[i+1:]...)
				safe = isSafe(temp)
				if safe {
					safeReports++
					break
				}
			}
		}
		levels = nil
	}

	fmt.Println(safeReports)
}

func part1() {
	f, err := os.Open("part1.txt")
	check(err)
	var report string
	var levels []int
	var safeReports int

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		report = scanner.Text()
		fields := strings.Fields(report)
		for _, f := range fields {
			num, err := strconv.Atoi(f)
			check(err)
			levels = append(levels, num)
		}

		if safe := isSafe(levels); safe {
			safeReports++
		}

		levels = nil
	}

	fmt.Println(safeReports)
}

func main() {
	part1()
	part2()
}
