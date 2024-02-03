package main

import (
	"fmt"
	"strconv"
)

func main() {
	result, err := strconv.ParseBool("true")

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result)

}
