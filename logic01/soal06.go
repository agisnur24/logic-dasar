package logic01

import "fmt"

func Logic01Soal06(n int) {
	kiri := 1
	iKuadrat := 1
	kanan := 0
	for i := 1; i <= n; i++ {
		if i%4 == 0 {
			iKuadrat = i * i
			fmt.Print(iKuadrat, "\t")
		} else if i <= 7 {
			kiri = i * 3
			fmt.Print(kiri, "\t")
		} else {
			kanan = kiri
			fmt.Print(kanan, "\t")
			kanan = kiri - 3
			kiri = kiri - 3
		}
	}
}
