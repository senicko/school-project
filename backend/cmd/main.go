package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/senicko/school-project-backend/pkg/users"
)

func connectToDatabase() (*pgxpool.Pool, error) {
	dbPool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		return nil, err
	}

	if err := dbPool.Ping(context.Background()); err != nil {
		return nil, err
	}

	return dbPool, nil
}

func main() {
	// load .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// connect with database
	dbPool, err := connectToDatabase()
	defer dbPool.Close()

	if err != nil {
		log.Fatalf("Failed to connect with database: %v", err)
	}

	// init chi router
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	userRepo := users.NewRepo(dbPool)
	userService := users.NewService(userRepo)
	userHandlers := users.NewHandler(userService)
	r.Route("/users", userHandlers.Routes)

	// start
	fmt.Println("Server starting http://localhost:3000")

	if err := http.ListenAndServe(":3000", r); err != nil {
		panic("Failed to start the server")
	}
}
