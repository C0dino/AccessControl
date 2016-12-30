package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	fmt.Println(`
	========================================================================================
	Welcome to the Card Decode Program.  Sometimes referred to as "long decode" this program
	takes a facility code, a card code and returns the "fully decoded" number which will
	show up in an Access control event log such as Lenel. (Good for EZ Pass Decodes)
	========================================================================================
	`)

	for {
		var num1, num2, moreCards int64
		var answer, additionalDecode string

		fmt.Print("Enter Facility Code: ")
		fmt.Scan(&num1) // Prompts user for input
		fac := dec2Hex(num1)

		fmt.Print("Enter Card Code: ")
		fmt.Scan(&num2) // Prompts user for input
		cardCode := dec2Hex(num2)

		fullDecode := hex2Dec(fac + cardCode) //concatenates strings fac + cardCode
		fmt.Println("Fully Decoded Card Number is:", fullDecode)
		fmt.Print("Would you like to decode additional consecutive cards? (y or n)")
		fmt.Scan(&additionalDecode)
		if additionalDecode == "y" {
			fmt.Print("How many additional cards? (enter #)")
			fmt.Scan(&moreCards)
			moreDecode := fullDecode + moreCards
			for fullDecode <= moreDecode {
				num2++
				fullDecode++
				fmt.Println(num2, " Decodes to: ", fullDecode)
			}
		}
		fmt.Print("Continue? (y or n): ")
		fmt.Scan(&answer)
		if answer == "n" {
			break
		}
		fmt.Println("--------------Next Card(s)--------------")
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
