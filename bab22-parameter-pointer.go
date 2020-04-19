package main

import "fmt"

// -- parameter pointer --
func mainParameter() {
	var num = 4
	fmt.Println(num)

	change(&num, 10)
	fmt.Println(num)
}

func change(original *int, value int) {
	*original = value // parameter pointer
}

// -- print nilai & alamat
/*
func main() {
	var a int = 10
	var sapi *int = &a
	fmt.Println(a)    // print nilai
	fmt.Println(sapi) // print alamat
} */
