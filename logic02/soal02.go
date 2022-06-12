package logic02

import "fmt"

func Logic02Soal02(n int) {
	//looping baris
	for i := 0; i < n; i++ {
		a := 3
		for j := 0; j < n; j++ {
			fmt.Print(a, "\t")
			a += 3
		}
		fmt.Println("\n")
	}
}
