package main

import (
	"Tugas-15/mahasiswa"
	"Tugas-15/matakuliah"
	"Tugas-15/models"
	"Tugas-15/nilai"
	"Tugas-15/utils"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/nilai", GetNilai)
	router.GET("/mahasiswa", GetMahasiswa)
	router.GET("/matakuliah", GetMataKuliah)

	router.POST("/nilai/create", PostNilai)
	router.POST("/mahasiswa/create", PostMahasiswa)
	router.POST("/matakuliah/create", PostMataKuliah)

	router.PUT("/nilai/:id/update", UpdateNilai)
	router.PUT("/mahasiswa/:id/update", UpdateMahasiswa)
	router.PUT("/matakuliah/:id/update", UpdateMataKuliah)

	router.DELETE("/nilai/:id/delete", DeleteNilai)
	router.DELETE("/mahasiswa/:id/delete", DeleteMahasiswa)
	router.DELETE("/matakuliah/:id/delete", DeleteMataKuliah)

	fmt.Println("Server Running at Port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

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

// GetMataKuliah
func GetMataKuliah(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	mk, err := matakuliah.GetAll(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, mk, http.StatusOK)
}

// GetNilai
func GetNilai(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	nilai, err := nilai.GetAll(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, nilai, http.StatusOK)
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

	if err := mahasiswa.Insert(ctx, mhs); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)

}

// PostMataKuliah
func PostMataKuliah(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var mk models.MataKuliah

	if err := json.NewDecoder(r.Body).Decode(&mk); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := matakuliah.Insert(ctx, mk); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)

}

// PostNilai
func PostNilai(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var total models.Nilai

	if err := json.NewDecoder(r.Body).Decode(&total); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	total.Indeks = nilaiToIndeks(uint(total.Skor))
	if err := nilai.Insert(ctx, total); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)

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

// UpdateMataKuliah
func UpdateMataKuliah(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var mk models.MataKuliah

	if err := json.NewDecoder(r.Body).Decode(&mk); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idMataKuliah = ps.ByName("id")

	if err := matakuliah.Update(ctx, mk, idMataKuliah); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

// UpdateNilai
func UpdateNilai(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var total models.Nilai

	if err := json.NewDecoder(r.Body).Decode(&total); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	total.Indeks = nilaiToIndeks(uint(total.Skor))
	var idNilai = ps.ByName("id")

	if err := nilai.Update(ctx, total, idNilai); err != nil {
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
func DeleteMahasiswa(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
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

// DeleteMataKuliah
func DeleteMataKuliah(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idMataKuliah = ps.ByName("id")

	if err := matakuliah.Delete(ctx, idMataKuliah); err != nil {
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

// DeleteNilai
func DeleteNilai(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idNilai = ps.ByName("id")

	if err := nilai.Delete(ctx, idNilai); err != nil {
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
