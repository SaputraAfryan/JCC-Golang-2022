package main

import (
	"fmt"
)

func soal1() {
	var luasPersegiPanjang = func(panjang int, lebar int) int {
		return panjang * lebar
	}

	var kelilingPersegiPanjang = func(panjang int, lebar int) int {
		return 2 * (panjang + lebar)
	}

	var volumeBalok = func(panjang int, lebar int, tinggi int) int {
		return panjang * lebar * tinggi
	}

	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println("Soal 1")
	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)
}

func soal2() {
	fmt.Printf("\nSoal 2\n")

	var introduce = func(name string, gender string, profession string, age string) string {
		var temp string

		if gender == "laki-laki" {
			temp = "Pak " + name + " adalah seorang " + profession + " yang berusia " + age + " tahun"
		} else if gender == "perempuan" {
			temp = "Bu " + name + " adalah seorang " + profession + " yang berusia " + age + " tahun"
		}

		return temp
	}

	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"
}

func soal3() {
	fmt.Printf("\nSoal 3\n")
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}

	var buahFavorit = func(name string, fruits ...string) (temp string) {
		temp = "halo nama saya " + name + " dan buah favorit saya adalah "
		for i := 0; i < len(fruits); i++ {
			if i+1 != len(fruits) {
				temp += "\"" + fruits[i] + "\", "
			} else {
				temp += "\"" + fruits[i] + "\""
			}
		}
		return
	}

	var buahFavoritJohn = buahFavorit("john", buah...)

	fmt.Println(buahFavoritJohn)
	// halo nama saya john dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"
}

func soal4() {
	fmt.Printf("\nSoal 4\n")
	var dataFilm = []map[string]string{}
	// buatlah closure function disini

	var tambahDataFilm = func(title string, duration string, genre string, year string) {
		movie := map[string]string{}
		movie["genre"] = genre
		movie["jam"] = duration
		movie["tahun"] = year
		movie["title"] = title

		dataFilm = append(dataFilm, movie)
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
}

func main() {
	soal1()
	soal2()
	soal3()
	soal4()
}
