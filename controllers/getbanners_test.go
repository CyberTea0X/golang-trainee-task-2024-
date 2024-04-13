package controllers

import (
	"encoding/json"
	"gobanner/models"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBannersSucceed(t *testing.T) {
	pCtrl, router := SetupE2ETest(t)
	banners := []models.Banner{
		{TagIds: []int64{1, 2, 3}, FeatureId: 0, Content: "{}", IsActive: true},
		{TagIds: []int64{2, 1, 3}, FeatureId: 0, Content: "{}", IsActive: false},
		{TagIds: []int64{}, FeatureId: 2, Content: "{}", IsActive: true},
	}
	for _, b := range banners {
		_, err := b.InsertToDB(pCtrl.db)
		if err != nil {
			t.Fatal(err)
		}
	}

	req, _ := http.NewRequest("GET", "/banner", nil)
	q := req.URL.Query()
	q.Add("tag_id", "1")
	q.Add("feature_id", "0")
	q.Add("offset", "0")
	q.Add("limit", "2")
	req.URL.RawQuery = q.Encode()
	req.Header.Add("token", AdminToken)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	resBody, err := io.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}
	resBanners := make([]models.Banner, 3)
	err = json.Unmarshal(resBody, &resBanners)
	if err != nil {
		t.Fatal(err)
	}
	for i := range resBanners {
		banners[i].Id = resBanners[i].Id
	}
	assert.Equal(t, banners[:2], resBanners)
	err = models.CleanDatabase(pCtrl.db)
	if err != nil {
		t.Fatal(err)
	}
}
