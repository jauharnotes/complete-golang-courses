package main

import "fmt"

func endApp(){
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi Error")
	}
}

func main() {
	runApp(false)
}