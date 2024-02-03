package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSLice []User

func (s UserSLice) Len() int {
	return len(s)
}

func (s UserSLice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s UserSLice) Swap(i, j int) {
	//temp := s[i]
	//s[i] = s[j]
	//s[j] = temp

	s[i], s[j] = s[j], s[i]
}

func main() {

	users := []User{
		{"Adib", 22},
		{"Mala", 21},
		{"Gilang", 21},
		{"Rinat", 21},
	}

	sort.Sort(UserSLice(users))

	fmt.Println(users)

}
