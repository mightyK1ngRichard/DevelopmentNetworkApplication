package handler

import (
	"VikingsServer/internal/app/ds"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) HikesList(ctx *gin.Context) {
	hikes, err := h.Repository.HikesList()
	if hikeIdString := ctx.Query("hike"); hikeIdString != "" {
		hikeById(ctx, h, hikeIdString)
		return
	}

	if err != nil {
		h.errorHandler(ctx, http.StatusNoContent, err)
		return
	}
	h.successHandler(ctx, "hikes", hikes)
}

func hikeById(ctx *gin.Context, h *Handler, hikeStringID string) {
	hikeID, err := strconv.Atoi(hikeStringID)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	hike, errDB := h.Repository.HikeByID(uint(hikeID))
	if errDB != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, errDB)
		return
	}
	h.successHandler(ctx, "hike", hike)
}

func (h *Handler) AddHike(ctx *gin.Context) {
	var hike ds.Hike
	if err := ctx.BindJSON(&hike); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}

	if errorCode, err := h.createHike(&hike); err != nil {
		h.errorHandler(ctx, errorCode, err)
		return
	}

	h.successAddHandler(ctx, "hike_id", hike.ID)
}

func (h *Handler) createHike(hike *ds.Hike) (int, error) {
	if hike.ID != 0 {
		return http.StatusBadRequest, idMustBeEmpty
	}
	if err := h.Repository.AddHike(hike); err != nil {
		return http.StatusInternalServerError, err
	}
	return 0, nil
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
		"id":                       updatedHike.ID,
		"hike_name":                updatedHike.HikeName,
		"date_created":             updatedHike.DateCreated,
		"date_end":                 updatedHike.DateEnd,
		"date_start_of_processing": updatedHike.DateStartOfProcessing,
		"date_approve":             updatedHike.DateApprove,
		"date_start_hike":          updatedHike.DateStartHike,
		"user_id":                  updatedHike.UserID,
		"status_id":                updatedHike.StatusID,
		"description":              updatedHike.Description,
	})
}
