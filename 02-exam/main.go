package main

import "fmt"

func main() {
	coreTemp := 1000       // อุณหภูมิแกนกลาง (เคลวิน)
	pressure := 250        // แรงดัน (บาร์)
	coolantLevel := 50     // ระดับน้ำหล่อเย็น (0 ถึง 100%)
	isManualLock := false  // สวิตช์ล็อกฉุกเฉินด้วยมือ
	overrideCode := "BETA" // รหัสคำสั่งพิเศษ: "ALPHA", "BETA", "OMEGA", "NONE"
	action := ""           // ตัวแปรเก็บคำสั่งการทำงาน

	// กฎข้อ E (เซนเซอร์พัง - "SENSOR_ERR"
	if coreTemp < 0 || pressure < 0 || coolantLevel > 100 || coolantLevel < 0 {
		fmt.Println("ERR: เซนเซอร์ขัดข้อง ตัดเข้าโหมดสำรอง")
		action = "SENSOR_ERR"
		return
	}
	//กฎข้อ D (ตัดการทำงานฉุกเฉิน - "SCRAM")
	if coreTemp > 1500 || pressure < 20 {
		if isManualLock && overrideCode != "ALPHA" {

		} else {
			action = "SCRAM"
			fmt.Println("DANGER: ทิ้งแท่งควบคุมและปิดเตาปฏิกรณ์ทันที")
			return
		}

	}
	//กฎข้อ C (สภาวะวิกฤต - "CRITICAL")
	if coreTemp > 1000 || pressure >= 250 {
		if coreTemp <= 1200 && overrideCode != "OMEGA" && !isManualLock {

		} else {
			action = "CRITICAL"
			switch overrideCode {
			case "ALPHA", "BETA":
				fmt.Println("WARN: สภาวะวิกฤต เปิดวาล์วระบายแรงดันฉุกเฉิน")
			default:
				fmt.Println("WARN: สภาวะวิกฤต อพยพเจ้าหน้าที่ออกจากพื้นที่")
			}
			return
		}

	}
	//กฎข้อ B (สภาวะเร่งเครื่อง - "OVERDRIVE")
	if coreTemp >= 800 || pressure >= 200 {
		if coolantLevel < 40 {

		} else {
			action = "OVERDRIVE"
			if action == "OVERDRIVE" && coolantLevel == 50 {
				fmt.Println("STATUS: เดินเครื่องเต็มกำลัง (เตือน: น้ำหล่อเย็นอยู่ที่ขีดจำกัดครึ่งถัง)")

			} else {
				fmt.Println("STATUS: ระบบทำงานปกติ")

			}
			return
		}
	}
	//กฎข้อ A (สภาวะปกติ - "NORMAL")
	if coreTemp >= 500 || pressure > 100 {

		action = "NORMAL"
		if action == "NORMAL" && coolantLevel == 50 {
			fmt.Println("STATUS: เดินเครื่องเต็มกำลัง (เตือน: น้ำหล่อเย็นอยู่ที่ขีดจำกัดครึ่งถัง)")

		} else {
			fmt.Println("STATUS: ระบบทำงานปกติ")

		}
		return

	} else {
		action = "IDLE"
		fmt.Println("STATUS: พลังงานต่ำ เดินเครื่องรอคำสั่ง")
	}
}
