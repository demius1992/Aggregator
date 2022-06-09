package service

import (
	"Aggregator/models"
	"Aggregator/repository"
)

type CottagesService struct {
	repo repository.Cottages
}

func (c *CottagesService) GetAllHouses(page int) ([]models.Cottages, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CottagesService) GetAllAreas(page int) ([]models.Areas, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CottagesService) GetAllDacha(page int) ([]models.Cottages, error) {
	//TODO implement me
	panic("implement me")
}

func NewCottagesService(repo repository.Cottages) *CottagesService {
	return &CottagesService{repo: repo}
}
