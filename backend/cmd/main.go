package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/senicko/school-project-backend/pkg/users"
)

func connectToDatabase() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		return nil, err
	}

	if err := conn.Ping(context.Background()); err != nil {
		return nil, err
	}

	return conn, nil
}

func main() {
	// load .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// connect with database
	conn, err := connectToDatabase()

	if err != nil {
		log.Fatalf("Failed to connect with database: %v", err)
	}

	// init chi router
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	// register user repo
	userRepo := users.NewUserRepo(conn)
	userHandlers := users.NewUserHandlers(userRepo)

	r.Post("/users", userHandlers.RegisterUser)

	fmt.Println("Server starting http://localhost:3000")

	if err := http.ListenAndServe(":3000", r); err != nil {
		panic("Failed to start the server")
	}
}
