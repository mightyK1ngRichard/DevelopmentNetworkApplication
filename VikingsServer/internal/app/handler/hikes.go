package handler

import (
	"VikingsServer/internal/app/ds"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) HikesList(ctx *gin.Context) {
	hikes, err := h.Repository.HikesList()
	if err != nil {
		h.errorHandler(ctx, http.StatusNoContent, err)
		return
	}

	h.successHandler(ctx, "hikes", hikes)
}

func (h *Handler) AddHike(ctx *gin.Context) {
	var hike ds.Hike
	if err := ctx.BindJSON(&hike); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	if hike.ID != 0 {
		h.errorHandler(ctx, http.StatusBadRequest, idMustBeEmpty)
		return
	}
	if err := h.Repository.AddHike(&hike); err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	h.successHandler(ctx, "hike_id", hike.ID)
}

func (h *Handler) DeleteHike(ctx *gin.Context) {
	var request struct {
		ID uint `json:"id"`
	}
	if err := ctx.BindJSON(&request); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	if request.ID == 0 {
		h.errorHandler(ctx, http.StatusBadRequest, idNotFound)
		return
	}
	if err := h.Repository.DeleteHike(request.ID); err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	h.successHandler(ctx, "hike_id", request.ID)
}

func (h *Handler) UpdateHike(ctx *gin.Context) {
	var updatedHike ds.Hike
	if err := ctx.BindJSON(&updatedHike); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	if updatedHike.ID == 0 {
		h.errorHandler(ctx, http.StatusBadRequest, idNotFound)
		return
	}
	if err := h.Repository.UpdateHike(&updatedHike); err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}
	h.successHandler(ctx, "updated_hike", gin.H{
		"id":                   updatedHike.ID,
		"hike_name":            updatedHike.HikeName,
		"date_start":           updatedHike.DateStart,
		"date_end":             updatedHike.DateEnd,
		"date_start_preparing": updatedHike.DateStartPreparing,
		"author_id":            updatedHike.AuthorID,
		"status_id":            updatedHike.StatusID,
		"description":          updatedHike.Description,
	})
}
