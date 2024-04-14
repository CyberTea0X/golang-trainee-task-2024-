package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	"gobanners/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type patchBannerInput struct {
	TagIds    []int64 `json:"tag_ids"`
	FeatureId *int64  `json:"feature_id"`
	Content   *string `json:"content"`
	IsActive  *bool   `json:"is_active"`
}

func (p *PublicController) patchBanner(c *gin.Context) {
	var id int64
	param := c.Param("id")
	_, err := fmt.Sscan(param, &id)
	// if param is "", Sscan will return error EOF
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": models.ErrInvalidId.Error()})
		return
	}
	var i patchBannerInput
	if err := c.ShouldBindJSON(&i); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": models.ErrInvalidJson.Error()})
		return
	}
	patch := &models.BannerPatch{
		TagIds:    i.TagIds,
		FeatureId: i.FeatureId,
		Content:   i.Content,
		IsActive:  i.IsActive,
	}
	err = models.PatchBanner(p.db, id, patch)
	if errors.Is(err, sql.ErrNoRows) {
		c.Status(http.StatusNotFound)
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": models.ErrInternal.Error()})
		log.Println(err)
		return
	}
	c.Status(http.StatusOK)
}
