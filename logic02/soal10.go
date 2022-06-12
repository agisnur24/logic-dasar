package logic02

import "fmt"

func Logic02Soal10(n int) {

	for i := 0; i < n; i++ {
		a := 0
		for j := 0; j < n; j++ {
			a += 3
			if i+j == n-1 {
				fmt.Print(a, "\t")
			} else if i+j <= n-1 {
				fmt.Print(a, "\t")
			} else {
				fmt.Print(" ", "\t")
			}
		}
		fmt.Print("\n")
	}
}
