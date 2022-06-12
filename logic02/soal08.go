package logic02

import "fmt"

func Logic02Soal08(n int) {
	a := 3

	//Looping baris
	for i := 0; i < n; i++ {
		//Looping kolom
		for j := 0; j < n; j++ {
			//Kondisi untuk kotak kosong
			if i > j {
				fmt.Print(" ", "\t")
			} else {
				fmt.Print(a, "\t")
			}
		}
		//Pindah baris selanjutnya
		fmt.Println("\n")
		//nilai variabel a ditambah 3
		a += 3
	}
}
