package main

import "fmt"

func Random() interface{}{
	return true
}

func main() {
	result := Random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	// resultInt := result.(int)
	// fmt.Println(resultInt)

	switch value := result.(type) {
		case string:
			fmt.Println("Value", value, "value is string")
		case int:
			fmt.Println("Value", value, "value is int")
		default:
			fmt.Println("Unknown type")

	}
}