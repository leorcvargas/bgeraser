package cmd

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
	"github.com/leorcvargas/bgeraser/internal/producers"
	"github.com/leorcvargas/bgeraser/internal/queues"
	"go.uber.org/fx"
)

func Server() {
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
		queues.Module,
		producers.Module,
	)

	app.Run()
}
