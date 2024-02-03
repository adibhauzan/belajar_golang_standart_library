package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("adib", "ad"))
	fmt.Println(strings.Split("adib", "a"))
	fmt.Println(strings.ToLower("Adib"))
	fmt.Println(strings.ToUpper("adib"))
	fmt.Println(strings.Trim("                adib             ", " "))
	fmt.Println(strings.ReplaceAll("ADib hauzan adib HAAAA", "adib", "bujang"))
}
