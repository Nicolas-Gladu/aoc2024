package main

import (
	"Majha/AdvantOfCode/pkg"
	"fmt"
)

const (
	KEYWORD = "XMAS"
)

type wordFinder struct {
	data   []string
	total  int
	height int
	width  int
}

func newWordFinder(data []string) wordFinder {
	return wordFinder{
		data:   data,
		total:  0,
		height: len(data),
		width:  len(data[0]),
	}
}

func main() {
	challenge, err := pkg.NewDayChallenge("input.txt")
	if err != nil {
		panic(err)
	}
	result := solutionPart1(challenge.Inputs)
	fmt.Println("result1:", result)
	result = solutionPart2(challenge.Inputs)
	fmt.Println("result2:", result)
}

func solutionPart1(data []string) int {
	finder := newWordFinder(data)
	for y, row := range data {
		for x, char := range row {
			if char != rune(KEYWORD[0]) {
				continue
			}
			finder.forward(x, y)
			finder.backward(x, y)
			finder.upward(x, y)
			finder.downward(x, y)
			finder.upright(x, y)
			finder.downright(x, y)
			finder.upleft(x, y)
			finder.downleft(x, y)
		}
	}
	return finder.total
}

func solutionPart2(data []string) int {
	var total int
	for y := 1; y < len(data)-1; y++ {
		for x := 1; x < len(data[y])-1; x++ {
			if data[y][x] != 'A' {
				continue
			}

			if ((data[y-1][x-1] == 'M' && data[y+1][x+1] == 'S') ||
				(data[y-1][x-1] == 'S' && data[y+1][x+1] == 'M')) &&
				((data[y+1][x-1] == 'M' && data[y-1][x+1] == 'S') ||
					(data[y+1][x-1] == 'S' && data[y-1][x+1] == 'M')) {
				total++
			}
		}
	}
	return total
}

func (w *wordFinder) forward(x, y int) {
	if w.rightCheck(x) {
		return
	}

	if w.data[y][x:x+len(KEYWORD)] == KEYWORD {
		w.total++
	}
}

func (w *wordFinder) backward(x, y int) {
	if w.leftCheck(x) {
		return
	}

	for i := len(KEYWORD) - 1; i > 0; i-- {
		if w.data[y][x-i] != KEYWORD[i] {
			return
		}
	}
	w.total++
}

func (w *wordFinder) upward(x, y int) {
	if w.upCheck(y) {
		return
	}

	for i := len(KEYWORD) - 1; i > 0; i-- {
		if w.data[y-i][x] != KEYWORD[i] {
			return
		}
	}
	w.total++
}

func (w *wordFinder) downward(x, y int) {
	if w.downCheck(y) {
		return
	}

	for i := 1; i < len(KEYWORD); i++ {
		if w.data[y+i][x] != KEYWORD[i] {
			return
		}
	}
	w.total++
}

func (w *wordFinder) upright(x, y int) {
	if w.upCheck(y) || w.rightCheck(x) {
		return
	}

	for i := 1; i < len(KEYWORD); i++ {
		if w.data[y-i][x+i] != KEYWORD[i] {
			return
		}
	}
	w.total++
}

func (w *wordFinder) downright(x, y int) {
	if w.downCheck(y) || w.rightCheck(x) {
		return
	}

	for i := 1; i < len(KEYWORD); i++ {
		if w.data[y+i][x+i] != KEYWORD[i] {
			return
		}
	}
	w.total++
}

func (w *wordFinder) upleft(x, y int) {
	if w.upCheck(y) || w.leftCheck(x) {
		return
	}

	for i := 1; i < len(KEYWORD); i++ {
		if w.data[y-i][x-i] != KEYWORD[i] {
			return
		}
	}
	w.total++
}

func (w *wordFinder) downleft(x, y int) {
	if w.downCheck(y) || w.leftCheck(x) {
		return
	}

	for i := 1; i < len(KEYWORD); i++ {
		if w.data[y+i][x-i] != KEYWORD[i] {
			return
		}
	}
	w.total++
}

func (w *wordFinder) upCheck(y int) bool {
	return y-(len(KEYWORD)-1) < 0
}

func (w *wordFinder) downCheck(y int) bool {
	return y > w.height-len(KEYWORD)
}

func (w *wordFinder) leftCheck(x int) bool {
	return x-(len(KEYWORD)-1) < 0
}

func (w *wordFinder) rightCheck(x int) bool {
	return x > w.width-len(KEYWORD)
}
