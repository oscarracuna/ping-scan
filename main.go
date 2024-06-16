package main

import (
	"fmt"
	"strings"
	//"net"
)

// TODO:
// Figure out how to cast last octet from string to int so you can
// do something like: last_octet++ and ping until you reach 255
func main() {
	var ipRange int

	var firstIP string
	fmt.Println("welcome to the basic ping scanner")
	fmt.Println("1st IP: ")
	fmt.Scan(&firstIP)

	firstIP_octets := strings.Split(firstIP, ".")
	//firstIP_octets := splitInt()
	fmt.Println(firstIP_octets)
	fmt.Println(firstIP_octets[3])

	//This is crucial.
	//Remember that strings.Split gives you an array, so you can
	//use it to get the last octet by using its index [3]
	if firstIP_octets[3] != "0" {
		ipRange = 0
	} else { //This else doesn't make much sense but
		ipRange = 0
	}

	for j := ipRange; j <= 255; j++ {
		fmt.Println("Pinging IP ending in: ", j)

	}

}

//This could help, but not rn
/*
func splitInt(n int) []int {
	slc := []int{}
	for n > 0 {
		slc =append(slc, n%10)
		n = n / 10
	}
	return slc
}
*/
