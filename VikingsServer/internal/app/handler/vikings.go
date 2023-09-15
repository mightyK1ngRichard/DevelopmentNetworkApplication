package handler

import (
	"VikingsServer/internal/app/ds"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) AddViking(ctx *gin.Context) {
	viking := ds.Vikings{CityOfBirth: -1}
	if err := ctx.BindJSON(&viking); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	if err := h.Repository.AddViking(&viking); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"vikingID": viking.ID,
	})
}

func (h *Handler) UpdateViking(ctx *gin.Context) {
	viking := ds.Vikings{ID: -1}
	if err := ctx.BindJSON(&viking); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	if viking.ID == -1 {
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("param `id` not found"))
		return
	}
	if err := h.Repository.UpdateViking(&viking); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"vikingID": viking.ID,
	})
}
