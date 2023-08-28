package routers

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/leorcvargas/bgeraser/internal/infra/config"
)

type Router interface {
	Load(r *fiber.App)
}

func buildRouterConfig(config *config.Config) fiber.Config {
	routerConfig := fiber.Config{
		AppName:       "bgeraser by @leorcvargas",
		CaseSensitive: true,
		Prefork:       config.Server.Prefork,
	}

	if config.Server.UseSonic {
		log.Info("Loading Sonic JSON into the router")
		routerConfig.JSONEncoder = sonic.Marshal
		routerConfig.JSONDecoder = sonic.Unmarshal
	}

	return routerConfig
}

func MakeRouter(
	routers []Router,
	config *config.Config,
) *fiber.App {
	routerConfig := buildRouterConfig(config)
	r := fiber.New(routerConfig)

	for _, router := range routers {
		router.Load(r)
	}

	return r
}
