package controllers

import (
	"database/sql"
	"errors"
	"gobanner/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getBannerInput struct {
	TagId     *int64 `form:"tag_id" binding:"required"`
	FeatureId *int64 `form:"feature_id" binding:"required"`
	// Если не указан то будет присвоено стандартное значение bool - false
	UseLastRevision bool `form:"use_last_revision"`
}

// Пока что игнорируем use_last_revision
func (p *PublicController) getBanner(c *gin.Context) {
	var i getBannerInput
	if err := c.ShouldBindQuery(&i); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": models.ErrInvalidQuery.Error()})
		return
	}
	banner, err := models.GetBanner(p.db, *i.TagId, *i.FeatureId)
	if errors.Is(err, sql.ErrNoRows) {
		c.Status(http.StatusNotFound)
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": models.ErrInternal.Error()})
		log.Println(err)
		return
	}
	c.Data(http.StatusOK, "application/json", []byte(banner.Content))
}
