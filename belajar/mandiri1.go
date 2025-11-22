package main

import"fmt"

func main(){
	var uts, uas, tugas float64

	fmt.Print("input nilai UTS: ")
	fmt.Scanln(&uts)
	fmt.Print("input nilai UAS: ")
	fmt.Scanln(&uas)
	fmt.Print("input nilai tugas: ")
	fmt.Scanln(&tugas)

	nilai:= uts*0.3 + uas*0.4 + tugas*0.3

	fmt.Print("Nilai Akkhir Mahasiswa: ",nilai)
}