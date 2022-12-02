// Author: Steven Chavez
// FileName: ip-calculator.go
// Creation Date: 12/1/2022

package main

import "fmt"

// FUNCTION
// publicOrPrivateIP() determines if an IP is a public
// of private ip
func publicOrPrivateIP(ip string) {
	fmt.Println(ip)
}

func main() {
	var ip string

	fmt.Scanln(&ip)
	publicOrPrivateIP(ip)
}
