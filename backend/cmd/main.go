package main

import (
	"fmt"
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
	userController := rest.NewUserController(userService, sessionManager)

	r.Route("/users", func(r chi.Router) {
		r.Post("/login", userController.Login)
		r.Post("/register", userController.Register)
		r.Get("/me", userController.Me)
	})

	// Register word set controller
	learningSetRepo := postgres.NewLearningSetRepo(dbPool)
	learningSetService := service.NewLearningSetService()
	learningSetController := rest.NewLearningSetController(userService, learningSetRepo, learningSetService)

	r.Route("/word-set", func(r chi.Router) {
		r.Post("/", learningSetController.Create)
		r.Get("/", learningSetController.GetAll)
	})

	// Start the server
	fmt.Println("Server starting http://localhost:3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		panic("Failed to start the server")
	}
}
