// Author: Steven Chavez
// FileName: subnet-calculator.go
// Creation Date: 12/8/2022

package main

// CREATE FUNCTIONS FOR EACH TASK
// -sanitizer
// -spliter
// -CIDR
// -IP's in Range
// -CIDR IP Range
// -Subnet Mask

func sanitizeInput(ip string) bool {
	// variables
	var ip_bool bool
	length := len([]rune(ip))

	// First check to see if the input is within range
	// of the min and max lenth an IP can be.
	// min: 0.0.0.0/0 = 9
	// max: 255.255.255.255/00 = 18
	if length >= 9 && length <= 18 {

	}
}

func main() {
	sanitizeInput()
}
