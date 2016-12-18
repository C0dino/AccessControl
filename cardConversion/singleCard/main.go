package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var num1, num2 int64
	fmt.Print("Enter Facility Code: ")
	fmt.Scan(&num1) // Prompts user for input
	fac := strconv.FormatInt(num1, 16)
	fmt.Println("Facility Code Decimal to Hex:", fac)
	fmt.Print("Enter Card Code: ")
	fmt.Scan(&num2) // Prompts user for input
	cardCode := strconv.FormatInt(num2, 16)
	fmt.Println("Card Code Decimal to Hex:", cardCode)
	fullDecode := fac + cardCode
	fmt.Println("Combined FAC + Card Code Hex:", fullDecode)
	i, err := strconv.ParseInt(fullDecode, 16, 0)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Fully Decoded Card Number is:", i)
}
