package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/senicko/school-project-backend/pkg/postgres"
	"github.com/senicko/school-project-backend/pkg/rest"
	"github.com/senicko/school-project-backend/pkg/service"
	"github.com/senicko/school-project-backend/pkg/session"
)

// initRedis initializes redis connection.
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
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// connect to redis
	redisClient, err := initRedis()
	if err != nil {
		log.Fatalf("failed to init redis client: %v", err)
	}

	// connect to postgres
	dbPool, err := postgres.Connect()
	defer dbPool.Close()
	if err != nil {
		log.Fatalf("failed to connect with database: %v", err)
	}

	sessionManager := session.NewManager(redisClient)

	// Create new router & register middlewares
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowCredentials: true,
	}))

	// Register user controller
	userRepo := postgres.NewUserRepo(dbPool)
	userService := service.NewUserService(userRepo, sessionManager)
	userController := rest.NewUserController(userRepo, userService, sessionManager)

	r.Route("/users", func(r chi.Router) {
		r.Post("/login", userController.Login)
		r.Post("/register", userController.Register)
		r.Get("/logout", userController.Logout)

		r.Get("/me", userController.Me)
		r.Post("/me/jokes", userController.SaveJoke)
	})

	// api.yomomma.info proxy
	r.Get("/joke", func(w http.ResponseWriter, r *http.Request) {
		res, err := http.Get("https://api.yomomma.info")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(res.StatusCode)
		w.Write(body)
	})

	// Start the server
	fmt.Println("Server starting http://localhost:3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		panic("Failed to start the server")
	}
}
