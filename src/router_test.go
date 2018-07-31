package smith

import (
	"testing"
	"reflect"
)

func TestCreateRouter(t *testing.T) {
	if reflect.TypeOf(createRouter()).String() != "*mux.Router" {
		t.Error("Assertion error")
	}
}
