package main

import "fmt"

func main() {
	var i int = 21
	var j bool = true

	// Menampilkan nilai i
	fmt.Printf("%v\n", i)

	// Menampilkan tipe data variable i
	fmt.Printf("%T\n", i)

	// Menampilkan tanda %
	fmt.Printf("%%\n")

	// Menampilkan nilai boolean j (Output: true)
	fmt.Printf("%t\n\n", j)

	// Menampilkan unicode russia : Я (ya)
	fmt.Printf("\u042F\n")

	// Menampilkan nilai base 10 (Output: 21)
	fmt.Printf("%d\n", 21)

	// Menampilkan nilai base 8 (Output: 25)
	fmt.Printf("%o\n", 21)

	// Menampilkan nilai base 16 (Output: f)
	fmt.Printf("%x\n", 15)

	// Menampilkan nilai base 16 (Output: F)
	fmt.Printf("%X\n", 15)

	// Menampilkan unicode karakter Я (Output: U+042F)
	// desimal untuk bilangan heksa 042F = 1071
	fmt.Printf("%U\n\n", 1071)

	// Menampilkan bilangan float
	var k float64 = 123.456
	fmt.Printf("%f\n", k)

	// Menampilkan bilangan float dalam bentuk notasi scientific
	fmt.Printf("%E\n", k)
}
