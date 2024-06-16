package main

import (
	"fmt"
	"strings"
	//"net"
)

func main() {
	var ipRange int

	firstIP := "192.168.0.1"
	fmt.Println("welcome to the basic ping scanner")
	fmt.Println("1st IP: ")
	fmt.Scan()

	firstIP_octets := strings.Split(firstIP, ".")
	fmt.Println(firstIP_octets)
	fmt.Println(firstIP_octets[3])

	if firstIP_octets[3] != "0" {
		ipRange = 0
	}

	for i := firstIP_octets[3]; i == "0"; {
		fmt.Print("Pinging IP addresses...")
	}
	for j := ipRange; j <= 255; j++ {
		fmt.Println("Pinging IP ending in: ", j)

	}

}
