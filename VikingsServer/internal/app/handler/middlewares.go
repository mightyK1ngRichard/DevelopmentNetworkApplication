package handler

import (
	"VikingsServer/internal/app/ds"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
)

const jwtPrefix = "Bearer "

func (h *Handler) WithAuthCheck(ctx *gin.Context) {
	jwtStr := ctx.GetHeader("Authorization")
	if !strings.HasPrefix(jwtStr, jwtPrefix) {
		h.errorHandler(ctx, http.StatusForbidden, prefixIsNil)
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}

	// отрезаем префикс
	jwtStr = jwtStr[len(jwtPrefix):]

	_, err := jwt.ParseWithClaims(jwtStr, &ds.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(h.Config.JWT.Token), nil
	})
	if err != nil {
		h.errorHandler(ctx, http.StatusForbidden, err)
		h.Logger.Error(err)
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}
}
