package crud

import (
	"Quiz-3/books"
	"Quiz-3/category"
	"Quiz-3/models"
	"Quiz-3/utils"
	"context"
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/julienschmidt/httprouter"
)

// Put Category
func PutCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var ctg models.Category

	if err := json.NewDecoder(r.Body).Decode(&ctg); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idCategory = ps.ByName("id")

	if err := category.Update(ctx, ctg, idCategory); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Put Books
func PutBooks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var book models.Books

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	regex, _ := regexp.Compile(`/^(https:|http:|www\.)\S*/`)

	if !regex.MatchString(book.Image_url) {
		http.Error(w, "not a valid url!", http.StatusBadRequest)
		return
	} else if book.Release_Year < 1980 || book.Release_Year > 2021{
		http.Error(w, "year not valid!", http.StatusBadRequest)
		return
	} else if !regex.MatchString(book.Image_url) && book.Release_Year < 1980 || book.Release_Year > 2021{
		http.Error(w, "year and url not valid!", http.StatusBadRequest)
		return
	}

	book.Thickness = thick(int(book.Total_Page))

	var idBook = ps.ByName("id")

	if err := books.Update(ctx, book, idBook); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}