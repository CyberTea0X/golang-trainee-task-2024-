package controllers

import (
	"fmt"
	"gobanners/models"
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
		t.Run(fmt.Sprintf("test_%d_mustfail=%t", i+1, tc.mustfail), func(t *testing.T) {
			pCtrl, router := SetupE2ETest(t)
			_, err := tc.banner.InsertToDB(pCtrl.db)
			if err != nil {
				t.Fatal(err)
			}
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/user_banner", nil)
			q := req.URL.Query()
			q.Add("tag_id", fmt.Sprint(tc.banner.TagIds[0]))
			q.Add("feature_id", fmt.Sprint(tc.banner.FeatureId))
			q.Add("use_last_revision", "true")
			req.URL.RawQuery = q.Encode()
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
