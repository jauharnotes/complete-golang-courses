package main

import "fmt"

func main() {
	persons := map[string]string {
		"name": "Johar",
		"address": "Depok",
	}

	persons["Title"] = "Programmer"
	
	fmt.Println(persons)
	fmt.Println(persons["name"])
	fmt.Println(persons["address"])
	fmt.Println(persons["Title"])

	var book = make(map[string]string)
	book["Title"] = "Belajar Golang"
	book["Audhor"] = "Johar"
	book["Terbit"] = "2021"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "Terbit")

	fmt.Println(book)
	fmt.Println(len(book))

	// iterasi item map menggunakan for range
	var printXerox = map[string]int{
		"art carton": 6500,
		"linen": 7000,
		"akasia": 7500,
	}

	for key,val := range printXerox {
		fmt.Println(key, ":", val)
	}

	// Deteksi Keberadaan Item Dengan Key Tertentu
	var fotocopy = map[string]int{"hvs a4": 300, "bookpaper": 350}
	var value, isExits = fotocopy["bookpaper"]

	if isExits {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}
}