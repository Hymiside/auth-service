package handler

import (
	"github.com/Hymiside/auth-microservice/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("signup", h.SignUp)
		auth.GET("signin", h.SignIn)
	}

	return router
}
