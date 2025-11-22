package main
 import (
	"fmt"
	"math"
)

func distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2 - x1, 2) + math.Pow(y2 - y1, 2))
}

func main() {
	var (x1, y1, x2, y2, x3, y3 float64
	
	)
	fmt.Print("masukkan x1 dan y1")
	fmt.Scanln(&x1, &y1)

	fmt.Print("masukkan x2 dan y2")
	fmt.Scanln(&x2, &y2)

	fmt.Print("masukkan x3 dan y3")
	fmt.Scanln(&x3, &y3)
	
	ab := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	bc := math.Sqrt(math.Pow(x3-x2, 2) + math.Pow(y3-y2, 2))
	ca := math.Sqrt(math.Pow(x1-x3, 2) + math.Pow(y3-y2, 2))

	
	sisipanjang := math.Max(ab, math.Max(bc, ca))

	
	fmt.Println(" hasil = ", sisipanjang)
}