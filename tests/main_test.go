package tests

import (
	"testing"
	"github.com/moriorgames/agent-smith/lib"
)

func TestSum(t *testing.T) {
	if smith.Add(2, 5) != 7 {
		t.Fail()
	}
	if smith.Add(2, 100) != 102 {
		t.Fail()
	}
	if smith.Add(222, 100) != 322 {
		t.Fail()
	}
}
