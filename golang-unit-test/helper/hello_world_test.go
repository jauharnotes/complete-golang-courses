package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Johar")
	if result != "Hallo Johar"{
		// init test failed
		panic("Result is not Hallo Johar")
	}
}