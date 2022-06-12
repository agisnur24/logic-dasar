package logic02

import "fmt"

func Logic02Soal04(n int) {
	a := n * 3

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(a, "\t")
		}
		fmt.Println("\n")
		a -= 3
	}
}
