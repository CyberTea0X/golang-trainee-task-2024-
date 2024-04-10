package controllers

import (
	"fmt"
	"gobanner/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p *PublicController) deleteBanner(c *gin.Context) {
	var id int64
	param := c.Param("id")
	_, err := fmt.Sscan(param, &id)
	// if param is "", Sscan will return error EOF
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": models.ErrInvalidId.Error()})
		return
	}
	err = models.DeleteBannerFromDB(p.db, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": models.ErrInternal.Error()})
		log.Println(err)
		return
	}
	c.Status(http.StatusNoContent)
}
