package main

import (
	"fmt"

	"github.com/Schidstorm/aoc/pkg/util"
)

func main() {
	lines := util.LoadDayFile(4).Lines()

	var board [][]rune
	for _, line := range lines {
		board = append(board, []rune(line))
	}

	var sum int64
	var lastRemoved int64 = 1
	for lastRemoved != 0 {
		board, lastRemoved = removeRolls(board)
		sum += lastRemoved
	}

	fmt.Println(sum)
}

func removeRolls(board [][]rune) ([][]rune, int64) {
	boardCopy := copyBoard(board)
	height := len(board)
	width := len(board[0])

	var sum int64
	for x := range width {
		for y := range height {
			if board[y][x] == '@' && neighbours(xy{x, y}, board) < 4 {
				sum++
				boardCopy[y][x] = '.'
			}
		}
	}

	return boardCopy, sum
}

func copyBoard(src [][]rune) [][]rune {
	dest := make([][]rune, len(src))
	for i := range src {
		dest[i] = make([]rune, len(src[i]))
		copy(dest[i], src[i])
	}

	return dest
}

func neighbours(pos xy, i [][]rune) int64 {
	height := len(i)
	width := len(i[0])

	positions := []xy{
		{-1, -1},
		{0, -1},
		{1, -1},
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
	}

	var sum int64
	for _, p := range positions {
		np := add(p, pos)
		if isIn(np, width, height) {
			if i[np.Y][np.X] == '@' {
				sum++
			}
		}
	}

	return sum
}

type xy struct {
	X, Y int
}

func add(a, b xy) xy {
	return xy{
		a.X + b.X,
		a.Y + b.Y,
	}
}

func isIn(p xy, width, height int) bool {
	return p.X >= 0 && p.X < width && p.Y >= 0 && p.Y < height
}
