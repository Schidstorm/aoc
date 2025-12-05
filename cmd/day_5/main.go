package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Schidstorm/aoc/pkg/util"
)

func main() {
	lines := strings.Split(util.LoadDayFile(5).Content(), "\n")

	i := 0
	var freshRanges []numberRange
	for {
		if lines[i] == "" {
			break
		}

		line := lines[i]
		parts := strings.Split(line, "-")
		freshRanges = append(freshRanges, numberRange{util.ParseUint64(parts[0]), util.ParseUint64(parts[1])})
		i++
	}

	sort.Sort(rangesSort(freshRanges))

	freshRanges = combineSortedRanges(freshRanges)

	var sum uint64
	for _, r := range freshRanges {
		rangeNum := r.end - r.begin + 1
		sum += rangeNum
	}

	fmt.Println(sum)

}

type rangesSort []numberRange

func (rs rangesSort) Len() int {
	return len(rs)
}
func (rs rangesSort) Less(i, j int) bool {
	return rs[i].begin < rs[j].begin
}
func (rs rangesSort) Swap(i, j int) {
	tmp := rs[i]
	rs[i] = rs[j]
	rs[j] = tmp
}

func combineSortedRanges(ranges []numberRange) []numberRange {
	for i := len(ranges) - 2; i >= 0; i-- {
		if ranges[i+1].begin <= ranges[i].end {
			ranges[i].end = max(ranges[i].end, ranges[i+1].end)
			ranges[i+1].begin = 0
			ranges[i+1].end = 0
		}
	}

	var currentEnd int
	for i := range len(ranges) {
		if ranges[i].begin == 0 && ranges[i].end == 0 {
			continue
		}

		ranges[currentEnd] = ranges[i]
		currentEnd++
	}

	return ranges[:currentEnd]
}

func max(a, b uint64) uint64 {
	if a > b {
		return a
	}

	return b
}

func isNumberInRanges(n uint64, ranges []numberRange) bool {
	for _, r := range ranges {
		if n >= r.begin && n <= r.end {
			return true
		}
	}

	return false
}

type numberRange struct {
	begin, end uint64
}
