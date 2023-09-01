package main

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"github.com/leorcvargas/bgeraser/internal/domain/images"
	"github.com/leorcvargas/bgeraser/internal/infra/config"
	"github.com/leorcvargas/bgeraser/internal/infra/database"
	"github.com/leorcvargas/bgeraser/internal/infra/httpapi"
	"github.com/leorcvargas/bgeraser/internal/infra/httpapi/controllers"
	"github.com/leorcvargas/bgeraser/internal/infra/httpapi/routers"
	"github.com/leorcvargas/bgeraser/internal/infra/storage"
	"github.com/leorcvargas/bgeraser/internal/infra/worker"
	"go.uber.org/fx"

	_ "go.uber.org/automaxprocs"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Warn("Couldn't load .env file")
	}

	app := fx.New(
		controllers.Module,
		config.Module,
		routers.Module,
		database.Module,
		images.Module,
		httpapi.Module,
		storage.Module,
		worker.Module,
	)

	app.Run()
}
