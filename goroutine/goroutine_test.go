package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloOWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloOWorld()
	fmt.Println("Upss")

	time.Sleep(1*time.Second)
}

func DisplyNumber(number int) {
	fmt.Println("Disply", number)
}

func TestManyGoroutines(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplyNumber(i)
	}

	time.Sleep(5*time.Second)
}