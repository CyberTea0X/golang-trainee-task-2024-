package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gobanners/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCasesPatchBanner = []struct {
	expectedCode int
	banner       *models.Banner
	patch        patchBannerInput
}{
	{
		200, &models.Banner{
			TagIds:    []int64{1, 2, 3},
			FeatureId: 0,
			Content:   "{}",
			IsActive:  true,
		},
		patchBannerInput{
			TagIds:    []int64{2, 3},
			FeatureId: addr(int64(1)),
		},
	},
	{
		404, nil,
		patchBannerInput{},
	},
	{
		404, nil,
		patchBannerInput{
			TagIds:    []int64{2, 3},
			FeatureId: addr(int64(1)),
		},
	},
}

func TestPatchBanner(t *testing.T) {
	for i, tc := range testCasesPatchBanner {
		t.Run(fmt.Sprintf("test_%d", i+1), func(t *testing.T) {
			pCtrl, router := SetupE2ETest(t)
			var id int64
			var err error
			if tc.banner != nil {
				banner := tc.banner
				id, err = banner.InsertToDB(pCtrl.db)
			} else {
				id = -1
			}
			if err != nil {
				t.Fatal(err)
			}
			w := httptest.NewRecorder()
			input := tc.patch
			body, err := json.Marshal(input)
			if err != nil {
				t.Fatal(err)
			}
			t.Log(string(body))
			r := bytes.NewReader(body)
			req, _ := http.NewRequest("PATCH", fmt.Sprintf("/banner/%d", id), r)
			req.Header.Add("token", AdminToken)

			router.ServeHTTP(w, req)
			assert.Equal(t, tc.expectedCode, w.Code)
			err = models.CleanDatabase(pCtrl.db)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}
