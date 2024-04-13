package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gobanner/models"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCasesGetBanner = []struct {
	mustfail     bool
	expectedCode int
	banner       models.Banner
}{
	{
		false, 200, models.Banner{
			TagIds:    []int64{1, 2, 3},
			FeatureId: 0,
			Content:   "{}",
			IsActive:  true,
		},
	},
	{
		true, 404, models.Banner{
			TagIds:    []int64{1, 2, 3},
			FeatureId: 0,
			Content:   "{}",
			IsActive:  false,
		},
	},
}

func TestGetBanner(t *testing.T) {
	for i, tc := range testCasesGetBanner {
		t.Run(fmt.Sprintf("test_%d_mustfail=%t", i, tc.mustfail), func(t *testing.T) {
			pCtrl, router := SetupE2ETest(t)
			_, err := tc.banner.InsertToDB(pCtrl.db)
			if err != nil {
				t.Fatal(err)
			}
			i := new(getBannerInput)
			i.TagId = addr(tc.banner.TagIds[0])
			i.FeatureId = addr(tc.banner.FeatureId)
			i.UseLastRevision = true
			body, err := json.Marshal(i)
			if err != nil {
				t.Fatal(err)
			}
			r := bytes.NewReader(body)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/user_banner", r)
			req.Header.Add("token", UserToken)
			router.ServeHTTP(w, req)
			assert.Equal(t, tc.expectedCode, w.Code)
			if tc.mustfail {
				err = models.CleanDatabase(pCtrl.db)
				if err != nil {
					t.Fatal(err)
				}
				return
			}
			resBody, err := io.ReadAll(w.Body)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tc.banner.Content, string(resBody))
			err = models.CleanDatabase(pCtrl.db)
			if err != nil {
				t.Fatal(err)
			}

		})
	}
}
