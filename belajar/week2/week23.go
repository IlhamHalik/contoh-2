package main

import "fmt"

func main() {
	var (
		tahun int
	)
	fmt.Print("masukkan tahun = ")
	fmt.Scanln(&tahun)

	if tahun%400 == 0 || (tahun%4 == 0 && tahun%100 != 0) {
		fmt.Println("kabisat = true ")
	} else {
		fmt.Println(" kabisat = false")
	}

}
