package smith

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
)

func TestStaticCssIsAbleToParseFile(t *testing.T) {
	req, err := http.NewRequest("GET", "/static/css/custom.css", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := CreateRouter()
	router.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
}

func TestStaticCssIsNOTAbleToParseFile(t *testing.T) {
	req, err := http.NewRequest("GET", "/static/css/fake.css", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := CreateRouter()
	router.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusNotFound)
}
