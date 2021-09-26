package main

import "fmt"

func change(original *int, value int) {
    *original = value
}

func main(){
  var numberA int = 4
var numberB *int = &numberA

fmt.Println("numberA (value)   :", numberA)  // 4
fmt.Println("numberA (address) :", &numberA) // 0xc20800a220

fmt.Println("numberB (value)   :", *numberB) // 4
fmt.Println("numberB (address) :", numberB)  // 0xc20800a220

// effect perubahan pointer
var numA int = 4
var numB *int = &numberA

fmt.Println("numberA (value)   :", numA)
fmt.Println("numberA (address) :", &numA)
fmt.Println("numberB (value)   :", *numB)
fmt.Println("numberB (address) :", numB)

fmt.Println("")

numberA = 5

fmt.Println("numberA (value)   :", numA)
fmt.Println("numberA (address) :", &numA)
fmt.Println("numberB (value)   :", *numB)
fmt.Println("numberB (address) :", numB)

//parameter pointer
var number = 4
    fmt.Println("before :", number) // 4

    change(&number, 10)
    fmt.Println("after  :", number) // 10
}


