package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`a([a-z])b`)

	fmt.Println(regex.MatchString("adib"))
	fmt.Println(regex.MatchString("aib"))
	fmt.Println(regex.MatchString("aDib"))

	fmt.Println(regex.FindAllString("adib ajib aib usap bujang asu", 10))
}
