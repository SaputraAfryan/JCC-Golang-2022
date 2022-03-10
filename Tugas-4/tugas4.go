package main

import "fmt"

func main() {
	// soal 1
	fmt.Println("Soal 1")
	for i := 1; i < 21; i++ {
		if i%3 == 0 && i%2 == 1 {
			fmt.Printf("%d - I Love Coding\n", i)
		} else if i%2 == 1 {
			fmt.Printf("%d - JCC\n", i)
		} else if i%2 == 0 {
			fmt.Printf("%d - Candradimuka\n", i)
		}
	}

	// soal 2
	fmt.Printf("\nSoal 2\n")
	for i := 1; i <= 7; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("#")
		}
		fmt.Println("")
	}

	// soal 3
	fmt.Printf("\nSoal 3\n")
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}

	fmt.Println(kalimat[2:])

	// soal 4
	fmt.Printf("\nSoal 4\n")
	var sayuran = []string{}

	sayuran = append(sayuran,
		"Bayam",
		"Buncis",
		"Kangkung",
		"Kubis",
		"Seledri",
		"Tauge",
		"Timun",
	)

	for i := 0; i < len(sayuran); i++ {
		fmt.Printf("%d. %s\n", i+1, sayuran[i])
	}

	// soal 5
	fmt.Printf("\nSoal 5\n")
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}

	volume := satuan["panjang"] * satuan["lebar"] * satuan["tinggi"]

	for key, elm := range satuan {
		fmt.Printf("%s = %d\n", key, elm)
	}

	fmt.Printf("Volume Balok = %d\n", volume)
}
