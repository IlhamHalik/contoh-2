package main

import "fmt"

func main(){
	var(
		C, F, K, R int
	)

	fmt.Print("masukkan nilai Celcius = ")
	fmt.Scanln(&C)
	 F = (C * 9 / 5) + 32
	 R = C * 4/5
	 K = (C +273)

	 fmt.Println("hasil nilai Fahrenheit = ", F)
	 fmt.Println("hasil nilai Reamur = ", R)
	 fmt.Println( "hasil nilai kelvin = ", K)
} 