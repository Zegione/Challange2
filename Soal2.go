package main

import (
	"fmt"
)

func calculateBMI(berat, tinggi float64) float64 {
	return berat / (tinggi * tinggi)
	/*Fungsi ini mengambil dua parameter: massa (berat dalam kilogram) dan tinggi (tinggi dalam meter).
	Fungsi ini menghitung BMI (Indeks Massa Tubuh) dengan menggunakan rumus: BMI = massa / (tinggi * tinggi).
	Fungsi ini mengembalikan BMI yang dihitung sebagai nilai float64.*/
}

func main() {

	// Data 1
	BeratMark1 := 78.0
	TinggiMark1 := 1.69

	BeratJohn1 := 92.0
	TinggiJohn1 := 1.95

	// Calculate BMIs
	markBMI1 := calculateBMI(BeratMark1, TinggiMark1)
	johnBMI1 := calculateBMI(BeratJohn1, TinggiJohn1)

	// Compare BMIs
	markHigherBMI1 := markBMI1 > johnBMI1

	fmt.Printf("Data 1:\n")
	fmt.Printf("Mark BMI: %.2f\n", markBMI1)
	fmt.Printf("John BMI: %.2f\n", johnBMI1)
	fmt.Printf("Apakah BMI Markk lebih tinggi daripada John? %t\n\n", markHigherBMI1)

	// Data 2
	BeratMark2 := 95.0
	TinggiMark2 := 1.88

	BeratJohn2 := 85.0
	TinggiJohn2 := 1.76

	// Calculate BMIs
	markBMI2 := calculateBMI(BeratMark2, TinggiMark2)
	johnBMI2 := calculateBMI(BeratJohn2, TinggiJohn2)

	// Compare BMIs
	markHigherBMI2 := markBMI2 > johnBMI2

	fmt.Printf("Data 2:\n")
	fmt.Printf("Mark's BMI: %.2f\n", markBMI2)
	fmt.Printf("John's BMI: %.2f\n", johnBMI2)
	fmt.Printf("Apakah BMI Markk lebih tinggi daripada John? %t\n", markHigherBMI2)
	/*  - Di dalam fungsi `maIn()`,  Mendefinisikan variabel `BeratMark1`, `TinggiMark1`, `BeratJohn1`, dan `TinggiJohn1`
		untuk menyimpan berat dan tinggi badan Mark dan John untuk Data 1 dan 2.

		kemudian memanggil fungsi `calculateBMI` dua kali,
	memberikan massa dan tinggi badan Mark dan John sebagai argumen untuk menghitung BMI masing-masing untuk Data 1 dan 2.

	Setelah menghitung BMI, membandingkannya menggunakan operator `>` untuk
	menentukan apakah BMI Mark lebih tinggi dari BMI John. Hasilnya disimpan dalam variabel `markHigherBMI1`.

	Terakhir,  mencetak hasilnya termasuk BMI Mark, BMI John,
	dan apakah BMI Mark lebih tinggi dari BMI John untuk Data 1 dan 2.*/
}
