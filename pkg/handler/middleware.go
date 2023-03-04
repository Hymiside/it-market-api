package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	userCtx             = "userId"
	authorizationHeader = "Authorization"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		responseWithError(c, http.StatusUnauthorized, "empty auth header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		responseWithError(c, http.StatusUnauthorized, "invalid auth header")
		return
	}
	if len(headerParts[1]) == 0 {
		responseWithError(c, http.StatusUnauthorized, "token is empty")
		return
	}
	userId, err := h.services.Auth.ParseToken(headerParts[1])
	if err != nil {
		responseWithError(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (string, error) {
	data, ok := c.Get(userCtx)
	if !ok {
		return "", ErrUserIdNotFound
	}
	userId := data.(string)
	return userId, nil
}
