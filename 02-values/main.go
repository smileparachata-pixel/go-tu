package main

import "fmt"

func main() {
	const shop = "Go Shop"
	price := 500
	quantity := 2
	total := price * quantity
	var paid bool
	fmt.Println(shop)
	fmt.Println("ยอดรวม:", total)
	fmt.Println("ชำระเงินแล้ว:", paid)
	quantity = 3
	fmt.Println("ยอดใหม่:", price*quantity)

}
