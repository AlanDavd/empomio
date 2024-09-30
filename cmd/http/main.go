package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/alandavd/empomio/internal/config"
	"github.com/alandavd/empomio/internal/adapters/repositories"
	"github.com/alandavd/empomio/internal/adapters/repositories/postgres"
	"github.com/alandavd/empomio/internal/core/services"
	"github.com/alandavd/empomio/internal/adapters/handlers"
)

func main() {
	log.Println("Starting server")

	log.Println("Loading configuration")
	config := config.LoadConfigFlags()

	log.Println("Loading database")
	conn, err := postgres.NewPostgresDB(context.Background(), config)
	if err != nil {
		log.Fatalf("could not connect to database: %s\n", err)
	}
	_= conn

	log.Println("Loading repositories")
	userRepo, err := repositories.NewUserRepository()
	if err != nil {
		log.Fatalf("could not load user repo: %s\n", err)
	}

	log.Println("Loading services")
	userService := services.NewUserService(userRepo)

	log.Println("Loading handlers")
	userHandler := handlers.NewUserHandler(*userService)

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	v1 := router.Group("/v1")

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	v1.GET("/users", userHandler.CreateUser)

	server := &http.Server{
		Addr:    config.ServerAddress(),
		Handler: router,
	}

	// Initializing the server in a goroutine so that it won't block the graceful
	// shutdown handling below.
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server listening failed: %s\n", err)
		}
	}()

	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish the request it
	// is currently handling.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
