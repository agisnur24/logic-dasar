package logic02

import "fmt"

func Logic02Soal01(n int) {
	a := 3

	//looping baris
	for i := 0; i < n; i++ {
		//looping kolom
		for j := 0; j < n; j++ {
			fmt.Print(a, "\t")
		}
		//looping kolom selesai
		fmt.Println("\n")
		//pindah baris selanjutnya
		//dan update nilai a
		a += 3
	}
}

func Logic02Soal02(n int) {
	//looping baris
	for i := 0; i < n; i++ {
		a := 3
		for j := 0; j < n; j++ {
			fmt.Print(a, "\t")
			a += 3
		}
		fmt.Println("\n")
	}
}

func Logic02Soal03(n int) {
	//Looping baris
	for i := 0; i < n; i++ {
		//Setiap kali pindah baris maka a = n * 3
		a := n * 3
		//Looping kolom
		for j := 0; j < n; j++ {
			//Menampilkan a setiap kolom
			fmt.Print(a, "\t")
			//Nilai a dikurangi 3 setiap kali pindah kolom
			a -= 3
		}
		//Pindah baris baru
		fmt.Println("\n")
	}
}

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

func Logic02Soal05(n int) {
	nt := n / 2
	a := 3

	// looping baris
	for i := 0; i < n; i++ {
		// looping kolom
		for j := 0; j < n; j++ {
			if j <= nt {
				fmt.Print(a, "\t")
				a += 3
			} else {
				a -= 3
				fmt.Print(a, "\t")
			}
		}
		fmt.Println("\n")
	}
}

func Logic02Soal06(n int) {
	nt := n / 2

	//Looping baris
	for i := 0; i < n; i++ {
		a := 3

		//Looping kolom
		for j := 0; j < n; j++ {
			if j < nt {
				fmt.Print(a, "\t")
				a += 3
			} else {
				fmt.Print(a, "\t")
				a -= 3
			}

		}
		//Pindah baris selanjutnya
		fmt.Println("\n")
	}
}

func Logic02Soal07(n int) {
	a := 3

	//Looping baris
	for i := 0; i < n; i++ {
		//Looping kolom
		for j := 0; j < n; j++ {
			if i < j {
				fmt.Print(" ", "\t")
			} else {
				fmt.Print(a, "\t")
			}
		}
		//Pindah baris selanjutnya
		fmt.Print("\n")
		//Saat pindah baris variable a ditambah 3
		a += 3
	}
}

func Logic02Soal08(n int) {
	a := 3

	//Looping baris
	for i := 0; i < n; i++ {
		//Looping kolom
		for j := 0; j < n; j++ {
			//Kondisi untuk kotak kosong
			if i > j {
				fmt.Print(" ", "\t")
			} else {
				fmt.Print(a, "\t")
			}
		}
		//Pindah baris selanjutnya
		fmt.Print("\n")
		//nilai variabel a ditambah 3
		a += 3
	}
}

func Logic02Soal09(n int) {

	for i := 0; i < n; i++ {
		a := 0
		for j := 0; j < n; j++ {
			a += 3
			if i+j == n-1 {
				fmt.Print(a, "\t")
			} else if i+j >= n-1 {
				fmt.Print(a, "\t")
			} else {
				fmt.Print(" ", "\t")
			}
		}
		fmt.Print("\n")
	}
}

func Logic02Soal10(n int) {

	for i := 0; i < n; i++ {
		a := 0
		for j := 0; j < n; j++ {
			a += 3
			if i+j == n-1 {
				fmt.Print(a, "\t")
			} else if i+j <= n-1 {
				fmt.Print(a, "\t")
			} else {
				fmt.Print(" ", "\t")
			}
		}
		fmt.Print("\n")
	}
}
