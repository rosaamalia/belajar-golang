package main

import "fmt"

func conditions() {
	// bisa membuat temporary variable di dalam kondisi
	// dan hanya bisa diakses di dalam scope tersebut

	var currentYear = 2023

	if age := currentYear - 2002; age < 17 {
		fmt.Println("Kamu belum boleh membuat kartu sim")
	} else {
		fmt.Println("Kamu sudah boleh membuat kartu sim")
	}
}

func switchFallthrough() {
	var score = 6
	// Expected output: Not Bad, Please study harder

	switch {
	case score == 8:
		fmt.Println("Perfect")
		fallthrough
	case (score < 8) && (score > 3):
		fmt.Println("Not Bad")
		fallthrough
	case score < 5:
		fmt.Println("Please study harder")
	default:
		{
			fmt.Println("Ayok belajar lagi")
		}
	}
}

func looping() {
	for i := 0; i < 3; i++ {
		fmt.Println("Looping cara pertama", i)
	}

	var j = 0
	for j < 3 {
		fmt.Println("Looping cara kedua", j)
		j++
	}

	var k = 0
	for {
		fmt.Println("Looping cara ketiga", k)

		k++
		if k == 3 {
			break
		}
	}

	// loop bisa diberi label
}

func array() {
	// array memiliki fixed-length
	var numbers [4]int
	numbers = [4]int{1, 2, 3, 4}

	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Angka numbers[%d] %d\n", i, numbers[i])
	}

	for j, number := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", j, number)
	}

	for _, number := range numbers {
		fmt.Printf("Value: %d\n", number)
	}
}

func main() {
	// conditions()
	// switchFallthrough()
	// looping()
	array()
}
