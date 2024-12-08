package main

import (
	"bufio"
	"fmt"
	"os"
)

type coord struct {
	x int
	y int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, e := os.Open("part1.txt")
	check(e)

	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	Ans1 := part1(lines)
	Ans2 := part2(lines)
	fmt.Println("Part1:", Ans1)
	fmt.Println("Part2:", Ans2)
}

func part2(lines []string) int {
	answer := 0
	for i := range lines {
		for j := range lines[i] {
			if string(lines[i][j]) == "A" {
				if i-1 >= 0 && i+1 < len(lines) && j-1 >= 0 && j+1 < len(lines[i]) {
					diag1 := string(fmt.Sprintf("%s%s%s", string(lines[i-1][j-1]), string(lines[i][j]), string(lines[i+1][j+1])))
					fmt.Println(diag1)
					diag2 := string(fmt.Sprintf("%s%s%s", string(lines[i-1][j+1]), string(lines[i][j]), string(lines[i+1][j-1])))
					if (diag1 == "MAS" || diag1 == "SAM") && (diag2 == "MAS" || diag2 == "SAM") {
						answer++
					}
				}
			}
		}
	}
	return answer
}

func part1(lines []string) int {
	var dirs []coord

	for _, i := range []int{-1, 0, 1} {
		for _, j := range []int{-1, 0, 1} {
			if i == 0 && i == j {
				continue
			}
			dir := coord{x: i, y: j}
			dirs = append(dirs, dir)
		}
	}

	ans := 0
	word := "XMAS"
	for i := range lines {
		for j := range lines[i] {
			for _, dir := range dirs {
				count := 0
				for k := range 4 {
					if i+dir.x*k >= 0 && i+dir.x*k < len(lines) && j+dir.y*k >= 0 && j+dir.y*k < len(lines[i]) {
						if lines[i+dir.x*k][j+dir.y*k] == word[k] {
							count++
						}
					}
				}
				if count == 4 {
					ans++
				}
			}
		}
	}

	return ans
}
