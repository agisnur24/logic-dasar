package logic02

import "fmt"

func Logic02Soal01(n int) {
	a := 3

	//looping baris
	for i := 0; i < n; i++ {
		//looping kolom
		for j := 0; j < n; j++ {
			fmt.Print(a, "\t")
		}
		//looping kolom selesai
		fmt.Println("\n")
		//pindah baris selanjutnya
		//dan update nilai a
		a += 3
	}
}
