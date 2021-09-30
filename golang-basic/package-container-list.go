package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Jauhar")
	data.PushBack("Uddin")
	data.PushBack("Batubara")
	data.PushFront("Juned")

	// fmt.Println(data.Front().Value)
	// fmt.Println(data.Back().Value)

	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	fmt.Println("===================")


	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}