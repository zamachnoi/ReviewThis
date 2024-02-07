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
)

func main() {
	lib.InitDB()
	lib.InitRD()
	// CORS middleware to handle cross-origin requests
	r := chi.NewRouter()

	if os.Getenv("DEV") == "true" {
		log.Println("Running in development mode")
		corsHandler := cors.New(cors.Options{
			AllowedOrigins:   []string{"http://localhost:3000"}, // replace with your frontend's origin
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			AllowCredentials: true,
			MaxAge:           300, // Maximum age for the preflight request
		})
	
		// Use the CORS handler
		r.Use(corsHandler.Handler)
	}
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Prefix all routes with "/api"
	api := chi.NewRouter()

	api.Route("/queues/{queueID}/submissions", func(r chi.Router) {
		r.Get("/", handlers.GetSubmissionsByQueueIDHandler)      // Get all submissions for a queue
		r.Post("/", handlers.CreateSubmissionHandler)
		r.Delete("/{id}", handlers.DeleteSubmissionByIDHandler)  // Delete submission by ID
		r.Patch("/{id}", handlers.UpdateSubmissionHandler)
		r.Get("/{id}", handlers.GetSubmissionByIDHandler)       // Update submission by ID
	})
	api.Route("/auth", func(r chi.Router) {
		r.Get("/discord/login", handlers.DiscordAuthLoginHandler)
		r.Get("/discord/callback", handlers.DiscordAuthCallbackHandler)
		r.Get("/discord/logout", handlers.DiscordAuthLogoutHandler)
	})
	

    api.Group(func(r chi.Router) {
        r.Use(auth.JWTAuthMiddleware)
        
        r.Route("/queues", func(r chi.Router) {
            r.Post("/", handlers.CreateQueueHandler)            // Create a new queue
            r.Delete("/{id}", handlers.DeleteQueueHandler)  // Delete queue by ID
            r.Get("/{id}", handlers.GetQueueByIDHandler)         // Get queue by ID
            r.Patch("/{id}", handlers.UpdateQueueHandler)        // Update queue by ID
            r.Patch("/{id}/clear", handlers.ClearQueueByIDHandler)  // Clear queue by ID TODO: setup check jwt for this
        })
	})

	api.Get("/queues", handlers.GetAllQueuesHandler) // Get all queues


	// Mount the API router under "/api" prefix
	r.Mount("/api", api)

	log.Println("Server starting on port 3001...")
	if err := http.ListenAndServe(":3001", r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

// type loggingResponseWriter struct {
//     http.ResponseWriter
//     statusCode int
//     body       bytes.Buffer
// }

// func (lrw *loggingResponseWriter) WriteHeader(statusCode int) {
//     lrw.statusCode = statusCode
//     lrw.ResponseWriter.WriteHeader(statusCode)
// }

// func (lrw *loggingResponseWriter) Write(data []byte) (int, error) {
//     lrw.body.Write(data)
//     return lrw.ResponseWriter.Write(data)
// }

// func loggingMiddleware(next http.Handler) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//         lrw := &loggingResponseWriter{ResponseWriter: w}

//         next.ServeHTTP(lrw, r)

//         // Log the response body and status code
//         log.Printf("Response body: %s\n", lrw.body.String())
//         log.Printf("Response status code: %d\n", lrw.statusCode)
//     })
// }