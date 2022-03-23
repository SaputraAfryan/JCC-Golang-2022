package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type NilaiMahasiswa struct {
	ID          uint   `json:"id"`
	Nama        string `json:"nama"`
	MataKuliah  string `json:"matakuliah"`
	Nilai       uint   `json:"nilai"`
	IndeksNilai string `json:"indeks"`
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}

func nilaiToIndeks(nilai uint) (indeks string) {
	if nilai >= 80 {
		indeks = "A"
	} else if nilai >= 70 && nilai < 80 {
		indeks = "B"
	} else if nilai >= 60 && nilai < 70 {
		indeks = "C"
	} else if nilai >= 50 && nilai < 60 {
		indeks = "D"
	} else if nilai < 50 {
		indeks = "E"
	}
	return
}

func giveID() int {
	return len(nilaiNilaiMahasiswa) + 1
}

// PostNilai
func PostNilai(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var mhs NilaiMahasiswa
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			// parse dari json
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&mhs); err != nil {
				log.Fatal(err)
			}
			mhs.ID = uint(giveID())
			mhs.IndeksNilai = nilaiToIndeks(uint(mhs.Nilai))

			nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, mhs)
		} else {
			// parse dari form
			nama := r.PostFormValue("nama")
			mataKuliah := r.PostFormValue("matakuliah")
			getNilai := r.PostFormValue("nilai")
			nilai, _ := strconv.Atoi(getNilai)
			indeks := nilaiToIndeks(uint(nilai))
			id := giveID()
			mhs = NilaiMahasiswa{
				ID:          uint(id),
				Nama:        nama,
				MataKuliah:  mataKuliah,
				Nilai:       uint(nilai),
				IndeksNilai: indeks,
			}
			nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, mhs)
		}
		dataMahasiswa, _ := json.Marshal(mhs)
		w.Write(dataMahasiswa)
		return
	}
	http.Error(w, "NOT FOUND", http.StatusNotFound)
	return
}

// Post Authentication
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uname, pwd, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("Username atau Password tidak boleh kosong"))
			return
		}

		if uname == "admin" && pwd == "admin" {
			next.ServeHTTP(w, r)
			return
		}
		w.Write([]byte("Username atau Password tidak sesuai"))
		return
	})
}

// GetNilai
func getNilai(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		dataMahasiswa, err := json.Marshal(nilaiNilaiMahasiswa)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(dataMahasiswa)
		return
	}
	http.Error(w, "ERROR....", http.StatusNotFound)
}

func main() {
	// konfigurasi server
	server := &http.Server{
		Addr: ":8080",
	}

	// routing
	http.Handle("/post_nilai", Auth(http.HandlerFunc(PostNilai)))
	http.HandleFunc("/get_nilai", getNilai)

	// jalankan server
	fmt.Println("server running at http://localhost:8080")
	server.ListenAndServe()
}
