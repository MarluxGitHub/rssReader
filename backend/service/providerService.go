package service

import "rssReader/models"

var ProviderServiceInstance ProviderService = ProviderService{}

type ProviderService struct{}

func (service *ProviderService) Add(provider models.Provider) models.Provider {
	DatabaseServiceInstance.Connection.Save(&provider)

	return provider
}

func (service *ProviderService) GetAll() []models.Provider {
	var providers []models.Provider
	DatabaseServiceInstance.Connection.Find(&providers)

	return providers
}

func (service *ProviderService) Get(id int) *models.Provider{
	provider := models.Provider{}
	DatabaseServiceInstance.Connection.First(&provider, id)
	return &provider
}

func (service *ProviderService) Delete(id int) bool {
	provider := models.Provider{}
	DatabaseServiceInstance.Connection.First(&provider, id)
	DatabaseServiceInstance.Connection.Delete(&provider)

	return true
}



