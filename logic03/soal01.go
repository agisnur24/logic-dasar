package logic03

import "fmt"

func Logic03Soal01(n int) {
	nt := n / 2
	a := 3

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i >= j && i+j <= n-1 {
				fmt.Print(a, "\t")
			} else if i <= j && i+j >= n-1 {
				fmt.Print(a, "\t")
			} else {
				fmt.Print(" ", "\t")
			}
		}
		fmt.Print("\n")
		if i < nt {
			a += 3
		} else {
			a -= 3
		}
	}
}
