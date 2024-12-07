package main

import (
	"Majha/AdvantOfCode/pkg"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	challenge, err := pkg.NewDayChallenge("input.txt")
	if err != nil {
		panic(err)
	}

	total, totalPt2 := solution(challenge)
	fmt.Println("part1:", total)
	fmt.Println("part2:", totalPt2)
}

func solution(challenge *pkg.DayChallenge) (int, int) {
	var list1, list2 = make([]int, len(challenge.Inputs)), make([]int, len(challenge.Inputs))
	for idx, input := range challenge.Inputs {
		left, right := splitList(input)
		list1[idx] = left
		list2[idx] = right
	}
	slices.Sort(list1)
	slices.Sort(list2)
	var total int
	var rep int
	for i, e1 := range list1 {
		total += abs(list1[i] - list2[i])
		occurance := 0
		for _, e2 := range list2 {
			if e1 == e2 {
				occurance++
			}
		}
		rep += e1 * occurance
	}
	return total, rep
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func splitList(in string) (int, int) {
	nums := strings.Split(in, "   ")
	if len(nums) != 2 {
		panic(fmt.Sprintf("need 2 row. Have %d", len(nums)))
	}
	left, err := strconv.Atoi(nums[0])
	if err != nil {
		panic(err)
	}
	right, err := strconv.Atoi(nums[1])
	if err != nil {
		panic(err)
	}
	return left, right
}
