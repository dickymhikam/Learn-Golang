package main

import (
	"container/list"
	"fmt"
)

func main()  {
	var data *list.List = list.New()

	data.PushBack("Dicky")
	data.PushBack("M")
	data.PushBack("Hikam")

	var head *list.Element = data.Front();
	fmt.Println(head.Value)

	var next = head.Next()
	fmt.Println(next.Value)

	next = next.Next()
	fmt.Println(next.Value)
}