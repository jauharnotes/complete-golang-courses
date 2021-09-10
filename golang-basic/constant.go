package main

import "fmt"

func main() {
	const firsName string = "Jauhar"
	const lastName = "Uddin"
	const fullName = "jauharuddin"

	fmt.Println(firsName)
	fmt.Println(lastName)
	fmt.Println(fullName)

	const (
		city string = "Depok"
		street = "Karya Bhakti"
		no = 125
	)

	fmt.Println(city)
	fmt.Println(street)
	fmt.Println(no)
}