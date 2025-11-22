package main

import"fmt"
 func main(){
	var( 
		ms int
		detik float64

	)
	fmt.Print("masukkan waktu: ")
	fmt.Scanln(&ms)

	detik = float64 (ms) /1000

	fmt.Println("waktu dalam detik = ", detik)

 }
