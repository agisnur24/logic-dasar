package latihan

import (
	"tugas-logic/array_f"
)

func Logic04Soal05(n int) {
	// create array
	array := array_f.NewNumberArray(n, n)
	// initial angka dari 1
	angka := 1
	// for untuk sisi dari 1 s.d
	for sisi := 1; sisi <= 4; sisi++ {
		for a := 1; a < n; a++ {
			if sisi == 1 { // sisi kiri dari bawah ke atas
				array[n-1-a][0] = int32(angka)
			} else if sisi == 2 { //sisi atas dari kiri kanan
				array[0][a] = int32(angka)
			} else if sisi == 3 { //sisi kanan dari atas ke bawah
				array[a][n-1] = int32(angka)
			} else if sisi == 4 { // sisi bawah dari kanan ke kiri
				array[n-1][n-1-a] = int32(angka)
			}
			angka++
		}
	}

	// print array
	array_f.PrintNumberArrayZero(array)
}

func Sample05Step01(n int) {
	kotak := array_f.NewNumberArray(n, n)

	angka := 1
	for a := 1; a < n; a++ {
		kotak[n-1-a][0] = int32(angka)
		angka++
	}
	array_f.PrintNumberArrayZero(kotak)
}

func Sample05Step02(n int) {
	kotak := array_f.NewNumberArray(n, n)

	angka := 1
	// buat untuk diulang sebanyak 4 sisi
	for sisi := 1; sisi <= 4; sisi++ {
		// isi array di ulang sebanyak sisi
		for a := 1; a < n; a++ {
			kotak[n-1-a][0] = int32(angka)
			angka++
		}
		// isi array selesai
	}
	array_f.PrintNumberArrayZero(kotak)
}

func Logic04Soal06(n int) {
	array := array_f.NewNumberArray(n, n)

	angka := 1

	for sisi := 1; sisi <= 4; sisi++ {
		for a := 1; a < n; a++ {
			if sisi == 1 { // sisi kiri dari atas ke bawah
				array[a][0] = int32(angka)
			} else if sisi == 2 { //sisi bawah dari kiri ke kanan
				array[n-1][a] = int32(angka)
			} else if sisi == 3 { //sisi kanan dari bawah ke atas
				array[n-1-a][n-1] = int32(angka)
			} else if sisi == 4 { // sisi atas dari kanan ke kiri
				array[0][n-1-a] = int32(angka)
			}
			angka++
		}
	}

	array_f.PrintNumberArrayZero(array)
}
