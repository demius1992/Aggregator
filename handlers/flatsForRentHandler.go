package handlers

import (
	"Aggregator/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getFlatsForDay(c *gin.Context) {
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid page param")
		return
	}
	flats, err := h.services.FlatsForRent.GetFlatsForDay(page)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, models.GetAllFlats{
		Data: flats,
	})
}

func (h *Handler) getFlatsLongRent(c *gin.Context) {
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid page param")
		return
	}
	flats, err := h.services.FlatsForRent.GetFlatsLongRent(page)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, models.GetAllFlats{
		Data: flats,
	})
}

func (h *Handler) getCottagesForDay(c *gin.Context) {
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid page param")
		return
	}
	cottages, err := h.services.FlatsForRent.GetCottagesForDay(page)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, models.GetAllCottages{
		Data: cottages,
	})
}

func (h *Handler) getCottagesLongRent(c *gin.Context) {
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid page param")
		return
	}
	cottages, err := h.services.FlatsForRent.GetCottagesLongRent(page)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, models.GetAllCottages{
		Data: cottages,
	})
}
