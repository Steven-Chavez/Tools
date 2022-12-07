// Author: Steven Chavez
// FileName: ip-calculator.go
// Creation Date: 12/1/2022

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
// Determines the class of an IP
func determineIPClass(ip [4]int) {

	// Variables
	class_A := "IP Class - A"
	class_B := "IP Class - B"
	class_C := "IP Class - C"
	oct1 := ip[0]

	// IP Classes
	// A - 0-127
	// B - 128-191
	// C - 192-223
	if oct1 >= 0 && oct1 <= 127 {
		fmt.Println(class_A)
	} else if oct1 >= 128 && oct1 <= 191 {
		fmt.Println(class_B)
	} else if oct1 >= 192 && oct1 <= 223 {
		fmt.Println(class_C)
	}

}

// FUNCTION
// publicOrPrivateIP() determines if an IP is a public
// of private ip
func isPrivate(ip [4]int) bool {

	// Variables
	private_str := "IP Status (public/private) - Private"
	public_str := "IP Status (public/private) - Public"
	oct1 := ip[0]
	oct2 := ip[1]

	// Private IP ranges
	// 10.0.0.0 - 10.255.255.255
	// 192.168.0.0 - 192.168.255.255
	// 172.16.0.0 - 172.31.255.255
	if oct1 == 10 {
		fmt.Println(private_str)
		return true
	} else if oct1 == 192 {
		if oct2 == 168 {
			fmt.Println(private_str)
			return true
		} else {
			fmt.Println(public_str)
			return false
		}
	} else if oct1 == 172 {
		if oct2 >= 16 && oct2 <= 31 {
			fmt.Println(private_str)
			return true
		} else {
			fmt.Println(public_str)
			return false
		}
	} else {
		fmt.Println(public_str)
		return false
	}
}

// FUNCTION
// Sanatizes input to ensure it's actually an IP.
func sanatizeInput(ip string) bool {
	// variables
	var ip_bool bool
	length := len([]rune(ip))

	// First check to see if the input is within range
	// of the min and max lenth an IP can be.
	// min: 0.0.0.0 = 7
	// max: 255.255.255.255 = 15
	if length >= 7 && length <= 15 {

		// split ip into octets
		octet_str := strings.Split(ip, ".")

		// Obtain the length of the new array
		// Should be 4 because there are 4 octets
		// in an IP address.
		length = len(octet_str)

		// Loop through each octet to see if the
		// values are numeric and within range
		for i := 0; i < length; i++ {

			// Variable to hold error message
			var err error

			// Convert each octet into an int
			_, err = strconv.Atoi(octet_str[i])

			if err == nil { // The octet was a number

				// Obtain the length of the octet
				var length int = len([]rune(octet_str[i]))

				// Each octet must be between 1 and 3 digits
				if length >= 1 && length <= 3 {
					ip_bool = true
				} else {
					return false
				}
			} else { // The octet was not a number
				return false
			}
		}
	} else { // Input is below 7 or higher than 15
		return false
	}

	return ip_bool
}

func getIpInfo(ip string) {
	var err error
	var token []byte
	var url string
	var resp *http.Response

	token, err = os.ReadFile("/home/steven/api/ipinfo.txt")

	if err != nil {
		log.Fatalln(err)
	}

	url = "https://ipinfo.io/" + ip + "?token=" + string(token)
	url = strings.Replace(url, "\n", "", -1)

	resp, err = http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)
}

func main() {
	var ip string

	fmt.Print("Enter an IP: ")
	fmt.Scanln(&ip)
	var realIP bool = sanatizeInput(ip)

	if realIP == true {
		fmt.Print("\n")
		fmt.Println("---------- IP INFO ----------")
		fmt.Print("\n")
		split := breakIntoOctets(ip)
		fmt.Println("IP Entered -", ip)
		ipstatus := isPrivate(split)
		determineIPClass(split)
		fmt.Print("\n")
		if ipstatus == false {
			getIpInfo(ip)
		}
	} else {
		fmt.Println("ERROR: Not an IP!")
	}
}
