// Author: Steven Chavez
// FileName: subnet-calculator.go
// Creation Date: 12/8/2022

package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

// CREATE FUNCTIONS FOR EACH TASK
// xsanitizer
// -spliter
// -CIDR
// -IP's in Range
// -CIDR IP Range
// -Subnet Mask

func splitIP(ip string) [5]int {
	//var err error

	var split_str []string
	var split_int [5]int

	split_str = strings.Split(ip, ".")
	split_cidr := strings.Split(split_str[3], "/")

	split_str[3] = split_cidr[0]
	split_str = append(split_str, split_cidr[1])

	fmt.Println(split_str)

	return split_int
}

// FUNCTION
// Sanitize input to verify is the input is an IP with a CIDR
//
// EXAMPLE
// 0.0.0.0/0
// 000.000.000.000/00
func sanitizeInput(ip string) bool {

	// Regex pattern used to determine valid IP w/CIDR
	var regex string = "^[\\d]{1,3}\\.[\\d]{1,3}\\.[\\d]{1,3}\\.[\\d]{1,3}\\/[\\d]{1,2}$"

	// Use regexp to see if regex matches input
	found, err := regexp.MatchString(regex, ip)

	// Log error
	if err != nil {
		log.Fatalln(err)
	}

	// If regex matches return true if not return false
	if found == true {
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

	if realIP == true {
		test := splitIP(ip)
		fmt.Println(test)
	} else {
		fmt.Println("Not an IP address with CIDR")
	}
}
