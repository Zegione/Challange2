package main

import "fmt"

// calculateAverage calculates the average score from a slice of scores.
func calculateAverage(scores []int) float64 {
	total := 0
	for _, score := range scores {
		total += score
	}
	return float64(total) / float64(len(scores))
	/*Fungsi calculateAverage, mengambil sepotong bilangan bulat yang mewakili skor sebagai masukan
	dan mengembalikan skor rata-rata sebagai float64. Fungsi ini mengulangi setiap skor dalam potongan, menjumlahkannya,
	dan kemudian membagi total dengan jumlah skor untuk menghitung rata-rata. */
}

// compareScores compares the average scores of two teams and prints the result.
func compareScores(teamName string, average float64, minScore int) {
	if average >= float64(minScore) {
		fmt.Printf("%s wins with an average score of %.2f!\n", teamName, average)
	} else {
		fmt.Printf("%s does not meet the minimum score requirement of %d, so no winner!\n", teamName, minScore)
	}
	/* Fungsi compareScores, mengambil tiga argumen: nama tim, skor rata-rata tim, dan skor minimum yang dibutuhkan untuk menang.
	Fungsi ini mencetak apakah tim menang dengan skor rata-rata atau tidak memenuhi persyaratan skor minimum.*/
}

func main() {

	// Data Uji 1
	ScoreLumbaLumba := []int{96, 108, 89}
	ScoreKoala := []int{88, 91, 110}
	// Di sini, kita mendefinisikan data tes untuk skor dua tim, lumba-lumba dan koala. Skor disimpan dalam irisan bilangan bulat.

	// Menghitung rata rata
	LumbaLumbaAverage := calculateAverage(ScoreLumbaLumba)
	koalasAverage := calculateAverage(ScoreKoala)
	/*memanggil fungsi calculateAverage untuk kedua tim untuk menghitung skor rata-rata mereka.
	Fungsi ini mengambil sepotong skor sebagai masukan dan mengembalikan skor rata-rata sebagai float64.*/

	fmt.Println("===== Data 1 =====")

	// Menampilkan Score Rata-rata
	fmt.Printf("Lumba Lumba average score: %.2f\n", LumbaLumbaAverage)
	fmt.Printf("Koalas average score: %.2f\n", koalasAverage)
	/*Mencetak header "Data 1 " untuk menandakan dimulainya bagian data.
	Kemudian, kita mencetak skor rata-rata dari kedua tim menggunakan fmt.Printf.
	Penentu format %.2f digunakan untuk mencetak angka floating-point dengan dua tempat desimal. */

	// Pengecekan pemenang
	if LumbaLumbaAverage > koalasAverage {
		fmt.Println("Lumba Lumba win!")
	} else if LumbaLumbaAverage < koalasAverage {
		fmt.Println("Koalas win!")
	} else {
		fmt.Println("It's a tie!")

		/*membandingkan skor rata-rata lumba-lumba dan koala untuk menentukan pemenangnya.
		Jika skor rata-rata lumba-lumba lebih besar dari koala,akan mencetak "Lumba-lumba menang!".
		Jika skor rata-rata koala lebih besar,  mencetak "Koala menang!".
		Jika kedua tim memiliki nilai rata-rata yang sama, akan mencetak "Seri!". */
	}

	// Bonus Data 1: Minimum score requirement
	const minimumScore = 100
	compareScores("Lumba Lumba", LumbaLumbaAverage, minimumScore)
	compareScores("Koalas", koalasAverage, minimumScore)
	/*mendefinisikan sebuah konstanta minimumSkor yang mewakili skor minimum yang dibutuhkan untuk menang.
	memanggil fungsi compareScores untuk kedua tim untuk memeriksa apakah skor rata-rata mereka memenuhi persyaratan skor minimum.
	Fungsi ini akan mencetak apakah masing-masing tim menang atau tidak berdasarkan skor minimum. */

	// Bonus Data 2: Comparison
	fmt.Println("===== Data Bonus 1 =====")
	runBonus("Lumba-lumba", []int{97, 112, 101}, []int{109, 95, 123}, minimumScore)
	fmt.Println("===== Data Bonus 2 =====")
	runBonus("Lumba-lumba", []int{97, 112, 101}, []int{109, 95, 106}, minimumScore)
}

// runBonus calculates average scores for two teams and compares them
func runBonus(teamName string, team1Scores, team2Scores []int, minScore int) {
	team1Average := calculateAverage(team1Scores)
	team2Average := calculateAverage(team2Scores)

	fmt.Printf("Average score for Team %s: %.2f\n", teamName, team1Average)
	fmt.Printf("Average score for Team Koala: %.2f\n", team2Average)

	compareScores(teamName, team1Average, minScore)
	compareScores("Koala", team2Average, minScore)

}
