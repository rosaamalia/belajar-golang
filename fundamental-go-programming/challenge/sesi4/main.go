package main

import (
	"fmt"
	"os"
)

type Siswa struct {
	Nama string
	Alamat string
	Pekerjaan string
	Alasan string
}

func main() {
	// Input pada terminal merupakan daftar angka untuk
	// menampilkan data sesuai dengan angka yang diinginkan
	// Contoh command untuk menjalankan: `go run main.go 1 2 3 4`

	// Mendapatkan argumen key dari terminal
	var keys = os.Args[1:]

	// Map untuk menyimpan data siswa
	var Kelas map[string]Siswa
	Kelas = map[string]Siswa{}

	Kelas = map[string]Siswa {
		"1": {Nama: "Alip", Alamat: "Jl. Hangtuah", Pekerjaan: "Mahasiswa", Alasan: "Ingin belajar hal baru"},
		"2" : {Nama: "Syarif", Alamat: "Jl. Tribrata", Pekerjaan: "Pegawai Swasta", Alasan: "Ingin bisa golang"},
		"3" : {Nama: "Reza", Alamat: "Jl. Sudirman", Pekerjaan: "Pegawai Swasta", Alasan: "Mau switch career"},
	} 
	
	for _, key := range keys {
		if _, exist := Kelas[key]; exist {
			fmt.Printf("Key\t\t: %v\n", key)
			fmt.Printf("Nama\t\t: %v\nAlamat\t\t: %v\nPekerjaan\t: %v\nAlasan\t\t: %v\n", Kelas[key].Nama, Kelas[key].Alamat, Kelas[key].Pekerjaan, Kelas[key].Alasan)
			fmt.Printf("\n")
		} else {
			fmt.Printf("Data dengan key %v tidak ditemukan\n\n", key)
		}
	}
}