package logic01

import "fmt"

func Logic01Soal07(n int) {
	//Nilai statis
	s := 4
	//Nilai kelipatan 3
	kt := 0
	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			fmt.Print(s, "\t")
		} else {
			kt += 3
			fmt.Print(kt, "\t")
		}
	}
}
