package main

/*กิ่งไม้มีความยาวเป็น N ต้องการตัดกิ่งไม้เป็นท่อนเล็กๆ ตามเงื่อนไขสองข้อดังต่อไปนี้:

กิ่งไม้แต่ละท่อนต้องมีความยาว A, B หรือ C
จำนวนท่อนไม้ควรมีมากที่สุด
Input บรรทัดแรกจะเป็นค่า N A B C โดยที่ N คือ ความยาวของกิ่งไม้ (1 ≤ N) และ A, B, C คือความยาวของท่อนไม้ที่กำหนด (A, B, C ≤ 4000) Output แสดงจำนวนท่อนไม้ทั้งหมดที่ตัดได้*/

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
