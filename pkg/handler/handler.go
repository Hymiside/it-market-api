package handler

import (
	"errors"
	"fmt"
	"github.com/Hymiside/it-market-api/pkg/service"
	"github.com/gin-gonic/gin"
	"time"
)

var (
	ErrUserIdNotFound = errors.New("userId not found")
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s]	REQUEST: %s %s    STATUS-CODE: %d    LATENSY: %s\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
		)
	}))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	_ = router.Group("/api", h.userIdentity)

	return router
}
