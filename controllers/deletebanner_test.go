package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteBannerSucceed(t *testing.T) {
	pCtrl, router := SetupE2ETest(t)
	banner := newTestBanner()
	id, err := banner.Insert(pCtrl.db)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/banner/%d", id), nil)
	req.Header.Add("token", AdminToken)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)
}
