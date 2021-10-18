package goroutine

import (
	"fmt"
	"strconv"
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

func OnlyIn(channel chan<- string) {
	time.Sleep(2*time.Second)
	channel <- "Jauharuddin"
}

func OnlyOut(channel<- chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5*time.Second)
}

// buffered
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Johar"
		channel <- "Uddin"
		channel <- "Khanedy"
	}()
		
	go func() {
		fmt.Println(<- channel)
		fmt.Println(<- channel)
		fmt.Println(<- channel)
	}()
			
	fmt.Println(cap(channel))
	fmt.Println(len(channel))

	time.Sleep(2*time.Second)
	fmt.Println("selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i <= 10; i++ {
			channel <- "Perulangan ke" + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima", data)
	}

	fmt.Println("selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GivMeResponse(channel1)
	go GivMeResponse(channel2)

	counter := 0
	for {
		select {
			case data := <-channel1:
				fmt.Println("Data dari channel 1", data)
				counter++
			case data := <-channel2:
				fmt.Println("Data dari channel 2", data)
				counter++
		}
		if counter == 2 {
			break
		}
	}
}

// default select
func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GivMeResponse(channel1)
	go GivMeResponse(channel2)

	counter := 0
	for {
		select {
			case data := <-channel1:
				fmt.Println("Data dari channel 1", data)
				counter++
			case data := <-channel2:
				fmt.Println("Data dari channel 2", data)
				counter++
			default:
				fmt.Println("Menunggu data")
		}
		if counter == 2 {
			break
		}
	}
}