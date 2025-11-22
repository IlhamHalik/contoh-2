package main

import "fmt"

func main(){
	var (
		fx, hasil float32
		x int
	)
	
	fmt.Print("masukkan bilangan fx =")
	fmt.Scanln(&fx)

	hasil = 2/(fx - 5) - 5

	x := int(hasil)

	fmt.Println("keluarkan x =", x)

}
