package smith

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
)

//func TestStaticCssIsAbleToParseFile(t *testing.T) {
//	req, err := http.NewRequest("GET", "custom.css", nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	rr := httptest.NewRecorder()
//	handler := http.HandlerFunc(StaticCss)
//	handler.ServeHTTP(rr, req)
//
//	assert.Equal(t, rr.Code, http.StatusOK)
//}

func TestStaticCssIsNOTAbleToParseFile(t *testing.T) {
	req, err := http.NewRequest("GET", "fake.css", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(StaticCss)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusNotFound)
}
