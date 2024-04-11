package controllers

import (
	"bytes"
	"encoding/json"
	"gobanner/models"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBannerSucceed(t *testing.T) {
	pCtrl, router := SetupE2ETest(t)
	banner := newTestBanner()
	_, err := banner.InsertToDB(pCtrl.db)
	if err != nil {
		t.Fatal(err)
	}
	i := new(getBannerInput)
	i.TagId = addr(banner.TagIds[0])
	i.FeatureId = addr(banner.FeatureId)
	i.UseLastRevision = banner.IsActive
	body, err := json.Marshal(i)
	if err != nil {
		t.Fatal(err)
	}
	r := bytes.NewReader(body)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user_banner", r)
	req.Header.Add("token", AdminToken)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	resBody, err := io.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, banner.Content, string(resBody))
	_, err = models.CleanDatabase(pCtrl.db)
	if err != nil {
		t.Fatal(err)
	}
}
