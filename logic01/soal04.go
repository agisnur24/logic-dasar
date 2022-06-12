package logic01

import "fmt"

func Logic01Soal04(n int) {
	angka := 1
	x := 777
	for i := 0; i < n; i++ {
		if i%2 == 1 {
			fmt.Print(angka, "\t")
			angka++
		} else {
			fmt.Print(x, "\t")
		}
	}
}
