package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Dir("hello/helloworld.go"))
	fmt.Println(path.Base("hello/helloworld.go"))
	fmt.Println(path.Ext("hello/helloworld.go"))
	fmt.Println(path.Join("hello", "world", "main.go"))
}
