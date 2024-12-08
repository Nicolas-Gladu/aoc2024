package main

import (
	"Majha/AdvantOfCode/pkg"
	"fmt"
	"strconv"
	"strings"
)

const (
	MUL_KEYWORD   = "mul("
	START_KEYWORD = "do()"
	END_KEYWORD   = "don't()"
)

func main() {
	challenge, err := pkg.NewDayChallenge("input.txt")
	if err != nil {
		panic(err)
	}

	part1 := solution2(challenge.Inputs)
	fmt.Println("part1:", part1)
	part2 := solutionPt2(challenge.Inputs)
	fmt.Println("part2:", part2)
}

func solution(inputs []string) (int, int) {
	var part1 int
	var part2 int
	data := strings.Join(inputs, "")
	for {
		var mulResult int
		data, mulResult = findToken(data)
		part1 += mulResult
		if data == "" {
			break
		}
	}
	return part1, part2
}

func solution2(inputs []string) int {
	var part1 int
	data := strings.Join(inputs, "")
	for {
		var mulResult int
		data, mulResult = findToken2(data)
		part1 += mulResult
		if data == "" {
			break
		}
	}

	return part1
}

func solutionPt2(inputs []string) int {
	var total int
	data := strings.Join(inputs, "")

	var startIdx int = 0
	var endIdx int = 0
	for {
		endIdx = strings.Index(data, END_KEYWORD)
		if endIdx < 0 {
			endIdx = len(data) - 1
		}
		scope := data[:endIdx]
		var mulResult int
		for {
			scope, mulResult = findToken2(scope)
			total += mulResult
			if scope == "" {
				break
			}
		}
		startIdx = strings.Index(data[endIdx:], START_KEYWORD)
		if startIdx < 0 {
			break
		}
		startIdx += len(START_KEYWORD) + endIdx
		if startIdx > len(data) {
			break
		}
		data = data[startIdx:]
	}
	return total
}

func findToken(data string) (string, int) {
	idx := strings.Index(data, MUL_KEYWORD)
	if idx < 0 {
		return "", 0
	}
	data = data[idx+len(MUL_KEYWORD):]
	endIdx := strings.Index(data, ")")
	if endIdx < 0 {
		return "", 0
	}

	// if we find another keyword inside the closing braket it's mean it's a bad token
	badTokenIdx := strings.Index(data[:endIdx], MUL_KEYWORD)
	if badTokenIdx >= 0 {
		return data[badTokenIdx:], 0
	}

	nums := strings.Split(data[:endIdx], ",")
	if len(nums) != 2 {
		return data[endIdx:], 0
	}
	total := 1
	for _, numstr := range nums {
		num, err := strconv.Atoi(numstr)
		if err != nil {
			return data[endIdx:], 0
		}
		total *= num
	}

	return data[endIdx:], total
}

func findToken2(data string) (string, int) {
	idx := strings.Index(data, MUL_KEYWORD)
	if idx < 0 {
		return "", 0
	}
	data = data[idx+len(MUL_KEYWORD):]
	endIdx := strings.Index(data, ")")
	if endIdx < 0 {
		return "", 0
	}

	// if we find another keyword inside the closing braket it's mean it's a bad token
	badTokenIdx := strings.Index(data[:endIdx], MUL_KEYWORD)
	if badTokenIdx >= 0 {
		return data[badTokenIdx:], 0
	}

	commaIdx := strings.Index(data[:endIdx], ",")
	if commaIdx < 0 {
		return data[endIdx:], 0
	}

	num1, str := atoi(data[:commaIdx])
	if str != "" {
		return data[endIdx:], 0
	}
	num2, str := atoi(data[commaIdx+1 : endIdx])
	if str != "" {
		return data[endIdx:], 0
	}

	return data[endIdx:], num1 * num2
}

func atoi(input string) (int, string) {
	var n int
	for i, b := range []byte(input) {
		b -= '0'
		if b > 9 {
			return n, input[i:]
		}
		n = n*10 + int(b)
	}
	return n, ""
}
