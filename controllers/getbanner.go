package controllers

import (
	"gobanner/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getBannerInput struct {
	TagId     *int64 `json:"tag_id" binding:"required"`
	FeatureId *int64 `json:"feature_id" binding:"required"`
	// Если не указан то будет присвоено стандартное значение bool - false
	UseLastRevision bool `json:"use_last_revision"`
}

func (p *PublicController) getBanner(c *gin.Context) {
	var i getBannerInput
	if err := c.ShouldBindJSON(&i); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": models.ErrInvalidJson.Error()})
		return
	}
	banner, err := models.GetBanner(p.db, *i.TagId, *i.FeatureId, i.UseLastRevision)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": models.ErrInternal.Error()})
		log.Println(err)
		return
	}
	c.Data(http.StatusOK, "application/json", []byte(banner.Content))
}
