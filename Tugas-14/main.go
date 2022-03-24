package main

import (
	"Tugas-14/config"
	"Tugas-14/mahasiswa"
	"Tugas-14/models"
	"Tugas-14/utils"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

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

// Create
// PostMahasiswa
func PostMahasiswa(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var mhs models.Mahasiswa
	if err := json.NewDecoder(r.Body).Decode(&mhs); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	mhs.IndeksNilai = nilaiToIndeks(uint(mhs.Nilai))
	if err := mahasiswa.Insert(ctx, mhs); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Read
// GetMahasiswa
func GetMahasiswa(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	mhs, err := mahasiswa.GetAll(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, mhs, http.StatusOK)
}

// Update
// UpdateMahasiswa
func UpdateMahasiswa(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var mhs models.Mahasiswa

	if err := json.NewDecoder(r.Body).Decode(&mhs); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	mhs.IndeksNilai = nilaiToIndeks(uint(mhs.Nilai))
	var idMahasiswa = ps.ByName("id")

	if err := mahasiswa.Update(ctx, mhs, idMahasiswa); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Delete
// DeleteMahasiswa
func DeleteMahasiswa(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var idMahasiswa = ps.ByName("id")
	if err := mahasiswa.Delete(ctx, idMahasiswa); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusOK)
}

func main() {
	db, e := config.MySQL()

	if e != nil {
		log.Fatal(e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}

	fmt.Println("Success")

	router := httprouter.New()
	router.GET("/mahasiswa", GetMahasiswa)
	router.POST("/mahasiswa/create", PostMahasiswa)
	router.PUT("/mahasiswa/:id/update", UpdateMahasiswa)
	router.DELETE("/mahasiswa/:id/delete", DeleteMahasiswa)

	fmt.Println("Server Running at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
