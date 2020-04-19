package main

import "fmt"

func mainArr() {
	// --- cara 1 ---
	// var fruits = [4]string{"apple", "grape", "banana"}

	// fmt.Println("Jumlah elemen: \t\t", len(fruits))
	// fmt.Println("Isi semua elemen: \t", fruits)

	// --- cara 2 --- | Array tanpa jumlah elemen
	// var angkas = [...]string{"sapi", "bebek"}

	// fmt.Println(angkas)

	// --- cara 3 --- | Array multidimensi
	// var angkas1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	// var angkas2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	// fmt.Println("angkas1", angkas1)
	// fmt.Println("angkas2", angkas2)

	// --- cara 4 --- | range
	var buahs = [3]string{"apel", "pisang", "kinca"}

	for i, buah := range buahs {
		fmt.Printf("elemen ke: %d %s\n", i, buah)
	}

}
