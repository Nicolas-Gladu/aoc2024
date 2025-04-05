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

	part1 := solutionPart1(challenge.Inputs)
	assert.Equal(t, 18, part1)
	part2 := solutionPart2(challenge.Inputs)
	assert.Equal(t, 9, part2)
}

func BenchmarkSolution1(b *testing.B) {
	challenge, err := pkg.NewDayChallenge("input.txt")
	if err != nil {
		panic(err)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solutionPart1(challenge.Inputs)
	}
}

func BenchmarkSolution2(b *testing.B) {
	challenge, err := pkg.NewDayChallenge("input.txt")
	if err != nil {
		panic(err)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solutionPart2(challenge.Inputs)
	}
}
