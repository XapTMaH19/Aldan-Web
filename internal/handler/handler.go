package handler

import (
	"github.com/XapTMaH19/Aldan-Web/internal/service"
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
		lists := api.Group("/schemes")
		{
			lists.POST("/", h.createScheme)
			lists.GET("/", h.getAllSchemes)
			lists.GET("/:id", h.getSchemeById)
			lists.PUT("/:id", h.updateScheme)
			lists.DELETE("/:id", h.deleteScheme)

		}
	}
	return router
}
