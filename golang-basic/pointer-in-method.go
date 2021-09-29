package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main(){
	johar := Man{"Johar"}
	johar.Married()
	fmt.Println(johar.Name)
}