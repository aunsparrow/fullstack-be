package main

import (
	"fmt"
	"strings"
)

func main() {

	inputQuantityStone := 5
	_ = inputQuantityStone
	inputStones := "RRRRR"
	_ = inputStones
	count := 0

	fmt.Println("Input")
	fmt.Println(inputQuantityStone)
	fmt.Println(inputStones)

	s := strings.Split(inputStones, "")
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 {
			if s[i] == s[i+1] {
				count++
			}
		}
	}

	fmt.Println(count)
}
