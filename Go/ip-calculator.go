// Author: Steven Chavez
// FileName: ip-calculator.go
// Creation Date: 12/1/2022

package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// FUNCTION
// Takes string IP, breaks it into octets, converts
// each octet into integers, assigns the octets to a
// 4 integer array, and returns.
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

	// Check for error
	if err != nil {
		log.Fatal(err)
	}

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
func publicOrPrivateIP(ip [4]int) {

	private_str := "IP Status (public/private) - Private"
	public_str := "IP Status (public/private) - Public"
	oct1 := ip[0]
	oct2 := ip[1]
	//oct3 := ip[2]
	//oct4 := ip[3]

	if oct1 == 10 {
		fmt.Println(private_str)
	} else if oct1 == 192 {
		if oct2 == 168 {
			fmt.Println(private_str)
		} else {
			fmt.Println(public_str)
		}
	} else if oct1 == 172 {

		if oct2 >= 16 && oct2 <= 31 {
			fmt.Println(private_str)
		} else {
			fmt.Println(public_str)
		}
	}

}

func main() {
	var ip string

	fmt.Print("Enter an IP: ")
	fmt.Scanln(&ip)
	split := breakIntoOctets(ip)
	publicOrPrivateIP(split)
}
