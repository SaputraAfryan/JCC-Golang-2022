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

// convert thickness
func thick(jumlah int) (params string) {
	if jumlah <= 100 {
		params = "tipis"
	} else if jumlah >= 101 && jumlah <= 200 {
		params = "sedang"
	} else if jumlah >= 201 {
		params = "tebal"
	}
	return
}

// Post Category
func PostCategory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

	if err := category.Insert(ctx, ctg); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)

}

// Post Books
func PostBooks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

	regex, _ := regexp.Compile(`^(ht|f)tp(s?)\:\/\/[0-9a-zA-Z]([-.\w]*[0-9a-zA-Z])*(:(0-9)*)*(\/?)([a-zA-Z0-9\-\.\?\,\'\/\\\+&amp;%\$#_]*)?$`)

	if regex.MatchString(book.Image_url) {
		http.Error(w, "not a valid url!", http.StatusBadRequest)
		return
	} else if book.Release_Year < 1980 || book.Release_Year > 2021 {
		http.Error(w, "year not valid!", http.StatusBadRequest)
		return
	} else if !regex.MatchString(book.Image_url) && book.Release_Year < 1980 || book.Release_Year > 2021 {
		http.Error(w, "year and url not valid!", http.StatusBadRequest)
		return
	}

	book.Thickness = thick(int(book.Total_Page))

	if err := books.Insert(ctx, book); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)

}
