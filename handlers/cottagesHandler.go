package handlers

import (
	"Aggregator/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getAllHouses(c *gin.Context) {
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid page param")
		return
	}
	cottages, err := h.services.Cottages.GetAllHouses(page)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, models.GetAllCottages{
		Data: cottages,
	})
}

func (h *Handler) getAllAreas(c *gin.Context) {
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid page param")
		return
	}
	areas, err := h.services.Cottages.GetAllAreas(page)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, models.GetAllAreas{
		Data: areas,
	})
}

func (h *Handler) getAllDacha(c *gin.Context) {
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid page param")
		return
	}
	dacha, err := h.services.Cottages.GetAllDacha(page)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, models.GetAllCottages{
		Data: dacha,
	})
}
