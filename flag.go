package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "Put your database host")
	var username *string = flag.String("username", "root", "Put your database username")
	var password *string = flag.String("password", "root", "Put your database password")

	flag.Parse()

	fmt.Println(*host, *username, *password)
}
