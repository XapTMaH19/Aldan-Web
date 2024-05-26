package handler

import (
	"github.com/XapTMaH19/AldanWeb/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services}
}

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("sign-up", h.signUp)
		auth.POST("sign-in", h.signIn)
	}
	api := router.Group("/api", h.userIdentity)
	{
		schemes := api.Group("/schemes")
		{
			schemes.POST("/", h.createScheme)
			schemes.GET("/", h.getAllSchemes)
			schemes.GET("/:id", h.getSchemeById)
			schemes.PUT("/:id", h.updateScheme)
			schemes.DELETE("/:id", h.deleteScheme)

		}
	}
	return router
}
