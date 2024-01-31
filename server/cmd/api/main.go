package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/zamachnoi/viewthis/db"
	"github.com/zamachnoi/viewthis/handlers"
)

func main() {
	
	db.Init()
	db.AutoMigrate(db.DB)


	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/users", func(r chi.Router) {
		r.Get("/{id}", handlers.UserHandler) // Get user by ID
		r.Post("/", handlers.UserHandler)    // Create user
	})

	r.Route("/delete", func(r chi.Router) {
        r.Delete("/submissions", handlers.DeleteAllSubmissionsHandler)
        r.Delete("/feedback", handlers.DeleteAllFeedbackHandler)
        r.Delete("/queues", handlers.DeleteAllQueuesHandler)
        r.Delete("/users", handlers.DeleteAllUsersHandler)
        r.Delete("/data", handlers.DeleteAllDataHandler)
    })

	http.ListenAndServe(":3000", r)
}