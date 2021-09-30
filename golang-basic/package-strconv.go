package main

import (
	"fmt"
	"strconv"
)

func main() {
	// persing to boolean
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	// parsing to integer
	integer, err := strconv.ParseInt("1000000", 10, 64)
	if err == nil {
		fmt.Println(integer)
	} else {
		fmt.Println(err.Error())
	}

	value := strconv.FormatInt(1000000, 10)
	fmt.Println(value)
}