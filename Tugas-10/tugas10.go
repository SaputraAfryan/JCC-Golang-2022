package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"sort"
	"time"
)

// soal 1

func printSentence(kalimat *string, tahun *int) {
	fmt.Printf("%s %d\n", *kalimat, *tahun)
}

func soal1() {
	kalimat := "Golang Backend Development"
	tahun := 2021
	defer printSentence(&kalimat, &tahun)
}

// soal 2

func recoverKeliling() {
	if r := recover(); r != nil {
		fmt.Printf("Maaf anda belum menginput sisi dari segitiga sama sisi\n")
		fmt.Println("recovered from :", r)
	}
}

func checkError(numbers float64, err bool) (string, error) {
	var output string
	var er error
	switch {
	case numbers > 0 && err:
		output = "keliling segitiga sama sisinya dengan sisi " + fmt.Sprintf("%.1f", numbers/3) + " cm adalah " + fmt.Sprintf("%.1f", numbers) + " cm"
		er = nil
	case numbers > 0 && !err:
		output = fmt.Sprintf("%.1v", numbers)
		er = nil
	case numbers == 0:
		output = fmt.Sprintf("%.1v", numbers)
		er = errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
		if !err {
			panic("Missing input and false condition")
		}
	}
	return output, er
}

func kelilingSegitigaSamaSisi(numbers float64, tf bool) {
	defer recoverKeliling()
	output, err := checkError(numbers, tf)
	if err == nil {
		fmt.Printf("%v\n", output)
	} else {
		fmt.Printf("%v\n", err.Error())
	}
}

func soal2() {

	kelilingSegitigaSamaSisi(4, true)

	kelilingSegitigaSamaSisi(8, false)

	kelilingSegitigaSamaSisi(0, true)

	kelilingSegitigaSamaSisi(0, false)
}

// soal 3

func tambahAngka(numbers int, base *int) {
	*base += numbers
}

func soal3() {
	fmt.Printf("\nSoal 3:\n")
	// deklarasi variabel angka ini simpan di baris pertama func main
	angka := 1

	tambahAngka(7, &angka)

	tambahAngka(6, &angka)

	tambahAngka(-1, &angka)

	tambahAngka(9, &angka)

	fmt.Printf("%d\n", angka)
}

func soal4() {
	fmt.Printf("\nSoal 4:\n")
	var phones = []string{}
	var input = func(phones *[]string, data ...string) {

		for _, value := range data {
			*phones = append(*phones, value)
		}
	}

	input(
		&phones,
		"Xiaomi",
		"Asus",
		"IPhone",
		"Samsung",
		"Oppo",
		"Realme",
		"Vivo",
	)

	var output = func(phones *[]string) {
		sort.Strings(*phones)
		for i, elm := range *phones {
			fmt.Printf("%d. %s\n", i+1, elm)
			time.Sleep(time.Second)
		}
	}

	output(&phones)
}

func soal5() {
	fmt.Printf("\nSoal 5:\n")
	var r = []float64{7, 10, 15}
	var luasLingkaran = func(r float64) float64 {
		return math.Round(math.Phi * math.Pow(2, r))
	}
	var kelilingLingkaran = func(r float64) float64 {
		return math.Round(2 * math.Phi * r)
	}

	fmt.Printf("Luas lingkaran\n")
	for _, value := range r {
		fmt.Printf("r : %.1f = %.1f\n", value, luasLingkaran(value))
	}

	fmt.Printf("\nkeliling lingkaran\n")
	for _, value := range r {
		fmt.Printf("r : %.1f = %.1f\n", value, kelilingLingkaran(value))
	}

}

func soal6() {
	fmt.Printf("\nSoal 6:\n")
	var panjang = flag.Int("panjang", 5, "type length value")
	var lebar = flag.Int("lebar", 7, "type width value")
	var luasPersegiPanjang = func(p int, l int) int {
		return p * l
	}
	var kelilingPersegiPanjang = func(p int, l int) int {
		return (p + l) * 2
	}

	flag.Parse()
	fmt.Printf("Luas Persegi Panjang dengan p : %d dan l : %d = %d\n", *panjang, *lebar, luasPersegiPanjang(*panjang, *lebar))
	fmt.Printf("Keliling Persegi Panjang dengan p : %d dan l : %d = %d\n", *panjang, *lebar, kelilingPersegiPanjang(*panjang, *lebar))

}

func main() {
	defer soal6()
	defer soal5()
	defer soal4()
	defer soal3()
	fmt.Printf("Soal 1:\n")
	soal1()
	fmt.Printf("\nSoal 2:\n")
	soal2()

}
