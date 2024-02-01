package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	"github.com/zamachnoi/viewthis/db"
	"github.com/zamachnoi/viewthis/handlers"
	auth "github.com/zamachnoi/viewthis/middleware"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Init()
	// add middleware to handle spam requests
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	rAuth := r.With(auth.JWTAuthMiddleware)

	r.Route("/users", func(r chi.Router) {
		r.Get("/{id}", handlers.GetUserByIDHandler) // Get user by ID
		r.Post("/", handlers.CreateUserHandler)    // Create user
	})

	rAuth.Route("/delete", func(r chi.Router) {
        r.Delete("/submissions", handlers.DeleteAllSubmissionsHandler)
        r.Delete("/feedback", handlers.DeleteAllFeedbackHandler)
        r.Delete("/queues", handlers.DeleteAllQueuesHandler)
        r.Delete("/users", handlers.DeleteAllUsersHandler)
        r.Delete("/data", handlers.DeleteAllDataHandler)
    })

	r.Route("/queues", func(r chi.Router) {
        r.Get("/", handlers.GetAllQueuesHandler) // Get all queues
        r.Post("/", handlers.CreateQueueHandler) // Create a new queue
        r.Patch("/{id}", handlers.UpdateQueueHandler) // Update queue by ID
        r.Delete("/{id}", handlers.DeleteQueueHandler) // Delete queue by ID
		r.Patch("/{id}/clear", handlers.ClearQueueByIDHandler) // Clear queue by ID
    })

	r.Route("/queues/{queueID}/submissions", func(r chi.Router) {
		r.Get("/", handlers.GetSubmissionsByQueueIDHandler) // Get all submissions for a queue
		r.Post("/", handlers.CreateSubmissionHandler) // Create a new submission
		r.Delete("/{id}", handlers.DeleteSubmissionByIDHandler) // Delete submission by ID
		r.Patch("/{id}", handlers.UpdateSubmissionHandler) // Update submission by ID
	})

	r.Route("/auth/discord", func(r chi.Router) {
		r.Get("/login", handlers.DiscordAuthLoginHandler)
		r.Get("/callback", handlers.DiscordAuthCallbackHandler)
	})

	rAuth.Route("/test", func(r chi.Router) {
		r.Get("/", handlers.TestingHandler)
	})
	
	log.Println("Server starting on port 3001...")
	if err := http.ListenAndServe(":3001", r); err != nil {
        log.Fatalf("Error starting server: %v", err)
	}
}

