package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/senicko/school-project-backend/pkg/session"
	"github.com/senicko/school-project-backend/pkg/users"
)

func initDatabase() (*pgxpool.Pool, error) {
	dbPool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		return nil, err
	}

	if err := dbPool.Ping(context.Background()); err != nil {
		return nil, err
	}

	return dbPool, nil
}

func initRedis() (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "",
		DB:       0,
	})

	if _, err := redisClient.Ping(redisClient.Context()).Result(); err != nil {
		return nil, err
	}

	return redisClient, nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	redisClient, err := initRedis()
	if err != nil {
		log.Fatalf("failed to init redis client: %v", err)
	}

	dbPool, err := initDatabase()
	defer dbPool.Close()
	if err != nil {
		log.Fatalf("failed to connect with database: %v", err)
	}

	sessionManager := session.NewManager(redisClient)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	userRepo := users.NewRepo(dbPool)
	userService := users.NewService(userRepo)
	userHandlers := users.NewHandler(userService, sessionManager)
	r.Route("/users", userHandlers.Routes)

	// start
	fmt.Println("Server starting http://localhost:3000")

	if err := http.ListenAndServe(":3000", r); err != nil {
		panic("Failed to start the server")
	}
}
