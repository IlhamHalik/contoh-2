package main

import "fmt"

func main(){
	var (
		x, y, z int// input sebagai integer
	
	)

	fmt.Print("masukkan tiga nilai : ")//ini buat nampilin perintah input di terminal
	fmt.Scan(&x, &y, &z)//ini buat ngebaca yang kita input dan nyimpen di variabel
	
	total := x + y + z //perhitungan yg akan dilakukan oleh komputer

	ratarata := float64(total)/3.0 //perhitungan yg akan dilakukan oleh komputer

	fmt.Println("total = ", total)// buat nampilin hasil dari yang suudah di input 
	
	fmt.Println("hasil = ", ratarata)// buat nampilin hasil dari yang suudah di input

}
