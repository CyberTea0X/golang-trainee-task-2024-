package controllers

import (
	"database/sql"
	middleware "gobanner/middlewares"

	"github.com/gin-gonic/gin"
)

type PublicController struct {
	db *sql.DB
}

func NewPublicController(db *sql.DB) *PublicController {
	return &PublicController{db}
}

// We need separate function for router setup to do testing properly
func SetupRouter(p *PublicController) *gin.Engine {
	r := gin.Default()
	admin := r.Group("/")
	admin.Use(middleware.AdminAuth())
	admin.GET("/banner", p.getBanners)
	admin.POST("/banner", p.createBanner)
	admin.DELETE("/banner/:id", p.deleteBanner)
	admin.PATCH("/banner/:id", p.patchBanner)
	user := r.Group("/")
	user.Use(middleware.UserAuth())
	user.GET("/user_banner", p.getBanner)
	return r
}
