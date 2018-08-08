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

func TestStaticJsIsAbleToParseFile(t *testing.T) {
	req, err := http.NewRequest("GET", "/static/js/custom.js", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := CreateRouter()
	router.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
}

func TestStaticJsIsNOTAbleToParseFile(t *testing.T) {
	req, err := http.NewRequest("GET", "/static/js/fake.js", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := CreateRouter()
	router.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusNotFound)
}

func TestStaticImgIsAbleToParseFile(t *testing.T) {
	req, err := http.NewRequest("GET", "/static/img/smith-logo.jpg", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := CreateRouter()
	router.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
}

func TestStaticImgIsNOTAbleToParseFile(t *testing.T) {
	req, err := http.NewRequest("GET", "/static/img/smith-logo.svg", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := CreateRouter()
	router.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusNotFound)
}
