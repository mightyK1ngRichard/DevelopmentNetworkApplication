package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) AuthorsList(ctx *gin.Context) {
	authors, err := h.Repository.AuthorsList()
	if err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	h.successHandler(ctx, "authors", authors)
}
