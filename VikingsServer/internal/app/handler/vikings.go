package handler

import (
	"VikingsServer/internal/app/ds"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) AddViking(ctx *gin.Context) {
	viking := ds.Vikings{
		CityOfBirth: -1,
	}
	err := ctx.BindJSON(&viking)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	errDB := h.Repository.AddViking(&viking)
	if errDB != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errDB.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"vikingID": viking.ID,
	})
}
