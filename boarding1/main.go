package main

import "fmt"

func main() {
	age := 56            // อายุ (ปี)
	hasPassport := true  // มีพาสปอร์ตหรือไม่
	lugageWeight := 25   // น้ำหนักกระเป๋า (กิโลกรัม)
	ticketClass := "eco" // ประเภทตั๋ว (first,biz,eco)
	isMember := false    // เป็นสมาชิกหรือไม่
	status := ""         // สถานะ (NO_PASSPORT, UNACCOMPANIED_MINOR, ERROR,OVERwEIGHT_REJECT,OVERWEIGHT_FEE,BOARDING)

	if lugageWeight < 0 || age < 0 || age > 120 {
		status = "ERROR"
		fmt.Println("error")
		return
	}
	//ตรวจสอบพาสปอร์ต
	if hasPassport == false {
		status = "NO_PASSPORT"
		fmt.Println("ไม่สามารถขึ้นเครื่องได้เนื่องจากไม่มีพาสปอร์ต")
		return
	}
	//ตรวจสอบอายุ
	if age < 12 && ticketClass != "first" {
		status = "UNACCOMPANIED_MINOR"
		fmt.Println("เด็กเล็กเดินทางคนเดียวได้เฉพาะชั้น First Class ที่มีพนักงานดูแลพิเศษเท่านั้น")
		return
	}
	//ตรวจสอบน้ำหนักกระเป๋า
	if lugageWeight > 30 {
		status = "OVERWEIGHT_REJECT"
		fmt.Println("ไม่สามารถขึ้นเครื่องได้เนื่องจากน้ำหนักกระเป๋าเกิน 30 กิโลกรัม")
		return
	} else if lugageWeight > 20 && ticketClass == "eco" && isMember == false {
		status = "OVERWEIGHT_FEE"
		fmt.Println("สามารถขึ้นเครื่องได้แต่ต้องเสียค่าธรรมเนียมเพิ่มเนื่องจากน้ำหนักกระเป๋าเกิน 20 กิโลกรัม")
		return
	} else {
		status = "BOARDING"
		switch status {
		case "BOARDING":
			switch ticketClass {
			case "first":
				fmt.Println("สามารถขึ้นเครื่องได้ชั้น First Class")

			case "biz":
				fmt.Println("สามารถขึ้นเครื่องได้ชั้น Business Class")

			case "eco":
				fmt.Println("สามารถขึ้นเครื่องได้ชั้น Economy Class")
			default:
				fmt.Println("ข้อมูลผิดพลาด")
			}

		}
	}

}
