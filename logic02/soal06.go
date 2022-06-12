package logic02

import "fmt"

func Logic02Soal06(n int) {
	nt := n / 2

	//Looping baris
	for i := 0; i < n; i++ {
		a := 3

		//Looping kolom
		for j := 0; j < n; j++ {
			if j < nt {
				fmt.Print(a, "\t")
				a += 3
			} else {
				fmt.Print(a, "\t")
				a -= 3
			}

		}
		//Pindah baris selanjutnya
		fmt.Println("\n")
	}
}
