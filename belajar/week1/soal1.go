package main
 import "fmt"

func main() {
	var  a, b int
	 
	fmt.Print("masukkan harga barang =")
	fmt.Scanln(&a)
	fmt.Print("masukkan diskon = ")
	fmt.Scanln(&b)

	diskon := a * b/100
 
	total := a - diskon

	fmt.Print("total belanja akhir =", total)
}
