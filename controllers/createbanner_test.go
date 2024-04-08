package controllers

import (
	"bytes"
	"encoding/json"
	"gobanner/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	pCtrl, router := SetupE2ETest(t)

	w := httptest.NewRecorder()
	banner := new(models.Banner)
	banner.Content = "{\"text\": \"Только сегодня и только у нас, скидка 99.9%...\"}"
	banner.IsActive = true
	banner.FeatureId = 1
	banner.TagIds = []int{1, 2, 3, 4}
	body, err := json.Marshal(&banner)
	if err != nil {
		t.Fatal(err)
	}
	r := bytes.NewReader(body)
	req, _ := http.NewRequest("POST", "/banner", r)
	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)
	_, err = models.CleanDatabase(pCtrl.db)
	if err != nil {
		t.Fatal(err)
	}
}
