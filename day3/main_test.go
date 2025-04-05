package main

import (
	"Majha/AdvantOfCode/pkg"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	challenge, err := pkg.NewDayChallenge("test.txt")
	if err != nil {
		panic(err)
	}

	part1 := solution2(strings.Join(challenge.Inputs, ""))
	assert.Equal(t, 161, part1)
	part2 := solutionPt2(strings.Join(challenge.Inputs, ""))
	assert.Equal(t, 48, part2)

}

func BenchmarkSolution(b *testing.B) {
	// Generate test data
	challenge, err := pkg.NewDayChallenge("input.txt")
	if err != nil {
		panic(err)
	}
	data := strings.Join(challenge.Inputs, "")
	b.ReportAllocs()
	b.ResetTimer() // Reset the timer to exclude setup time
	for i := 0; i < b.N; i++ {
		_, _ = solution(data)
	}
}

func BenchmarkSolution2(b *testing.B) {
	// Generate test data
	challenge, err := pkg.NewDayChallenge("input.txt")
	if err != nil {
		panic(err)
	}
	data := strings.Join(challenge.Inputs, "")
	b.ReportAllocs()
	b.ResetTimer() // Reset the timer to exclude setup time
	for i := 0; i < b.N; i++ {
		_ = solution2(data)
	}
}

func BenchmarkSolution3(b *testing.B) {
	// Generate test data
	challenge, err := pkg.NewDayChallenge("input.txt")
	if err != nil {
		panic(err)
	}
	data := strings.Join(challenge.Inputs, "")
	b.ReportAllocs()
	b.ResetTimer() // Reset the timer to exclude setup time
	for i := 0; i < b.N; i++ {
		_ = solution3(data)
	}
}

func BenchmarkSolution4(b *testing.B) {
	// Generate test data
	challenge, err := pkg.NewDayChallenge("input.txt")
	if err != nil {
		panic(err)
	}
	data := strings.Join(challenge.Inputs, "")
	b.ReportAllocs()
	b.ResetTimer() // Reset the timer to exclude setup time
	for i := 0; i < b.N; i++ {
		_ = solution4(data)
	}
}

func BenchmarkSolutionPt2(b *testing.B) {
	// Generate test data
	challenge, err := pkg.NewDayChallenge("input.txt")
	if err != nil {
		panic(err)
	}
	data := strings.Join(challenge.Inputs, "")
	b.ReportAllocs()
	b.ResetTimer() // Reset the timer to exclude setup time
	for i := 0; i < b.N; i++ {
		_ = solutionPt2(data)
	}
}
