package main

import (
	"flag"
	"net/http"
	"os"

	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/lokicodess/CatalogX/internal/handler"
	"github.com/lokicodess/CatalogX/internal/middleware"
	"github.com/lokicodess/CatalogX/internal/repository/postgres"
	config "github.com/lokicodess/CatalogX/pkg/config"
	"github.com/lokicodess/CatalogX/pkg/database"
)

func main() {

	var cfg config.Config

	flag.IntVar(&cfg.Port, "port", 8080, "API server port")
	flag.StringVar(&cfg.Env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.DB_Config.Dsn, "dsn", "postgresql://postgres:postgres@localhost:5432/product_catalog", "Data Source Name")

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

	defer db.Close()

	r := gin.New()
	r.Use(middleware.LogRequest(app))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/products", productHandler.CreateProduct)
	r.GET("/products/:id", productHandler.GetProduct)
	r.GET("/products", productHandler.ListProducts)

	logger.Info("starting server", "addr", cfg.Port, "env", cfg.Env)
	if err := r.Run(); err != nil {
		app.Logger.Error(err.Error())
		os.Exit(1)
	}
}
