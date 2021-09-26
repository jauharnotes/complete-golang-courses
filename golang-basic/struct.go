package main

import "fmt"

type Customer struct {
	name, address string
	age int
}

func (customer Customer) sayHi(name string){
	fmt.Println("Hello", name, "My name is", customer.name)
}

func (customer Customer) sayBay(){
	fmt.Println("Selamat tinggal", customer.name,)
}

func main() {
	var johar  Customer
	johar.name = "Johar"
	johar.address = "Depok"
	johar.age = 28

	fmt.Println(johar.name)
	fmt.Println(johar.address)
	fmt.Println(johar.age)

	// cara lain
	juned := Customer{"Juned", "Depok", 26}
	fmt.Println(juned)

	johar.sayHi("Johar")
	juned.sayBay()
}