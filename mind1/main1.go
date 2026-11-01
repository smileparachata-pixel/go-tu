package main

import "fmt"

func main() {
	weigth := 3.5
	serviceType := "same-day" // "standart","express","same-day"
	isMember := false
	destinationZone := "bkk" // โซนปลายทาง "bkk","upcountry","remote"

	if weigth <= 0 || weight > 50 {
		fmt.Println("INVALID_WEIGHT")
		return
	} else if destinationZone == "remote" {
		fmt.Println("SAME_DAY_NOT_AVAILABLE")
		return
	} else if destinationZone == "same-day" && weight > 5 {
		fmt.Println("WEIGHT_EXCEED_LIMIT")
		return
	} else {
		switch destinationZone {
		case "bkk" && seviceType == "standard":
			fmt.Printf("FEE_50")
		case "bkk" && seviceType == "express":
			fmt.Printf("FEE_80")
		case "bkk" && seviceType == "same-day":
			fmt.Printf("FEE_120")
		case "remote" || "upcountry" && seviceType == "standard":
			fmt.Printf("FEE_100")
		case "remote" || "upcountry" && seviceType == "express":
			fmt.Printf("FEE_150")
		case "upcountry" && seviceType == "same-day":
			fmt.Printf("FEE_250")

		}

	}
}
