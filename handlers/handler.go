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

	flats := router.Group("/flats/sale")
	{
		flats.GET("/:page", h.getAllFlats) // get all flats

		oneRoomFlats := flats.Group("/one-room-flats")
		{
			oneRoomFlats.GET("/:page", h.getOneRoomFlats) //get all one-room flats
		}
		twoRoomFlats := flats.Group("/two-room-flats")
		{
			twoRoomFlats.GET("/:page", h.getTwoRoomFlats) //get all two-room flats
		}
		threeRoomFlats := flats.Group("/three-room-flats")
		{
			threeRoomFlats.GET("/:page", h.getThreeRoomFlats) //get all three-room flats
		}
		moreThanThreeRoomFlats := flats.Group("/three-more-room-flats")
		{
			moreThanThreeRoomFlats.GET("/:page", h.getMoreThanThreeRoomFlats) //get all more than three-room flats
		}
	}

	cottages := router.Group("/cottages/sale")
	{
		houses := cottages.Group("/houses")
		{
			houses.GET("/:page", h.getAllHouses) //GET all houses
		}
		areas := cottages.Group("areas")
		{
			areas.GET("/:page", h.getAllAreas) //GET all areas
		}
		dacha := cottages.Group("vitebskiyi-rayion/dacha")
		{
			dacha.GET("/:page", h.getAllDacha) // GET all country houses
		}
	}

	flatsForRent := router.Group("/flats/rent")
	{
		flatsForDay := flatsForRent.Group("/flats-for-day")
		{
			flatsForDay.GET("/:page", h.getFlatsForDay) //GET available for day flats for rent
		}
		flatsLongRent := flatsForRent.Group("/flats-for-long")
		{
			flatsLongRent.GET("/:page", h.getFlatsLongRent) //Get available flats for long rent
		}
		cottagesForDay := flatsForRent.Group("/cottages-for-day")
		{
			cottagesForDay.GET("/:page", h.getCottagesForDay) //Get available cottages for day rent
		}
		cottagesLongRent := flatsForRent.Group("/cottages-for-long")
		{
			cottagesLongRent.GET("/:page", h.getCottagesLongRent) //Get available cottages for long rent
		}
	}
	return router
}
