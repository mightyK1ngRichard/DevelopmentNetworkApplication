package handler

import (
	"VikingsServer/internal/app/ds"
	"VikingsServer/internal/app/role"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
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
	user, err := h.Repository.GetUserByLogin(req.Login)
	if err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	if req.Login == user.Login && user.Password == generateHashString(req.Password) {
		token := jwt.NewWithClaims(cfg.JWT.SigningMethod, &ds.JWTClaims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(cfg.JWT.ExpiresIn).Unix(),
				IssuedAt:  time.Now().Unix(),
				Issuer:    "bitop-admin",
			},
			UserUUID: uuid.New(),
			Role:     user.Role,
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

func (h *Handler) UsersList(ctx *gin.Context) {
	users, err := h.Repository.UsersList()
	if err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	h.successHandler(ctx, "users", users)
}

func (h *Handler) Register(ctx *gin.Context) {
	type registerReq struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	type registerResp struct {
		Ok bool `json:"ok"`
	}

	req := &registerReq{}

	err := json.NewDecoder(ctx.Request.Body).Decode(req)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}

	if req.Password == "" {
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("pass is empty"))
		return
	}

	if req.Login == "" {
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("name is empty"))
		return
	}

	err = h.Repository.Register(&ds.User{
		Role:     role.Buyer,
		Login:    req.Login,
		Password: generateHashString(req.Password),
	})

	if err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, &registerResp{
		Ok: true,
	})
}

// MARK: - Inner functions

func generateHashString(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
