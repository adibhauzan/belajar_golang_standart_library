package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Hauzan", "Sofyan", "Adib"}
	age := []int{21, 22, 23}

	fmt.Println(slices.Max(age))
	fmt.Println(slices.Min(age))
	fmt.Println(slices.Contains(names, "Sofyan"))
	fmt.Println(slices.Index(names, "Adib"))
}
