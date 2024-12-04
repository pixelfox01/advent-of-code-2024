package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part1() {
	var sum int
	f, e := os.Open("part1.txt")
	check(e)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		r := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
		mulStrings := r.FindAllString(line, -1)

		r = regexp.MustCompile(`[0-9]{1,3}`)
		for _, v := range mulStrings {
			numStrings := r.FindAllString(v, -1)
			num1, err := strconv.Atoi(numStrings[0])
			check(err)
			num2, err := strconv.Atoi(numStrings[1])
			check(err)
			sum += num1 * num2
		}
	}

	fmt.Println(sum)
}

func part2() {
	var sum int
	do := true
	f, e := os.Open("part2.txt")
	check(e)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		r := regexp.MustCompile(`(don't\(\)|do\(\)|mul\([0-9]{1,3},[0-9]{1,3}\))`)
		matches := r.FindAllString(line, -1)

		// fmt.Println(matches)

		for _, v := range matches {
			if v == "don't()" {
				do = false
			} else if v == "do()" {
				do = true
			} else if do {
				r = regexp.MustCompile(`[0-9]{1,3}`)
				numStrings := r.FindAllString(v, -1)
				num1, err := strconv.Atoi(numStrings[0])
				check(err)
				num2, err := strconv.Atoi(numStrings[1])
				check(err)
				sum += num1 * num2
				// fmt.Println(num1, num2, "\tsum:", sum)
			}
		}
	}
	fmt.Println(sum)
}

func main() {
	part1()
	part2()
}
