package controllers

import (
	"gobanner/models"

	"github.com/gin-gonic/gin"
)

type PublicController struct {
	db models.Database
}

func NewPublicController(db models.Database) *PublicController {
	return &PublicController{db}
}

// We need separate function for router setup to do testing properly
func SetupRouter(p *PublicController) *gin.Engine {
	r := gin.Default()

	r.POST("/banner", p.createBanner)
	return r
}
