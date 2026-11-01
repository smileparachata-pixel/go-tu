package main

import "fmt"

func main() {
	quantity := 2
	stock := 5
	if quantity <= 0 {
		fmt.Println("จำนวนต้องมากกว่า 0")
	} else if quantity > stock {
		fmt.Println("จำนวนสินค้าไม่พอ")
	} else {
		fmt.Println("สั่งซื้อได้")
	}
	status := "paid"
	switch status {
	case "paid":
		fmt.Println("ชำระเงินแล้ว")
	default:
		fmt.Println("รอชำระ")
	}
}
