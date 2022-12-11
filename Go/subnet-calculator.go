// Author: Steven Chavez
// FileName: subnet-calculator.go
// Creation Date: 12/8/2022

package main

import (
	"fmt"
	"log"
	"regexp"
)

// CREATE FUNCTIONS FOR EACH TASK
// -sanitizer
// -spliter
// -CIDR
// -IP's in Range
// -CIDR IP Range
// -Subnet Mask

// FUNCTION
// Sanitize input to verify is the input is an IP with a CIDR
//
// EXAMPLE
// 0.0.0.0/0
// 000.000.000.000/00
func sanitizeInput(ip string) bool {

	var regex string = "^[\\d]{1,3}\\.[\\d]{1,3}\\.[\\d]{1,3}\\.[\\d]{1,3}\\/[\\d]{1,2}$"

	found, err := regexp.MatchString(regex, ip)

	fmt.Println("Error", err)
	fmt.Println("Found", found)

	if err != nil {
		log.Fatalln(err)
	}

	if found == true {
		fmt.Println(found)
		return true
	} else {
		return false
	}
}

func main() {
	// Variable
	var ip string

	// Print menu and scan input
	fmt.Print("Enter an IP: ")
	fmt.Scanln(&ip)

	// Sanatize to insure inpute is an IP
	var realIP bool = sanitizeInput(ip)

	fmt.Println(realIP)
}
