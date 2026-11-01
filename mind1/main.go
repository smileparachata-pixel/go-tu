package main

import "fmt"

func main() {
	age := 20           // อายุ (ปี)
	isMember := true    // เป็นสมาชิกหรือไม่
	movieType := "imac" // ประเภทภาพยนตร์ (normai,4dx,IMAX)
	ticketCount := 2    // จำนวนตั๋ว
	hasCoupon := true   // มีคูปองหรือไม่

	if age < 0 || age > 120 || ticketCount <= 0 {
		fmt.Println("ERROR")
		return
	} else if movieType == "4dx" && age < 15 {
		fmt.Println("4DX_AGE_RESTRICTED")
		return
	} else if ticketCount > 10 {
		fmt.Println("BULK_PURCHASE_LIMIT")
		return
	} else if isMember && hasCoupon {
		fmt.Println("VIP_DISCOUNT")
		return
	} else if isMember && !hasCoupon {
		fmt.Println("MEMBER_DISCOUNT")
		return
	} else if !isMember && hasCoupon {
		fmt.Println("COUPON_DISCOUNT")
		return
	} else {
		fmt.Println("STANDARD_PRICE")
		return
	}

}
