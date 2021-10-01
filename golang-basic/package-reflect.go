package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string
}

func main() {
	sample := Sample{"Johar"}
	sampleType := reflect.TypeOf(sample)
	structField := sampleType.Field(0)
	// required := structField.Tag.Get("requared")
	fmt.Println(structField)
}