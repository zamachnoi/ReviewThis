package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/zamachnoi/viewthis/handlers"
	"github.com/zamachnoi/viewthis/lib"
	auth "github.com/zamachnoi/viewthis/middleware"
	// "github.com/danielgtaylor/huma/v2"
	// "github.com/danielgtaylor/huma/v2/adapters/humachi"
)

func main() {
	lib.InitDB()
	lib.InitRD()

	r := chi.NewRouter()

	var allowedOrigins []string = []string{"https://viewthis.app"}
	if os.Getenv("DEV") == "true" {
		allowedOrigins = []string{"http://localhost:3000"}
		log.Println("Running in development mode")
	}

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	})

	r.Use(corsHandler.Handler)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	api := chi.NewRouter()

	// Auth routes
	api.Route("/auth", func(r chi.Router) {
		r.Get("/discord/login", handlers.DiscordAuthLoginHandler)
		r.Get("/discord/callback", handlers.DiscordAuthCallbackHandler)
		r.Get("/discord/logout", handlers.DiscordAuthLogoutHandler)
	})

	// Routes that do not require authentication
	api.Get("/queues", handlers.GetAllQueuesHandler) // Get all queues
	api.Route("/queues/{queueID}/submissions", func(r chi.Router) {
		r.Get("/", handlers.GetSubmissionsByQueueIDHandler)      // Get all submissions for a queue
		r.Post("/", handlers.CreateSubmissionHandler)
		r.Delete("/{id}", handlers.DeleteSubmissionByIDHandler)  // Delete submission by ID
		r.Patch("/{id}", handlers.UpdateSubmissionHandler)
		r.Get("/{id}", handlers.GetSubmissionByIDHandler)        // Get submission by ID
	})


	// Group routes that require JWT authentication
	api.Route("/protected", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(auth.JWTAuthMiddleware)
	
			r.Route("/queues", func(r chi.Router) {
				r.Post("/", handlers.CreateQueueHandler)               // Create a new queue
				r.Delete("/{id}", handlers.DeleteQueueHandler)         // Delete queue by ID
				r.Get("/{id}", handlers.GetQueueByIDHandler)           // Get queue by ID
				r.Patch("/{id}", handlers.UpdateQueueHandler)          // Update queue by ID
				r.Patch("/{id}/clear", handlers.ClearQueueByIDHandler) // Clear queue by ID
			})
	
			// Add more routes here, they will all be under /protected and use the JWT auth middleware
		})
	})


	r.Mount("/api", api)

	log.Println("Server starting on port 3001...")
	if err := http.ListenAndServe(":3001", r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
