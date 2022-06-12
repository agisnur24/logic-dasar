package logic03

import (
	"fmt"
	"tugas-logic/array_f"
)

func Logic03Soal03(n int) {
	a := 3

	for i := 0; i < n; i++ {
		nt := n / 2
		for j := 0; j < n; j++ {
			if i < j && j <= nt {
				fmt.Print(a, "\t")
			} else if i > j && j >= nt {
				fmt.Print(a, "\t")
			} else if i+j < n-1 && i >= nt {
				fmt.Print(a, "\t")
			} else if i+j > n-1 && i <= nt {
				fmt.Print(a, "\t")
			} else {
				fmt.Print(" ", "\t")
			}
		}
		println("\n")
		if i < nt {
			a += 3
		} else {
			a -= 3
		}
	}
}

func Logic03Soal03WithArray(n int) {
	array := array_f.NewNumberArray(n, n) //line ini untuk membuat array

	//mengisi array
	a := 3 //Tipe variable int
	nt := n / 2
	for i := 0; i < len(array); i++ { // Looping baris
		for j := 0; j < len(array[i]); j++ { //Looping kolom
			if i < j && j <= nt {
				array[i][j] = int32(a) //Konvert variable a ke int32
			} else if i > j && j >= nt {
				array[i][j] = int32(a) //Konvert variable a ke int32
			} else if i+j < n-1 && i >= nt {
				array[i][j] = int32(a) //Konvert variable a ke int32
			} else if i+j > n-1 && i <= nt {
				array[i][j] = int32(a) //Konvert variable a ke int32
			}
		} //Looping kolom selesai
		if i < nt {
			a += 3
		} else {
			a -= 3
		}
	} //Looping baris selesai
	array_f.PrintNumberArrayZero(array) //Print array
}
