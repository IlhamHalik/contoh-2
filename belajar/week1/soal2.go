package main
 import "fmt"

func main(){
	var(
		BMI, Tb, BB float64
	)
	fmt.Print("masukkan BMI = ")
	fmt.Scanln(&BMI)
	fmt.Print("masukkan tinggi badan = ")
	fmt.Scanln(&Tb)

	BB = BMI * Tb * Tb

	fmt.Printf("Berat Badan = %.0f", BB)
}