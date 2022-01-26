package main

import "fmt"

func main() {
	nilai := 50

	if nilai == 100 {
		fmt.Println("Excellent")
	} else if nilai >= 80 {
		fmt.Println("Good")
	} else {
		fmt.Println("it's ok")
	}
}
