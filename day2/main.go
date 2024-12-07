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
	part1, part2 := solution2(challenge.Inputs) //solution(challenge.Inputs)
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

func solution2(reports []string) (int, int) {
	var part1 int
	var part2 int
	for _, report := range reports {
		levels := convertToInts(strings.Split(report, " "))
		unsafeAt := unsafeIdx(levels)
		if unsafeAt == -1 {
			part1++
			part2++
			continue
		}
		replacement1, replacement2 := unsafeAt-1, unsafeAt
		change1Safe := unsafeIdx(deleteLevelAt(replacement1, levels)) == -1
		change2Safe := unsafeIdx(deleteLevelAt(replacement2, levels)) == -1
		if change1Safe || change2Safe {
			part2++
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

// Returns the index of level that made record unsafe
// returns -1 if safe (no unsafe idx found)
func unsafeIdx(levels []int) int {
	if len(levels) <= 1 {
		return -1
	}
	increasing := levels[1] > levels[0]
	for i := 1; i < len(levels); i++ {
		diff := abs(levels[i] - levels[i-1])
		isSequential := (increasing && levels[i] > levels[i-1]) ||
			(!increasing && levels[i] < levels[i-1])
		validDiff := 1 <= diff && diff <= 3
		if !isSequential || !validDiff {
			return i
		}
	}
	return -1
}

func deleteLevelAt(idx int, levels []int) []int {
	deleted := make([]int, len(levels)-1)
	copy(deleted[:idx], levels[:idx])
	copy(deleted[idx:], levels[idx+1:])
	return deleted
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
