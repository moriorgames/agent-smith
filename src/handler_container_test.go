package smith

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
)

func TestIsAbleToFindNewContainerRoute(t *testing.T) {
	req, err := http.NewRequest("GET", "/container", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := CreateRouter()
	router.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
}

func TestIsAbleToFindEditContainerRoute(t *testing.T) {
	req, err := http.NewRequest("GET", "/container/1234", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := CreateRouter()
	router.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
}

func TestUpdateContainerRedirectsToEditView(t *testing.T) {
	req, err := http.NewRequest("POST", "/container/1234", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := CreateRouter()
	router.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusSeeOther)
}

func TestNewContainerCreatesUuidKey(t *testing.T) {
	req, err := http.NewRequest("POST", "/container", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := CreateRouter()
	router.ServeHTTP(rr, req)
	//
	//assert.Equal(t, rr.Code, http.StatusSeeOther)
}
