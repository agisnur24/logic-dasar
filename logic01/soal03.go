package logic01

import "fmt"

func Logic01Soal03(n int) {
	angka := 1
	x := 99
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(angka, "\t")
			angka++
		} else {
			fmt.Print(x, "\t")
		}
	}
}
