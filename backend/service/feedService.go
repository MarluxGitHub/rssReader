package service

import "rssReader/models"

var FeedServiceInstance FeedService = FeedService{}

type FeedService struct{}

func (feedService *FeedService) GetAllFeeds() []models.Rss2 {
	providers := ProviderServiceInstance.GetAll()
	feeds := make([]models.Rss2, len(providers))

	for i, provider := range providers {
		feeds[i] = getFeedsOfProvider(&provider)
	}

	return feeds
}

func (feedService *FeedService) GetAllFeedsOfProvider(providerId int) models.Rss2{
	provider := ProviderServiceInstance.Get(providerId)
	feeds := getFeedsOfProvider(provider)

	return feeds
}

func (feedService *FeedService) GetFeedOfProvider(providerId int, id int) models.Item {
	provider := ProviderServiceInstance.Get(providerId)
	feeds := getFeedsOfProvider(provider)

	return feeds.Channel.Item[id]
}

func getFeedsOfProvider(provider *models.Provider) models.Rss2 {
	return RssReaderInstance.Read(provider)
}


