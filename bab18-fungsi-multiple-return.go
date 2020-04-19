package main

import (
	"fmt"
	"math"
)

func mainFuncMultiple() {
	var diameter float64 = 15
	var area, circumference = calculated(diameter)

	fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)
}

//  Fungsi Dengan Predefined Return Value
func calculated(d float64) (area float64, circumference float64) {
	area = math.Pi * math.Pow(d/2, 2)
	circumference = math.Pi * d

	return
}

// --- fungsi biasa ---
// func calculate(d float64) (float64, float64) {
// 	// hitung luas
// 	var area = math.Pi * math.Pow(d/2, 2)
// 	// hitung keliling
// 	var circumference = math.Pi * d

// 	// kembalikan 2 nilai
// 	return area, circumference
// }
