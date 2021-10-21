package goroutine

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGomaxprocs(t *testing.T) {
	totalCpu := runtime.NumCPU()
	fmt.Println("Total Cpu:", totalCpu)

	totalTread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Tread:", totalTread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Totalgoroutine:", totalGoroutine)
}