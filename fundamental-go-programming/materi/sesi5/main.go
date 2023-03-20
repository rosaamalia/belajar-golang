package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/**
 * Goroutine
 */
func prosesPertama(index int) {
	fmt.Println("Proses pertama dimulai")
	for i := 0; i < index; i++ {
        fmt.Println("index ke-", i)
    }
	fmt.Println("Proses pertama berakhir")
}

func prosesKedua(index int) {
	fmt.Println("Proses kedua dimulai")
    for i := 0; i < index; i++ {
        fmt.Println("index ke-", i)
    }
    fmt.Println("Proses kedua berakhir")
}

func goroutine() {
	/**
	 * Goroutine adalah thread pada bahasa Go
	 * untuk melakukan concurrency yang bersifat asinkronus
	 */

	fmt.Println("Proses main goroutine dimulai")

	// Untuk menggunakan goroutine, dengan
	// menambahkan keyword `go` ke fungsi yang ingin
	// dijalankan sebagai goroutine
	go prosesPertama(8)

	prosesKedua(8)

	// Menghitung jumlah goroutine yang berjalan
	// Fungsi main() termasuk goroutine
	fmt.Println("Jumlah Goroutine ", runtime.NumGoroutine())

	// Karena sifatnya asinkronus, fungsi main()
	// tidak akan menunggu goroutine prosesPertama selesai
	// Jadi prosesPertama bisa saja belum selesai tereksekusi
	// setelah fungsi main() selesai

	// Untuk menahan fungsi main() berjalan lebih lama
	time.Sleep(time.Second * 2)

	fmt.Println("Proses main goroutine berakhir")
}

/**
 * Sync-waitgroup
 */
func sync_waitgroup() {
	/**
	 * Waitgroup digunakan untuk membantu proses
	 * sinkronisasi Goroutine
	 */

	fruits := []string {"apel", "mangga", "durian", "rambutan"}

	// Deklarasi variabel sync-waitgroup
	var wg sync.WaitGroup

	// Membuat goroutine
	for index, fruit := range fruits {
		// Counter untuk menghitung jumlah goroutine
		// yang harus ditunggu
		wg.Add(1)

		// variabel wg menggunakan pointer agar
		// sync-waitgroup menggunakan memori yang sama
		go printFruit(index, fruit, &wg)
	}

	fmt.Println("Jumlah Goroutine ", runtime.NumGoroutine())

	// Wait() digunakan untuk menahan fungsi main
	// menunggu proses goroutine selesai
	wg.Wait()
}

func printFruit(index int, fruit string, wg *sync.WaitGroup) {
	fmt.Printf("index ke-%d, buah %s\n", index, fruit)
	wg.Done() // Tanda proses goroutine yg ditunggu selesai
}

/**
 * Channel
 */
func channel() {
	/**
	 * Channel adalah mekanisme untuk Goroutine
	 * saling berkomunikasi dengan Goroutine lainnya
	 * (mengirim data dari satu Goroutine ke Goroutine lain)
	 */

	// Deklarasi variabel untuk channel
	// Menambahkan keyword `chan`
	c := make(chan string) // variabel c merupakan channel bertipe string

	/**
	 * Operator yang digunakan berkomunikasi
	 * Mengirim data ke channel:
	 * c <- value
	 * Menerima data dari channel:
	 * result := <- c
	 */
	
	// Membuat goroutine
	go introduce("Ariel", c)
	go introduce("Cinderella", c)
	go introduce("Daniel", c)

	// Menerima data dari goroutine introduce() melalui channel
	value1 := <- c
	fmt.Println(value1)

	value2 := <- c
	fmt.Println(value2)

	value3 := <- c
	fmt.Println(value3)

	// Karena sifatnya asinkronus, value yang akan
	// ditampilkan sesuai dengan proses goroutine yang
	// selesai duluan dieksekusi

	// Menutup channel yang tidak digunakan
	close(c)
}

func introduce(student string, c chan string) {
	result := fmt.Sprintf("Halo, nama saya %s", student)

	// Variabel result dikirim ke channel c
	c <- result
}

func channel_anonymous() {
	/**
	 * Channel dengan anonymous function
	 */

	c := make(chan string)
	students := []string {"Ariel", "Cinderella", "Daniel"}

	for _, v := range students {
		// Fungsi anonymous untuk membuat goroutine
		go func(student string) {
			fmt.Println("Siswa: ", student)
			result := fmt.Sprintf("Halo, nama saya %s", student)
			c <- result // Mengirim data ke channel
		} (v)
	}

	for i:= 0; i < len(students); i++ {
		// Data yang dikirim melalui channel tidak perlu ditampung
		// terlebih dahulu, jadi langsung dapat menggunakan `<-c`
		fmt.Println(<-c)
    }

	close(c)
}

func main() {
	// goroutine()
	// sync_waitgroup()
	// channel()
	// channel_anonymous()
}