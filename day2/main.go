package main

import (
	"Majha/AdvantOfCode/pkg"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	challenge, err := pkg.NewDayChallenge("input.txt")
	if err != nil {
		panic(err)
	}
	now := time.Now()
	part1, part2 := solution(challenge.Inputs) //solution(challenge.Inputs)
	fmt.Println("time: ", time.Since(now))
	fmt.Println("part1: ", part1)
	fmt.Println("part2: ", part2)
}

func solution(reports []string) (int, int) {
	var part1 int
	var part2 int
	for _, report := range reports {
		values := convertToInts(strings.Split(report, " "))
		errors := findError(values)
		if len(errors) == 0 {
			part1++
			part2++
		} else {
			for i := range values {
				temp := slices.Clone(values)
				temp = append(temp[:i], temp[i+1:]...)
				errors := findError(temp)
				if len(errors) == 0 {
					part2++
					break
				}
			}

		}
	}
	return part1, part2
}

func convertToInts(values []string) []int {
	out := make([]int, len(values))
	for i, value := range values {
		nb, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		out[i] = nb
	}
	return out
}

func findError(values []int) []int {
	var increasing *bool
	var errors []int
	for k := 1; k < len(values); k++ {
		var delta int
		elem1 := values[k-1]
		elem2 := values[k]

		if increasing == nil {
			inc := elem1 < elem2
			increasing = &inc
		}
		if *increasing {
			delta = elem2 - elem1
		} else {
			delta = elem1 - elem2
		}
		if delta <= 0 || delta > 3 {
			errors = append(errors, k)
		}
	}
	return errors
}
