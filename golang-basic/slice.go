package main

import "fmt"

func main() {
	var months = [...]string {"januari", "februari", "maret", "april", "mei", "juni", "juli", "august", "september", "oktober", "november", "december"}

	var slice = months[4:7]

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	months[5] = "Diubah"
	fmt.Println(slice)

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Johar")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)
	
	fmt.Println(slice2)
	fmt.Println(months)

	newSlice := make([]string,2,5)

	newSlice[0] = "Juned"
	newSlice[1] = "Agung"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copyClices := make([]string, len(newSlice), cap(newSlice))
	copy(copyClices, newSlice)
	fmt.Println(copyClices)

}