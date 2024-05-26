package handler

import (
	"github.com/XapTMaH19/AldanWeb/internal/models"
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
	var input models.Schema

	/*if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})*/
}

func (h *Handler) getSchemeById(c *gin.Context) {

}

func (h *Handler) updateScheme(c *gin.Context) {

}

func (h *Handler) deleteScheme(c *gin.Context) {

}
