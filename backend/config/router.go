package config

import (
	"rssReader/controller"
	"rssReader/service"

	"github.com/gofiber/fiber/v2"
)



func GetRouting(app *fiber.App) {
	service.DatabaseServiceInstance.Connect()
	service.DatabaseServiceInstance.Migrate()

	feeds := controller.FeedsController{}
	provider := controller.ProviderController{}
	swagger := controller.SwaggerController{}

	feeds.GetRouting(app)
	provider.GetRouting(app)
	swagger.GetRouting(app)
}