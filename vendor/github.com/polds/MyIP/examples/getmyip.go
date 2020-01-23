package main

import (
	"fmt"

	"github.com/polds/MyIP"
)

func main() {
	ip, err := myip.GetMyIP()
	if err != nil {
		fmt.Printf("There was an error retrieving my IP Address: %s\n", err.Error())
		return
	}

	fmt.Printf("My IP Address is: %s\n", ip)
}
