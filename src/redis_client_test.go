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
				err:    errors.New("dial tcp host:port: connect: connection refused"),
			},
			"",
			errors.New("dial tcp host:port: connect: connection refused"),
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

func TestIsNotAbleToConnectToRedis(t *testing.T) {

	redisClient := ConnectToRedis()
	result, err := redisClient.Ping()
	t.Logf("Result: %s Error %s", result, err)
	if err == nil {
		t.Fail()
	}
}

func TestIsAbleToConnectToRedis(t *testing.T) {

	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	redisClient := ConnectToRedis()
	result, err := redisClient.Ping()
	t.Logf("Result: %s Error %s", result, err)
	if err != nil && result != "PONG" {
		t.Fail()
	}
}
