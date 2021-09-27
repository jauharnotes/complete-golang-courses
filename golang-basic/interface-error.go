package main

import (
	"errors"
	"fmt"
)

func Pembagi(nilai, pembagi int) (int, error) {
	if nilai == 0 {
		return 0, errors.New("Tidak dapat menjalankan pembagian dengan NOL")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	hasil, err := Pembagi(100, 0)
	if err == nil {
		fmt.Println("Hasil",hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}