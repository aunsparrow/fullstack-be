/*1.2 ประเทศ XYZ มีธนบัตรทั้งหมดอยู่ 5 ชนิดคือ 1, 5, 10, 20 และ 100 ลุงพรชัยมีเงินอยู่ในธนาคาร N บาท หากถอนเงินออกมาทั้งหมดจะได้ธนบัตรจำนวนน้อยที่สุดจำนวนเท่าไร

Input บรรทัดแรกคือค่า N ซึ่งแสดงถึงจำนวนเงินของลุงพรชัย Output แสดงจำนวนธนบัตรที่น้อยที่สุดที่ลุงพรชัยจะได้รับ*/

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

		} else {
			count += 1
		}

	}

	fmt.Println(countBankNote)
}
