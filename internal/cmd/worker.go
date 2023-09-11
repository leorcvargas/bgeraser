package cmd

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"github.com/leorcvargas/bgeraser/internal/domain/entities"
	"github.com/leorcvargas/bgeraser/internal/domain/imageprocesses"
	"github.com/leorcvargas/bgeraser/internal/domain/images"
	"github.com/leorcvargas/bgeraser/internal/infra/config"
	"github.com/leorcvargas/bgeraser/internal/infra/database"
	"github.com/leorcvargas/bgeraser/internal/infra/storage"
	"github.com/leorcvargas/bgeraser/internal/queues"
	"go.uber.org/fx"
)

func Worker() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Warn("Couldn't load .env file")
	}

	app := fx.New(
		config.Module,
		database.Module,
		images.Module,
		imageprocesses.Module,
		storage.Module,
		queues.Module,
		fx.Invoke(func(queue queues.Queue[*entities.ImageProcess]) {
			err := queue.StartConsumers()
			if err != nil {
				panic(err)
			}
		}),
	)

	app.Run()
}
