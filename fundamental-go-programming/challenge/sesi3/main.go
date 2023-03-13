package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string = "Selamat Malam"
	var hitung = make(map[any]int)

	// Ubah string input menjadi lowercase
	input = strings.ToLower(input)

	// Hitung setiap karakter dan simpan ke dalam map
	for i := 0; i < len(input); i++ {
		fmt.Printf("%c\n", input[i])

		hitung[string(input[i])]++
	}

	fmt.Println(hitung)
}
