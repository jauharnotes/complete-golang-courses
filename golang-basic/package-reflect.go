package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `requared:"true" max:"10"`
}

func main() {
	sample := Sample{"Johar"}
	var sampleType reflect.Type = reflect.TypeOf(sample)
	
	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)
	fmt.Println(sampleType.Field(0).Tag.Get("requared"))
	fmt.Println(sampleType.Field(0).Tag.Get("max"))

	// sample.Name = "" 
	fmt.Println(IsValid(sample))
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("requared") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}