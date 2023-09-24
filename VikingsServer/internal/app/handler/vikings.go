package handler

import (
	"VikingsServer/internal/app/ds"
	"VikingsServer/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) VikingsList(ctx *gin.Context) {
	vikings, err := h.Repository.VikingList()

	if idStr := ctx.Query("viking"); idStr != "" {
		vikingById(ctx, h, idStr)
		return
	}

	if err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	h.successHandler(ctx, "vikings", vikings)
}

func vikingById(ctx *gin.Context, h *Handler, idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}

	viking, errBD := h.Repository.VikingById(uint(id))
	if errBD != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, errBD)
		return
	}

	h.successHandler(ctx, "viking", viking)
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

	h.successAddHandler(ctx, "viking_id", viking.ID)
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
