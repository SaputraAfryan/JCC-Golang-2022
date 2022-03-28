package crud

import (
	"Quiz-3/books"
	"Quiz-3/category"
	"Quiz-3/utils"
	"context"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Get Category
func GetCategory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	categories, err := category.GetAll(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, categories, http.StatusOK)
}

// Get Books
func GetBooks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	title := r.URL.Query().Get("title")
	minYear := r.URL.Query().Get("minYear")
	maxYear := r.URL.Query().Get("maxYear")
	minPage := r.URL.Query().Get("minPage")
	maxPage := r.URL.Query().Get("maxPage")
	sortByTitle := r.URL.Query().Get("sortByTitle")

	book, err := books.GetAll(ctx, title, minYear, maxYear, minPage, maxPage, sortByTitle)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, book, http.StatusOK)
}

// Get Books by Category
func GetBooksbyCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	title := r.URL.Query().Get("title")
	minYear := r.URL.Query().Get("minYear")
	maxYear := r.URL.Query().Get("maxYear")
	minPage := r.URL.Query().Get("minPage")
	maxPage := r.URL.Query().Get("maxPage")
	sortByTitle := r.URL.Query().Get("sortByTitle")

	var idCategory = ps.ByName("id")
	book, err := books.GetBooks(ctx, idCategory, title, minYear, maxYear, minPage, maxPage, sortByTitle)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, book, http.StatusOK)
}
