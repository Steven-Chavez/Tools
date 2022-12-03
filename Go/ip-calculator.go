// Author: Steven Chavez
// FileName: ip-calculator.go
// Creation Date: 12/1/2022

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func breakIntoOctets(ip string) [4]int {

	// VARIABLES

	// Error var for strconv
	var err error

	// Slice and array var for IP conversion
	var octet_str []string
	var octet_int [4]int

	// Used to seperate IP into octets for eady conversion
	var oct1 int
	var oct2 int
	var oct3 int
	var oct4 int

	// Split string by "."
	octet_str = strings.Split(ip, ".")

	// Convert each split string var into an int
	oct1, err = strconv.Atoi(octet_str[0])
	oct2, err = strconv.Atoi(octet_str[1])
	oct3, err = strconv.Atoi(octet_str[2])
	oct4, err = strconv.Atoi(octet_str[3])

	fmt.Println(err)

	// Populate int array with octets
	octet_int[0] = oct1
	octet_int[1] = oct2
	octet_int[2] = oct3
	octet_int[3] = oct4

	// return converted IP
	return octet_int
}

// FUNCTION
// publicOrPrivateIP() determines if an IP is a public
// of private ip
func publicOrPrivateIP(ip string) {
	fmt.Println(ip)
}

func main() {
	var ip string

	fmt.Scanln(&ip)
	split := breakIntoOctets(ip)
	fmt.Println(split)
}
