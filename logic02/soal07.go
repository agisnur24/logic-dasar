package logic02

import "fmt"

func Logic02Soal07(n int) {
	a := 3

	//Looping baris
	for i := 0; i < n; i++ {
		//Looping kolom
		for j := 0; j < n; j++ {
			if i < j {
				fmt.Print(" ", "\t")
			} else {
				fmt.Print(a, "\t")
			}
		}
		//Pindah baris selanjutnya
		fmt.Print("\n")
		//Saat pindah baris variable a ditambah 3
		a += 3
	}
}
