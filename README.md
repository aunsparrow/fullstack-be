# fullstack-be
1.Solve Problem
1.1******************************************************************
package main

import (
	"fmt"
)

func main() {
	n := 157 //input
	count := 0
	countBankNote := 0
	bankNotes := []int{100, 50, 20, 10, 5, 1}
	for i := 0; i <= count; i++ {
		if n >= bankNotes[i] {
			current := n % bankNotes[i]
			countBankNote += n / bankNotes[i]
			if current != 0 {
				count += 1
			}
			n -= (n / bankNotes[i]) * bankNotes[i]

			if n == 0 {
				break
			}

		} else{
			count += 1
		}
		

	}

	fmt.Println(countBankNote)
}

1.2******************************************************************
package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 4//input
	a := 4//input
	b := 4//input
	c := 4//input
	count :=0
	logCount :=0
	intSlice := []int{a,b,c}
    	sort.Sort(sort.IntSlice(intSlice))
	
	for i := 0; i <= count; i++ {
		if n>= intSlice [i]{
			n -= intSlice [i]
			logCount ++
			count++
		}
		
		 
	}
	
	
    	fmt.Println(logCount )
}

1.3******************************************************************(ทำไม่ทัน)

how to run api
-go root project 
-go get github.com/gofiber/fiber/v2 v2.10.0
-go get gorm.io/driver/postgres
-gp get gorm.io/gorm

collection postman อยู่ใรโปรเจ็ค



