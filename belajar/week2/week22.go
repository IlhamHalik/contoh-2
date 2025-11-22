package main

import "fmt"

func main() {
	var (
		r, Vbola, Lbola float64
	)
	const pi = 3.1415926535
	fmt.Print("masukkan nilai jari-jari =")
	fmt.Scanln(&r)
	Vbola = 4 * pi * (r * r * r) / 3
	Lbola = 4 * pi * (r * r)
	fmt.Println("hasil dari Volume bola", Vbola, "hasil dari Luas bola", Lbola)

}
