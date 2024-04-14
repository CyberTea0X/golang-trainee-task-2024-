package controllers

import (
	"gobanners/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const (
	UserToken  = "user_token"
	AdminToken = "admin_token"
)

func addr[T any](v T) *T { return &v }

func newTestBanner() *models.Banner {
	banner := new(models.Banner)
	banner.TagIds = []int64{1, 2, 3}
	banner.Content = "{\"text\": \"Только сегодня и только у нас, скидка 99.9%...\"}"
	banner.FeatureId = 0
	banner.IsActive = true
	return banner
}

func SetupE2ETest(t *testing.T) (*PublicController, *gin.Engine) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Log(err)
	}
	dbconf, err := models.DBConfigFromEnv("../test.env")
	if err != nil {
		t.Fatal(err)
	}
	db, err := models.SetupDatabase(dbconf)
	if err != nil {
		t.Fatal(err)
	}
	if err = models.MigrateDatabase(db); err != nil {
		t.Fatal(err)
	}
	pCtrl := NewPublicController(db)
	return pCtrl, SetupRouter(pCtrl)
}
