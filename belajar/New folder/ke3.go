package main

import "fmt"

func main() {
	const pi = 3.14//pi itu bernilai tetap sehingga menggunakan const
	var r float64//mendeklarasi variabel dalam bentuk float
	fmt.Print("masukan jari-jari lingkaran: ")//ini buat nampilin perintah input di terminal
	fmt.Scanln(&r)// ini buat ngebaca yang kita input dan nyimpen di variabel r

	luaslingkaran := pi * r * r // := agar langsung mengisi nilai ke variabel

	fmt.Println("Luas lingkaran =", luaslingkaran)// buat nampilin hasil dari yang suudah di input

}
