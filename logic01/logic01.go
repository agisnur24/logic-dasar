package logic01

import (
	"fmt"
)

func Logic01Soal01(n int) {
	angka := 1
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(angka, "\t")
			angka++
		} else {
			fmt.Print(3, "\t")
		}
	}
}

func Logic01Soal02(n int) {
	angka := 1
	for i := 0; i < n; i++ {
		if i%2 == 1 {
			fmt.Print(angka, "\t")
			angka++
		} else {
			fmt.Print(3, "\t")

		}
	}
}

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

func Logic01Soal10(n int) {
	a := 1
	b := 2
	c := 3

	for i := 1; i < n; i++ {
		if i%4 == 1 {
			fmt.Print(c, "\t")
			c += 3
		} else if i%4 == 2 {
			fmt.Print(b, "\t")
			b += 2
		} else if i%4 == 3 {
			fmt.Print(a, "\t")
			a += 1
		} else {
			fmt.Print(999, "\t")
		}
	}
}
