package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createScheme(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllSchemes(c *gin.Context) {

}

func (h *Handler) getSchemeById(c *gin.Context) {

}

func (h *Handler) updateScheme(c *gin.Context) {

}

func (h *Handler) deleteScheme(c *gin.Context) {

}
