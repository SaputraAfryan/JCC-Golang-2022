package main

import (
	"fmt"
	"strconv"
)

func main() {
	// soal 1
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	pPersegiPanjang, _ := strconv.ParseInt(panjangPersegiPanjang, 10, 64)
	lPersegiPanjang, _ := strconv.ParseInt(lebarPersegiPanjang, 10, 64)
	aSegitiga, _ := strconv.ParseInt(alasSegitiga, 10, 64)
	tSegitiga, _ := strconv.ParseInt(tinggiSegitiga, 10, 64)

	var kelilingPersegiPanjang int = (int(pPersegiPanjang) + int(lPersegiPanjang)) * 2
	var luasSegitiga int = int(aSegitiga) * int(tSegitiga) / 2

	fmt.Println("Soal 1")
	fmt.Printf("Keliling Persegi Panjang = %d\nLLuas Segitiga = %d\n", kelilingPersegiPanjang, luasSegitiga)

	// soal 2
	fmt.Printf("\nSoal 2\n")
	var nilaiJohn = 80
	var nilaiDoe = 50

	if nilaiJohn >= 80 {
		fmt.Println("John = A")
	} else if nilaiJohn >= 70 && nilaiJohn < 80 {
		fmt.Println("John = B")
	} else if nilaiJohn >= 60 && nilaiJohn < 70 {
		fmt.Println("John = C")
	} else if nilaiJohn >= 50 && nilaiJohn < 60 {
		fmt.Println("John = D")
	} else if nilaiJohn < 50 {
		fmt.Println("John = E")
	}

	if nilaiDoe >= 80 {
		fmt.Println("Doe = A")
	} else if nilaiDoe >= 70 && nilaiDoe < 80 {
		fmt.Println("Doe = B")
	} else if nilaiDoe >= 60 && nilaiDoe < 70 {
		fmt.Println("Doe = C")
	} else if nilaiDoe >= 50 && nilaiDoe < 60 {
		fmt.Println("Doe = D")
	} else if nilaiDoe < 50 {
		fmt.Println("Doe = E")
	}

	// soal 3
	fmt.Printf("\nSoal 3\n")
	var tanggal = 21
	var bulan = 4
	var tahun = 2002

	var temp string

	switch bulan {
	case 1:
		temp = "Januari"
	case 2:
		temp = "Februari"
	case 3:
		temp = "Maret"
	case 4:
		temp = "April"
	case 5:
		temp = "Mei"
	case 6:
		temp = "Juni"
	case 7:
		temp = "Juli"
	case 8:
		temp = "Agustus"
	case 9:
		temp = "September"
	case 10:
		temp = "Oktober"
	case 11:
		temp = "November"
	case 12:
		temp = "Desember"
	}

	fmt.Printf("%d %s %d\n", tanggal, temp, tahun)

	// soal 4
	fmt.Printf("\nSoal 4\n")
	if tahun >= 1944 && tahun <= 1964 {
		fmt.Printf("%d, Baby boomer\n", tahun)
	} else if tahun >= 1965 && tahun <= 1979 {
		fmt.Printf("%d, Generasi X\n", tahun)
	} else if tahun >= 1980 && tahun <= 1994 {
		fmt.Printf("%d, Generasi Y (Millenials)\n", tahun)
	} else if tahun >= 1995 && tahun <= 2015 {
		fmt.Printf("%d, Generasi Z\n", tahun)
	}
}
