package smith

import (
	"testing"
	"errors"
)

type RedisMock struct {
	result string
	err    error
}

func (r *RedisMock) Ping() (string, error) {

	return r.result, r.err
}

func TestInterfaceRedisConnectable(t *testing.T) {

	tests := map[string]struct {
		mockClient     RedisConnectable
		expectedResult string
		expectedErr    error
	}{
		"successful": {
			&RedisMock{
				result: "PONG",
				err:    nil,
			},
			"PONG",
			nil,
		},
		"failure": {
			&RedisMock{
				result: "",
				err:    errors.New("dial tcp 127.0.0.1:6379: connect: connection refused"),
			},
			"",
			errors.New("dial tcp 127.0.0.1:6379: connect: connection refused"),
		},
	}

	for name, test := range tests {
		t.Logf("Running test case: %s", name)
		result, err := test.mockClient.Ping()
		if result != test.expectedResult && err != test.expectedErr {
			t.Fail()
		}
	}
}
