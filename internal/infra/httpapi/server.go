package httpapi

import (
	"context"
	"fmt"
	"os"
	"runtime/pprof"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/leorcvargas/bgeraser/ent"
	"github.com/leorcvargas/bgeraser/internal/infra/config"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"go.uber.org/fx"
)

const finishProfAfter = 3 * time.Minute

func startProfiling(config *config.Config) {
	log.Infof(
		"Starting CPU and Memory profiling on %s and %s",
		config.Profiling.CPU,
		config.Profiling.Mem,
	)

	cpuProfFile, err := os.Create(config.Profiling.CPU)
	if err != nil {
		log.Fatal(err)
	}

	if err = pprof.StartCPUProfile(cpuProfFile); err != nil {
		log.Fatal(err)
	}

	memoryProfFile, err := os.Create(config.Profiling.Mem)
	if err != nil {
		log.Fatal(err)
	}

	if err = pprof.WriteHeapProfile(memoryProfFile); err != nil {
		log.Fatal(err)
	}

	after := time.After(finishProfAfter)

	go func() {
		<-after
		log.Info("Stopping CPU and Memory profiling")
		pprof.StopCPUProfile()
		_ = cpuProfFile.Close()
		_ = memoryProfFile.Close()
	}()
}

// NewServer starts the fiber application.
func NewServer(
	lifecycle fx.Lifecycle,
	router *fiber.App,
	config *config.Config,
	db *ent.Client,
) *fasthttp.Server {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				log.Info("Starting the server...")

				if config.Profiling.Enabled {
					startProfiling(config)
				}

				addr := fmt.Sprintf(":%s", config.Server.Port)
				if err := router.Listen(addr); err != nil {
					log.Fatalf("Error starting the server: %s\n", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			defer func() {
				if err := db.Close(); err != nil {
					log.Warn("Error closing the database connection: %s", err)
				}
			}()

			log.Info("Stopping the server...")

			return router.ShutdownWithContext(ctx)
		},
	})

	return router.Server()
}
