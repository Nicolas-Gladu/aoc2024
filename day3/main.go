package main

import (
	"Majha/AdvantOfCode/pkg"
	"fmt"
	"strconv"
	"strings"
)

const (
	MulKeyword   = "mul("
	StartKeyword = "do()"
	EndKeyword   = "don't()"
)

func main() {
	challenge, err := pkg.NewDayChallenge("day3/input.txt")
	if err != nil {
		panic(err)
	}

	part1 := solution4(strings.Join(challenge.Inputs, ""))
	fmt.Println("part1:", part1)
	part2 := solutionPt2(strings.Join(challenge.Inputs, ""))
	fmt.Println("part2:", part2)
}

func solution(data string) (int, int) {
	var part1 int
	var part2 int
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

func solution2(data string) int {
	var part1 int
	var mulResult int
	for {
		data, mulResult = findToken2(data)
		part1 += mulResult
		if data == "" {
			break
		}
	}

	return part1
}

func solution3(data string) int {
	var part1 int
	var mulResult int
	for {
		data, mulResult = findToken3(data)
		part1 += mulResult
		if data == "" {
			break
		}
	}

	return part1
}

func solution4(data string) int {
	total := 0
	n := len(data)
	i := 0

	for i < n-4 {
		// Check for literal "mul("
		if data[i] != 'm' || data[i+1] != 'u' || data[i+2] != 'l' || data[i+3] != '(' {
			i++
			continue
		}
		j := i + 4 // Start of the number
		start := j

		// Read first number (X)
		for j < n && data[j] >= '0' && data[j] <= '9' {
			j++
		}
		if j == start { // No digits
			i += 4
			continue
		}
		num1 := parseIntRange(data, start, j)

		// Expect comma
		if j >= n || data[j] != ',' {
			i += 4
			continue
		}
		j++
		start = j

		// Read second number (Y)
		for j < n && data[j] >= '0' && data[j] <= '9' {
			j++
		}
		if j == start { // No digits
			i += 4
			continue
		}

		num2 := parseIntRange(data, start, j)

		// Expect closing parenthesis
		if j >= n || data[j] != ')' {
			i += 4
			continue
		}

		total += num1 * num2

		i = j + 1 // Move past the closing ')'
	}

	return total
}

func solutionPt2(data string) int {
	var total int
	var startIdx = 0
	var endIdx = 0
	for {
		endIdx = strings.Index(data, EndKeyword)
		if endIdx < 0 {
			endIdx = len(data) - 1
		}
		scope := data[:endIdx]
		var mulResult int
		for {
			scope, mulResult = findToken3(scope)
			total += mulResult
			if scope == "" {
				break
			}
		}
		startIdx = strings.Index(data[endIdx:], StartKeyword)
		if startIdx < 0 {
			break
		}
		startIdx += len(StartKeyword) + endIdx
		if startIdx > len(data) {
			break
		}
		data = data[startIdx:]
	}
	return total
}

func findToken(data string) (string, int) {
	idx := strings.Index(data, MulKeyword)
	if idx < 0 {
		return "", 0
	}
	data = data[idx+len(MulKeyword):]
	endIdx := strings.Index(data, ")")
	if endIdx < 0 {
		return "", 0
	}

	// if we find another keyword inside the closing braket it's mean it's a bad token
	badTokenIdx := strings.Index(data[:endIdx], MulKeyword)
	if badTokenIdx >= 0 {
		return data[badTokenIdx:], 0
	}

	nums := strings.Split(data[:endIdx], ",")
	if len(nums) != 2 {
		return data[endIdx:], 0
	}
	total := 1
	for _, numStr := range nums {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return data[endIdx:], 0
		}
		total *= num
	}

	return data[endIdx:], total
}

func findToken2(data string) (string, int) {
	idx := strings.Index(data, MulKeyword)
	if idx < 0 {
		return "", 0
	}
	data = data[idx+len(MulKeyword):]
	endIdx := strings.Index(data, ")")
	if endIdx < 0 {
		return "", 0
	}

	// if we find another keyword inside the closing braket it's mean it's a bad token
	badTokenIdx := strings.Index(data[:endIdx], MulKeyword)
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

func findToken3(data string) (string, int) {
	idx := strings.Index(data, MulKeyword)
	if idx < 0 {
		return "", 0
	}
	start := idx + len(MulKeyword)
	data = data[start:]

	endIdx := strings.IndexByte(data, ')')
	if endIdx < 0 {
		return "", 0
	}
	args := data[:endIdx]

	// Check for nested MulKeyword (bad token)
	if strings.Contains(args, MulKeyword) {
		return data[strings.Index(args, MulKeyword):], 0
	}

	// Find comma in args
	commaIdx := strings.IndexByte(args, ',')
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
func parseIntRange(data string, start, end int) int {
	n := 0
	for i := start; i < end; i++ {
		n = n*10 + int(data[i]-'0')
	}
	return n
}
