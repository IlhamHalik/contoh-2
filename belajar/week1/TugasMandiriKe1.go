package main

import "fmt"

func main() {
	var F int       // input sebagai integer
	var C float64   // hasil konversi pakai float

	fmt.Print("Masukkan suhu Fahrenheit: ")//ini buat nampilin perintah input di terminal
	fmt.Scanln(&F) //ini buat ngebaca yang kita input dan nyimpen di variabel

	// konversi ke float biar hasilnya bilangan riil
	C = float64(F-32) * 5.0 / 9.0

	fmt.Println("Suhu dalam Celcius = ", C)// buat nampilin hasil dari yang suudah di input,
}
