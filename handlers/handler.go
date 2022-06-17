package handlers

import (
	"Aggregator/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// relative paths for flats for sale
	router.GET("/flats/sale/:page", h.getAllFlats)
	router.GET("/flats/sale/one-room-flats/:page", h.getOneRoomFlats)
	router.GET("/flats/sale/two-room-flats/:page", h.getTwoRoomFlats)
	router.GET("/flats/sale/three-room-flats/:page", h.getThreeRoomFlats)
	router.GET("/flats/sale/three-more-room-flats/:page", h.getMoreThanThreeRoomFlats)

	// relative paths for cottages, areas, country houses for sale
	router.GET("/cottages/sale/houses/:page", h.getAllHouses)
	router.GET("/cottages/sale/areas/:page", h.getAllAreas)
	router.GET("/cottages/sale/vitebskiyi-rayion/dacha/:page", h.getAllDacha)

	// relative paths for cottages and flats for rent
	router.GET("/flats/rent/flats-for-day/:page", h.getFlatsForDay)
	router.GET("/flats/rent/flats-for-long/:page", h.getFlatsLongRent)
	router.GET("/flats/rent/cottages-for-day/:page", h.getCottagesForDay)
	router.GET("/flats/rent/cottages-for-long/:page", h.getCottagesLongRent)

	return router
}
