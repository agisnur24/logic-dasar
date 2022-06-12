package logic01

import "fmt"

func Logic01Soal10(n int) {
	a := 1
	b := 2
	c := 3

	for i := 1; i <= n; i++ {
		if i%4 == 1 {
			fmt.Print(c, " | ")
			c += 3
		} else if i%4 == 2 {
			fmt.Print(b, " | ")
			b += 2
		} else if i%4 == 3 {
			fmt.Print(a, " | ")
			a += 1
		} else {
			fmt.Print(999, " | ")
		}
	}
}
