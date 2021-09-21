package main

import "fmt"

func main() {
	var point = 6

	switch point {
	case 8:
		fmt.Println("Perfect")
	case 7:
		fmt.Println("Awesome")
	default:
		fmt.Println("Not Bad")
	}

	//dengan banyak kondisi
	var point2 = 6

	switch point2 {
	case 8:
		fmt.Println("Perfect")
	case 7, 6, 5, 4:
		fmt.Println("Awesome")
	default:
		fmt.Println("Not Bad")
	}

	var point3 = 2

	switch point3 {
	case 8:
		fmt.Println("Perfect")
	case 7, 6, 5, 4:
		fmt.Println("Awesome")
	default:
		{
			fmt.Println("Not Bad")
			fmt.Println("you can be better!")
		}
	}

	var point4 = 6

	switch {
	case point4 == 8:
		fmt.Println("Perfect")
	case (point4 < 8) && (point4 > 3):
		fmt.Println("Awesome")
	default:
		{
			fmt.Println("Not Bad")
			fmt.Println("you can be better!")
		}
	}

	// Penggunaan Keyword fallthrough Dalam switch
	var point5 = 6

	switch {
	case point5 == 8:
		fmt.Println("Perfect")
	case (point5 < 8) && (point5 > 3):
		fmt.Println("Awesome")
		fallthrough
	case point5 < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("Not Bad")
			fmt.Println("you can be better!")
		}
	}

	// seleksi komdisi bersarang
	var point6 = 0

	if point6 > 7 {
		switch point6 {
		case 10:
			fmt.Println("perfect")
		default:
			fmt.Println("nice")
		}
	} else {
		if point6 == 5 {
			fmt.Println("not bad")
		} else if point6 == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if point6 == 0 {
				fmt.Println("try herder!")
			}
		}
	}

}
