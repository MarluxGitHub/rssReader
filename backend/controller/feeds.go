package controller

import (
	"rssReader/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type FeedsController struct {}

func (*FeedsController) GetRouting(app *fiber.App) {
	app.Get("/feeds", getAllFeeds)
	app.Get("/feeds/:providerId", getFeedsOfProvider)
	app.Get("/feeds/:providerId/:feedId", getFeed)
}

// GetRouting is the function that is called by the main controller to register the routes
// @Summary      Get all feeds
// @Description  Get all feeds
// @Accept       json
// @Produce      json
// @Success      200	{object} []models.Rss2
// @Router       /feeds [get]
func getAllFeeds(c *fiber.Ctx) error {
	feeds := service.FeedServiceInstance.GetAllFeeds()
	return c.JSON(feeds)
}

// GetRouting is the function that is called by the main controller to register the routes
// @Summary      Get all feeds From a Provider
// @Description  Get all feeds From a Provider
// @Accept       json
// @Produce      json
// @Param 		 providerId path	int	true	"Provider Id"
// @Success      200	{object} models.Rss2
// @Router       /feeds/{providerId} [get]
func getFeedsOfProvider(c *fiber.Ctx) error {
	providerId, _ := strconv.Atoi(c.Params("providerId"))
	feeds := service.FeedServiceInstance.GetAllFeedsOfProvider(providerId)
	return c.JSON(feeds)
}

// GetRouting is the function that is called by the main controller to register the routes
// @Summary      Get one feed From a Provider
// @Description  Get one feed From a Provider
// @Accept       json
// @Produce      json
// @Param 		 providerId path	int	true	"Provider Id"
// @Param 		 feedId path	int	true	"Feed Id"
// @Success      200	{object} models.Item
// @Router       /feeds/{providerId}/{feedId} [get]
func getFeed(c *fiber.Ctx) error {
	providerId, _ := strconv.Atoi(c.Params("providerId"))
	feedId, _ := strconv.Atoi(c.Params("feedId"))
	feed := service.FeedServiceInstance.GetFeedOfProvider(providerId, feedId)
	return c.JSON(feed)
}