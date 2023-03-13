package main

import "fmt"

func main() {
	// Perulangan
	for i := 0; i < 5; i++ {
		fmt.Printf("Nilai i = %d\n", i)
	}

	// Perulangan dan if-else
	for j := 0; j < 11; j++ {
		if j == 5 {
			for pos, char := range "САШАРВО" {
				fmt.Printf("character %#U starts at byte position %d\n", char, pos)
			}
		} else {
			fmt.Printf("Nilai j = %d\n", j)
		}
	}
}
