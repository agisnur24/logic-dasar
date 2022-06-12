package logic02

import "fmt"

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
