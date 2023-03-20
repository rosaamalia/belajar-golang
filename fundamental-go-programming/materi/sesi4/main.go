package main

import (
	"fmt"
	"math"
)

/**
 * Interface adalah tipe data dari kumpulan
 * definsi satu atau lebih method
 */

// Interface shape digunakan untuk menghitung
// luas dan keliling dari bangun datar
// tanpa perlu tahu jenis bangun datarnya
type shape interface {
	area() float64
	perimeter() float64
}

// Struct yang akan mengimplementasikan interface shape
// Rectangle
type rectangle struct {
	width float64
	height float64
}

func (r rectangle) area() float64 {
    return r.width * r.height
}

func (r rectangle) perimeter() float64 {
    return 2 * r.width + 2 * r.height
}

// Circle
type circle struct {
	radius float64
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
    return 2 * math.Pi * c.radius
}

func main() {
	/**
	 * Implementasi interface
	 */

	var c1 shape = circle{radius: 10.0}
	fmt.Println("Luas lingkaran:", c1.area())
	fmt.Println("Keliling lingkaran:", c1.perimeter())

	var r1 shape = rectangle{width: 5.0, height: 12.0}
	fmt.Println("Luas persegi panjang:", r1.area())
	fmt.Println("Keliling persegi panjang:", r1.perimeter())

	// Hasil area() dan perimeter() dari dua struct yang berbeda
	// dapat disimpan ke dalam interface shape
	// karena memiliki method yang sama

	// Kalau kesimpulan saya, interface ini fungsinya untuk
	// membuat banyak objek dengan jenis yang sama (method yang sama)
	// sehingga menjadi konsisten.
}