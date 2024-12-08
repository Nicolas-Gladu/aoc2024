package main

import (
	"Majha/AdvantOfCode/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	challenge, err := pkg.NewDayChallenge("test.txt")
	if err != nil {
		panic(err)
	}

	part1 := solution2(challenge.Inputs)
	assert.Equal(t, 161, part1)
	part2 := solutionPt2(challenge.Inputs)
	assert.Equal(t, 48, part2)

}

func BenchmarkSolution(b *testing.B) {
	// Generate test data
	challenge, err := pkg.NewDayChallenge("input.txt")
	if err != nil {
		panic(err)
	}
	b.ReportAllocs()
	b.ResetTimer() // Reset the timer to exclude setup time
	for i := 0; i < b.N; i++ {
		_, _ = solution(challenge.Inputs)
	}
}

func BenchmarkSolution2(b *testing.B) {
	// Generate test data
	challenge, err := pkg.NewDayChallenge("input.txt")
	if err != nil {
		panic(err)
	}
	b.ReportAllocs()
	b.ResetTimer() // Reset the timer to exclude setup time
	for i := 0; i < b.N; i++ {
		_ = solution2(challenge.Inputs)
	}
}

func BenchmarkSolutionPt2(b *testing.B) {
	// Generate test data
	challenge, err := pkg.NewDayChallenge("input.txt")
	if err != nil {
		panic(err)
	}
	b.ReportAllocs()
	b.ResetTimer() // Reset the timer to exclude setup time
	for i := 0; i < b.N; i++ {
		_ = solutionPt2(challenge.Inputs)
	}
}
