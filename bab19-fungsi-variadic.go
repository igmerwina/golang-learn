package main

import "fmt"

/*
	fungsi variadic = fungsi yang pake parameter sejenis yang tak terbatas
	kaya contoh fungsi(angka, angka, angka..... ~)
*/

func mainFuncVariadic() {
	var avg = hitung(2, 1, 2, 1, 2, 1, 2)
	var msg = fmt.Sprintf("%.2f", avg)
	fmt.Println(msg)
}

// fungsi variadic kena ...
func hitung(numbers ...int) float64 {
	jumlah := 0
	for _, number := range numbers {
		jumlah += number
	}

	var avg = float64(jumlah) / float64(len(numbers))
	return avg
}
