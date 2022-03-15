package main

import (
	"fmt"
)

func soal1() {
	type buah struct {
		nama       string
		warna      string
		adaBijinya bool
		harga      int
	}

	var inputBuah = func(data *[]buah, name string, colour string, beans bool, price int) {
		var fruits buah
		fruits.nama = name
		fruits.warna = colour
		fruits.adaBijinya = beans
		fruits.harga = price

		*data = append(*data, fruits)
	}

	var fruits []buah

	inputBuah(&fruits, "Nanas", "Kuning", false, 9000)
	inputBuah(&fruits, "Jeruk", "Oranye", true, 8000)
	inputBuah(&fruits, "Semangka", "Hijau & Merah", true, 10000)
	inputBuah(&fruits, "Pisang", "Kuning", false, 5000)

	var printBuah = func(fruits *[]buah) {
		fmt.Printf("Nama    \t|\tWarna      \t|\tAda Bijinya\t|\tHarga\n")
		fruit := *fruits
		for i := 0; i < len(fruit); i++ {
			if fruit[i].adaBijinya == true && fruit[i].nama != "Semangka" {
				beans := "Ada"
				fmt.Printf("%s    \t|\t%s      \t|\t%s\t\t|\t%d\n", fruit[i].nama, fruit[i].warna, beans, fruit[i].harga)
			} else if fruit[i].adaBijinya == false && fruit[i].nama != "Semangka" {
				beans := "Tidak"
				fmt.Printf("%s    \t|\t%s      \t|\t%s\t\t|\t%d\n", fruit[i].nama, fruit[i].warna, beans, fruit[i].harga)
			} else if fruit[i].adaBijinya == true && fruit[i].nama == "Semangka" {
				beans := "Ada"
				fmt.Printf("%s    \t|\t%s   |\t%s\t\t|\t%d\n", fruit[i].nama, fruit[i].warna, beans, fruit[i].harga)
			}
		}
	}

	printBuah(&fruits)
}

// soal no 2

type segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

func (s segitiga) luasSegitiga() {
	var luas int
	luas = s.alas * s.tinggi / 2
	fmt.Printf("Luas Segitiga : %d\n", luas)
}

func (p persegi) luasPersegi() {
	var luas int
	luas = p.sisi * p.sisi
	fmt.Printf("Luas Persegi : %d\n", luas)
}

func (pp persegiPanjang) luasPersegiPanjang() {
	var luas int
	luas = pp.panjang * pp.lebar
	fmt.Printf("Luas Persegi Panjang : %d\n", luas)
}

func soal2() {
	var s3 = segitiga{3, 3}
	s3.luasSegitiga()
	var pss = persegi{4}
	pss.luasPersegi()
	var pp = persegiPanjang{5, 4}
	pp.luasPersegiPanjang()
}

// soal no 3

type phone struct {
	name, brand string
	year        int
	colors      []string
}

func (s *phone) addColors(warna ...string) {
	for _, elem := range warna {
		s.colors = append(s.colors, elem)
	}
}

func soal3() {
	var hp phone
	hp.name = "Galaxy S20"
	hp.brand = "Samsung"
	hp.year = 2020
	hp.addColors("Ocean Blue", "Gunmetal Grey")

	fmt.Printf("1. %s %s built %d with color option : ", hp.brand, hp.name, hp.year)
	for i := 0; i < len(hp.colors); i++ {
		if i != len(hp.colors)-1 {
			fmt.Printf("%s, ", hp.colors[i])
		} else {
			fmt.Printf("%s\n", hp.colors[i])
		}
	}
}

// soal no 4

type movie struct {
	title, genre   string
	duration, year int
}

func soal4() {
	var dataFilm = []movie{}

	var tambahDataFilm = func(tittle string, duration int, genre string, year int, movies *[]movie) {
		var data movie
		data.genre = genre
		data.duration = duration
		data.year = year / 2
		data.title = tittle
		*movies = append(*movies, data)
	}

	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	// isi dengan jawaban anda untuk menampilkan data
	for i := 0; i < len(dataFilm); i++ {
		fmt.Printf("%d. ", i+1)
		fmt.Printf("tittle : %s\n", dataFilm[i].title)
		fmt.Printf("   duration : %d jam\n", dataFilm[i].duration)
		fmt.Printf("   genre : %s\n", dataFilm[i].genre)
		fmt.Printf("   year : %d\n\n", dataFilm[i].year)
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
