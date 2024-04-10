package controllers

import (
	"gobanner/models"
	"testing"

	"github.com/gin-gonic/gin"
)

const (
	UserToken  = "user_token"
	AdminToken = "admin_token"
)

func newTestBanner() *models.Banner {
	return &models.Banner{
		TagIds:    []int{1, 2, 3},
		Content:   "{\"text\": \"Только сегодня и только у нас, скидка 99.9%...\"}",
		FeatureId: 0,
		IsActive:  false,
	}
}

func SetupE2ETest(t *testing.T) (*PublicController, *gin.Engine) {
	dbconf, err := models.DBConfigFromEnv("../test.env")
	if err != nil {
		t.Fatal(err)
	}
	db, err := models.SetupDatabase(dbconf)
	if err != nil {
		t.Fatal(err)
	}
	pCtrl := NewPublicController(db)
	return pCtrl, SetupRouter(pCtrl)
}
