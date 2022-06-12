package logic01

import "fmt"

func Logic01Soal08(n int) {
	a := 1
	b := 1

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			a *= 2
			fmt.Print(a, "\t")
		} else {
			b *= 5
			fmt.Print(b, "\t")
		}
	}
}
