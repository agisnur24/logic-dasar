package logic01

import "fmt"

func Logic01Soal09(n int) {
	a := 1
	b := 2
	c := 3

	for i := 1; i < n; i++ {
		if i%3 == 2 {
			fmt.Print(b, "\t")
			b *= 2
		} else if i%3 == 0 {
			fmt.Print(c, "\t")
			c *= 3
		} else {
			fmt.Print(a, "\t")
		}
	}
}
