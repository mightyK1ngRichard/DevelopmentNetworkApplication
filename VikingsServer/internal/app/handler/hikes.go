package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) HikesList(ctx *gin.Context) {
	hikes, err := h.Repository.HikesList()
	if err != nil {
		ctx.JSON(http.StatusNoContent, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"hikes": hikes,
	})
}
