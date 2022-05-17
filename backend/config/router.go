package config

import (
	"rssReader/controller"
	"rssReader/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func GetRouting(app *fiber.App) {
	app.Use(cors.New(cors.ConfigDefault))

	service.DatabaseServiceInstance.Connect()
	service.DatabaseServiceInstance.Migrate()

	feeds := controller.FeedsController{}
	provider := controller.ProviderController{}
	swagger := controller.SwaggerController{}

	feeds.GetRouting(app)
	provider.GetRouting(app)
	swagger.GetRouting(app)
}