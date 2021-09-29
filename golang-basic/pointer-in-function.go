package main

import "fmt"

type Address struct {
	city, provinsi, country string
}

func ChangeAddressToIndonesia(address *Address) {
	address.country = "Indonesia"
}

func main() {
	address := Address{"Depok", "Jawa Barat", ""}
	ChangeAddressToIndonesia(&address)

	fmt.Println(address)
}