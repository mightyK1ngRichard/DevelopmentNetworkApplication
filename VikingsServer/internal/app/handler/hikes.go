package handler

import (
	"VikingsServer/internal/app/ds"
	"VikingsServer/internal/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// HikesList godoc
// @Summary Список походов
// @Tags Походы
// @Description Получение списка походов с фильтрами по статусу, дате начала и дате окончания.
// @Produce json
// @Param hike query string false "Идентификатор конкретного похода для получения информации"
// @Param status_id query string false "Статус похода. Возможные значения: 1, 2, 3, 4."
// @Param start_date query string false "Дата начала периода фильтрации в формате '2006-01-02'. Если не указана, используется '0001-01-01'."
// @Param end_date query string false "Дата окончания периода фильтрации в формате '2006-01-02'. Если не указана, используется текущая дата."
// @Success 200 {array} ds.HikesListRes "Список походов"
// @Success 200 {array} ds.HikesListRes2 "Список походов"
// @Failure 400 {object} errorResp "Неверный запрос"
// @Failure 204 {object} errorResp "Нет данных"
// @Router /api/v3/hikes [get]
func (h *Handler) HikesList(ctx *gin.Context) {
	if hikeIdString := ctx.Query("hike"); hikeIdString != "" {
		hikeById(ctx, h, hikeIdString)
		return
	}
	statusID := ctx.Query("status_id")
	startDateStr := ctx.Query("start_date")
	endDateStr := ctx.Query("end_date")

	if startDateStr == "" {
		startDateStr = "0001-01-01"
	}
	if endDateStr == "" {
		endDateStr = time.Now().Format("2006-01-02")
	}

	startDate, errStart := utils.ParseDateString(startDateStr)
	endDate, errEnd := utils.ParseDateString(endDateStr)
	h.Logger.Info(startDate, endDate)
	if errEnd != nil || errStart != nil {
		h.errorHandler(ctx, http.StatusBadRequest, errors.New("incorrect `start_date` or `end_date`"))
		return
	}

	if statusID == "" {
		statusID = "3"
	}
	if isOk := utils.Contains([]string{"1", "2", "3", "4"}, statusID); !isOk {
		h.errorHandler(ctx, http.StatusBadRequest, errors.New("param `status_id` not contains into [1, 2, 3, 4]"))
		return
	}
	hikes, err := h.Repository.HikesList(statusID, startDate, endDate)
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

// DeleteHike godoc
// @Summary Удаление похода
// @Security ApiKeyAuth
// @Tags Походы
// @Description Удаление похода по идентификатору.
// @Accept json
// @Produce json
// @Param request body ds.DeleteHikeReq true "Идентификатор похода для удаления"
// @Success 200 {object} ds.DeleteHikeRes "Успешное удаление похода"
// @Failure 400 {object} errorResp "Неверный запрос"
// @Failure 500 {object} errorResp "Внутренняя ошибка сервера"
// @Router /api/v3/hikes [delete]
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

// UpdateStatusForUser godoc
// @Summary Обновление статуса похода для пользователя. Т.е сформировать поход
// @Tags Походы
// @Description Обновление статуса похода для пользователя. Можно только сформировать(2)
// @Accept json
// @Produce json
// @Param body body ds.UpdateStatusForUserReq true "Детали обновления статуса [2]"
// @Success 200 {object} string "Успешное обновление статуса"
// @Failure 400 {object} errorResp "Неверный запрос"
// @Failure 500 {object} errorResp "Внутренняя ошибка сервера"
// @Router /api/v3/hikes/update/status-for-user [put]
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

	if err := h.Repository.UpdateStatusForUser(HikeID, body.StatusID); err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
}

// UpdateStatusForModerator godoc
// @Summary Обновление статуса похода для модератора
// @Tags Походы
// @Description Обновление статуса похода для модератора. Можно только принять(3) отказать(4)
// @Accept json
// @Produce json
// @Param body body ds.UpdateStatusForModeratorReq true "Детали обновления статуса [3, 4]"
// @Success 200 {object} string "Успешное обновление статуса"
// @Failure 400 {object} errorResp "Неверный запрос"
// @Failure 500 {object} errorResp "Внутренняя ошибка сервера"
// @Router /api/v3/hikes/update/status-for-moderator [put]
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

// UpdateHike godoc
// @Summary Обновление данных о походе
// @Tags Походы
// @Description Обновление данных о походе.
// @Accept json
// @Produce json
// @Param updatedHike body ds.UpdateHikeReq true "Данные для обновления похода"
// @Success 200 {object} ds.UpdatedHikeRes "Успешное обновление данных о походе"
// @Failure 400 {object} errorResp "Неверный запрос"
// @Failure 500 {object} errorResp "Внутренняя ошибка сервера"
// @Router /api/v3/hikes [put]
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

// MARK: - OLD

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
