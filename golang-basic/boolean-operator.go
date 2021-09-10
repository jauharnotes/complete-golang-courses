package main

import "fmt"

func main() {
	var ujian = 90
	var absensi = 85

	var lulusujian = ujian > 80
	var lulusAbsensi = absensi > 80
	fmt.Println(lulusujian, lulusAbsensi)

	var lulus = lulusujian && lulusAbsensi
	fmt.Println(lulus)

	fmt.Println(ujian > 80 && absensi > 80)
}