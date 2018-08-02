package smith

import "testing"

func TestRedisClientIsAbleToConnect(t *testing.T) {

	ExampleNewClient()

	if add(2, 5) != 7 {
		t.Fail()
	}
}
