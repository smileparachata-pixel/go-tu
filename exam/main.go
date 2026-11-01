package main

import "fmt"

func main() {
	const exam = "ระบบตัดเกรดและเกียรตินิยม"
	score := 95
	attendance := 81

	var validation bool
	if score >= 0 && score <= 100 && attendance >= 0 && attendance <= 100 {
		fmt.Println("ข้อมูลถูกต้อง")
		validation = true

	} else {
		fmt.Println("ข้อมูลผิดพลาด")
		validation = false
	}
	// ตรวจสอบสิทธิ์การสอบ
	if validation == true {
		if attendance < 80 {

			fmt.Println("ติด F (เวลาเรียนไม่พอ)")
		}
	}

	// ตรวจสอบเกรด
	if validation == true && attendance >= 80 {
		if score < 50 {
			fmt.Println("F")
		} else if score >= 50 && score < 80 {
			fmt.Println("Pass")
		} else if score >= 80 && score < 90 {
			fmt.Println("A")
		} else {
			fmt.Println("A (เกรียตินิยม)")
		}

	}
}
