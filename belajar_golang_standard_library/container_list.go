package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()

	data.PushBack("Ilham")
	data.PushBack("Muhammad")
	data.PushBack("Haikal")

	var head *list.Element = data.Front()
	fmt.Println(head.Value)

	next := head.Next()
	fmt.Println(next.Value)

	next = next.Next()
	fmt.Println(next.Value)

	// pemanggilan menggunakan for lopp
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
