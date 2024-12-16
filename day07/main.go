package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func permutations(n int, operands []string) []string {
	var result []string
	var backtrack func(cur string)
	backtrack = func(cur string) {
		if len(cur) == n {
			result = append(result, cur)
			return
		}
		// backtrack(fmt.Sprintf("%s%s", cur, "+"))
		// backtrack(fmt.Sprintf("%s%s", cur, "*"))
		for _, op := range operands {
			backtrack(fmt.Sprintf("%s%s", cur, op))
		}
	}
	backtrack("")
	return result
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part1(eqs [][]int) int {
	sum := 0

	for _, e := range eqs {
		ans := e[0]
		operands := e[1:]
		perms := permutations(len(operands)-1, []string{"+", "*"})
		for _, p := range perms {
			res := operands[0]
			for i, o := range p {
				if o == '+' {
					res += operands[i+1]
				} else {
					res *= operands[i+1]
				}
			}
			if res == ans {
				sum += ans
				break
			}
		}
	}

	return sum
}

func part2(eqs [][]int) int {
	sum := 0

	for _, e := range eqs {
		ans := e[0]
		operands := e[1:]
		perms := permutations(len(operands)-1, []string{"+", "*", "|"})
		for _, p := range perms {
			res := operands[0]
			for i, o := range p {
				if o == '+' {
					res += operands[i+1]
				} else if o == '*' {
					res *= operands[i+1]
				} else {
					temp := fmt.Sprintf("%d%d", res, operands[i+1])
					tempRes, e := strconv.Atoi(temp)
					check(e)
					res = tempRes
				}
			}
			if res == ans {
				sum += ans
				break
			}
		}
	}

	return sum
}

func main() {
	f, e := os.Open("input.txt")
	check(e)

	s := bufio.NewScanner(f)

	var eqs [][]int

	for s.Scan() {
		line := s.Text()
		var curEq []int
		temp, e := strconv.Atoi(strings.Split(line, ":")[0])
		check(e)
		curEq = append(curEq, temp)

		fields := strings.Fields(strings.Split(line, ":")[1])
		for _, f := range fields {
			temp, e := strconv.Atoi(string(f))
			check(e)
			curEq = append(curEq, temp)
		}
		eqs = append(eqs, curEq)
	}

	part1 := part1(eqs)
	fmt.Println("Part1 answer:", part1)
	part2 := part2(eqs)
	fmt.Println("Part2 answer:", part2)
}
