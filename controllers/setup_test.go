package controllers

import (
	"gobanner/models"
	"testing"

	"github.com/gin-gonic/gin"
)

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
