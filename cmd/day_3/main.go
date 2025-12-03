package main

import (
	"fmt"

	"github.com/Schidstorm/aoc/pkg/util"
)

func main() {
	lines := util.LoadDayFile(3).Lines()

	var sum int64
	for _, line := range lines {
		sum += getMaxJoltage(line)
	}

	fmt.Println(sum)
}

func getMaxJoltage(line string) int64 {
	digits := []byte(line)
	l := len(digits)

	var sum int64
	var nextIndex int
	const num = 12
	for i := range num {
		digitIndex, digit := largest(digits[nextIndex : l-(num-i-1)])
		digitIndex += nextIndex

		fmt.Printf("%s\033[1m%s\033[0m", digits[nextIndex:digitIndex], digits[digitIndex:digitIndex+1])
		sum *= 10
		sum += int64(digit - '0')
		nextIndex = digitIndex + 1
	}

	fmt.Printf("%s\n", digits[nextIndex:])

	return sum
}

func largest(digits []byte) (int, byte) {
	var largest byte
	var index int

	for i, d := range digits {
		if d > largest {
			largest = d
			index = i
		}
	}

	return index, largest
}
