package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	for _, a := range args {
		fmt.Println(a)
	}

	hostName, err := os.Hostname()
	if err != nil {
		fmt.Println("Error", err.Error())
	}
	fmt.Println(hostName)
}
