package controller

import (
	"rssReader/models"
	"rssReader/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProviderController struct {}

func (*ProviderController) GetRouting(app *fiber.App) {
	app.Get("/providers", getProviders)
	app.Get("/providers/:id", getProvider)
	app.Post("/providers", createProvider)
	app.Delete("/providers/:id", deleteProvider)
}

func getProviders(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
			"providers" : service.ProviderServiceInstance.GetAll(),
		})
}

func getProvider(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	provider := service.ProviderServiceInstance.Get(id)
	if provider == nil {
		return c.Status(404).SendString("No Provider")
	}
	return c.JSON(fiber.Map{
		"provider": *provider,
	})
}

func createProvider(c *fiber.Ctx) error {
	var provider models.Provider
	c.BodyParser(&provider)
	provider = service.ProviderServiceInstance.Add(provider)
	return c.JSON(fiber.Map{
		"provider": provider,
	})
}

func deleteProvider(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if(err != nil) {
		return c.Status(404).SendString("Error")
	}

	if service.ProviderServiceInstance.Delete(id) {
		return c.SendString("Provider Deleted")
	} else {
		return c.Status(404).SendString("No Provider")
	}
}


