package smith

import (
	"testing"
	"fmt"
	"errors"
)

type RedisMock struct {
	PingFunc  func() (string, error)
	CloseFunc func() error
}

func (r *RedisMock) Ping() (string, error) {
	if r.PingFunc != nil {
		return r.Ping()
	}

	return "", fmt.Errorf("error")
}

func (r *RedisMock) Close() error {
	return fmt.Errorf("error")
}

func TestRedisClientIsAbleToConnect(t *testing.T) {

	tests := map[string]struct {
		redisClient RedisConnectable
		item        string
		err         error
	}{
		"successful": {
			&RedisMock{},
			"foo",
			errors.New("dial tcp 127.0.0.1:6379: connect: connection refused"),
		},
	}

	for name, test := range tests {
		t.Logf("Running test case: %s", name)

		result, _ := test.redisClient.Ping()
		t.Logf("Result: %s", result)
		//if result != "foo" {
		//	t.Fail()
		//}
	}
}
