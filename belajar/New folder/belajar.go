package main
 
import "fmt"

func main() {
	var ( //mendeklarasi variabel yang tipenya string
		satu, dua, tiga string
		temp            string
	)
	fmt.Print("Masukan input string: ")//ini buat nampilin di terminal 
	fmt.Scanln(&satu)// ini buat ngebaca yang kita input dan nyimpen di variabel satu
	fmt.Print("Masukan input string: ")//ini buat nampilin di terminal
	fmt.Scanln(&dua)// ini buat ngebaca yang kita input dan nyimpen di variabel dua
	fmt.Print("Masukan input string: ")//ini buat nampilin di terminal
	fmt.Scanln(&tiga)// ini buat ngebaca yang kita input dan nyimpen di variabel tiga
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)//buat nampilin teks sama variabelnya
	temp = satu//ini buat nnuker yang di variabel satu jadi 3 trs variabel 2 jadi satu dan tiga jadi 2
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)// buat nampilin hasil pertukarannya

}
