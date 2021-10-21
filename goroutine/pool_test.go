package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPools(t *testing.T) {
	var pool = sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("Jauhar")
	pool.Put("Uddin")
	pool.Put("Batubara")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1* time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)

}