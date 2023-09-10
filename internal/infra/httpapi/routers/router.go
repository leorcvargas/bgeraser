package routers

import (
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"

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

func loadMiddlewares(r *fiber.App) {
	// Security
	r.Use(recover.New())
	r.Use(csrf.New())
	r.Use(cors.New(cors.Config{
		AllowMethods: "*",
		AllowOrigins: "https://kamui.app, https://www.kamui.app",
	}))
	r.Use(etag.New())
	r.Use(helmet.New())
	r.Use(limiter.New(limiter.Config{Max: 50}))

	// Access log
	r.Use(requestid.New())
	r.Use(logger.New(logger.Config{
		Format: "${pid} | ${ip} | ${locals:requestid} | ${status} | ${method} ${path}\n",
	}))

	// Static server
	r.Use("/i", filesystem.New(filesystem.Config{
		Root: http.Dir("./data/images"),
	}))
}

func MakeRouter(
	routers []Router,
	config *config.Config,
) *fiber.App {
	routerConfig := buildRouterConfig(config)

	r := fiber.New(routerConfig)
	loadMiddlewares(r)

	for _, router := range routers {
		router.Load(r)
	}

	return r
}
