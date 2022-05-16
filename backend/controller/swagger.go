package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	_ "rssReader/docs"
)

type SwaggerController struct {}


func (*SwaggerController) GetRouting(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)
}
