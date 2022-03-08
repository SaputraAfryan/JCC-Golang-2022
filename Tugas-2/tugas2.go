package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// soal 1
	fmt.Printf("Soal 1\n")
	fmt.Println("Jabar Coding Camp Golang 2022")

	// soal 2
	halo := "Halo Dunia"
	halo = strings.Replace(halo, "Dunia", "Golang", 1)
	fmt.Printf("\nSoal 2\n")
	fmt.Println(halo)

	// soal 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	kataKedua = strings.ToTitle(kataKedua)
	kataKetiga = strings.Replace(kataKetiga, "r", "R", 1)
	kataKeempat = strings.ToUpper(kataKeempat)

	fmt.Printf("\nSoal 3\n")
	fmt.Printf("%s %s %s %s\n", kataPertama, kataKedua, kataKetiga, kataKeempat)

	// soal 4
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	pertama, _ := strconv.ParseInt(angkaPertama, 10, 64)
	Kedua, _ := strconv.ParseInt(angkaKedua, 10, 64)
	Ketiga, _ := strconv.ParseInt(angkaKetiga, 10, 64)
	Keempat, _ := strconv.ParseInt(angkaKeempat, 10, 64)

	fmt.Printf("\nSoal 4\n")
	fmt.Println(pertama + Kedua + Ketiga + Keempat)

	// soal 5
	kalimat := "halo halo bandung"
	angka := 2022

	fmt.Printf("\nSoal 5\n")
	fmt.Printf("\"%s\" - %s\n", strings.Replace(kalimat, "halo", "Hi", 2), strconv.Itoa(angka))

}
