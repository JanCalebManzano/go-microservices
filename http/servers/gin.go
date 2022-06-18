package servers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/JanCalebManzano/go-microservices/http/handlers"
	"github.com/JanCalebManzano/go-microservices/http/middlewares"
	"github.com/gin-gonic/gin"
)

type ginServer struct {
	router *gin.Engine
	srv    http.Server
}

func NewGinServer(port string) *ginServer {
	r := gin.Default()

	// r.Use(gin.Logger())
	// r.Use(gin.Recovery())
	r.Use(middlewares.SetJSONContentType())

	userRoutes := r.Group("/users")
	{
		userHandler := handlers.NewUserHandler()
		userRoutes.GET("/", userHandler.GetUsers())
		userRoutes.POST("/", userHandler.ValidateUser(), userHandler.AddUser())
		userRoutes.PUT("/:id", userHandler.ValidateUser(), userHandler.UpdateUser())
	}

	return &ginServer{
		router: r,
		srv: http.Server{
			Addr:         fmt.Sprintf(":%s", port),
			Handler:      r,
			IdleTimeout:  120 * time.Second,
			ReadTimeout:  1 * time.Second,
			WriteTimeout: 1 * time.Second,
		},
	}
}

func (s *ginServer) Start() error {
	return s.srv.ListenAndServe()
}

func (s *ginServer) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return s.srv.Shutdown(ctx)
}
