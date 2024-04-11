package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPatchBannerSucceed(t *testing.T) {
	pCtrl, router := SetupE2ETest(t)
	banner := newTestBanner()
	id, err := banner.InsertToDB(pCtrl.db)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	input := patchBannerInput{
		TagIds:   []int64{0, 1},
		IsActive: addr(true),
	}
	body, err := json.Marshal(input)
	if err != nil {
		t.Fatal(err)
	}
	r := bytes.NewReader(body)
	req, _ := http.NewRequest("PATCH", fmt.Sprintf("/banner/%d", id), r)
	req.Header.Add("token", AdminToken)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
