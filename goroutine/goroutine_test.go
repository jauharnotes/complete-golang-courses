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

// channel
func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2*time.Second)
		channel <- "Jauharuddin"
		fmt.Println("Selesai Mengirim Data ke Channel")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5*time.Second)
}

func GivMeResponse(channel chan string) {
	time.Sleep(2*time.Second)
	channel <- "Jauharuddin"
}

func TestChannelParameters(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GivMeResponse(channel)

	data := <-channel
	fmt.Println(data)
}