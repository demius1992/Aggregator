package service

import (
	"Aggregator/models"
	"Aggregator/repository"
)

type CottagesService struct {
	repo repository.Cottages
}

func (c *CottagesService) GetAllHouses(page int) ([]models.Cottages, error) {
	return c.repo.GetAllHouses(page)
}

func (c *CottagesService) GetAllAreas(page int) ([]models.Areas, error) {
	return c.repo.GetAllAreas(page)
}

func (c *CottagesService) GetAllDacha(page int) ([]models.Cottages, error) {
	return c.repo.GetAllDacha(page)
}

func NewCottagesService(repo repository.Cottages) *CottagesService {
	return &CottagesService{repo: repo}
}
