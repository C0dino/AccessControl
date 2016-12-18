package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	for {
		var num1, num2 int64
		var answer string

		fmt.Print("Enter Facility Code: ")
		fmt.Scan(&num1) // Prompts user for input
		fac := dec2Hex(num1)

		fmt.Print("Enter Card Code: ")
		fmt.Scan(&num2) // Prompts user for input
		cardCode := dec2Hex(num2)

		fullDecode := hex2Dec(fac + cardCode) //concatenates strings fac + cardCode
		fmt.Println("Fully Decoded Card Number is:", fullDecode)
		fmt.Print("Continue? (y or n): ")
		fmt.Scan(&answer)
		if answer == "n" {
			break
		}
		fmt.Println("--------------Next Card--------------")
	}
}

func dec2Hex(i int64) string {
	return strconv.FormatInt(i, 16)
}
func hex2Dec(s string) int64 {
	i, err := strconv.ParseInt(s, 16, 0)
	if err != nil {
		log.Fatalln(err)
	}
	return i
}
