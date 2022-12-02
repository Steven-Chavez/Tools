// Author: Steven Chavez
// FileName: ip-calculator.go
// Creation Date: 12/1/2022

package main

import (
	"fmt"
	"strings"
)

func breakIntoOctets(ip string) {
	split := strings.Split(ip, ".")
	fmt.Println(split[0])
	fmt.Println(split[1])
	fmt.Println(split[2])
	fmt.Println(split[3])
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
	breakIntoOctets(ip)
}
