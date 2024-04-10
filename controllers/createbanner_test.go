package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gobanner/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBannerSucceed(t *testing.T) {
	pCtrl, router := SetupE2ETest(t)

	w := httptest.NewRecorder()
	banner := newTestBanner()
	body, err := json.Marshal(banner)
	fmt.Println(string(body))
	if err != nil {
		t.Fatal(err)
	}
	r := bytes.NewReader(body)
	req, _ := http.NewRequest("POST", "/banner", r)
	req.Header.Add("token", AdminToken)
	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)
	_, err = models.CleanDatabase(pCtrl.db)
	if err != nil {
		t.Fatal(err)
	}
}
