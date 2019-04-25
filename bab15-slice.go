package main

import "fmt"

func main() {
	var fruits = []string{"apel", "salak", "tomat"}
	var apel = fruits[:1] // akses elemen dari 0 sampai 1

	fmt.Println(apel)
	fmt.Println(len(fruits))

	dst := make([]string, 2)
	src := []string{"melon", "semangka"}
	n := copy(dst, src)

	fmt.Println(n) // print jumlah elemen dari dst
}
