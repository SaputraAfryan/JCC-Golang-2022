package main

import (
	"fmt"
	"math"
	"strconv"
	"sync"
	"time"
)

// soal 1
func printPhones(index int, text string, wg *sync.WaitGroup) {
	fmt.Printf("%d. %s\n", index, text)
	time.Sleep(time.Second)
	wg.Done()
}

// soal 2
func getMovies(mov chan<- string, list ...string) {
	var text string
	mov <- "List Movies:"
	for key, elm := range list {
		text = strconv.Itoa(key+1) + ". " + elm
		mov <- text
	}
	close(mov)
}

// soal 3
func luasLingkaran(r []int, ch chan<- float64) {
	for _, elm := range r {
		ch <- math.Pi * math.Pow(float64(elm), 2)
	}
	close(ch)
}
func kelilingLingkaran(r []int, ch chan<- float64) {
	for _, elm := range r {
		ch <- 2 * math.Pi * float64(elm)
	}
	close(ch)
}
func volumeTabung(r []int, ch chan<- float64) {
	for _, elm := range r {
		ch <- math.Pi * math.Pow(float64(elm), 2) * 10
	}
	close(ch)
}

// soal 4
func luasPersegiPanjang(p int, l int, lc chan int) {
	luas := p * l
	lc <- luas
	close(lc)
}

func kelilingPersegiPanjang(p int, l int, kc chan int) {
	keliling := (p + l) * 2
	kc <- keliling
	close(kc)
}
func volumeBalok(p int, l int, t int, vc chan int) {
	volume := p * l * t
	vc <- volume
	close(vc)
}
func main() {
	// soal 1
	var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}

	var wg sync.WaitGroup

	for key, elm := range phones {
		wg.Add(1)
		go printPhones(key+1, elm, &wg)
		wg.Wait()
	}

	// soal 2
	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies...)

	for value := range moviesChannel {
		fmt.Println(value)
	}

	// soal 3
	var r = []int{8, 14, 20}

	var ch1 = make(chan float64)
	go luasLingkaran(r, ch1)

	var ch2 = make(chan float64)
	go kelilingLingkaran(r, ch2)

	var ch3 = make(chan float64)
	go volumeTabung(r, ch3)

	fmt.Printf("Luas Lingkaran: \n")
	for value := range ch1 {
		fmt.Printf("%2.f\n", value)
	}
	fmt.Printf("Keliling Lingkaran: \n")
	for value := range ch2 {
		fmt.Printf("%2.f\n", value)
	}
	fmt.Printf("Volume Tabung: \n")
	for value := range ch3 {
		fmt.Printf("%2.f\n", value)
	}

	// soal 4
	var lc = make(chan int)
	go luasPersegiPanjang(12, 6, lc)

	var kc = make(chan int)
	go kelilingPersegiPanjang(12, 6, kc)

	var vc = make(chan int)
	go volumeBalok(12, 6, 9, vc)

	for i := 0; i < 3; i++ {
		select {
		case luas := <-lc:
			fmt.Printf("Luas Persegi Panjang: %d\n", luas)
		case keliling := <-kc:
			fmt.Printf("Keliling Persegi Panjang: %d\n", keliling)
		case volume := <-vc:
			fmt.Printf("Volume Balok: %d\n", volume)
		}
	}
}
