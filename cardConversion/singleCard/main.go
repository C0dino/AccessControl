package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1, num2 int64
	fmt.Print("Enter Facility Code: ")
	fmt.Scan(&num1) // Prompts user for input
	fac := strconv.FormatInt(num1, 16)
	fmt.Println(fac)
	fmt.Print("Enter Card Code: ")
	fmt.Scan(&num2) // Prompts user for input
	cardCode := strconv.FormatInt(num2, 16)
	fmt.Println(cardCode)
	fullDecode := fac + cardCode
	fmt.Println(fullDecode)
	i, err := strconv.ParseInt(fullDecode, 16, 0)
	fmt.Println("Card Number is: ", i, err)
}
