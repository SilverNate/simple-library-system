package main

import (
	"github.com/go-chi/chi/v5"
	"library-system/config"
	"library-system/middlewares"
	"library-system/repositories"
	"library-system/services"
	"log"
	"net/http"
	"os"

	"library-system/handlers"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	db := config.InitializeDB()

	// Initialize repositories, services, and handlers
	historyRepo := repositories.NewHistoryRepository(db)
	historyService := services.NewHistoryService(historyRepo)
	historyHandler := handlers.NewHistoryHandler(historyService)

	bookRepo := repositories.NewBookRepository(db)
	bookService := services.NewBookService(bookRepo)
	bookHandler := handlers.NewBookHandler(bookService)

	borrowerRepo := repositories.NewBorrowerRepository(db)
	borrowerService := services.NewBorrowerService(borrowerRepo)
	borrowerHandler := handlers.NewBorrowerHandler(borrowerService)

	r := chi.NewRouter()
	//r.Use(middleware.Logger)
	r.Use(middlewares.LogrusMiddleware) // we use sirupsen logger as middlewares

	// use this api to get token
	r.Post("/token", handlers.LoginHandler)

	r.Group(func(r chi.Router) {
		r.Use(middlewares.JWTMiddleware)
		// Register routes
		r.Post("/borrow_history", historyHandler.CreateHistory)
		r.Put("/borrow_history/return", historyHandler.UpdateReturnDate)
		r.Get("/borrow_history", historyHandler.GetHistory)
		r.Get("/borrow_history/overdue", historyHandler.GetOverdueBooks)
		r.Get("/borrow_history/most_borrowed", historyHandler.FindMostBorrowedBooks)

		r.Post("/books", bookHandler.CreateBook)

		r.Post("/borrowers", borrowerHandler.CreateBorrower)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Database Postgresql connected")
	log.Println("Server is running on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
