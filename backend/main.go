package main

import (
	"log"
	"rssReader/config"

	"github.com/gofiber/fiber/v2"
)

// @title RSS READER API by Marlux
// @version 1.0
// @description This is a simple RSS Reader API
// @termsOfService http://swagger.io/terms/

// @contact.name -Marlux-

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	app := fiber.New()

	config.GetRouting(app)

	log.Fatal(app.Listen(":8080"))
}