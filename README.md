package main 
import"fmt"
func main(){
	var gram, total2, total int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scanln(&gram)

	berat := gram/1000
	sisa := gram%1000
	total1 := berat * 10000

	if berat>10{
		total2 = 0
	}else if sisa>=500{
		total2 = sisa*5
	}else{
		total2 = sisa*15
	}
	total = total1+total2

	fmt.Printf("Detail berat: %d kg + %d gr\n", berat, sisa)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n",total1, total2)
	fmt.Printf("Total biaya: Rp.%d\n",total)
}
