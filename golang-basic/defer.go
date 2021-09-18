package main

import "fmt"

 func logging() {
	fmt.Println("Selesai memanggil function")
 }

 func runAplication(value int) {
	 defer logging()
	 fmt.Println("Run aplication")
	 result := 10/ value
	 fmt.Println("Result", result)
 }

 func main() {
	 runAplication(0)
 }