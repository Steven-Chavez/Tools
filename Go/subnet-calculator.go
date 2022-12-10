// Author: Steven Chavez
// FileName: subnet-calculator.go
// Creation Date: 12/8/2022

package main

import "fmt"

// CREATE FUNCTIONS FOR EACH TASK
// -sanitizer
// -spliter
// -CIDR
// -IP's in Range
// -CIDR IP Range
// -Subnet Mask

func sanitizeInput(ip string) bool {
	// variables
	//var ip_bool bool
	length := len([]rune(ip))

	// First check to see if the input is within range
	// of the min and max lenth an IP can be.
	// min: 0.0.0.0/0 = 9
	// max: 255.255.255.255/00 = 18
	if length >= 9 && length <= 18 {
		return true
	}
	return false
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
