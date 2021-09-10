package main

import "fmt"

func main() {
	// convert to int
	var nilai32 int32 = 127
	var nilai64 int64 =int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32, nilai64, nilai8)

	// convert to string
	var name = "Jauhar"
	var e = name[0]
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}