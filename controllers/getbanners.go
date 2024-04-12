package controllers

import (
	"gobanner/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getBannersInput struct {
	FeatureId *int64 `form:"feature_id"`
	TagId     *int64 `form:"tag_id"`
	Limit     *int64 `form:"limit"`
	Offset    *int64 `form:"offset"`
}

func (p *PublicController) getBanners(c *gin.Context) {
	var i getBannersInput
	if err := c.ShouldBindQuery(&i); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": models.ErrInvalidQuery})
		return
	}
	filter := &models.BannerFilter{FeatureId: i.FeatureId, TagId: i.TagId}
	banners, err := models.GetBanners(p.db, filter, i.Limit, i.Offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": models.ErrInternal})
	}
	c.JSON(http.StatusOK, banners)
}
