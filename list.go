package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()

	data.PushBack("Adib")
	data.PushBack("Hauzan")
	data.PushBack("Sofyan")

	var head *list.Element = data.Front()
	fmt.Println(head.Value)

	head = head.Next()
	if head != nil {
		fmt.Println(head.Value)
	}

	next := head.Next()
	if next != nil {
		fmt.Println(next.Value)
	}

	next = next.Next()
	if next != nil {
		fmt.Println(next.Value)
	}

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}
