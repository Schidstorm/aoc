package util

import "strconv"

func ParseInt64(s string) int64 {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return n
}

func ParseInt(s string) int {
	n, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		panic(err)
	}
	return int(n)
}
func ParseUint64(s string) uint64 {
	n, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return n
}

func ParseUint(s string) uint {
	n, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		panic(err)
	}
	return uint(n)
}
