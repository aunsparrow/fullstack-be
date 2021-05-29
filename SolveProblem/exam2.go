package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 4 //input
	a := 4 //input
	b := 4 //input
	c := 4 //input
	count := 0
	logCount := 0
	intSlice := []int{a, b, c}
	sort.Sort(sort.IntSlice(intSlice))

	for i := 0; i <= count; i++ {
		if n >= intSlice[i] {
			n -= intSlice[i]
			logCount++
			count++
		}

	}

	fmt.Println(logCount)
}
