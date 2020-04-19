package main

import "fmt"

func mainLoop() {
	// cara 1
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Sapi ke:", i)
	// }

	// cara 2
	// var j = 0
	// for j < 5 {
	// 	fmt.Println("Sapi ke:", j)
	// 	j++
	// }

	// cara 3
	// lakukan perulangan dari 0 - 10
	// for i := 0; i <= 10; i++ {
	// 	// cari angka ganjil [%2 == 1 --> ganjil]
	// 	if i%2 == 1 {
	// 		continue
	// 	}

	// 	// i > 8 berhenti
	// 	if i > 8 {
	// 		break
	// 	}
	// 	// print nilai i yang bukan ganjil
	// 	fmt.Println("Sapi ke:", i)
	// }

	// cara 4, outerloop
outerloop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if i == 2 {
				break outerloop
			}
			fmt.Print("matrix [", i, "][", j, "]", "\n")
		}
	}
}
