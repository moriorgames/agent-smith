package smith

import (
	"testing"
	"reflect"
)

func TestCreateRouter(t *testing.T) {
	if reflect.TypeOf(CreateRouter()).String() != "*mux.Router" {
		t.Error("Assertion error")
	}
}
