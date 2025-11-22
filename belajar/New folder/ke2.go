package main 

import "fmt"

func main() {

	var (//mendeklarasi variabel yang tipenya string
		satu, dua, tiga string
		temp            string
	)
	fmt.Print("Masukan kelas: ")//ini buat nampilin input di terminal
	fmt.Scanln(&satu)// ini buat ngebaca yang kita input dan nyimpen di variabel satu
	fmt.Print("Masukan nama: ")//ini buat nampilin input di terminal
	fmt.Scanln(&dua)// ini buat ngebaca yang kita input dan nyimpen di variabel dua
	fmt.Print("Masukan NIM: ")//ini buat nampilin input di terminal
	fmt.Scanln(&tiga)// ini buat ngebaca yang kita input dan nyimpen di variabel tiga
	temp = satu//ini buat nnuker yang di variabel satu jadi 3 trs variabel 2 jadi satu dan tiga jadi 2
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("nama saya  " + satu + " dengan NIM " + dua + " dari kelas " + tiga)// buat nampilin hasil dari yang suudah di input dalam bentuk pertukarannya

}
