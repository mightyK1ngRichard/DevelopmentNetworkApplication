package handler

import (
	"VikingsServer/internal/app/ds"
	"VikingsServer/internal/app/role"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"strings"
)

const jwtPrefix = "Bearer "

func (h *Handler) WithAuthCheck(assignedRoles ...role.Role) func(ctx *gin.Context) {
	return func(gCtx *gin.Context) {
		jwtStr := gCtx.GetHeader("Authorization")
		if !strings.HasPrefix(jwtStr, jwtPrefix) {
			h.errorHandler(gCtx, http.StatusForbidden, prefixIsNil)
			gCtx.AbortWithStatus(http.StatusForbidden)
			return
		}

		jwtStr = jwtStr[len(jwtPrefix):]
		token, err := jwt.ParseWithClaims(jwtStr, &ds.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(h.Config.JWT.Token), nil
		})
		if err != nil {
			h.errorHandler(gCtx, http.StatusForbidden, prefixIsNil)
			gCtx.AbortWithStatus(http.StatusForbidden)
			return
		}

		myClaims := token.Claims.(*ds.JWTClaims)

		for _, oneOfAssignedRole := range assignedRoles {
			if myClaims.Role == oneOfAssignedRole {
				gCtx.AbortWithStatus(http.StatusForbidden)
				log.Printf("role %s is not assigned in %s", myClaims.Role, assignedRoles)
				return
			}
		}
	}
}
