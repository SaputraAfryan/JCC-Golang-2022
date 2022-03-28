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

// Delete Category
func DeleteCategory(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idCategory = ps.ByName("id")

	if err := category.Delete(ctx, idCategory); err != nil {
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

// Delete Books
func DeleteBooks(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idCategory = ps.ByName("id")

	if err := books.Delete(ctx, idCategory); err != nil {
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