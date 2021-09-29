package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args
	fmt.Println("Argumenet:")
	fmt.Println(arg)

	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname:", name)
	} else {
		fmt.Println("Error", err.Error())
	}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	fmt.Println(username, password)
}