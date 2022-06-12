package logic02

import "fmt"

func Logic02Soal05(n int) {
	nt := n / 2
	a := 3

	// looping baris
	for i := 0; i < n; i++ {
		// looping kolom
		for j := 0; j < n; j++ {
			if j <= nt {
				fmt.Print(a, "\t")
				a += 3
			} else {
				a -= 3
				fmt.Print(a, "\t")
			}
		}
		fmt.Println("\n")
	}
}
