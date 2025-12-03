package main

import (
	"fmt"
	"strings"

	"github.com/Schidstorm/aoc/pkg/util"
)

func main() {
	lines := util.LoadDayFile(1).Lines()

	var counter int
	state := 50
	for _, line := range lines {
		delta := 0
		if after, ok := strings.CutPrefix(line, "L"); ok {
			delta = -util.ParseInt(after)
		} else if after, ok := strings.CutPrefix(line, "R"); ok {
			delta = util.ParseInt(after)
		} else {
			continue
		}

		res, count := rotate(state, delta, 100)
		if res >= 100 || res < 0 {
			panic(res)
		}
		counter += count
		state = res
	}

	fmt.Println(counter)
}

func rotate(value, delta, modulo int) (result, count int) {
	sum := value + delta
	count = abs(sum) / modulo
	if sum == 0 {
		count++
	} else if sum < 0 && value != 0 {
		count++
	}

	shiftAmount := (count + 1) * modulo
	result = (sum + shiftAmount) % modulo

	return
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
