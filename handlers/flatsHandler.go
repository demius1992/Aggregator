package handlers

import (
	"Aggregator/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getAllFlats(c *gin.Context) {
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid page param")
		return
	}
	flats, err := h.services.Flats.GetAllFlats(page)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, models.GetAllFlats{
		Data: flats,
	})
}

func (h *Handler) getOneRoomFlats(c *gin.Context) {
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid page param")
		return
	}
	flats, err := h.services.Flats.GetOneRoomFlats(page)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, models.GetAllFlats{
		Data: flats,
	})
}

func (h *Handler) getTwoRoomFlats(c *gin.Context) {
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid page param")
		return
	}
	flats, err := h.services.Flats.GetTwoRoomFlats(page)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, models.GetAllFlats{
		Data: flats,
	})
}

func (h *Handler) getThreeRoomFlats(c *gin.Context) {
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid page param")
		return
	}
	flats, err := h.services.Flats.GetThreeRoomFlats(page)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, models.GetAllFlats{
		Data: flats,
	})
}

func (h *Handler) getMoreThanThreeRoomFlats(c *gin.Context) {
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid page param")
		return
	}
	flats, err := h.services.Flats.GetMoreThanThreeRoomFlats(page)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, models.GetAllFlats{
		Data: flats,
	})
}
