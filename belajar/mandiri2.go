package main

import "fmt"

func main(){
	var( nama string
		persenpajak, gaji float64
	)
	fmt.Print("input nama: ")
	fmt.Scanln(&nama)
	fmt.Print("input gaji: ")
	fmt.Scanln(&gaji)
	fmt.Print("input pajak: ")
	fmt.Scanln(&persenpajak)

	 pajak:=(persenpajak/100)*gaji
	 gajibersih := gaji - pajak

	 fmt.Println("Nama Karyawan:", nama)
	 fmt.Printf("Gaji Pokok: %.f\n", gaji)
	 fmt.Printf("Pajak:%.f\n", pajak)
	 fmt.Printf("Gaji Bersih: %.f \n", gajibersih)
}