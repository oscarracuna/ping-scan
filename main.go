package main

import (
	"fmt"
	"log"
)

var (
	firstIP  string
	secondIP string
)

type ips struct {
	octet1 string
	octet2 string
	octet3 string
	octet4 string
}

func main() {
	fmt.Println("welcome to the basic ping scanner")
	fmt.Println("Type 1st IP: ")
	fmt.Scan(&firstIP)
	fmt.Println("Type 2nd IP: ")
	fmt.Scan(&secondIP)
	log.Println(firstIP, secondIP)

}
