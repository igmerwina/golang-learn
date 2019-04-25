package main

import "fmt"

func main() {
	var ayam = map[string]int{
		"ayam goreng":   100,
		"ayam panggang": 50,
		"ayam mentah":   200,
	}

	for key, val := range ayam {
		fmt.Println(key, ": ", val)
	}
	delete(ayam, "ayam mentah")
	fmt.Println(ayam)

}
