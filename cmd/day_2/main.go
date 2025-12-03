package main

import (
	"fmt"
	"strings"

	"github.com/Schidstorm/aoc/pkg/util"
)

func main() {
	ranges := util.LoadDayFile(2).Split(",")

	var counter int64
	for _, r := range ranges {
		from := util.ParseInt64(strings.Split(r, "-")[0])
		to := util.ParseInt64(strings.Split(r, "-")[1])

		for n := from; ; {
			l := log10(n)
			lowest := nextInvalidId(n, 2)
			for i := int64(3); i <= l; i++ {
				lowest = min(lowest, nextInvalidId(n, i))
			}
			if lowest > to {
				break
			}
			fmt.Println(lowest)
			counter += lowest
			n = lowest + 1
		}
	}

	fmt.Println(counter)
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func nextInvalidId(n int64, parts int64) int64 {
	l10 := log10(n)
	if l10%parts != 0 {
		return nextInvalidId(pow10(l10), parts)
	}

	left := getLeftestPart(n, parts)
	if partsEqual(n, parts) {
		return n
	} else {
		a := sameParts(left, parts)
		if a > n {
			return a
		}
		return sameParts(left+1, parts)
	}
}

func getLeftestPart(n, parts int64) int64 {
	shiftAmount := log10(n) - (log10(n) / parts)
	return n / pow10(shiftAmount)
}

func partsEqual(n, parts int64) bool {
	left := getLeftestPart(n, parts)
	shiftAmount := pow10(log10(n) / parts)
	tmpN := n
	for i := int64(1); i < parts; i++ {
		if tmpN%shiftAmount != left {
			return false
		}
		tmpN /= shiftAmount
	}

	return true
}

func sameParts(n, parts int64) int64 {
	shiftAmount := pow10(log10(n))
	var result int64
	for ; parts > 0; parts-- {
		result *= shiftAmount
		result += n
	}

	return result
}

func log10(n int64) int64 {
	var count int64
	for n > 0 {
		count++
		n /= 10
	}
	return count
}

func pow10(count int64) int64 {
	n := int64(1)
	for i := int64(0); i < count; i++ {
		n *= 10
	}
	return n
}
