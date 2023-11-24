package main

import (
	"container/list"
	"fmt"
)

func main() {
	mylist := list.New()

	mylist.PushBack(1)
	mylist.PushBack(5)
	elem3 := mylist.PushBack(3)
	mylist.PushFront(10)
	mylist.PushFront(20)
	mylist.Remove(elem3)

	mylist.Remove(mylist.Front())
	mylist.Remove(mylist.Back())

	for temp := mylist.Front(); temp != nil; temp = temp.Next() {
		fmt.Println(temp.Value)
	}
}