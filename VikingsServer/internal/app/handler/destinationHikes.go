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
	var destinationHike ds.DestinationHikes
	if err := ctx.BindJSON(&destinationHike); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	if destinationHike.ID != 0 {
		h.errorHandler(ctx, http.StatusBadRequest, idMustBeEmpty)
		return
	}
	if destinationHike.SerialNumber == 0 {
		h.errorHandler(ctx, http.StatusBadRequest, serialNumberCannotBeEmpty)
		return
	}
	if destinationHike.HikeID == 0 || destinationHike.CityID == 0 {
		h.errorHandler(ctx, http.StatusBadRequest, destinationOrCityIsEmpty)
		return
	}
	if err := h.Repository.AddDestinationToHike(&destinationHike); err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	h.successAddHandler(ctx, "updated_destination_hike", destinationHike)
}
