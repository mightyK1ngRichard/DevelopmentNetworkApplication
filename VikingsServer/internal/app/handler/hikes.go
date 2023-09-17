package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) HikesList(ctx *gin.Context) {
	hikes, err := h.Repository.HikesList()
	if err != nil {
		h.errorHandler(ctx, http.StatusNoContent, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"hikes":  hikes,
	})
}
