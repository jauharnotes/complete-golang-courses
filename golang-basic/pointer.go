package main

import "fmt"

type Address struct {
  kota, provinsi, negara string
}

func change(original *int, value int) {
    *original = value
}

func main(){
  address1 := Address{"Depok", "Jawa Barat", "Indonesia"}
  address2 := &address1

  address2.provinsi = "Jawa Timur"

  *address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

  fmt.Println(address1)
  fmt.Println(address2)

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


