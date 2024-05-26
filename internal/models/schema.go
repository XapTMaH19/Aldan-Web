package models

type Schema struct {
	Description string `json:"description" binding:"required"`
}
