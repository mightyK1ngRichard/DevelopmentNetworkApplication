package handler

import (
	"VikingsServer/internal/app/ds"
	"VikingsServer/internal/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) HikesList(ctx *gin.Context) {
	if hikeIdString := ctx.Query("hike"); hikeIdString != "" {
		hikeById(ctx, h, hikeIdString)
		return
	}

	statusID := ctx.Query("status_id")
	if statusID == "" {
		statusID = "3"
	}
	if isOk := utils.Contains([]string{"1", "2", "3", "4"}, statusID); !isOk {
		h.errorHandler(ctx, http.StatusBadRequest, errors.New("param `status_id` not contains into [1, 2, 3, 4]"))
		return
	}
	hikes, err := h.Repository.HikesList(statusID)
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

func (h *Handler) UpdateStatusForUser(ctx *gin.Context) {
	var body struct {
		StatusID uint `json:"status_id"`
	}

	if err := ctx.BindJSON(&body); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}

	if isOk := utils.Contains([]string{"2"}, strconv.Itoa(int(body.StatusID))); !isOk {
		h.errorHandler(ctx, http.StatusBadRequest, errors.New("param `status_id` not contains into [2]"))
		return
	}

	HikeID, err := h.Repository.HikeBasketId()
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, errors.New("basket is not created"))
		return
	}

	if err := h.Repository.UpdateHikeForModerator(HikeID, body.StatusID); err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) UpdateStatusForModerator(ctx *gin.Context) {
	var body struct {
		HikeID   uint `json:"hike_id"`
		StatusID uint `json:"status_id"`
	}

	if err := ctx.BindJSON(&body); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}

	if isOk := utils.Contains([]string{"3", "4"}, strconv.Itoa(int(body.StatusID))); !isOk {
		h.errorHandler(ctx, http.StatusBadRequest, errors.New("param `status_id` not contains into [3, 4]"))
		return
	}

	if err := h.Repository.UpdateHikeForModerator(body.HikeID, body.StatusID); err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) UpdateHikeStatus(ctx *gin.Context) {
	var updatedHike ds.Hike
	if err := ctx.BindJSON(&updatedHike); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	if updatedHike.ID != 0 ||
		updatedHike.HikeName != "" ||
		updatedHike.DateCreated.String() != utils.EmptyDate ||
		updatedHike.DateEnd.String() != utils.EmptyDate ||
		updatedHike.DateStartOfProcessing.String() != utils.EmptyDate ||
		updatedHike.DateApprove.String() != utils.EmptyDate ||
		updatedHike.DateStartHike.String() != utils.EmptyDate ||
		updatedHike.UserID != 0 ||
		updatedHike.Description != "" ||
		updatedHike.Leader != "" {
		h.errorHandler(ctx, http.StatusBadRequest, mustBeJustStatus)
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
