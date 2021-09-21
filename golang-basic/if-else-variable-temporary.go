package main

import "fmt"

func  main() {
	var nilaiR = 10

	if ujian := nilaiR / 2; ujian >= 10 {
		fmt.Println("Great!. Anda lulus dengan nilai", ujian)
	} else if ujian >= 7 {
		fmt.Println("Good!. Anda lulus dengan nilai", ujian)
	} else {
		fmt.Println("Sorry!. Anda tidak lulus dengan nilai", ujian)
	}

	var point = 50000.0

	if persen := point/ 100; persen >= 100 {
		fmt.Printf("%.1f%s perfect!\n", persen, "%")
	} else if persen > 70 {
		fmt.Printf("%.1f%s good\n", persen, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", persen, "%")
	}
}