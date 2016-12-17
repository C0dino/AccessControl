package main

import (
	"fmt"
	"strconv"
)

func main() {
	var fac, cardCode int64
	fmt.Print("Enter Facility Code: ")
	fmt.Scan(&fac) // Prompts user for input
	strconv.FormatInt(fac, 16)
	fmt.Println(fac)
	fmt.Print("Enter Card Code: ")
	fmt.Scan(&cardCode) // Prompts user for input
	strconv.FormatInt(cardCode, 16)
	fmt.Println(cardCode)
	fullDecode := fac + cardCode
	fmt.Println(fullDecode)
	i, err := ParseInt(stringconv(fullDecode), 16, 0)
	fmt.Println("Card Number is: ", i, err)
}
