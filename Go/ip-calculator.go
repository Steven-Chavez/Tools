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

	var octet_str []string

	var octet_int [4]int
	var oct1 string
	var oct2 string
	var oct3 string
	var oct4 string

	octet_str = strings.Split(ip, ".")

	oct1 = octet_str[0]
	oct2 = octet_str[1]
	oct3 = octet_str[2]
	oct4 = octet_str[3]

	octet_int[0], err = strconv.Atoi(oct1)
	octet_int[1], err = strconv.Atoi(oct2)
	octet_int[2], err = strconv.Atoi(oct3)
	octet_int[3], err = strconv.Atoi(oct4)

	fmt.Println(err)

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
