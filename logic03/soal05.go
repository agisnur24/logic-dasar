package logic03

import "fmt"

func Logic03Soal05(n int) {
	a := 3
	nt := n / 2

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i%4 == 0 {
				fmt.Print(a, "\t")
			} else if i%4 == 1 && j == 0 {
				fmt.Print(a, "\t")
			} else if i%4 == 2 {
				fmt.Print(a, "\t")
			} else if i%4 == 3 && j == n-1 {
				fmt.Print(a, "\t")
			} else {
				fmt.Print(" ", "\t")
			}
		}
		fmt.Println("\n")
		if i < nt {
			a += 3
		} else {
			a -= 3
		}
	}
}
