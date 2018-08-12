package smith

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
)

func TestHandlerHomeIsAbleToListDockerContainers(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := CreateRouter()
	router.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
}
