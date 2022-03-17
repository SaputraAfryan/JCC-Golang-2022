package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func (l segitigaSamaSisi) luas() int {
	return l.alas * l.tinggi / 2
}

func (l persegiPanjang) luas() int {
	return l.panjang * l.lebar
}

func (k segitigaSamaSisi) keliling() int {
	return k.alas * 3
}

func (k persegiPanjang) keliling() int {
	return (k.lebar + k.panjang) * 2
}

func (v tabung) volume() float64 {
	return math.Pi * math.Pow(v.jariJari, 2) * v.tinggi
}

func (v balok) volume() float64 {
	return float64(v.panjang * v.lebar * v.tinggi)
}

func (lp tabung) luasPermukaan() float64 {
	return (lp.jariJari + lp.tinggi) * 2 * math.Pi * lp.jariJari
}

func (lp balok) luasPermukaan() float64 {
	pl := lp.panjang * lp.lebar
	pt := lp.panjang * lp.tinggi
	lt := lp.lebar * lp.tinggi
	sum := pl + pt + lt
	return float64(2 * sum)
}

func soal1() {
	var bangunDatar hitungBangunDatar
	var bangunRuang hitungBangunRuang

	bangunDatar = segitigaSamaSisi{10, 5}
	fmt.Printf("===== segitiga sama sisi\n")
	fmt.Printf("luas\t\t: %d\n", bangunDatar.luas())
	fmt.Printf("Keliling\t: %d\n", bangunDatar.keliling())

	bangunDatar = persegiPanjang{12, 6}
	fmt.Printf("===== persegi panjang\n")
	fmt.Printf("luas\t\t: %d\n", bangunDatar.luas())
	fmt.Printf("Keliling\t: %d\n", bangunDatar.keliling())

	bangunRuang = tabung{7, 14}
	fmt.Printf("===== tabung\n")
	fmt.Printf("volume\t\t: %f\n", bangunRuang.volume())
	fmt.Printf("luas permukaan\t: %f\n", bangunRuang.luasPermukaan())

	bangunRuang = balok{2, 3, 5}
	fmt.Printf("===== balok\n")
	fmt.Printf("volume\t\t: %f\n", bangunRuang.volume())
	fmt.Printf("luas permukaan\t: %f\n", bangunRuang.luasPermukaan())
}

//soal 2

type phone struct {
	name, brand string
	year        int
	colors      []string
}

type phones interface {
	printAll() string
}

func (s *phone) printAll() (output string) {
	nama := s.name
	merk := s.brand
	tahun := strconv.Itoa(s.year)
	warna := strings.Join(s.colors, ", ")
	output = "name : " + nama + "\nbrand : " + merk + "\ntahun : " + tahun + "\ncolors : " + warna

	return output
}

func soal2() {
	hp := phone{
		name:  "Samsung Galaxy Note 20",
		brand: "Samsung Galaxy Note 20",
		year:  2020,
	}
	hp.colors = append(hp.colors, "Mystic Bronze", "Mystic White", "Mystic Black")
	fmt.Println(hp.printAll())
}

// Soal 3
func luasPersegi(numbers int, ok bool) (sentence interface{}) {
	var sisi string
	var l string
	switch {
	case numbers > 0 && ok:
		sisi = strconv.Itoa(numbers / 2)
		l = strconv.Itoa(numbers)
		sentence = "luas persegi dengan sisi " + sisi + " cm adalah " + l + " cm"
	case numbers > 0 && !ok:
		sentence = numbers
	case numbers == 0 && ok:
		sentence = "Maaf anda belum mengisi menginput sisi dari persegi"
	case numbers == 0 && !ok:
		sentence = nil
	}
	return sentence
}

func soal3() {

	fmt.Println(luasPersegi(4, true))

	fmt.Println(luasPersegi(8, false))

	fmt.Println(luasPersegi(0, true))

	fmt.Println(luasPersegi(0, false))
}

// Soal 4
func soal4() {
	var prefix interface{} = "hasil penjumlahan dari "

	var kumpulanAngkaPertama interface{} = []int{6, 8}

	var kumpulanAngkaKedua interface{} = []int{12, 14}

	// tulis jawaban anda disini
	s := prefix.(string)
	i := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(kumpulanAngkaPertama.([]int))), " + "), "[]")
	ii := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(kumpulanAngkaKedua.([]int))), " + "), "[]")
	var sum int
	for _, elem := range kumpulanAngkaPertama.([]int) {
		sum += elem
	}
	for _, elem := range kumpulanAngkaKedua.([]int) {
		sum += elem
	}
	fmt.Printf("%q", s+i+" + "+ii+" = "+strconv.Itoa(sum))
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
