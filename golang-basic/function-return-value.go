package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getHello(name string)string {
	// return "Hello "+ name
	if name == "" {
		return "Hello Bro"
	} else {
		return "Hello " + name
	}
}

func randomWithRange(min, max int) int {
	var value = rand.Int() % (max - min +1) + min
	return value
}

//penggunaan return untuk menghentikan proses dalam funsi
func devideNumber(m, n int) {
	if n == 0 {
		fmt.Printf("invalid devider. %d cannot devided br %d\n", m, n)
		return
	}

	var res = m / n
	fmt.Printf("%d / %d = %d\n", m, n, res)
}

func main() {
	result := getHello("Johar")
	fmt.Println(result)

	fmt.Println(getHello(""))

	rand.Seed(time.Now().Unix())
	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)

	devideNumber(10, 2)
	devideNumber(4, 0)
	devideNumber(8, -4)

}