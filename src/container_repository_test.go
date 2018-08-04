package smith

import (
	"testing"
	"errors"
)

func TestRepositoryIsAbleToPersistModel(t *testing.T) {

	mockClient := new(RedisMock)
	mockClient.result = "test_key"
	mockClient.err = errors.New("error")

	containerRepository := NewContainerRepository(mockClient)
	containerRepository.FindByID("test_key")
}
