package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Jauhar Udding", "Jauhar"))
	fmt.Println(strings.Split("Jauhar Uddin", " "))
	fmt.Println(strings.ToLower("Jauhar Uddin"))
	fmt.Println(strings.ToUpper("Jauhar Uddin"))
	fmt.Println(strings.Trim("   Jauhar Uddin"  , " "))
	fmt.Println(strings.ReplaceAll("jauhar jauhar budi juned", "jauhar", "johar"))
}