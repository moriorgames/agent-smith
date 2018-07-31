package smith

import "testing"

func TestSum(t *testing.T) {
	if add(2, 5) != 7 {
		t.Fail()
	}
	if add(2, 100) != 102 {
		t.Fail()
	}
	if add(222, 100) != 322 {
		t.Fail()
	}
}
