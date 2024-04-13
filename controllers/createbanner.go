package controllers

import (
	"gobanners/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createBannerInput struct {
	TagIds    []int64 `json:"tag_ids" binding:"required"`
	FeatureId *int64  `json:"feature_id" binding:"required"`
	Content   *string `json:"content" binding:"required"`
	IsActive  *bool   `json:"is_active" binding:"required"`
}

func (p *PublicController) createBanner(c *gin.Context) {
	var i createBannerInput
	if err := c.ShouldBindJSON(&i); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": models.ErrInvalidJson.Error()})
		return
	}
	banner := new(models.Banner)
	banner.TagIds = i.TagIds
	banner.FeatureId = *i.FeatureId
	banner.Content = *i.Content
	banner.IsActive = *i.IsActive
	id, err := banner.InsertToDB(p.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": models.ErrInternal.Error()})
		log.Println(err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"banner_id": id})
	return
}
