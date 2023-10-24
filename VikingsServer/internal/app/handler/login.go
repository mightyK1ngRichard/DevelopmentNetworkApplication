package handler

import (
	"VikingsServer/internal/app/ds"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (h *Handler) Login(ctx *gin.Context) {
	cfg := h.Config
	req := &ds.LoginReq{}

	if err := json.NewDecoder(ctx.Request.Body).Decode(req); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}

	login := "king"
	password := "king"
	if req.Login == login && req.Password == password {
		// генерируем ему jwt
		token := jwt.NewWithClaims(cfg.JWT.SigningMethod, &ds.JWTClaims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(cfg.JWT.ExpiresIn).Unix(),
				IssuedAt:  time.Now().Unix(),
				Issuer:    "bitop-admin",
			},
			UserUUID: uuid.New(),
			Scopes:   []string{},
		})

		if token == nil {
			h.errorHandler(ctx, http.StatusInternalServerError, tokenIsNil)
			return
		}

		strToken, err := token.SignedString([]byte(cfg.JWT.Token))
		if err != nil {
			h.errorHandler(ctx, http.StatusInternalServerError, cannotCreateToken)
			return
		}

		ctx.JSON(http.StatusOK, ds.LoginResp{
			ExpiresIn:   cfg.JWT.ExpiresIn,
			AccessToken: strToken,
			TokenType:   "Bearer",
		})
	}

	ctx.AbortWithStatus(http.StatusForbidden)
}
