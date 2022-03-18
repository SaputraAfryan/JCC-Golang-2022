package main

import (
	"errors"
	"fmt"
)

// soal 1

func printSentence(kalimat *string, tahun *int) {
	fmt.Printf("%s %d\n", *kalimat, *tahun)
}

func soal1() {
	kalimat := "Candradimuka Jabar Coding Camp"
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
	// deklarasi variabel angka ini simpan di baris pertama func main
	angka := 1

	tambahAngka(7, &angka)

	tambahAngka(6, &angka)

	tambahAngka(-1, &angka)

	tambahAngka(9, &angka)

	fmt.Printf("%d\n", angka)
}

func main() {
	defer soal3()
	defer fmt.Printf("\nSoal 3:\n")
	fmt.Printf("Soal 1:\n")
	soal1()
	fmt.Printf("\nSoal 2:\n")
	soal2()

}
