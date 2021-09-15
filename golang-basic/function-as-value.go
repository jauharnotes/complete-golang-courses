package main

import "fmt"

func goodBy(name string) string {
	return "Selamat tinggal " + name
}

func main() {
	sayGoodBy := goodBy
	name := sayGoodBy("Joko")
	fmt.Println(name)
}