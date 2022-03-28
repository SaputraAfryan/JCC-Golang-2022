package main

import (
	"Quiz-3/crud"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Middleware Basic Auth
func Auth(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		uname, pwd, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("Username atau Password tidak boleh kosong"))
			return
		}

		if uname == "admin" && pwd == "password" {
			next(w, r, ps)
			return
		} else if uname == "editor" && pwd == "secret" {
			next(w, r, ps)
			return
		} else if uname == "trainer" && pwd == "rahasia" {
			next(w, r, ps)
			return
		}

		w.Write([]byte("Username atau Password tidak sesuai"))
		return
	}
}

func main() {
	router := httprouter.New()

	router.GET("/categories", crud.GetCategory)
	router.POST("/categories", Auth(crud.PostCategory))
	router.PUT("/categories/:id", Auth(crud.PutCategory))
	router.DELETE("/categories/:id", Auth(crud.DeleteCategory))
	router.GET("/categories/:id/books", crud.GetBooksbyCategory)

	router.GET("/books", crud.GetBooks)
	router.POST("/books", Auth(crud.PostBooks))
	router.PUT("/books/:id", Auth(crud.PutBooks))
	router.DELETE("/books/:id", Auth(crud.DeleteBooks))

	fmt.Println("Server Running at Port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
