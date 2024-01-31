package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/zamachnoi/viewthis/db"
)

func main() {
	
	db.Init()
	db.AutoMigrate(db.DB)


	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("welcome"))
    })

	http.ListenAndServe(":3000", r)
}