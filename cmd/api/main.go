package main

import (
	"flag"
	"net/http"
	"os"

	"log/slog"

	"github.com/gin-gonic/gin"
	config "github.com/lokicodess/CatalogX/internal/config"
	"github.com/lokicodess/CatalogX/internal/handler"
	"github.com/lokicodess/CatalogX/internal/middleware"
	"github.com/lokicodess/CatalogX/internal/repository/postgres"
	"github.com/lokicodess/CatalogX/internal/service"
	"github.com/lokicodess/CatalogX/pkg/database"
)

func main() {

	var cfg config.Config

	flag.IntVar(&cfg.Port, "port", 8080, "API server port")
	flag.StringVar(&cfg.Env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.DB_Config.Dsn, "dsn", "postgresql://postgres:postgres@localhost:5432/product_catalog", "Data Source Name")

	// TODO: Remove the default token before pushing in PROD
	flag.StringVar(&cfg.JWTSecret, "jwt_secret", "dev-secret", "JWT Secret Token")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &config.Application{
		Config: cfg,
		Logger: logger,
	}

	db, err := database.OpenDB(app, cfg.DB_Config.Dsn)
	if err != nil {
		app.Logger.Error(err.Error())
		os.Exit(1)
	}

	logger.Info("connected database", "port", 5432)

	productRepo := postgres.NewPostgresProductRepository(db)
	productHandler := handler.NewProductHandler(productRepo)

	userRepo := postgres.NewPostgresUserRepository(db)
	userHandler := handler.NewUserHandler(userRepo)

	authService := service.NewAuthService(userRepo, cfg.JWTSecret)
	authHandler := handler.NewAuthHandler(authService)

	defer db.Close()

	r := gin.New()
	r.Use(middleware.LogRequest(app))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Public routes
	r.POST("/auth/register", userHandler.CreateUser)
	r.POST("/auth/login", authHandler.Login)

	// Protected routes
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware(app))
	{
		protected.POST("/products", productHandler.CreateProduct)
		protected.GET("/products", productHandler.ListProducts)
		protected.GET("/products/:id", productHandler.GetProduct)
	}
	logger.Info("starting server", "addr", cfg.Port, "env", cfg.Env)
	if err := r.Run(); err != nil {
		app.Logger.Error(err.Error())
		os.Exit(1)
	}
}
