package handler

import (
	"VikingsServer/internal/app/ds"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) DestinationHikesList(ctx *gin.Context) {
	destinationHikes, err := h.Repository.DestinationHikesList()
	if err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	h.successHandler(ctx, "destination_hikes", destinationHikes)
}

func (h *Handler) AddDestinationToHike(ctx *gin.Context) {
	var body struct {
		Hike         ds.Hike `json:"hike"`
		City         ds.City `json:"city"`
		SerialNumber int     `json:"serial_number"`
	}
	if err := ctx.BindJSON(&body); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	if body.SerialNumber == 0 {
		h.errorHandler(ctx, http.StatusBadRequest, serialNumberCannotBeEmpty)
		return
	}

	/// Определяем, какой метод использовать. С id или без
	if body.Hike.ID == 0 {
		if errorCode, err := h.createHike(&body.Hike); err != nil {
			h.errorHandler(ctx, errorCode, err)
			return
		}
	}

	/// Определяем, какой метод использовать. С id или без
	if body.City.ID == 0 {
		if errorCode, err := h.createCity(&body.City); err != nil {
			h.errorHandler(ctx, errorCode, err)
			return
		}
	}

	if body.Hike.ID == 0 || body.City.ID == 0 {
		h.errorHandler(ctx, http.StatusBadRequest, destinationOrCityIsEmpty)
		return
	}

	destinationHike := ds.DestinationHikes{
		City:         body.City,
		Hike:         body.Hike,
		SerialNumber: body.SerialNumber,
	}

	if err := h.Repository.AddDestinationToHike(&destinationHike); err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	h.successAddHandler(ctx, "updated_destination_hike", destinationHike)
}
