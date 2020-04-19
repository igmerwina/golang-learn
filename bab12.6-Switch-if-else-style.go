package main

import "fmt"

func mainCondition() {
	point := 6

	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad sapi")
			fmt.Println("you need to learn more babi")
		}
	}
}
