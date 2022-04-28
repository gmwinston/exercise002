package main

import (
	"fmt"
	"log"
	"net"

)
func main() {
	/*if len(os.Args) != 2 {
		log.Fatal("No IP address argument provided.")
	}*/
	//arg := os.Args[1]
     ips:= "210.59.110.4"
	// Parse the IP for validation
	ip := net.ParseIP(ips)
	if ip == nil {
		log.Fatal("Valid IP not detected. Value provided: " + ips)
	}

	fmt.Println("Looking up hostnames for IP address: " + ips)
	hostnames, err := net.LookupAddr(ip.String())
	if err != nil {
		log.Fatal(err)
	}
	for _, hostnames := range hostnames {
		fmt.Println(hostnames)
	}
}