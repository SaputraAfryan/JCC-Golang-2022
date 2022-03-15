package main

import (
	"fmt"
)

func soal1() {
	var luasLigkaran float64
	var kelilingLingkaran float64

	var r float64
	r = 7

	var luas = func(lingkaran *float64, r float64) {
		*lingkaran = r * r * 3.14
	}
	var keliling = func(lingkaran *float64, r float64) {
		*lingkaran = 2 * 3.14 * r
	}
	fmt.Printf("luas before : %f\n", luasLigkaran)
	luas(&luasLigkaran, r)
	fmt.Printf("luas after dengan r = %f : %f\n\n", r, luasLigkaran)

	fmt.Printf("keliling before : %f\n", kelilingLingkaran)
	keliling(&kelilingLingkaran, r)
	fmt.Printf("keliling after dengan r = %f : %f\n", r, kelilingLingkaran)
}

func soal2() {
	var sentence string
	var introduce = func(sentence *string, name string, gender string, occupation string, age string) {
		if gender == "laki-laki" {
			*sentence = "\"Pak " + name + " adalah seorang " + occupation + " yang berusia " + age + " tahun\""
		} else if gender == "perempuan" {
			*sentence = "\"Bu " + name + " adalah seorang " + occupation + " yang berusia " + age + " tahun\""
		}
	}

	introduce(&sentence, "John", "laki-laki", "penulis", "30")

	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
	introduce(&sentence, "Sarah", "perempuan", "model", "28")

	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"
}

func soal3() {
	var buah = []string{}
	var newBuah *[]string = &buah

	var inputData = func(fruits *[]string, buah ...string) []string {
		for _, elem := range buah {
			*fruits = append(*fruits, elem)
		}
		return *fruits
	}
	buah = inputData(newBuah, "Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat")

	for i := 0; i < len(buah); i++ {
		fmt.Printf("%d. %s\n", i+1, buah[i])
	}
}

func soal4() {
	var dataFilm = []map[string]string{}

	var tambahDataFilm = func(tittle string, duration string, genre string, year string, movies *[]map[string]string) {
		data := map[string]string{}
		data["genre"] = genre
		data["jam"] = duration
		data["tahun"] = year
		data["tittle"] = tittle
		*movies = append(*movies, data)
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	// isi dengan jawaban anda untuk menampilkan data
	for i := 0; i < len(dataFilm); i++ {
		fmt.Printf("%d. ", i+1)
		fmt.Printf("tittle : %s\n", dataFilm[i]["tittle"])
		fmt.Printf("   duration : %s\n", dataFilm[i]["jam"])
		fmt.Printf("   genre : %s\n", dataFilm[i]["genre"])
		fmt.Printf("   year : %s\n\n", dataFilm[i]["tahun"])
	}
}

func main() {
	fmt.Printf("Soal 1:\n")
	soal1()
	fmt.Printf("\nSoal 2:\n")
	soal2()
	fmt.Printf("\nSoal 3:\n")
	soal3()
	fmt.Printf("\nSoal 4:\n")
	soal4()
}
