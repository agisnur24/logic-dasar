package logic01

import "fmt"

func Logic01Soal05(n int) {
	a := 1
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			fmt.Print("fizz", "\t")
			a++
		} else if i%5 == 0 {
			fmt.Print("buzz", "\t")
			a++
		} else if i%3 == 0 && i%5 == 0 {
			fmt.Print("fizz,buzz", "\t")
		} else {
			fmt.Print(a, "\t")
			a++
		}
	}
}
