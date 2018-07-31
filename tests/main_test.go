package tests

import (
	"testing"
	"github.com/moriorgames/agent-smith/lib"
)


// 	"github.com/drone/drone/version"
// https://github.com/drone/drone/version
// https://github.com/moriorgames/agent-smith/smith

func TestSum(t *testing.T) {
	if smith.add(2, 5) != 7 {
		t.Fail()
	}
	if add(2, 100) != 102 {
		t.Fail()
	}
	if add(222, 100) != 322 {
		t.Fail()
	}
}
