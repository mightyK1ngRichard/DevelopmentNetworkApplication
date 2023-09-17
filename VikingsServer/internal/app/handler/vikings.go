package handler

import (
	"VikingsServer/internal/app/ds"
	"VikingsServer/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) VikingsList(ctx *gin.Context) {
	vikings, err := h.Repository.VikingList()
	if err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}
	h.successHandler(ctx, "vikings", vikings)
}

func (h *Handler) AddViking(ctx *gin.Context) {
	viking := ds.Viking{}
	if err := ctx.BindJSON(&viking); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	if viking.ID != utils.EmptyInt {
		h.errorHandler(ctx, http.StatusBadRequest, idMustBeEmpty)
		return
	}
	if viking.VikingName == utils.EmptyString {
		h.errorHandler(ctx, http.StatusBadRequest, vikingCannotBeEmpty)
		return
	}
	if err := h.Repository.AddViking(&viking); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}

	h.successHandler(ctx, "viking_id", viking.ID)
}

func (h *Handler) UpdateViking(ctx *gin.Context) {
	viking := ds.Viking{}
	if err := ctx.BindJSON(&viking); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	if viking.ID == utils.EmptyInt {
		h.errorHandler(ctx, http.StatusBadRequest, idNotFound)
		return
	}
	if err := h.Repository.UpdateViking(&viking); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}

	h.successHandler(ctx, "viking_id", viking.ID)
}
