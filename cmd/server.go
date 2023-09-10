package main

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"github.com/leorcvargas/bgeraser/internal/domain/imageprocesses"
	"github.com/leorcvargas/bgeraser/internal/domain/images"
	"github.com/leorcvargas/bgeraser/internal/infra/config"
	"github.com/leorcvargas/bgeraser/internal/infra/database"
	"github.com/leorcvargas/bgeraser/internal/infra/httpapi"
	"github.com/leorcvargas/bgeraser/internal/infra/httpapi/controllers"
	"github.com/leorcvargas/bgeraser/internal/infra/httpapi/routers"
	"github.com/leorcvargas/bgeraser/internal/infra/storage"
	processinworker "github.com/leorcvargas/bgeraser/internal/infra/worker/process_in"
	processoutworker "github.com/leorcvargas/bgeraser/internal/infra/worker/process_out"

	_ "go.uber.org/automaxprocs"
	"go.uber.org/fx"
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
		imageprocesses.Module,
		httpapi.Module,
		storage.Module,
		processinworker.Module,
		processoutworker.Module,
		fx.NopLogger,
	)

	app.Run()
}
