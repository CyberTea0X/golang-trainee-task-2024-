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

	i := new(getBannersInput)
	i.TagId = addr(int64(1))
	i.FeatureId = addr(int64(0))
	i.Offset = addr(int64(0))
	i.Limit = addr(int64(2))
	body, err := json.Marshal(i)
	if err != nil {
		t.Fatal(err)
	}
	r := bytes.NewReader(body)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/banner", r)
	req.Header.Add("token", AdminToken)
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
	for i := range banners {
		banners[i].Id = resBanners[i].Id
	}
	assert.Equal(t, banners, resBanners)
	_, err = models.CleanDatabase(pCtrl.db)
	if err != nil {
		t.Fatal(err)
	}
}
